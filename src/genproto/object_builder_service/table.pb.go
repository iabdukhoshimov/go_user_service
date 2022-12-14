// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: table.proto

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

type CreateTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label             string                 `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Description       string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Slug              string                 `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Fields            []*CreateFieldsRequest `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	ShowInMenu        bool                   `protobuf:"varint,5,opt,name=show_in_menu,json=showInMenu,proto3" json:"show_in_menu,omitempty"`
	Icon              string                 `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	SubtitleFieldSlug string                 `protobuf:"bytes,7,opt,name=subtitle_field_slug,json=subtitleFieldSlug,proto3" json:"subtitle_field_slug,omitempty"`
	Sections          []*Section             `protobuf:"bytes,8,rep,name=sections,proto3" json:"sections,omitempty"`
	AppId             string                 `protobuf:"bytes,9,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	IncrementId       *IncrementID           `protobuf:"bytes,10,opt,name=increment_id,json=incrementId,proto3" json:"increment_id,omitempty"`
}

func (x *CreateTableRequest) Reset() {
	*x = CreateTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTableRequest) ProtoMessage() {}

func (x *CreateTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTableRequest.ProtoReflect.Descriptor instead.
func (*CreateTableRequest) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTableRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CreateTableRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTableRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CreateTableRequest) GetFields() []*CreateFieldsRequest {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *CreateTableRequest) GetShowInMenu() bool {
	if x != nil {
		return x.ShowInMenu
	}
	return false
}

func (x *CreateTableRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *CreateTableRequest) GetSubtitleFieldSlug() string {
	if x != nil {
		return x.SubtitleFieldSlug
	}
	return ""
}

func (x *CreateTableRequest) GetSections() []*Section {
	if x != nil {
		return x.Sections
	}
	return nil
}

func (x *CreateTableRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *CreateTableRequest) GetIncrementId() *IncrementID {
	if x != nil {
		return x.IncrementId
	}
	return nil
}

type IncrementID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithIncrementId bool   `protobuf:"varint,1,opt,name=with_increment_id,json=withIncrementId,proto3" json:"with_increment_id,omitempty"`
	DigitNumber     int32  `protobuf:"varint,2,opt,name=digit_number,json=digitNumber,proto3" json:"digit_number,omitempty"`
	Prefix          string `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *IncrementID) Reset() {
	*x = IncrementID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementID) ProtoMessage() {}

func (x *IncrementID) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementID.ProtoReflect.Descriptor instead.
func (*IncrementID) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{1}
}

func (x *IncrementID) GetWithIncrementId() bool {
	if x != nil {
		return x.WithIncrementId
	}
	return false
}

func (x *IncrementID) GetDigitNumber() int32 {
	if x != nil {
		return x.DigitNumber
	}
	return 0
}

func (x *IncrementID) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type CreateTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label             string       `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description       string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Slug              string       `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	Fields            []*Field     `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	ShowInMenu        bool         `protobuf:"varint,6,opt,name=show_in_menu,json=showInMenu,proto3" json:"show_in_menu,omitempty"`
	Icon              string       `protobuf:"bytes,7,opt,name=icon,proto3" json:"icon,omitempty"`
	SubtitleFieldSlug string       `protobuf:"bytes,8,opt,name=subtitle_field_slug,json=subtitleFieldSlug,proto3" json:"subtitle_field_slug,omitempty"`
	Sections          []*Section   `protobuf:"bytes,9,rep,name=sections,proto3" json:"sections,omitempty"`
	IncrementId       *IncrementID `protobuf:"bytes,10,opt,name=increment_id,json=incrementId,proto3" json:"increment_id,omitempty"`
}

func (x *CreateTableResponse) Reset() {
	*x = CreateTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTableResponse) ProtoMessage() {}

func (x *CreateTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTableResponse.ProtoReflect.Descriptor instead.
func (*CreateTableResponse) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTableResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateTableResponse) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CreateTableResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTableResponse) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CreateTableResponse) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *CreateTableResponse) GetShowInMenu() bool {
	if x != nil {
		return x.ShowInMenu
	}
	return false
}

func (x *CreateTableResponse) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *CreateTableResponse) GetSubtitleFieldSlug() string {
	if x != nil {
		return x.SubtitleFieldSlug
	}
	return ""
}

func (x *CreateTableResponse) GetSections() []*Section {
	if x != nil {
		return x.Sections
	}
	return nil
}

func (x *CreateTableResponse) GetIncrementId() *IncrementID {
	if x != nil {
		return x.IncrementId
	}
	return nil
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label             string       `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description       string       `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Slug              string       `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	ShowInMenu        bool         `protobuf:"varint,5,opt,name=show_in_menu,json=showInMenu,proto3" json:"show_in_menu,omitempty"`
	Icon              string       `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	SubtitleFieldSlug string       `protobuf:"bytes,7,opt,name=subtitle_field_slug,json=subtitleFieldSlug,proto3" json:"subtitle_field_slug,omitempty"`
	IsVisible         bool         `protobuf:"varint,8,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`
	IsOwnTable        bool         `protobuf:"varint,9,opt,name=is_own_table,json=isOwnTable,proto3" json:"is_own_table,omitempty"`
	IncrementId       *IncrementID `protobuf:"bytes,10,opt,name=increment_id,json=incrementId,proto3" json:"increment_id,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{3}
}

func (x *Table) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Table) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Table) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Table) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Table) GetShowInMenu() bool {
	if x != nil {
		return x.ShowInMenu
	}
	return false
}

func (x *Table) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Table) GetSubtitleFieldSlug() string {
	if x != nil {
		return x.SubtitleFieldSlug
	}
	return ""
}

func (x *Table) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

func (x *Table) GetIsOwnTable() bool {
	if x != nil {
		return x.IsOwnTable
	}
	return false
}

func (x *Table) GetIncrementId() *IncrementID {
	if x != nil {
		return x.IncrementId
	}
	return nil
}

type GetAllTablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables []*Table `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
	Count  int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllTablesResponse) Reset() {
	*x = GetAllTablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTablesResponse) ProtoMessage() {}

func (x *GetAllTablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTablesResponse.ProtoReflect.Descriptor instead.
func (*GetAllTablesResponse) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTablesResponse) GetTables() []*Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

func (x *GetAllTablesResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TablePrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TablePrimaryKey) Reset() {
	*x = TablePrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TablePrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TablePrimaryKey) ProtoMessage() {}

func (x *TablePrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TablePrimaryKey.ProtoReflect.Descriptor instead.
func (*TablePrimaryKey) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{5}
}

func (x *TablePrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllTablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllTablesRequest) Reset() {
	*x = GetAllTablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTablesRequest) ProtoMessage() {}

func (x *GetAllTablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTablesRequest.ProtoReflect.Descriptor instead.
func (*GetAllTablesRequest) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllTablesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllTablesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllTablesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

var File_table_proto protoreflect.FileDescriptor

var file_table_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7,
	0x03, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x12, 0x43, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x69,
	0x6e, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x68,
	0x6f, 0x77, 0x49, 0x6e, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x3b, 0x0a, 0x08,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x46, 0x0a, 0x0c, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0b, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x69, 0x74, 0x68, 0x5f,
	0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x67, 0x69, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x69, 0x67, 0x69, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x93,
	0x03, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x12, 0x35, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x68, 0x6f,
	0x77, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x63, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x13, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x6c, 0x75, 0x67, 0x12,
	0x3b, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x46, 0x0a, 0x0c,
	0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0xd2, 0x02, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x68,
	0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x2e, 0x0a, 0x13, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73,
	0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x6c, 0x75, 0x67,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x46, 0x0a, 0x0c, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0b, 0x69, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x21,
	0x0a, 0x0f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x32, 0xbf,
	0x03, 0x0a, 0x0c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x63, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x27, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1d, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x2b, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x21, 0x5a, 0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_table_proto_rawDescOnce sync.Once
	file_table_proto_rawDescData = file_table_proto_rawDesc
)

func file_table_proto_rawDescGZIP() []byte {
	file_table_proto_rawDescOnce.Do(func() {
		file_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_table_proto_rawDescData)
	})
	return file_table_proto_rawDescData
}

var file_table_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_table_proto_goTypes = []interface{}{
	(*CreateTableRequest)(nil),   // 0: object_builder_service.CreateTableRequest
	(*IncrementID)(nil),          // 1: object_builder_service.IncrementID
	(*CreateTableResponse)(nil),  // 2: object_builder_service.CreateTableResponse
	(*Table)(nil),                // 3: object_builder_service.Table
	(*GetAllTablesResponse)(nil), // 4: object_builder_service.GetAllTablesResponse
	(*TablePrimaryKey)(nil),      // 5: object_builder_service.TablePrimaryKey
	(*GetAllTablesRequest)(nil),  // 6: object_builder_service.GetAllTablesRequest
	(*CreateFieldsRequest)(nil),  // 7: object_builder_service.CreateFieldsRequest
	(*Section)(nil),              // 8: object_builder_service.Section
	(*Field)(nil),                // 9: object_builder_service.Field
	(*emptypb.Empty)(nil),        // 10: google.protobuf.Empty
}
var file_table_proto_depIdxs = []int32{
	7,  // 0: object_builder_service.CreateTableRequest.fields:type_name -> object_builder_service.CreateFieldsRequest
	8,  // 1: object_builder_service.CreateTableRequest.sections:type_name -> object_builder_service.Section
	1,  // 2: object_builder_service.CreateTableRequest.increment_id:type_name -> object_builder_service.IncrementID
	9,  // 3: object_builder_service.CreateTableResponse.fields:type_name -> object_builder_service.Field
	8,  // 4: object_builder_service.CreateTableResponse.sections:type_name -> object_builder_service.Section
	1,  // 5: object_builder_service.CreateTableResponse.increment_id:type_name -> object_builder_service.IncrementID
	1,  // 6: object_builder_service.Table.increment_id:type_name -> object_builder_service.IncrementID
	3,  // 7: object_builder_service.GetAllTablesResponse.tables:type_name -> object_builder_service.Table
	0,  // 8: object_builder_service.TableService.Create:input_type -> object_builder_service.CreateTableRequest
	5,  // 9: object_builder_service.TableService.GetByID:input_type -> object_builder_service.TablePrimaryKey
	6,  // 10: object_builder_service.TableService.GetAll:input_type -> object_builder_service.GetAllTablesRequest
	3,  // 11: object_builder_service.TableService.Update:input_type -> object_builder_service.Table
	5,  // 12: object_builder_service.TableService.Delete:input_type -> object_builder_service.TablePrimaryKey
	2,  // 13: object_builder_service.TableService.Create:output_type -> object_builder_service.CreateTableResponse
	3,  // 14: object_builder_service.TableService.GetByID:output_type -> object_builder_service.Table
	4,  // 15: object_builder_service.TableService.GetAll:output_type -> object_builder_service.GetAllTablesResponse
	10, // 16: object_builder_service.TableService.Update:output_type -> google.protobuf.Empty
	10, // 17: object_builder_service.TableService.Delete:output_type -> google.protobuf.Empty
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_table_proto_init() }
func file_table_proto_init() {
	if File_table_proto != nil {
		return
	}
	file_field_proto_init()
	file_section_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTableRequest); i {
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
		file_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementID); i {
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
		file_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTableResponse); i {
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
		file_table_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_table_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTablesResponse); i {
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
		file_table_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TablePrimaryKey); i {
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
		file_table_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTablesRequest); i {
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
			RawDescriptor: file_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_table_proto_goTypes,
		DependencyIndexes: file_table_proto_depIdxs,
		MessageInfos:      file_table_proto_msgTypes,
	}.Build()
	File_table_proto = out.File
	file_table_proto_rawDesc = nil
	file_table_proto_goTypes = nil
	file_table_proto_depIdxs = nil
}
