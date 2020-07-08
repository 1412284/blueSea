package main

import (
	"fmt"
	"log"
	"net"

	otpHandler "netacom.vn/netachat/otp/delivery/grpc"

	"google.golang.org/grpc"
	otp "netacom.vn/netachat/otp/delivery/grpc/pb"
)

// type server struct{}

// func (*server) GetOTPSMS(ctx context.Context, req *otp.OTPMTRequest) (*otp.OTPMTResponse, error) {
// 	log.Println("GetOTPSMS called...")
// 	res, err := otpHandler.OTPService.GetOTPSMS(ctx, req)
// 	return res, err
// }

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}
	s := grpc.NewServer()
	server := otpHandler.OTPServer{}
	otp.RegisterOTPServiceServer(s, &server)
	fmt.Println("otp is running...")
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}

}
