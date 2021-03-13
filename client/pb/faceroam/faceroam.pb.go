// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: faceroam.proto

package faceroam

import (
	proto "github.com/golang/protobuf/proto"
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

type PlatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Implat *int64  `protobuf:"varint,1,opt,name=implat" json:"implat,omitempty"`
	Osver  *string `protobuf:"bytes,2,opt,name=osver" json:"osver,omitempty"`
	Mqqver *string `protobuf:"bytes,3,opt,name=mqqver" json:"mqqver,omitempty"`
}

func (x *PlatInfo) Reset() {
	*x = PlatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faceroam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatInfo) ProtoMessage() {}

func (x *PlatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_faceroam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatInfo.ProtoReflect.Descriptor instead.
func (*PlatInfo) Descriptor() ([]byte, []int) {
	return file_faceroam_proto_rawDescGZIP(), []int{0}
}

func (x *PlatInfo) GetImplat() int64 {
	if x != nil && x.Implat != nil {
		return *x.Implat
	}
	return 0
}

func (x *PlatInfo) GetOsver() string {
	if x != nil && x.Osver != nil {
		return *x.Osver
	}
	return ""
}

func (x *PlatInfo) GetMqqver() string {
	if x != nil && x.Mqqver != nil {
		return *x.Mqqver
	}
	return ""
}

type FaceroamReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comm          *PlatInfo      `protobuf:"bytes,1,opt,name=comm" json:"comm,omitempty"`
	Uin           *uint64        `protobuf:"varint,2,opt,name=uin" json:"uin,omitempty"`
	SubCmd        *uint32        `protobuf:"varint,3,opt,name=subCmd" json:"subCmd,omitempty"`
	ReqUserInfo   *ReqUserInfo   `protobuf:"bytes,4,opt,name=reqUserInfo" json:"reqUserInfo,omitempty"`
	ReqDeleteItem *ReqDeleteItem `protobuf:"bytes,5,opt,name=reqDeleteItem" json:"reqDeleteItem,omitempty"`
}

func (x *FaceroamReqBody) Reset() {
	*x = FaceroamReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faceroam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceroamReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceroamReqBody) ProtoMessage() {}

func (x *FaceroamReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_faceroam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceroamReqBody.ProtoReflect.Descriptor instead.
func (*FaceroamReqBody) Descriptor() ([]byte, []int) {
	return file_faceroam_proto_rawDescGZIP(), []int{1}
}

func (x *FaceroamReqBody) GetComm() *PlatInfo {
	if x != nil {
		return x.Comm
	}
	return nil
}

func (x *FaceroamReqBody) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *FaceroamReqBody) GetSubCmd() uint32 {
	if x != nil && x.SubCmd != nil {
		return *x.SubCmd
	}
	return 0
}

func (x *FaceroamReqBody) GetReqUserInfo() *ReqUserInfo {
	if x != nil {
		return x.ReqUserInfo
	}
	return nil
}

func (x *FaceroamReqBody) GetReqDeleteItem() *ReqDeleteItem {
	if x != nil {
		return x.ReqDeleteItem
	}
	return nil
}

type ReqDeleteItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename []string `protobuf:"bytes,1,rep,name=filename" json:"filename,omitempty"`
}

func (x *ReqDeleteItem) Reset() {
	*x = ReqDeleteItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faceroam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqDeleteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDeleteItem) ProtoMessage() {}

func (x *ReqDeleteItem) ProtoReflect() protoreflect.Message {
	mi := &file_faceroam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDeleteItem.ProtoReflect.Descriptor instead.
func (*ReqDeleteItem) Descriptor() ([]byte, []int) {
	return file_faceroam_proto_rawDescGZIP(), []int{2}
}

func (x *ReqDeleteItem) GetFilename() []string {
	if x != nil {
		return x.Filename
	}
	return nil
}

type ReqUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqUserInfo) Reset() {
	*x = ReqUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faceroam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUserInfo) ProtoMessage() {}

func (x *ReqUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_faceroam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUserInfo.ProtoReflect.Descriptor instead.
func (*ReqUserInfo) Descriptor() ([]byte, []int) {
	return file_faceroam_proto_rawDescGZIP(), []int{3}
}

type FaceroamRspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret           *int64         `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
	Errmsg        *string        `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	SubCmd        *uint32        `protobuf:"varint,3,opt,name=subCmd" json:"subCmd,omitempty"`
	RspUserInfo   *RspUserInfo   `protobuf:"bytes,4,opt,name=rspUserInfo" json:"rspUserInfo,omitempty"`
	RspDeleteItem *RspDeleteItem `protobuf:"bytes,5,opt,name=rspDeleteItem" json:"rspDeleteItem,omitempty"`
}

func (x *FaceroamRspBody) Reset() {
	*x = FaceroamRspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faceroam_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceroamRspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceroamRspBody) ProtoMessage() {}

func (x *FaceroamRspBody) ProtoReflect() protoreflect.Message {
	mi := &file_faceroam_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceroamRspBody.ProtoReflect.Descriptor instead.
func (*FaceroamRspBody) Descriptor() ([]byte, []int) {
	return file_faceroam_proto_rawDescGZIP(), []int{4}
}

func (x *FaceroamRspBody) GetRet() int64 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *FaceroamRspBody) GetErrmsg() string {
	if x != nil && x.Errmsg != nil {
		return *x.Errmsg
	}
	return ""
}

func (x *FaceroamRspBody) GetSubCmd() uint32 {
	if x != nil && x.SubCmd != nil {
		return *x.SubCmd
	}
	return 0
}

func (x *FaceroamRspBody) GetRspUserInfo() *RspUserInfo {
	if x != nil {
		return x.RspUserInfo
	}
	return nil
}

func (x *FaceroamRspBody) GetRspDeleteItem() *RspDeleteItem {
	if x != nil {
		return x.RspDeleteItem
	}
	return nil
}

type RspDeleteItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename []string `protobuf:"bytes,1,rep,name=filename" json:"filename,omitempty"`
	Ret      []int64  `protobuf:"varint,2,rep,name=ret" json:"ret,omitempty"`
}

func (x *RspDeleteItem) Reset() {
	*x = RspDeleteItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faceroam_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspDeleteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspDeleteItem) ProtoMessage() {}

func (x *RspDeleteItem) ProtoReflect() protoreflect.Message {
	mi := &file_faceroam_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspDeleteItem.ProtoReflect.Descriptor instead.
func (*RspDeleteItem) Descriptor() ([]byte, []int) {
	return file_faceroam_proto_rawDescGZIP(), []int{5}
}

func (x *RspDeleteItem) GetFilename() []string {
	if x != nil {
		return x.Filename
	}
	return nil
}

func (x *RspDeleteItem) GetRet() []int64 {
	if x != nil {
		return x.Ret
	}
	return nil
}

type RspUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename    []string `protobuf:"bytes,1,rep,name=filename" json:"filename,omitempty"`
	DeleteFile  []string `protobuf:"bytes,2,rep,name=deleteFile" json:"deleteFile,omitempty"`
	Bid         *string  `protobuf:"bytes,3,opt,name=bid" json:"bid,omitempty"`
	MaxRoamSize *uint32  `protobuf:"varint,4,opt,name=maxRoamSize" json:"maxRoamSize,omitempty"`
	EmojiType   []uint32 `protobuf:"varint,5,rep,name=emojiType" json:"emojiType,omitempty"`
}

func (x *RspUserInfo) Reset() {
	*x = RspUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faceroam_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspUserInfo) ProtoMessage() {}

func (x *RspUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_faceroam_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspUserInfo.ProtoReflect.Descriptor instead.
func (*RspUserInfo) Descriptor() ([]byte, []int) {
	return file_faceroam_proto_rawDescGZIP(), []int{6}
}

func (x *RspUserInfo) GetFilename() []string {
	if x != nil {
		return x.Filename
	}
	return nil
}

func (x *RspUserInfo) GetDeleteFile() []string {
	if x != nil {
		return x.DeleteFile
	}
	return nil
}

func (x *RspUserInfo) GetBid() string {
	if x != nil && x.Bid != nil {
		return *x.Bid
	}
	return ""
}

func (x *RspUserInfo) GetMaxRoamSize() uint32 {
	if x != nil && x.MaxRoamSize != nil {
		return *x.MaxRoamSize
	}
	return 0
}

func (x *RspUserInfo) GetEmojiType() []uint32 {
	if x != nil {
		return x.EmojiType
	}
	return nil
}

var File_faceroam_proto protoreflect.FileDescriptor

var file_faceroam_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x61, 0x63, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x50, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x73, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x73, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x71,
	0x71, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x71, 0x71, 0x76,
	0x65, 0x72, 0x22, 0xc0, 0x01, 0x0a, 0x0f, 0x46, 0x61, 0x63, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x0a, 0x04, 0x63, 0x6f, 0x6d, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x63, 0x6f, 0x6d, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x43, 0x6d,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x12,
	0x2e, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x34, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x71, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x2b, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x0f, 0x46, 0x61, 0x63, 0x65, 0x72, 0x6f, 0x61, 0x6d, 0x52, 0x73,
	0x70, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x73, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x72, 0x73, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52,
	0x73, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x72, 0x73, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0d, 0x72, 0x73, 0x70, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x52, 0x73, 0x70, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0d,
	0x72, 0x73, 0x70, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x3d, 0x0a,
	0x0d, 0x52, 0x73, 0x70, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x72, 0x65, 0x74, 0x22, 0x9b, 0x01, 0x0a,
	0x0b, 0x52, 0x73, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61,
	0x78, 0x52, 0x6f, 0x61, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x6d, 0x61, 0x78, 0x52, 0x6f, 0x61, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x09, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b,
	0x66, 0x61, 0x63, 0x65, 0x72, 0x6f, 0x61, 0x6d,
}

var (
	file_faceroam_proto_rawDescOnce sync.Once
	file_faceroam_proto_rawDescData = file_faceroam_proto_rawDesc
)

func file_faceroam_proto_rawDescGZIP() []byte {
	file_faceroam_proto_rawDescOnce.Do(func() {
		file_faceroam_proto_rawDescData = protoimpl.X.CompressGZIP(file_faceroam_proto_rawDescData)
	})
	return file_faceroam_proto_rawDescData
}

var file_faceroam_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_faceroam_proto_goTypes = []interface{}{
	(*PlatInfo)(nil),        // 0: PlatInfo
	(*FaceroamReqBody)(nil), // 1: FaceroamReqBody
	(*ReqDeleteItem)(nil),   // 2: ReqDeleteItem
	(*ReqUserInfo)(nil),     // 3: ReqUserInfo
	(*FaceroamRspBody)(nil), // 4: FaceroamRspBody
	(*RspDeleteItem)(nil),   // 5: RspDeleteItem
	(*RspUserInfo)(nil),     // 6: RspUserInfo
}
var file_faceroam_proto_depIdxs = []int32{
	0, // 0: FaceroamReqBody.comm:type_name -> PlatInfo
	3, // 1: FaceroamReqBody.reqUserInfo:type_name -> ReqUserInfo
	2, // 2: FaceroamReqBody.reqDeleteItem:type_name -> ReqDeleteItem
	6, // 3: FaceroamRspBody.rspUserInfo:type_name -> RspUserInfo
	5, // 4: FaceroamRspBody.rspDeleteItem:type_name -> RspDeleteItem
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_faceroam_proto_init() }
func file_faceroam_proto_init() {
	if File_faceroam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_faceroam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatInfo); i {
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
		file_faceroam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceroamReqBody); i {
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
		file_faceroam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqDeleteItem); i {
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
		file_faceroam_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUserInfo); i {
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
		file_faceroam_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceroamRspBody); i {
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
		file_faceroam_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspDeleteItem); i {
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
		file_faceroam_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspUserInfo); i {
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
			RawDescriptor: file_faceroam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_faceroam_proto_goTypes,
		DependencyIndexes: file_faceroam_proto_depIdxs,
		MessageInfos:      file_faceroam_proto_msgTypes,
	}.Build()
	File_faceroam_proto = out.File
	file_faceroam_proto_rawDesc = nil
	file_faceroam_proto_goTypes = nil
	file_faceroam_proto_depIdxs = nil
}
