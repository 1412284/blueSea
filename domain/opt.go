package domain

import (
	"encoding/xml"
	"net/http"
)

type ConfigMTSendOTP struct {
	URL      string
	USERNAME string
	PASSWORD string
}
type SendMessageRequest struct {
	Phone string
	Text  string
}

type Response struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapBody *SOAPBodyResponse
}

type SOAPBodyResponse struct {
	XMLName      xml.Name `xml:"Body"`
	Resp         *ResponseBody
	FaultDetails *Fault
}

type Fault struct {
	XMLName     xml.Name `xml:"Fault"`
	Faultcode   string   `xml:"faultcode"`
	Faultstring string   `xml:"faultstring"`
}

type ResponseBody struct {
	XMLName xml.Name `xml:"sendMTResponse"`
	Result  int      `xml:"result"`
}

type OTPUseCase interface {
	GenerateSOAPRequestMTOTP(config *ConfigMTSendOTP, req *SendMessageRequest) (*http.Request, error)
	SoapCallMTOTP(req *http.Request) (*Response, error)
}
