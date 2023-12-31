// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api/staff/staff_records.proto

package staff_records

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

// Enum value maps for HealthCheckResponse_ServingStatus.
var (
	HealthCheckResponse_ServingStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
	}
	HealthCheckResponse_ServingStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"SERVING":     1,
		"NOT_SERVING": 2,
	}
)

func (x HealthCheckResponse_ServingStatus) Enum() *HealthCheckResponse_ServingStatus {
	p := new(HealthCheckResponse_ServingStatus)
	*p = x
	return p
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckResponse_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_staff_staff_records_proto_enumTypes[0].Descriptor()
}

func (HealthCheckResponse_ServingStatus) Type() protoreflect.EnumType {
	return &file_api_staff_staff_records_proto_enumTypes[0]
}

func (x HealthCheckResponse_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckResponse_ServingStatus.Descriptor instead.
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{12, 0}
}

type StaffRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID     string `protobuf:"bytes,1,opt,name=staffID,proto3" json:"staffID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	JobTitle    string `protobuf:"bytes,3,opt,name=jobTitle,proto3" json:"jobTitle,omitempty"`
	Department  string `protobuf:"bytes,4,opt,name=department,proto3" json:"department,omitempty"`
	IsAvailable bool   `protobuf:"varint,5,opt,name=isAvailable,proto3" json:"isAvailable,omitempty"`
}

func (x *StaffRecord) Reset() {
	*x = StaffRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaffRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffRecord) ProtoMessage() {}

func (x *StaffRecord) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffRecord.ProtoReflect.Descriptor instead.
func (*StaffRecord) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{0}
}

func (x *StaffRecord) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *StaffRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StaffRecord) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *StaffRecord) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *StaffRecord) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

type CreateStaff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Staff *StaffRecord `protobuf:"bytes,1,opt,name=staff,proto3" json:"staff,omitempty"`
}

func (x *CreateStaff) Reset() {
	*x = CreateStaff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStaff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStaff) ProtoMessage() {}

func (x *CreateStaff) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStaff.ProtoReflect.Descriptor instead.
func (*CreateStaff) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStaff) GetStaff() *StaffRecord {
	if x != nil {
		return x.Staff
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID string `protobuf:"bytes,1,opt,name=staffID,proto3" json:"staffID,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *CreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetStaffRecords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStaffRecords) Reset() {
	*x = GetStaffRecords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaffRecords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaffRecords) ProtoMessage() {}

func (x *GetStaffRecords) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaffRecords.ProtoReflect.Descriptor instead.
func (*GetStaffRecords) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{3}
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffRecords []*StaffRecord `protobuf:"bytes,1,rep,name=staffRecords,proto3" json:"staffRecords,omitempty"`
	Error        string         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetStaffRecords() []*StaffRecord {
	if x != nil {
		return x.StaffRecords
	}
	return nil
}

func (x *GetResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetStaffAvailability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID string `protobuf:"bytes,1,opt,name=staffID,proto3" json:"staffID,omitempty"`
}

func (x *GetStaffAvailability) Reset() {
	*x = GetStaffAvailability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStaffAvailability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaffAvailability) ProtoMessage() {}

func (x *GetStaffAvailability) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaffAvailability.ProtoReflect.Descriptor instead.
func (*GetStaffAvailability) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{5}
}

func (x *GetStaffAvailability) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

type GetAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID     string `protobuf:"bytes,1,opt,name=staffID,proto3" json:"staffID,omitempty"`
	IsAvailable bool   `protobuf:"varint,2,opt,name=isAvailable,proto3" json:"isAvailable,omitempty"`
	Error       string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetAvailabilityResponse) Reset() {
	*x = GetAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailabilityResponse) ProtoMessage() {}

func (x *GetAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*GetAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{6}
}

func (x *GetAvailabilityResponse) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *GetAvailabilityResponse) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

func (x *GetAvailabilityResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateStaff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffRecord *StaffRecord `protobuf:"bytes,1,opt,name=staffRecord,proto3" json:"staffRecord,omitempty"`
}

func (x *UpdateStaff) Reset() {
	*x = UpdateStaff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStaff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStaff) ProtoMessage() {}

func (x *UpdateStaff) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStaff.ProtoReflect.Descriptor instead.
func (*UpdateStaff) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateStaff) GetStaffRecord() *StaffRecord {
	if x != nil {
		return x.StaffRecord
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID string `protobuf:"bytes,1,opt,name=staffID,proto3" json:"staffID,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateResponse) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *UpdateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteStaff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID string `protobuf:"bytes,1,opt,name=staffID,proto3" json:"staffID,omitempty"`
}

func (x *DeleteStaff) Reset() {
	*x = DeleteStaff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStaff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStaff) ProtoMessage() {}

func (x *DeleteStaff) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStaff.ProtoReflect.Descriptor instead.
func (*DeleteStaff) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteStaff) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaffID string `protobuf:"bytes,1,opt,name=staffID,proto3" json:"staffID,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteResponse) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *DeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{11}
}

func (x *HealthCheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=StaffRecords.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_staff_staff_records_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_staff_staff_records_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_api_staff_staff_records_proto_rawDescGZIP(), []int{12}
}

func (x *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if x != nil {
		return x.Status
	}
	return HealthCheckResponse_UNKNOWN
}

var File_api_staff_staff_records_proto protoreflect.FileDescriptor

var file_api_staff_staff_records_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x99, 0x01,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6a,
	0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a,
	0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x3e, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x22, 0x5a, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x62, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x30, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x22, 0x6b,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4a, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x5a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x27, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2e, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x13, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2f, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x0a, 0x0d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x32, 0xc4, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x1c, 0x2e, 0x53, 0x74,
	0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1d, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a,
	0x19, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x2e,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x1a, 0x25, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x1c, 0x2e,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66,
	0x1a, 0x1c, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13,
	0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2d, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_staff_staff_records_proto_rawDescOnce sync.Once
	file_api_staff_staff_records_proto_rawDescData = file_api_staff_staff_records_proto_rawDesc
)

func file_api_staff_staff_records_proto_rawDescGZIP() []byte {
	file_api_staff_staff_records_proto_rawDescOnce.Do(func() {
		file_api_staff_staff_records_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_staff_staff_records_proto_rawDescData)
	})
	return file_api_staff_staff_records_proto_rawDescData
}

var file_api_staff_staff_records_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_staff_staff_records_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_staff_staff_records_proto_goTypes = []interface{}{
	(HealthCheckResponse_ServingStatus)(0), // 0: StaffRecords.HealthCheckResponse.ServingStatus
	(*StaffRecord)(nil),                    // 1: StaffRecords.StaffRecord
	(*CreateStaff)(nil),                    // 2: StaffRecords.CreateStaff
	(*CreateResponse)(nil),                 // 3: StaffRecords.CreateResponse
	(*GetStaffRecords)(nil),                // 4: StaffRecords.GetStaffRecords
	(*GetResponse)(nil),                    // 5: StaffRecords.GetResponse
	(*GetStaffAvailability)(nil),           // 6: StaffRecords.GetStaffAvailability
	(*GetAvailabilityResponse)(nil),        // 7: StaffRecords.GetAvailabilityResponse
	(*UpdateStaff)(nil),                    // 8: StaffRecords.UpdateStaff
	(*UpdateResponse)(nil),                 // 9: StaffRecords.UpdateResponse
	(*DeleteStaff)(nil),                    // 10: StaffRecords.DeleteStaff
	(*DeleteResponse)(nil),                 // 11: StaffRecords.DeleteResponse
	(*HealthCheckRequest)(nil),             // 12: StaffRecords.HealthCheckRequest
	(*HealthCheckResponse)(nil),            // 13: StaffRecords.HealthCheckResponse
}
var file_api_staff_staff_records_proto_depIdxs = []int32{
	1,  // 0: StaffRecords.CreateStaff.staff:type_name -> StaffRecords.StaffRecord
	1,  // 1: StaffRecords.GetResponse.staffRecords:type_name -> StaffRecords.StaffRecord
	1,  // 2: StaffRecords.UpdateStaff.staffRecord:type_name -> StaffRecords.StaffRecord
	0,  // 3: StaffRecords.HealthCheckResponse.status:type_name -> StaffRecords.HealthCheckResponse.ServingStatus
	2,  // 4: StaffRecords.StaffRecords.Create:input_type -> StaffRecords.CreateStaff
	4,  // 5: StaffRecords.StaffRecords.Get:input_type -> StaffRecords.GetStaffRecords
	6,  // 6: StaffRecords.StaffRecords.GetAvailability:input_type -> StaffRecords.GetStaffAvailability
	8,  // 7: StaffRecords.StaffRecords.Update:input_type -> StaffRecords.UpdateStaff
	10, // 8: StaffRecords.StaffRecords.Delete:input_type -> StaffRecords.DeleteStaff
	12, // 9: StaffRecords.StaffRecords.Check:input_type -> StaffRecords.HealthCheckRequest
	3,  // 10: StaffRecords.StaffRecords.Create:output_type -> StaffRecords.CreateResponse
	5,  // 11: StaffRecords.StaffRecords.Get:output_type -> StaffRecords.GetResponse
	7,  // 12: StaffRecords.StaffRecords.GetAvailability:output_type -> StaffRecords.GetAvailabilityResponse
	9,  // 13: StaffRecords.StaffRecords.Update:output_type -> StaffRecords.UpdateResponse
	11, // 14: StaffRecords.StaffRecords.Delete:output_type -> StaffRecords.DeleteResponse
	13, // 15: StaffRecords.StaffRecords.Check:output_type -> StaffRecords.HealthCheckResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_staff_staff_records_proto_init() }
func file_api_staff_staff_records_proto_init() {
	if File_api_staff_staff_records_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_staff_staff_records_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaffRecord); i {
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
		file_api_staff_staff_records_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStaff); i {
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
		file_api_staff_staff_records_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_api_staff_staff_records_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaffRecords); i {
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
		file_api_staff_staff_records_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_api_staff_staff_records_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStaffAvailability); i {
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
		file_api_staff_staff_records_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAvailabilityResponse); i {
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
		file_api_staff_staff_records_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStaff); i {
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
		file_api_staff_staff_records_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_api_staff_staff_records_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStaff); i {
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
		file_api_staff_staff_records_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_api_staff_staff_records_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_api_staff_staff_records_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_api_staff_staff_records_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_staff_staff_records_proto_goTypes,
		DependencyIndexes: file_api_staff_staff_records_proto_depIdxs,
		EnumInfos:         file_api_staff_staff_records_proto_enumTypes,
		MessageInfos:      file_api_staff_staff_records_proto_msgTypes,
	}.Build()
	File_api_staff_staff_records_proto = out.File
	file_api_staff_staff_records_proto_rawDesc = nil
	file_api_staff_staff_records_proto_goTypes = nil
	file_api_staff_staff_records_proto_depIdxs = nil
}
