// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: OTP/delivery/grpc/pb/otp.proto

package otp

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OTPMTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Text  string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *OTPMTRequest) Reset() {
	*x = OTPMTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OTP_delivery_grpc_pb_otp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPMTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPMTRequest) ProtoMessage() {}

func (x *OTPMTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_OTP_delivery_grpc_pb_otp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPMTRequest.ProtoReflect.Descriptor instead.
func (*OTPMTRequest) Descriptor() ([]byte, []int) {
	return file_OTP_delivery_grpc_pb_otp_proto_rawDescGZIP(), []int{0}
}

func (x *OTPMTRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *OTPMTRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type OTPMTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *OTPMTResponse) Reset() {
	*x = OTPMTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OTP_delivery_grpc_pb_otp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPMTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPMTResponse) ProtoMessage() {}

func (x *OTPMTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_OTP_delivery_grpc_pb_otp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPMTResponse.ProtoReflect.Descriptor instead.
func (*OTPMTResponse) Descriptor() ([]byte, []int) {
	return file_OTP_delivery_grpc_pb_otp_proto_rawDescGZIP(), []int{1}
}

func (x *OTPMTResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_OTP_delivery_grpc_pb_otp_proto protoreflect.FileDescriptor

var file_OTP_delivery_grpc_pb_otp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x4f, 0x54, 0x50, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x6f, 0x74, 0x70, 0x73, 0x6d, 0x73, 0x22, 0x38, 0x0a, 0x0c, 0x4f, 0x54, 0x50, 0x4d,
	0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65,
	0x78, 0x74, 0x22, 0x27, 0x0a, 0x0d, 0x4f, 0x54, 0x50, 0x4d, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x48, 0x0a, 0x0a, 0x4f,
	0x54, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x4f, 0x54, 0x50, 0x53, 0x4d, 0x53, 0x12, 0x14, 0x2e, 0x6f, 0x74, 0x70, 0x73, 0x6d, 0x73, 0x2e,
	0x4f, 0x54, 0x50, 0x4d, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6f,
	0x74, 0x70, 0x73, 0x6d, 0x73, 0x2e, 0x4f, 0x54, 0x50, 0x4d, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x6e, 0x65, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x61, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6f, 0x74, 0x70,
	0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x64, 0x2f, 0x6f, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OTP_delivery_grpc_pb_otp_proto_rawDescOnce sync.Once
	file_OTP_delivery_grpc_pb_otp_proto_rawDescData = file_OTP_delivery_grpc_pb_otp_proto_rawDesc
)

func file_OTP_delivery_grpc_pb_otp_proto_rawDescGZIP() []byte {
	file_OTP_delivery_grpc_pb_otp_proto_rawDescOnce.Do(func() {
		file_OTP_delivery_grpc_pb_otp_proto_rawDescData = protoimpl.X.CompressGZIP(file_OTP_delivery_grpc_pb_otp_proto_rawDescData)
	})
	return file_OTP_delivery_grpc_pb_otp_proto_rawDescData
}

var file_OTP_delivery_grpc_pb_otp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_OTP_delivery_grpc_pb_otp_proto_goTypes = []interface{}{
	(*OTPMTRequest)(nil),  // 0: otpsms.OTPMTRequest
	(*OTPMTResponse)(nil), // 1: otpsms.OTPMTResponse
}
var file_OTP_delivery_grpc_pb_otp_proto_depIdxs = []int32{
	0, // 0: otpsms.OTPService.GetOTPSMS:input_type -> otpsms.OTPMTRequest
	1, // 1: otpsms.OTPService.GetOTPSMS:output_type -> otpsms.OTPMTResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_OTP_delivery_grpc_pb_otp_proto_init() }
func file_OTP_delivery_grpc_pb_otp_proto_init() {
	if File_OTP_delivery_grpc_pb_otp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OTP_delivery_grpc_pb_otp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPMTRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_OTP_delivery_grpc_pb_otp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPMTResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_OTP_delivery_grpc_pb_otp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_OTP_delivery_grpc_pb_otp_proto_goTypes,
		DependencyIndexes: file_OTP_delivery_grpc_pb_otp_proto_depIdxs,
		MessageInfos:      file_OTP_delivery_grpc_pb_otp_proto_msgTypes,
	}.Build()
	File_OTP_delivery_grpc_pb_otp_proto = out.File
	file_OTP_delivery_grpc_pb_otp_proto_rawDesc = nil
	file_OTP_delivery_grpc_pb_otp_proto_goTypes = nil
	file_OTP_delivery_grpc_pb_otp_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OTPServiceClient is the client API for OTPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OTPServiceClient interface {
	GetOTPSMS(ctx context.Context, in *OTPMTRequest, opts ...grpc.CallOption) (*OTPMTResponse, error)
}

type oTPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOTPServiceClient(cc grpc.ClientConnInterface) OTPServiceClient {
	return &oTPServiceClient{cc}
}

func (c *oTPServiceClient) GetOTPSMS(ctx context.Context, in *OTPMTRequest, opts ...grpc.CallOption) (*OTPMTResponse, error) {
	out := new(OTPMTResponse)
	err := c.cc.Invoke(ctx, "/otpsms.OTPService/GetOTPSMS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OTPServiceServer is the server API for OTPService service.
type OTPServiceServer interface {
	GetOTPSMS(context.Context, *OTPMTRequest) (*OTPMTResponse, error)
}

// UnimplementedOTPServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOTPServiceServer struct {
}

func (*UnimplementedOTPServiceServer) GetOTPSMS(context.Context, *OTPMTRequest) (*OTPMTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOTPSMS not implemented")
}

func RegisterOTPServiceServer(s *grpc.Server, srv OTPServiceServer) {
	s.RegisterService(&_OTPService_serviceDesc, srv)
}

func _OTPService_GetOTPSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPMTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OTPServiceServer).GetOTPSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otpsms.OTPService/GetOTPSMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OTPServiceServer).GetOTPSMS(ctx, req.(*OTPMTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OTPService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "otpsms.OTPService",
	HandlerType: (*OTPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOTPSMS",
			Handler:    _OTPService_GetOTPSMS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OTP/delivery/grpc/pb/otp.proto",
}
