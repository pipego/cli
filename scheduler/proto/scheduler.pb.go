// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: scheduler/proto/scheduler.proto

package scheduler

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

// The request message.
type ServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string    `protobuf:"bytes,1,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind       string    `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata   *Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec       *Spec     `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *ServerRequest) Reset() {
	*x = ServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerRequest) ProtoMessage() {}

func (x *ServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerRequest.ProtoReflect.Descriptor instead.
func (*ServerRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *ServerRequest) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *ServerRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ServerRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ServerRequest) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task  *Task   `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Nodes []*Node `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *Spec) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *Spec) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                   string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NodeName               string             `protobuf:"bytes,2,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	NodeSelectors          []string           `protobuf:"bytes,3,rep,name=nodeSelectors,proto3" json:"nodeSelectors,omitempty"`
	RequestedResource      *RequestedResource `protobuf:"bytes,4,opt,name=requestedResource,proto3" json:"requestedResource,omitempty"`
	ToleratesUnschedulable bool               `protobuf:"varint,5,opt,name=toleratesUnschedulable,proto3" json:"toleratesUnschedulable,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_scheduler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_scheduler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_scheduler_proto_rawDescGZIP(), []int{3}
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Task) GetNodeSelectors() []string {
	if x != nil {
		return x.NodeSelectors
	}
	return nil
}

func (x *Task) GetRequestedResource() *RequestedResource {
	if x != nil {
		return x.RequestedResource
	}
	return nil
}

func (x *Task) GetToleratesUnschedulable() bool {
	if x != nil {
		return x.ToleratesUnschedulable
	}
	return false
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Host                string               `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Label               string               `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	AllocatableResource *AllocatableResource `protobuf:"bytes,4,opt,name=allocatableResource,proto3" json:"allocatableResource,omitempty"`
	RequestedResource   *RequestedResource   `protobuf:"bytes,5,opt,name=requestedResource,proto3" json:"requestedResource,omitempty"`
	Unschedulable       bool                 `protobuf:"varint,6,opt,name=unschedulable,proto3" json:"unschedulable,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_scheduler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_scheduler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_scheduler_proto_rawDescGZIP(), []int{4}
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Node) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Node) GetAllocatableResource() *AllocatableResource {
	if x != nil {
		return x.AllocatableResource
	}
	return nil
}

func (x *Node) GetRequestedResource() *RequestedResource {
	if x != nil {
		return x.RequestedResource
	}
	return nil
}

func (x *Node) GetUnschedulable() bool {
	if x != nil {
		return x.Unschedulable
	}
	return false
}

type AllocatableResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MilliCPU int64 `protobuf:"varint,1,opt,name=milliCPU,proto3" json:"milliCPU,omitempty"`
	Memory   int64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage  int64 `protobuf:"varint,3,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *AllocatableResource) Reset() {
	*x = AllocatableResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_scheduler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocatableResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocatableResource) ProtoMessage() {}

func (x *AllocatableResource) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_scheduler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocatableResource.ProtoReflect.Descriptor instead.
func (*AllocatableResource) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_scheduler_proto_rawDescGZIP(), []int{5}
}

func (x *AllocatableResource) GetMilliCPU() int64 {
	if x != nil {
		return x.MilliCPU
	}
	return 0
}

func (x *AllocatableResource) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *AllocatableResource) GetStorage() int64 {
	if x != nil {
		return x.Storage
	}
	return 0
}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_scheduler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_scheduler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_scheduler_proto_rawDescGZIP(), []int{6}
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RequestedResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MilliCPU int64 `protobuf:"varint,1,opt,name=milliCPU,proto3" json:"milliCPU,omitempty"`
	Memory   int64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage  int64 `protobuf:"varint,3,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *RequestedResource) Reset() {
	*x = RequestedResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_scheduler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestedResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestedResource) ProtoMessage() {}

func (x *RequestedResource) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_scheduler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestedResource.ProtoReflect.Descriptor instead.
func (*RequestedResource) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_scheduler_proto_rawDescGZIP(), []int{7}
}

func (x *RequestedResource) GetMilliCPU() int64 {
	if x != nil {
		return x.MilliCPU
	}
	return 0
}

func (x *RequestedResource) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *RequestedResource) GetStorage() int64 {
	if x != nil {
		return x.Storage
	}
	return 0
}

// The response message.
type ServerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServerReply) Reset() {
	*x = ServerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_proto_scheduler_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReply) ProtoMessage() {}

func (x *ServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_proto_scheduler_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReply.ProtoReflect.Descriptor instead.
func (*ServerReply) Descriptor() ([]byte, []int) {
	return file_scheduler_proto_scheduler_proto_rawDescGZIP(), []int{8}
}

func (x *ServerReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_scheduler_proto_scheduler_proto protoreflect.FileDescriptor

var file_scheduler_proto_scheduler_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0x99, 0x01, 0x0a,
	0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x1e, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x23, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x25, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xe0, 0x01, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x6f,
	0x64, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x4a, 0x0a, 0x11, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x74, 0x6f, 0x6c, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x73, 0x55, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x73, 0x55, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x88, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x50, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x75, 0x6e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x63, 0x0a, 0x13, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x43, 0x50, 0x55, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x43, 0x50, 0x55, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22,
	0x1b, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x11,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x43, 0x50, 0x55, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x43, 0x50, 0x55, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22,
	0x37, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x4f, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x40, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x67, 0x6f, 0x2f, 0x63,
	0x6c, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scheduler_proto_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_proto_scheduler_proto_rawDescData = file_scheduler_proto_scheduler_proto_rawDesc
)

func file_scheduler_proto_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_proto_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_proto_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_proto_scheduler_proto_rawDescData)
	})
	return file_scheduler_proto_scheduler_proto_rawDescData
}

var file_scheduler_proto_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_scheduler_proto_scheduler_proto_goTypes = []interface{}{
	(*ServerRequest)(nil),       // 0: scheduler.ServerRequest
	(*Metadata)(nil),            // 1: scheduler.Metadata
	(*Spec)(nil),                // 2: scheduler.Spec
	(*Task)(nil),                // 3: scheduler.Task
	(*Node)(nil),                // 4: scheduler.Node
	(*AllocatableResource)(nil), // 5: scheduler.AllocatableResource
	(*Label)(nil),               // 6: scheduler.Label
	(*RequestedResource)(nil),   // 7: scheduler.RequestedResource
	(*ServerReply)(nil),         // 8: scheduler.ServerReply
}
var file_scheduler_proto_scheduler_proto_depIdxs = []int32{
	1, // 0: scheduler.ServerRequest.metadata:type_name -> scheduler.Metadata
	2, // 1: scheduler.ServerRequest.spec:type_name -> scheduler.Spec
	3, // 2: scheduler.Spec.task:type_name -> scheduler.Task
	4, // 3: scheduler.Spec.nodes:type_name -> scheduler.Node
	7, // 4: scheduler.Task.requestedResource:type_name -> scheduler.RequestedResource
	5, // 5: scheduler.Node.allocatableResource:type_name -> scheduler.AllocatableResource
	7, // 6: scheduler.Node.requestedResource:type_name -> scheduler.RequestedResource
	0, // 7: scheduler.ServerProto.SendServer:input_type -> scheduler.ServerRequest
	8, // 8: scheduler.ServerProto.SendServer:output_type -> scheduler.ServerReply
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_scheduler_proto_scheduler_proto_init() }
func file_scheduler_proto_scheduler_proto_init() {
	if File_scheduler_proto_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_proto_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerRequest); i {
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
		file_scheduler_proto_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_scheduler_proto_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Spec); i {
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
		file_scheduler_proto_scheduler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_scheduler_proto_scheduler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_scheduler_proto_scheduler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocatableResource); i {
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
		file_scheduler_proto_scheduler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_scheduler_proto_scheduler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestedResource); i {
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
		file_scheduler_proto_scheduler_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReply); i {
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
			RawDescriptor: file_scheduler_proto_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduler_proto_scheduler_proto_goTypes,
		DependencyIndexes: file_scheduler_proto_scheduler_proto_depIdxs,
		MessageInfos:      file_scheduler_proto_scheduler_proto_msgTypes,
	}.Build()
	File_scheduler_proto_scheduler_proto = out.File
	file_scheduler_proto_scheduler_proto_rawDesc = nil
	file_scheduler_proto_scheduler_proto_goTypes = nil
	file_scheduler_proto_scheduler_proto_depIdxs = nil
}
