// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: dashboard.proto

package object_builder_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateDashboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Icon string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *CreateDashboardRequest) Reset() {
	*x = CreateDashboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDashboardRequest) ProtoMessage() {}

func (x *CreateDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDashboardRequest.ProtoReflect.Descriptor instead.
func (*CreateDashboardRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDashboardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDashboardRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type Dashboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon      string      `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Panels    []*Panel    `protobuf:"bytes,4,rep,name=panels,proto3" json:"panels,omitempty"`
	Variables []*Variable `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty"`
}

func (x *Dashboard) Reset() {
	*x = Dashboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dashboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard) ProtoMessage() {}

func (x *Dashboard) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard.ProtoReflect.Descriptor instead.
func (*Dashboard) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *Dashboard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Dashboard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dashboard) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Dashboard) GetPanels() []*Panel {
	if x != nil {
		return x.Panels
	}
	return nil
}

func (x *Dashboard) GetVariables() []*Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

type GetAllDashboadsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dashboards []*Dashboard `protobuf:"bytes,1,rep,name=dashboards,proto3" json:"dashboards,omitempty"`
	Count      int32        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllDashboadsResponse) Reset() {
	*x = GetAllDashboadsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDashboadsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDashboadsResponse) ProtoMessage() {}

func (x *GetAllDashboadsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDashboadsResponse.ProtoReflect.Descriptor instead.
func (*GetAllDashboadsResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllDashboadsResponse) GetDashboards() []*Dashboard {
	if x != nil {
		return x.Dashboards
	}
	return nil
}

func (x *GetAllDashboadsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetAllDashboardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAllDashboardsRequest) Reset() {
	*x = GetAllDashboardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDashboardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDashboardsRequest) ProtoMessage() {}

func (x *GetAllDashboardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDashboardsRequest.ProtoReflect.Descriptor instead.
func (*GetAllDashboardsRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllDashboardsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DashboardPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DashboardPrimaryKey) Reset() {
	*x = DashboardPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DashboardPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardPrimaryKey) ProtoMessage() {}

func (x *DashboardPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardPrimaryKey.ProtoReflect.Descriptor instead.
func (*DashboardPrimaryKey) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{4}
}

func (x *DashboardPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_dashboard_proto protoreflect.FileDescriptor

var file_dashboard_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0xba, 0x01, 0x0a, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x06, 0x70,
	0x61, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x52, 0x06, 0x70, 0x61, 0x6e, 0x65,
	0x6c, 0x73, 0x12, 0x3e, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x22, 0x72, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a,
	0x0a, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x0a, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd7, 0x03, 0x0a,
	0x10, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x00,
	0x12, 0x6d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x21, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x00, 0x12, 0x45,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x2b, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dashboard_proto_rawDescOnce sync.Once
	file_dashboard_proto_rawDescData = file_dashboard_proto_rawDesc
)

func file_dashboard_proto_rawDescGZIP() []byte {
	file_dashboard_proto_rawDescOnce.Do(func() {
		file_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_dashboard_proto_rawDescData)
	})
	return file_dashboard_proto_rawDescData
}

var file_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dashboard_proto_goTypes = []interface{}{
	(*CreateDashboardRequest)(nil),  // 0: object_builder_service.CreateDashboardRequest
	(*Dashboard)(nil),               // 1: object_builder_service.Dashboard
	(*GetAllDashboadsResponse)(nil), // 2: object_builder_service.GetAllDashboadsResponse
	(*GetAllDashboardsRequest)(nil), // 3: object_builder_service.GetAllDashboardsRequest
	(*DashboardPrimaryKey)(nil),     // 4: object_builder_service.DashboardPrimaryKey
	(*Panel)(nil),                   // 5: object_builder_service.Panel
	(*Variable)(nil),                // 6: object_builder_service.Variable
	(*emptypb.Empty)(nil),           // 7: google.protobuf.Empty
}
var file_dashboard_proto_depIdxs = []int32{
	5, // 0: object_builder_service.Dashboard.panels:type_name -> object_builder_service.Panel
	6, // 1: object_builder_service.Dashboard.variables:type_name -> object_builder_service.Variable
	1, // 2: object_builder_service.GetAllDashboadsResponse.dashboards:type_name -> object_builder_service.Dashboard
	0, // 3: object_builder_service.DashboardService.Create:input_type -> object_builder_service.CreateDashboardRequest
	3, // 4: object_builder_service.DashboardService.GetList:input_type -> object_builder_service.GetAllDashboardsRequest
	4, // 5: object_builder_service.DashboardService.GetSingle:input_type -> object_builder_service.DashboardPrimaryKey
	1, // 6: object_builder_service.DashboardService.Update:input_type -> object_builder_service.Dashboard
	4, // 7: object_builder_service.DashboardService.Delete:input_type -> object_builder_service.DashboardPrimaryKey
	1, // 8: object_builder_service.DashboardService.Create:output_type -> object_builder_service.Dashboard
	2, // 9: object_builder_service.DashboardService.GetList:output_type -> object_builder_service.GetAllDashboadsResponse
	1, // 10: object_builder_service.DashboardService.GetSingle:output_type -> object_builder_service.Dashboard
	7, // 11: object_builder_service.DashboardService.Update:output_type -> google.protobuf.Empty
	7, // 12: object_builder_service.DashboardService.Delete:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dashboard_proto_init() }
func file_dashboard_proto_init() {
	if File_dashboard_proto != nil {
		return
	}
	file_panel_proto_init()
	file_variable_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDashboardRequest); i {
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
		file_dashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dashboard); i {
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
		file_dashboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDashboadsResponse); i {
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
		file_dashboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDashboardsRequest); i {
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
		file_dashboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DashboardPrimaryKey); i {
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
			RawDescriptor: file_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dashboard_proto_goTypes,
		DependencyIndexes: file_dashboard_proto_depIdxs,
		MessageInfos:      file_dashboard_proto_msgTypes,
	}.Build()
	File_dashboard_proto = out.File
	file_dashboard_proto_rawDesc = nil
	file_dashboard_proto_goTypes = nil
	file_dashboard_proto_depIdxs = nil
}
