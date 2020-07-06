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
	ID           int      `xml:"id,attr"`
	Name         string   `xml:"name"`
	UserID       string   `xml:"user_id"`
	Message      string   `xml:"message"`
	ServiceID    string   `xml:"service_id"`
	CommandCode  string   `xml:"command_code"`
	MessageType  string   `xml:"message_type"`
	RequestID    string   `xml:"request_id"`
	TotalMessage string   `xml:"total_message"`
	MessageIndex string   `xml:"message_index"`
	IsMore       string   `xml:"is_more"`
	ContentType  string   `xml:"content_type"`
}

// //KTResponse .......................
// type MTResponse struct {
// 	Status int `xml:status`
// }

func main() {
	var user string = "testapi"
	var password string = "testapi22446688()"
	var url string = " http://sms.8x77.vn:8077/mt-services/MTService"
	reqMT := &MTRequest{
		ID:           27,
		Name:         "MTResponse",
		UserID:       "84902866568",
		Message:      "abc",
		ServiceID:    "8077",
		CommandCode:  "CSKH",
		MessageType:  "0",
		RequestID:    "123456",
		TotalMessage: "1",
		MessageIndex: "1",
		IsMore:       "1",
		ContentType:  "1",
	}
	result := getResponseMT(user, password, url, reqMT)
	fmt.Println(result)
}

func getResponseMT(user string, password string, url string, reqMT *MTRequest) int {
	client := &http.Client{}
	body, _ := xml.MarshalIndent(reqMT, " ", "  ")
	fmt.Println(string(body))
	fmt.Println(xml.Header + string(body))
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
