// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: portfolios.proto

package __

import (
	context "context"
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

//-------------------------------获取照片类别 start----------------
type ImgCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ImgCategoryRequest) Reset() {
	*x = ImgCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portfolios_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgCategoryRequest) ProtoMessage() {}

func (x *ImgCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portfolios_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgCategoryRequest.ProtoReflect.Descriptor instead.
func (*ImgCategoryRequest) Descriptor() ([]byte, []int) {
	return file_portfolios_proto_rawDescGZIP(), []int{0}
}

func (x *ImgCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImgCategoryRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ImgCategoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*OutImgCategory `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ImgCategoryReply) Reset() {
	*x = ImgCategoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portfolios_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgCategoryReply) ProtoMessage() {}

func (x *ImgCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_portfolios_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgCategoryReply.ProtoReflect.Descriptor instead.
func (*ImgCategoryReply) Descriptor() ([]byte, []int) {
	return file_portfolios_proto_rawDescGZIP(), []int{1}
}

func (x *ImgCategoryReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImgCategoryReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ImgCategoryReply) GetData() []*OutImgCategory {
	if x != nil {
		return x.Data
	}
	return nil
}

type OutImgCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VCategoryName string `protobuf:"bytes,2,opt,name=v_category_name,json=vCategoryName,proto3" json:"v_category_name,omitempty"`
	VPhotoUrl     string `protobuf:"bytes,3,opt,name=v_photo_url,json=vPhotoUrl,proto3" json:"v_photo_url,omitempty"`
}

func (x *OutImgCategory) Reset() {
	*x = OutImgCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portfolios_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutImgCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutImgCategory) ProtoMessage() {}

func (x *OutImgCategory) ProtoReflect() protoreflect.Message {
	mi := &file_portfolios_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutImgCategory.ProtoReflect.Descriptor instead.
func (*OutImgCategory) Descriptor() ([]byte, []int) {
	return file_portfolios_proto_rawDescGZIP(), []int{2}
}

func (x *OutImgCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OutImgCategory) GetVCategoryName() string {
	if x != nil {
		return x.VCategoryName
	}
	return ""
}

func (x *OutImgCategory) GetVPhotoUrl() string {
	if x != nil {
		return x.VPhotoUrl
	}
	return ""
}

//-------------------------------根据类别Id获取图片 start----------------
type ImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message       string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ImgCategoryId string `protobuf:"bytes,3,opt,name=img_category_id,json=imgCategoryId,proto3" json:"img_category_id,omitempty"`
}

func (x *ImgRequest) Reset() {
	*x = ImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portfolios_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgRequest) ProtoMessage() {}

func (x *ImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portfolios_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgRequest.ProtoReflect.Descriptor instead.
func (*ImgRequest) Descriptor() ([]byte, []int) {
	return file_portfolios_proto_rawDescGZIP(), []int{3}
}

func (x *ImgRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImgRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ImgRequest) GetImgCategoryId() string {
	if x != nil {
		return x.ImgCategoryId
	}
	return ""
}

type ImgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*OutImg `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ImgReply) Reset() {
	*x = ImgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portfolios_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImgReply) ProtoMessage() {}

func (x *ImgReply) ProtoReflect() protoreflect.Message {
	mi := &file_portfolios_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImgReply.ProtoReflect.Descriptor instead.
func (*ImgReply) Descriptor() ([]byte, []int) {
	return file_portfolios_proto_rawDescGZIP(), []int{4}
}

func (x *ImgReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImgReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ImgReply) GetData() []*OutImg {
	if x != nil {
		return x.Data
	}
	return nil
}

type OutImg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PhotoCategoryId string `protobuf:"bytes,2,opt,name=photo_category_id,json=photoCategoryId,proto3" json:"photo_category_id,omitempty"`
	VPhotoAlt       string `protobuf:"bytes,3,opt,name=v_photo_alt,json=vPhotoAlt,proto3" json:"v_photo_alt,omitempty"`
	VPhotoUrl       string `protobuf:"bytes,4,opt,name=v_photo_url,json=vPhotoUrl,proto3" json:"v_photo_url,omitempty"`
	UploadTime      string `protobuf:"bytes,5,opt,name=upload_time,json=uploadTime,proto3" json:"upload_time,omitempty"`
	IsFirst         string `protobuf:"bytes,6,opt,name=is_first,json=isFirst,proto3" json:"is_first,omitempty"`
	DeleteTime      string `protobuf:"bytes,7,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
}

func (x *OutImg) Reset() {
	*x = OutImg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portfolios_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutImg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutImg) ProtoMessage() {}

func (x *OutImg) ProtoReflect() protoreflect.Message {
	mi := &file_portfolios_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutImg.ProtoReflect.Descriptor instead.
func (*OutImg) Descriptor() ([]byte, []int) {
	return file_portfolios_proto_rawDescGZIP(), []int{5}
}

func (x *OutImg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OutImg) GetPhotoCategoryId() string {
	if x != nil {
		return x.PhotoCategoryId
	}
	return ""
}

func (x *OutImg) GetVPhotoAlt() string {
	if x != nil {
		return x.VPhotoAlt
	}
	return ""
}

func (x *OutImg) GetVPhotoUrl() string {
	if x != nil {
		return x.VPhotoUrl
	}
	return ""
}

func (x *OutImg) GetUploadTime() string {
	if x != nil {
		return x.UploadTime
	}
	return ""
}

func (x *OutImg) GetIsFirst() string {
	if x != nil {
		return x.IsFirst
	}
	return ""
}

func (x *OutImg) GetDeleteTime() string {
	if x != nil {
		return x.DeleteTime
	}
	return ""
}

var File_portfolios_proto protoreflect.FileDescriptor

var file_portfolios_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x22, 0x42,
	0x0a, 0x12, 0x49, 0x6d, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x70, 0x0a, 0x10, 0x49, 0x6d, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x2e,
	0x4f, 0x75, 0x74, 0x49, 0x6d, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x0e, 0x4f, 0x75, 0x74, 0x49, 0x6d, 0x67, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x76, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x76, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0b, 0x76, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x22, 0x62,
	0x0a, 0x0a, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x6d,
	0x67, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6d, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x22, 0x60, 0x0a, 0x08, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x49, 0x6d, 0x67, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xe1, 0x01, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x49, 0x6d, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2a, 0x0a, 0x11, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x76,
	0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x41, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x76,
	0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xa2, 0x01, 0x0a, 0x0a, 0x50, 0x6f, 0x72,
	0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x12, 0x50, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6d,
	0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x2e, 0x49, 0x6d, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x2e, 0x49, 0x6d, 0x67, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x49, 0x6d, 0x67, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x73, 0x2e, 0x49, 0x6d, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69,
	0x6f, 0x73, 0x2e, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x04, 0x5a,
	0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portfolios_proto_rawDescOnce sync.Once
	file_portfolios_proto_rawDescData = file_portfolios_proto_rawDesc
)

func file_portfolios_proto_rawDescGZIP() []byte {
	file_portfolios_proto_rawDescOnce.Do(func() {
		file_portfolios_proto_rawDescData = protoimpl.X.CompressGZIP(file_portfolios_proto_rawDescData)
	})
	return file_portfolios_proto_rawDescData
}

var file_portfolios_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_portfolios_proto_goTypes = []interface{}{
	(*ImgCategoryRequest)(nil), // 0: portfolios.ImgCategoryRequest
	(*ImgCategoryReply)(nil),   // 1: portfolios.ImgCategoryReply
	(*OutImgCategory)(nil),     // 2: portfolios.OutImgCategory
	(*ImgRequest)(nil),         // 3: portfolios.ImgRequest
	(*ImgReply)(nil),           // 4: portfolios.ImgReply
	(*OutImg)(nil),             // 5: portfolios.OutImg
}
var file_portfolios_proto_depIdxs = []int32{
	2, // 0: portfolios.ImgCategoryReply.data:type_name -> portfolios.OutImgCategory
	5, // 1: portfolios.ImgReply.data:type_name -> portfolios.OutImg
	0, // 2: portfolios.Portfolios.GetImgCategory:input_type -> portfolios.ImgCategoryRequest
	3, // 3: portfolios.Portfolios.GetImgByCategory:input_type -> portfolios.ImgRequest
	1, // 4: portfolios.Portfolios.GetImgCategory:output_type -> portfolios.ImgCategoryReply
	4, // 5: portfolios.Portfolios.GetImgByCategory:output_type -> portfolios.ImgReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_portfolios_proto_init() }
func file_portfolios_proto_init() {
	if File_portfolios_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_portfolios_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgCategoryRequest); i {
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
		file_portfolios_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgCategoryReply); i {
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
		file_portfolios_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutImgCategory); i {
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
		file_portfolios_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgRequest); i {
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
		file_portfolios_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImgReply); i {
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
		file_portfolios_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutImg); i {
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
			RawDescriptor: file_portfolios_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portfolios_proto_goTypes,
		DependencyIndexes: file_portfolios_proto_depIdxs,
		MessageInfos:      file_portfolios_proto_msgTypes,
	}.Build()
	File_portfolios_proto = out.File
	file_portfolios_proto_rawDesc = nil
	file_portfolios_proto_goTypes = nil
	file_portfolios_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PortfoliosClient is the client API for Portfolios service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortfoliosClient interface {
	GetImgCategory(ctx context.Context, in *ImgCategoryRequest, opts ...grpc.CallOption) (*ImgCategoryReply, error)
	GetImgByCategory(ctx context.Context, in *ImgRequest, opts ...grpc.CallOption) (*ImgReply, error)
}

type portfoliosClient struct {
	cc grpc.ClientConnInterface
}

func NewPortfoliosClient(cc grpc.ClientConnInterface) PortfoliosClient {
	return &portfoliosClient{cc}
}

func (c *portfoliosClient) GetImgCategory(ctx context.Context, in *ImgCategoryRequest, opts ...grpc.CallOption) (*ImgCategoryReply, error) {
	out := new(ImgCategoryReply)
	err := c.cc.Invoke(ctx, "/portfolios.Portfolios/GetImgCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portfoliosClient) GetImgByCategory(ctx context.Context, in *ImgRequest, opts ...grpc.CallOption) (*ImgReply, error) {
	out := new(ImgReply)
	err := c.cc.Invoke(ctx, "/portfolios.Portfolios/GetImgByCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortfoliosServer is the server API for Portfolios service.
type PortfoliosServer interface {
	GetImgCategory(context.Context, *ImgCategoryRequest) (*ImgCategoryReply, error)
	GetImgByCategory(context.Context, *ImgRequest) (*ImgReply, error)
}

// UnimplementedPortfoliosServer can be embedded to have forward compatible implementations.
type UnimplementedPortfoliosServer struct {
}

func (*UnimplementedPortfoliosServer) GetImgCategory(context.Context, *ImgCategoryRequest) (*ImgCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImgCategory not implemented")
}
func (*UnimplementedPortfoliosServer) GetImgByCategory(context.Context, *ImgRequest) (*ImgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImgByCategory not implemented")
}

func RegisterPortfoliosServer(s *grpc.Server, srv PortfoliosServer) {
	s.RegisterService(&_Portfolios_serviceDesc, srv)
}

func _Portfolios_GetImgCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImgCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfoliosServer).GetImgCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portfolios.Portfolios/GetImgCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfoliosServer).GetImgCategory(ctx, req.(*ImgCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portfolios_GetImgByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfoliosServer).GetImgByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portfolios.Portfolios/GetImgByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfoliosServer).GetImgByCategory(ctx, req.(*ImgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Portfolios_serviceDesc = grpc.ServiceDesc{
	ServiceName: "portfolios.Portfolios",
	HandlerType: (*PortfoliosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImgCategory",
			Handler:    _Portfolios_GetImgCategory_Handler,
		},
		{
			MethodName: "GetImgByCategory",
			Handler:    _Portfolios_GetImgByCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portfolios.proto",
}