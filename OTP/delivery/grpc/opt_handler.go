package otpHandler

import (
	"context"
	"encoding/base64"
	"fmt"

	"netacom.vn/netachat/domain"
	otp "netacom.vn/netachat/otp/delivery/grpc/pb"
)

type OTPServer struct {
	OTPUsecase domain.OTPUseCase
}

func (otps *OTPServer) GetOTPSMS(ctx context.Context, req *otp.OTPMTRequest) (*otp.OTPMTResponse, error) {
	fmt.Println("GetOTPSMS ...")
	const HOST = "167.71.218.111"
	var url = fmt.Sprintf("http://%s:8077/mt-services/MTService", HOST)
	config := domain.ConfigMTSendOTP{
		URL:      url,
		USERNAME: "netacom",
		PASSWORD: "netacom@!NET789",
	}

	user := domain.SendMessageRequest{
		Phone: req.Phone,
		Text:  base64.StdEncoding.EncodeToString([]byte(req.Text)),
	}

	httpReq, err := otps.OTPUsecase.GenerateSOAPRequestMTOTP(&config, &user)
	if err != nil {
		fmt.Println("Some problem occurred in request generation")
	}

	defer httpReq.Body.Close()

	response, err := otps.OTPUsecase.SoapCallMTOTP(httpReq)
	if err != nil {
		fmt.Println("Problem occurred in making a SOAP call")
	}
	fmt.Printf("SOAP response: %v\n", response)
	if err == nil {
		return &otp.OTPMTResponse{
				Status: "Success",
			},
			nil
	} else {
		return nil,
			err
	}

}
