// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: rpc_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_rpc_service_proto protoreflect.FileDescriptor

var file_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbc, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x60, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_rpc_service_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),          // 0: service_payment.HelloRequest
	(*PaymentCreateRequest)(nil),  // 1: service_payment.PaymentCreateRequest
	(*HelloResponse)(nil),         // 2: service_payment.HelloResponse
	(*PaymentCreateResponse)(nil), // 3: service_payment.PaymentCreateResponse
}
var file_rpc_service_proto_depIdxs = []int32{
	0, // 0: service_payment.PaymentService.Hello:input_type -> service_payment.HelloRequest
	1, // 1: service_payment.PaymentService.PaymentCreate:input_type -> service_payment.PaymentCreateRequest
	2, // 2: service_payment.PaymentService.Hello:output_type -> service_payment.HelloResponse
	3, // 3: service_payment.PaymentService.PaymentCreate:output_type -> service_payment.PaymentCreateResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_service_proto_init() }
func file_rpc_service_proto_init() {
	if File_rpc_service_proto != nil {
		return
	}
	file_rpc_hello_proto_init()
	file_rpc_payment_information_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_service_proto_goTypes,
		DependencyIndexes: file_rpc_service_proto_depIdxs,
	}.Build()
	File_rpc_service_proto = out.File
	file_rpc_service_proto_rawDesc = nil
	file_rpc_service_proto_goTypes = nil
	file_rpc_service_proto_depIdxs = nil
}