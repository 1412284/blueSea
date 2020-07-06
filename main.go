package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

//MTRequest .........
type MTRequest struct {
	XMLName      xml.Name `xml:"MTRequest"`
	Id           int      `xml:"id,attr"`
	name         string   `xml:"name"`
	userID       string   `xml:"user_id"`
	message      string   `xml:"message"`
	serviceID    string   `xml:"service_id"`
	commandCode  string   `xml:"command_code"`
	messageType  string   `xml:"message_type"`
	requestID    string   `xml:"request_id"`
	totalMessage string   `xml:"total_message"`
	messageIndex string   `xml:"message_index"`
	isMore       string   `xml:"is_more"`
	contentType  string   `xml:"content_type"`
}

//KTResponse .......................
type MTResponse struct {
	status int `xml:status`
}

func main() {
	var user string = "testapi"
	var password string = "testapi22446688()"
	var url string = " http://sms.8x77.vn:8077/mt-services/MTService"
	reqMT := &MTRequest{
		Id:           27,
		name:         "MTResponse",
		userID:       "84902866568",
		message:      "abc",
		serviceID:    "8077",
		commandCode:  "CSKH",
		messageType:  "0",
		requestID:    "123456",
		totalMessage: "1",
		messageIndex: "1",
		isMore:       "1",
		contentType:  "1",
	}

	out, _ := xml.MarshalIndent(reqMT, " ", "  ")
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))
	result := getResponseMT(user, password, url, reqMT)
	fmt.Println(result)
}

func getResponseMT(user string, password string, url string, reqMT *MTRequest) int {
	client := &http.Client{}
	body, _ := xml.MarshalIndent(reqMT, " ", "  ")
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err)
	}
	// you can then set the Header here
	// I think the content-type should be "application/xml" like json...
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	req.SetBasicAuth(user, password)
	// now POST it
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	return 1
}
