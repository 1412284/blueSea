package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	otp "netacom.vn/netachat/otp/delivery/grpc/pb"
)

func main() {
	cc, err := grpc.Dial("localhost:5069", grpc.WithInsecure())
	if err != nil {
		log.Fatal("err while dial %v", err)
	}
	defer cc.Close()
	client := otp.NewOTPServiceClient(cc)
	callOTP(client)
	// log.Printf("service client %f", client)
}

func callOTP(o otp.OTPServiceClient) {
	log.Println("calling otp api")
	resp, err := o.GetOTPSMS(context.Background(), &otp.OTPMTRequest{
		Phone: "84974858367",
		Text:  "Ma OTP cua may ban la:",
	})

	if err != nil {
		log.Fatalf("call otp api err %v", err)
	}

	log.Printf("otp api response %v\n", resp)
}
