// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: creditService.proto

package creditService

import (
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

type GetCardByUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetCardByUserIDRequest) Reset() {
	*x = GetCardByUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCardByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCardByUserIDRequest) ProtoMessage() {}

func (x *GetCardByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_creditService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCardByUserIDRequest.ProtoReflect.Descriptor instead.
func (*GetCardByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_creditService_proto_rawDescGZIP(), []int{0}
}

func (x *GetCardByUserIDRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetCardByUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards []*Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *GetCardByUserIDResponse) Reset() {
	*x = GetCardByUserIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCardByUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCardByUserIDResponse) ProtoMessage() {}

func (x *GetCardByUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_creditService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCardByUserIDResponse.ProtoReflect.Descriptor instead.
func (*GetCardByUserIDResponse) Descriptor() ([]byte, []int) {
	return file_creditService_proto_rawDescGZIP(), []int{1}
}

func (x *GetCardByUserIDResponse) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Number    string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	Balance   string `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_creditService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_creditService_proto_rawDescGZIP(), []int{2}
}

func (x *Card) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Card) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Card) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Card) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *Card) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Card) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateCardBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Balance string `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *UpdateCardBalanceRequest) Reset() {
	*x = UpdateCardBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCardBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCardBalanceRequest) ProtoMessage() {}

func (x *UpdateCardBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_creditService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCardBalanceRequest.ProtoReflect.Descriptor instead.
func (*UpdateCardBalanceRequest) Descriptor() ([]byte, []int) {
	return file_creditService_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCardBalanceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCardBalanceRequest) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

type UpdateCardBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Number    string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	Balance   string `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UpdateCardBalanceResponse) Reset() {
	*x = UpdateCardBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creditService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCardBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCardBalanceResponse) ProtoMessage() {}

func (x *UpdateCardBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_creditService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCardBalanceResponse.ProtoReflect.Descriptor instead.
func (*UpdateCardBalanceResponse) Descriptor() ([]byte, []int) {
	return file_creditService_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCardBalanceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCardBalanceResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateCardBalanceResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *UpdateCardBalanceResponse) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *UpdateCardBalanceResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UpdateCardBalanceResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_creditService_proto protoreflect.FileDescriptor

var file_creditService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x31, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x22, 0x9f, 0x01,
	0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x44, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xdd, 0x01, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x25, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x68, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e,
	0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creditService_proto_rawDescOnce sync.Once
	file_creditService_proto_rawDescData = file_creditService_proto_rawDesc
)

func file_creditService_proto_rawDescGZIP() []byte {
	file_creditService_proto_rawDescOnce.Do(func() {
		file_creditService_proto_rawDescData = protoimpl.X.CompressGZIP(file_creditService_proto_rawDescData)
	})
	return file_creditService_proto_rawDescData
}

var file_creditService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_creditService_proto_goTypes = []interface{}{
	(*GetCardByUserIDRequest)(nil),    // 0: creditService.GetCardByUserIDRequest
	(*GetCardByUserIDResponse)(nil),   // 1: creditService.GetCardByUserIDResponse
	(*Card)(nil),                      // 2: creditService.Card
	(*UpdateCardBalanceRequest)(nil),  // 3: creditService.UpdateCardBalanceRequest
	(*UpdateCardBalanceResponse)(nil), // 4: creditService.UpdateCardBalanceResponse
}
var file_creditService_proto_depIdxs = []int32{
	2, // 0: creditService.GetCardByUserIDResponse.cards:type_name -> creditService.Card
	0, // 1: creditService.CreditService.GetCardByUserID:input_type -> creditService.GetCardByUserIDRequest
	3, // 2: creditService.CreditService.UpdateCardBalance:input_type -> creditService.UpdateCardBalanceRequest
	1, // 3: creditService.CreditService.GetCardByUserID:output_type -> creditService.GetCardByUserIDResponse
	4, // 4: creditService.CreditService.UpdateCardBalance:output_type -> creditService.UpdateCardBalanceResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_creditService_proto_init() }
func file_creditService_proto_init() {
	if File_creditService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_creditService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCardByUserIDRequest); i {
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
		file_creditService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCardByUserIDResponse); i {
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
		file_creditService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_creditService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCardBalanceRequest); i {
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
		file_creditService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCardBalanceResponse); i {
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
			RawDescriptor: file_creditService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_creditService_proto_goTypes,
		DependencyIndexes: file_creditService_proto_depIdxs,
		MessageInfos:      file_creditService_proto_msgTypes,
	}.Build()
	File_creditService_proto = out.File
	file_creditService_proto_rawDesc = nil
	file_creditService_proto_goTypes = nil
	file_creditService_proto_depIdxs = nil
}
