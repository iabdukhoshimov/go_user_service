// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: field.proto

package object_builder_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateFieldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Default       string           `protobuf:"bytes,1,opt,name=default,proto3" json:"default,omitempty"`
	Type          string           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Index         string           `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
	Label         string           `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Slug          string           `protobuf:"bytes,5,opt,name=slug,proto3" json:"slug,omitempty"`
	TableId       string           `protobuf:"bytes,6,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Attributes    *structpb.Struct `protobuf:"bytes,7,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Id            string           `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	IsVisible     bool             `protobuf:"varint,9,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`
	AutofillTable string           `protobuf:"bytes,10,opt,name=autofill_table,json=autofillTable,proto3" json:"autofill_table,omitempty"`
	AutofillField string           `protobuf:"bytes,11,opt,name=autofill_field,json=autofillField,proto3" json:"autofill_field,omitempty"`
}

func (x *CreateFieldRequest) Reset() {
	*x = CreateFieldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFieldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFieldRequest) ProtoMessage() {}

func (x *CreateFieldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFieldRequest.ProtoReflect.Descriptor instead.
func (*CreateFieldRequest) Descriptor() ([]byte, []int) {
	return file_field_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFieldRequest) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *CreateFieldRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateFieldRequest) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *CreateFieldRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CreateFieldRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CreateFieldRequest) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *CreateFieldRequest) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *CreateFieldRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateFieldRequest) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

func (x *CreateFieldRequest) GetAutofillTable() string {
	if x != nil {
		return x.AutofillTable
	}
	return ""
}

func (x *CreateFieldRequest) GetAutofillField() string {
	if x != nil {
		return x.AutofillField
	}
	return ""
}

type CreateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Default    string           `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Type       string           `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Index      string           `protobuf:"bytes,4,opt,name=index,proto3" json:"index,omitempty"`
	Label      string           `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Slug       string           `protobuf:"bytes,6,opt,name=slug,proto3" json:"slug,omitempty"`
	Required   bool             `protobuf:"varint,7,opt,name=required,proto3" json:"required,omitempty"`
	Attributes *structpb.Struct `protobuf:"bytes,8,opt,name=attributes,proto3" json:"attributes,omitempty"`
	IsVisible  bool             `protobuf:"varint,9,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`
}

func (x *CreateFieldsRequest) Reset() {
	*x = CreateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_field_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFieldsRequest) ProtoMessage() {}

func (x *CreateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_field_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFieldsRequest.ProtoReflect.Descriptor instead.
func (*CreateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_field_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFieldsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateFieldsRequest) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *CreateFieldsRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateFieldsRequest) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *CreateFieldsRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CreateFieldsRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *CreateFieldsRequest) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *CreateFieldsRequest) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *CreateFieldsRequest) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Default       string           `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Type          string           `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Index         string           `protobuf:"bytes,4,opt,name=index,proto3" json:"index,omitempty"`
	Label         string           `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Slug          string           `protobuf:"bytes,6,opt,name=slug,proto3" json:"slug,omitempty"`
	TableId       string           `protobuf:"bytes,7,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Required      bool             `protobuf:"varint,8,opt,name=required,proto3" json:"required,omitempty"`
	Attributes    *structpb.Struct `protobuf:"bytes,9,opt,name=attributes,proto3" json:"attributes,omitempty"`
	IsVisible     bool             `protobuf:"varint,10,opt,name=is_visible,json=isVisible,proto3" json:"is_visible,omitempty"`
	AutofillField string           `protobuf:"bytes,11,opt,name=autofill_field,json=autofillField,proto3" json:"autofill_field,omitempty"`
	AutofillTable string           `protobuf:"bytes,12,opt,name=autofill_table,json=autofillTable,proto3" json:"autofill_table,omitempty"`
	RelationId    string           `protobuf:"bytes,13,opt,name=relation_id,json=relationId,proto3" json:"relation_id,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_field_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_field_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_field_proto_rawDescGZIP(), []int{2}
}

func (x *Field) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Field) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *Field) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Field) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *Field) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Field) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Field) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *Field) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *Field) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Field) GetIsVisible() bool {
	if x != nil {
		return x.IsVisible
	}
	return false
}

func (x *Field) GetAutofillField() string {
	if x != nil {
		return x.AutofillField
	}
	return ""
}

func (x *Field) GetAutofillTable() string {
	if x != nil {
		return x.AutofillTable
	}
	return ""
}

func (x *Field) GetRelationId() string {
	if x != nil {
		return x.RelationId
	}
	return ""
}

type GetAllFieldsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*Field         `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	Count  int32            `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Data   *structpb.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllFieldsResponse) Reset() {
	*x = GetAllFieldsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_field_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFieldsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFieldsResponse) ProtoMessage() {}

func (x *GetAllFieldsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_field_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFieldsResponse.ProtoReflect.Descriptor instead.
func (*GetAllFieldsResponse) Descriptor() ([]byte, []int) {
	return file_field_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllFieldsResponse) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *GetAllFieldsResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetAllFieldsResponse) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableId          string `protobuf:"bytes,1,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	Limit            int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset           int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Search           string `protobuf:"bytes,4,opt,name=search,proto3" json:"search,omitempty"`
	TableSlug        string `protobuf:"bytes,5,opt,name=table_slug,json=tableSlug,proto3" json:"table_slug,omitempty"`
	WithManyRelation bool   `protobuf:"varint,6,opt,name=with_many_relation,json=withManyRelation,proto3" json:"with_many_relation,omitempty"`
	WithOneRelation  bool   `protobuf:"varint,7,opt,name=with_one_relation,json=withOneRelation,proto3" json:"with_one_relation,omitempty"`
}

func (x *GetAllFieldsRequest) Reset() {
	*x = GetAllFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_field_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFieldsRequest) ProtoMessage() {}

func (x *GetAllFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_field_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFieldsRequest.ProtoReflect.Descriptor instead.
func (*GetAllFieldsRequest) Descriptor() ([]byte, []int) {
	return file_field_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllFieldsRequest) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

func (x *GetAllFieldsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllFieldsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllFieldsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllFieldsRequest) GetTableSlug() string {
	if x != nil {
		return x.TableSlug
	}
	return ""
}

func (x *GetAllFieldsRequest) GetWithManyRelation() bool {
	if x != nil {
		return x.WithManyRelation
	}
	return false
}

func (x *GetAllFieldsRequest) GetWithOneRelation() bool {
	if x != nil {
		return x.WithOneRelation
	}
	return false
}

type FieldPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FieldPrimaryKey) Reset() {
	*x = FieldPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_field_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldPrimaryKey) ProtoMessage() {}

func (x *FieldPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_field_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldPrimaryKey.ProtoReflect.Descriptor instead.
func (*FieldPrimaryKey) Descriptor() ([]byte, []int) {
	return file_field_proto_rawDescGZIP(), []int{5}
}

func (x *FieldPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_field_proto protoreflect.FileDescriptor

var file_field_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd3, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f,
	0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c,
	0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x87, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x22, 0x83, 0x03, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x66,
	0x69, 0x6c, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xef, 0x01, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x75, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x75,
	0x67, 0x12, 0x2c, 0x0a, 0x12, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x77,
	0x69, 0x74, 0x68, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x11, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0f, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xdc,
	0x02, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x55, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x2b, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a,
	0x1f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_field_proto_rawDescOnce sync.Once
	file_field_proto_rawDescData = file_field_proto_rawDesc
)

func file_field_proto_rawDescGZIP() []byte {
	file_field_proto_rawDescOnce.Do(func() {
		file_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_field_proto_rawDescData)
	})
	return file_field_proto_rawDescData
}

var file_field_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_field_proto_goTypes = []interface{}{
	(*CreateFieldRequest)(nil),   // 0: object_builder_service.CreateFieldRequest
	(*CreateFieldsRequest)(nil),  // 1: object_builder_service.CreateFieldsRequest
	(*Field)(nil),                // 2: object_builder_service.Field
	(*GetAllFieldsResponse)(nil), // 3: object_builder_service.GetAllFieldsResponse
	(*GetAllFieldsRequest)(nil),  // 4: object_builder_service.GetAllFieldsRequest
	(*FieldPrimaryKey)(nil),      // 5: object_builder_service.FieldPrimaryKey
	(*structpb.Struct)(nil),      // 6: google.protobuf.Struct
	(*emptypb.Empty)(nil),        // 7: google.protobuf.Empty
}
var file_field_proto_depIdxs = []int32{
	6, // 0: object_builder_service.CreateFieldRequest.attributes:type_name -> google.protobuf.Struct
	6, // 1: object_builder_service.CreateFieldsRequest.attributes:type_name -> google.protobuf.Struct
	6, // 2: object_builder_service.Field.attributes:type_name -> google.protobuf.Struct
	2, // 3: object_builder_service.GetAllFieldsResponse.fields:type_name -> object_builder_service.Field
	6, // 4: object_builder_service.GetAllFieldsResponse.data:type_name -> google.protobuf.Struct
	0, // 5: object_builder_service.FieldService.Create:input_type -> object_builder_service.CreateFieldRequest
	4, // 6: object_builder_service.FieldService.GetAll:input_type -> object_builder_service.GetAllFieldsRequest
	2, // 7: object_builder_service.FieldService.Update:input_type -> object_builder_service.Field
	5, // 8: object_builder_service.FieldService.Delete:input_type -> object_builder_service.FieldPrimaryKey
	2, // 9: object_builder_service.FieldService.Create:output_type -> object_builder_service.Field
	3, // 10: object_builder_service.FieldService.GetAll:output_type -> object_builder_service.GetAllFieldsResponse
	7, // 11: object_builder_service.FieldService.Update:output_type -> google.protobuf.Empty
	7, // 12: object_builder_service.FieldService.Delete:output_type -> google.protobuf.Empty
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_field_proto_init() }
func file_field_proto_init() {
	if File_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFieldRequest); i {
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
		file_field_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFieldsRequest); i {
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
		file_field_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_field_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllFieldsResponse); i {
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
		file_field_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllFieldsRequest); i {
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
		file_field_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldPrimaryKey); i {
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
			RawDescriptor: file_field_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_field_proto_goTypes,
		DependencyIndexes: file_field_proto_depIdxs,
		MessageInfos:      file_field_proto_msgTypes,
	}.Build()
	File_field_proto = out.File
	file_field_proto_rawDesc = nil
	file_field_proto_goTypes = nil
	file_field_proto_depIdxs = nil
}
