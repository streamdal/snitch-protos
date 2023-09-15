// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: sp_internal.proto

package protos

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

type NewAudienceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The session that is performing this call
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Newly created audience.
	Audience *Audience `protobuf:"bytes,2,opt,name=audience,proto3" json:"audience,omitempty"`
}

func (x *NewAudienceRequest) Reset() {
	*x = NewAudienceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAudienceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAudienceRequest) ProtoMessage() {}

func (x *NewAudienceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAudienceRequest.ProtoReflect.Descriptor instead.
func (*NewAudienceRequest) Descriptor() ([]byte, []int) {
	return file_sp_internal_proto_rawDescGZIP(), []int{0}
}

func (x *NewAudienceRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *NewAudienceRequest) GetAudience() *Audience {
	if x != nil {
		return x.Audience
	}
	return nil
}

// Each consumer and producer should send periodic heartbeats to the server
// to let the server know that they are still active.
type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Session ID for this instance of the SDK.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_sp_internal_proto_rawDescGZIP(), []int{1}
}

func (x *HeartbeatRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type NotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineId          string    `protobuf:"bytes,1,opt,name=pipeline_id,json=pipelineId,proto3" json:"pipeline_id,omitempty"`
	StepName            string    `protobuf:"bytes,2,opt,name=step_name,json=stepName,proto3" json:"step_name,omitempty"`
	Audience            *Audience `protobuf:"bytes,3,opt,name=audience,proto3" json:"audience,omitempty"`
	OccurredAtUnixTsUtc int64     `protobuf:"varint,4,opt,name=occurred_at_unix_ts_utc,json=occurredAtUnixTsUtc,proto3" json:"occurred_at_unix_ts_utc,omitempty"`
}

func (x *NotifyRequest) Reset() {
	*x = NotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRequest) ProtoMessage() {}

func (x *NotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRequest.ProtoReflect.Descriptor instead.
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return file_sp_internal_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyRequest) GetPipelineId() string {
	if x != nil {
		return x.PipelineId
	}
	return ""
}

func (x *NotifyRequest) GetStepName() string {
	if x != nil {
		return x.StepName
	}
	return ""
}

func (x *NotifyRequest) GetAudience() *Audience {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *NotifyRequest) GetOccurredAtUnixTsUtc() int64 {
	if x != nil {
		return x.OccurredAtUnixTsUtc
	}
	return 0
}

type MetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *MetricsRequest) Reset() {
	*x = MetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_internal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsRequest) ProtoMessage() {}

func (x *MetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_internal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsRequest.ProtoReflect.Descriptor instead.
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return file_sp_internal_proto_rawDescGZIP(), []int{3}
}

func (x *MetricsRequest) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED -- Name of the service that is registering.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// REQUIRED -- Unique ID for this SDK instance.
	//
	// This should be generated every time the SDK is instantiated (oe. every
	// time a NEW registration is performed).
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// REQUIRED -- Info about the client (lib name, lang, os, arch, etc.)
	ClientInfo *ClientInfo `protobuf:"bytes,3,opt,name=client_info,json=clientInfo,proto3" json:"client_info,omitempty"`
	// OPTIONAL -- if these are defined, these will show up in the UI even if
	// there is no active .Process() call from the SDK.
	Audiences []*Audience `protobuf:"bytes,4,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// OPTIONAL -- If set, we know that any pipelines or steps executed in this
	// SDK will NOT modify the input/output data. As in, the SDK will log what it
	// _would_ do and always return the original data set.
	DryRun bool `protobuf:"varint,5,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_internal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_internal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_sp_internal_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RegisterRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *RegisterRequest) GetClientInfo() *ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *RegisterRequest) GetAudiences() []*Audience {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *RegisterRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

// Same as RegisterRequest - used for broadcasting a deregistration event
type DeregisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	SessionId   string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *DeregisterRequest) Reset() {
	*x = DeregisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_internal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeregisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterRequest) ProtoMessage() {}

func (x *DeregisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_internal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterRequest.ProtoReflect.Descriptor instead.
func (*DeregisterRequest) Descriptor() ([]byte, []int) {
	return file_sp_internal_proto_rawDescGZIP(), []int{5}
}

func (x *DeregisterRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *DeregisterRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type GetAttachCommandsByServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *GetAttachCommandsByServiceRequest) Reset() {
	*x = GetAttachCommandsByServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_internal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachCommandsByServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachCommandsByServiceRequest) ProtoMessage() {}

func (x *GetAttachCommandsByServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_internal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachCommandsByServiceRequest.ProtoReflect.Descriptor instead.
func (*GetAttachCommandsByServiceRequest) Descriptor() ([]byte, []int) {
	return file_sp_internal_proto_rawDescGZIP(), []int{6}
}

func (x *GetAttachCommandsByServiceRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type GetAttachCommandsByServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active []*Command `protobuf:"bytes,1,rep,name=active,proto3" json:"active,omitempty"`
	Paused []*Command `protobuf:"bytes,2,rep,name=paused,proto3" json:"paused,omitempty"`
}

func (x *GetAttachCommandsByServiceResponse) Reset() {
	*x = GetAttachCommandsByServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_internal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachCommandsByServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachCommandsByServiceResponse) ProtoMessage() {}

func (x *GetAttachCommandsByServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sp_internal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachCommandsByServiceResponse.ProtoReflect.Descriptor instead.
func (*GetAttachCommandsByServiceResponse) Descriptor() ([]byte, []int) {
	return file_sp_internal_proto_rawDescGZIP(), []int{7}
}

func (x *GetAttachCommandsByServiceResponse) GetActive() []*Command {
	if x != nil {
		return x.Active
	}
	return nil
}

func (x *GetAttachCommandsByServiceResponse) GetPaused() []*Command {
	if x != nil {
		return x.Paused
	}
	return nil
}

type SendSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audience *Audience `protobuf:"bytes,1,opt,name=audience,proto3" json:"audience,omitempty"`
	Schema   []byte    `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *SendSchemaRequest) Reset() {
	*x = SendSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sp_internal_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSchemaRequest) ProtoMessage() {}

func (x *SendSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sp_internal_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSchemaRequest.ProtoReflect.Descriptor instead.
func (*SendSchemaRequest) Descriptor() ([]byte, []int) {
	return file_sp_internal_proto_rawDescGZIP(), []int{8}
}

func (x *SendSchemaRequest) GetAudience() *Audience {
	if x != nil {
		return x.Audience
	}
	return nil
}

func (x *SendSchemaRequest) GetSchema() []byte {
	if x != nil {
		return x.Schema
	}
	return nil
}

var File_sp_internal_proto protoreflect.FileDescriptor

var file_sp_internal_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x10, 0x73, 0x70, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x73,
	0x70, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x73, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a,
	0x12, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0x31, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x34, 0x0a, 0x17, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x74, 0x73, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x13, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x41, 0x74, 0x55, 0x6e,
	0x69, 0x78, 0x54, 0x73, 0x55, 0x74, 0x63, 0x22, 0x3a, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0b, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e,
	0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x22, 0x55, 0x0a, 0x11, 0x44, 0x65, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x46,
	0x0a, 0x21, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x76, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x42, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x22, 0x59,
	0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x32, 0xb6, 0x04, 0x0a, 0x08, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x30, 0x01, 0x12, 0x43,
	0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x42, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12,
	0x41, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x6c, 0x2f, 0x73, 0x6e, 0x69, 0x74, 0x63,
	0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sp_internal_proto_rawDescOnce sync.Once
	file_sp_internal_proto_rawDescData = file_sp_internal_proto_rawDesc
)

func file_sp_internal_proto_rawDescGZIP() []byte {
	file_sp_internal_proto_rawDescOnce.Do(func() {
		file_sp_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_sp_internal_proto_rawDescData)
	})
	return file_sp_internal_proto_rawDescData
}

var file_sp_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_sp_internal_proto_goTypes = []interface{}{
	(*NewAudienceRequest)(nil),                 // 0: protos.NewAudienceRequest
	(*HeartbeatRequest)(nil),                   // 1: protos.HeartbeatRequest
	(*NotifyRequest)(nil),                      // 2: protos.NotifyRequest
	(*MetricsRequest)(nil),                     // 3: protos.MetricsRequest
	(*RegisterRequest)(nil),                    // 4: protos.RegisterRequest
	(*DeregisterRequest)(nil),                  // 5: protos.DeregisterRequest
	(*GetAttachCommandsByServiceRequest)(nil),  // 6: protos.GetAttachCommandsByServiceRequest
	(*GetAttachCommandsByServiceResponse)(nil), // 7: protos.GetAttachCommandsByServiceResponse
	(*SendSchemaRequest)(nil),                  // 8: protos.SendSchemaRequest
	(*Audience)(nil),                           // 9: protos.Audience
	(*Metric)(nil),                             // 10: protos.Metric
	(*ClientInfo)(nil),                         // 11: protos.ClientInfo
	(*Command)(nil),                            // 12: protos.Command
	(*TailResponse)(nil),                       // 13: protos.TailResponse
	(*StandardResponse)(nil),                   // 14: protos.StandardResponse
}
var file_sp_internal_proto_depIdxs = []int32{
	9,  // 0: protos.NewAudienceRequest.audience:type_name -> protos.Audience
	9,  // 1: protos.NotifyRequest.audience:type_name -> protos.Audience
	10, // 2: protos.MetricsRequest.metrics:type_name -> protos.Metric
	11, // 3: protos.RegisterRequest.client_info:type_name -> protos.ClientInfo
	9,  // 4: protos.RegisterRequest.audiences:type_name -> protos.Audience
	12, // 5: protos.GetAttachCommandsByServiceResponse.active:type_name -> protos.Command
	12, // 6: protos.GetAttachCommandsByServiceResponse.paused:type_name -> protos.Command
	9,  // 7: protos.SendSchemaRequest.audience:type_name -> protos.Audience
	4,  // 8: protos.Internal.Register:input_type -> protos.RegisterRequest
	0,  // 9: protos.Internal.NewAudience:input_type -> protos.NewAudienceRequest
	1,  // 10: protos.Internal.Heartbeat:input_type -> protos.HeartbeatRequest
	2,  // 11: protos.Internal.Notify:input_type -> protos.NotifyRequest
	3,  // 12: protos.Internal.Metrics:input_type -> protos.MetricsRequest
	6,  // 13: protos.Internal.GetAttachCommandsByService:input_type -> protos.GetAttachCommandsByServiceRequest
	13, // 14: protos.Internal.SendTail:input_type -> protos.TailResponse
	8,  // 15: protos.Internal.SendSchema:input_type -> protos.SendSchemaRequest
	12, // 16: protos.Internal.Register:output_type -> protos.Command
	14, // 17: protos.Internal.NewAudience:output_type -> protos.StandardResponse
	14, // 18: protos.Internal.Heartbeat:output_type -> protos.StandardResponse
	14, // 19: protos.Internal.Notify:output_type -> protos.StandardResponse
	14, // 20: protos.Internal.Metrics:output_type -> protos.StandardResponse
	7,  // 21: protos.Internal.GetAttachCommandsByService:output_type -> protos.GetAttachCommandsByServiceResponse
	14, // 22: protos.Internal.SendTail:output_type -> protos.StandardResponse
	14, // 23: protos.Internal.SendSchema:output_type -> protos.StandardResponse
	16, // [16:24] is the sub-list for method output_type
	8,  // [8:16] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_sp_internal_proto_init() }
func file_sp_internal_proto_init() {
	if File_sp_internal_proto != nil {
		return
	}
	file_sp_command_proto_init()
	file_sp_common_proto_init()
	file_sp_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sp_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAudienceRequest); i {
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
		file_sp_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_sp_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRequest); i {
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
		file_sp_internal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsRequest); i {
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
		file_sp_internal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_sp_internal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeregisterRequest); i {
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
		file_sp_internal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachCommandsByServiceRequest); i {
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
		file_sp_internal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachCommandsByServiceResponse); i {
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
		file_sp_internal_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSchemaRequest); i {
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
			RawDescriptor: file_sp_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sp_internal_proto_goTypes,
		DependencyIndexes: file_sp_internal_proto_depIdxs,
		MessageInfos:      file_sp_internal_proto_msgTypes,
	}.Build()
	File_sp_internal_proto = out.File
	file_sp_internal_proto_rawDesc = nil
	file_sp_internal_proto_goTypes = nil
	file_sp_internal_proto_depIdxs = nil
}
