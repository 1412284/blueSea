package main

import (
	"fmt"
	"log"
	"net"

	otpHandler "netacom.vn/netachat/otp/delivery/grpc"

	"google.golang.org/grpc"
	otp "netacom.vn/netachat/otp/delivery/grpc/pb"
)

func main() {
	log.Printf("Server started")
	lis, err := net.Listen("tcp", "0.0.0.0:5069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}
	serverGrpc := grpc.NewServer()
	server := otpHandler.OTPServer{}
	otp.RegisterOTPServiceServer(serverGrpc, &server)
	fmt.Printf("otp is running... %+v", &server)
	err = serverGrpc.Serve(lis)

	if err != nil {
		log.Fatalf("err while serve %v", err)
	}

}
