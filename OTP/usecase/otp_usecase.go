package usecase

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"netacom.vn/netachat/domain"
)

var getTemplate = `
<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
	xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" 
	xmlns:mts="MTService">
	<soapenv:Header/>
	<soapenv:Body>
		<mts:sendMT soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
		<string xsi:type="xsd:string">{{.Phone}}</string>
		<string0 xsi:type="xsd:string">{{.Text}}</string0>
		<string1 xsi:type="xsd:string">NETACOM</string1>
		<string2 xsi:type="xsd:string">NETACOM</string2>
		<string3 xsi:type="xsd:string">0</string3>
		<string4 xsi:type="xsd:string">0</string4>
		<string5 xsi:type="xsd:string">0</string5>
		<string6 xsi:type="xsd:string">0</string6>
		<string7 xsi:type="xsd:string">0</string7>
		<string8 xsi:type="xsd:string">0</string8>
	</mts:sendMT> 
	</soapenv:Body>
</soapenv:Envelope>"
`

type otpUsecase struct {
	optRepo domain.OTPRepository
}

func GenerateSOAPRequestMTOTP(config *domain.ConfigMTSendOTP, req *domain.SendMessageRequest) (*http.Request, error) {
	// Using the var getTemplate to construct request
	template, err := template.New("InputRequest").Parse(getTemplate)
	if err != nil {
		fmt.Println("Error while marshling object. %s ", err.Error())
		return nil, err
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, req)
	if err != nil {
		fmt.Println("template.Execute error. %s ", err.Error())
		return nil, err
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		fmt.Println("encoder.Encode error. %s ", err.Error())
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, config.URL, bytes.NewBuffer([]byte(doc.String())))
	if err != nil {
		fmt.Println("Error making a request. %s ", err.Error())
		return nil, err
	}

	r.Header.Set("Content-Type", "text/xml;charset=utf-8")
	r.SetBasicAuth(config.USERNAME, config.PASSWORD)

	return r, nil
}

func SoapCallMTOTP(req *http.Request) (*domain.Response, error) {
	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("Soap response: %v\n", string(body))

	r := &domain.Response{}
	err = xml.Unmarshal(body, &r)

	if err != nil {
		return nil, err
	}

	if r.SoapBody.Resp.Result != 0 {
		return nil, err
	}

	return r, nil
}
