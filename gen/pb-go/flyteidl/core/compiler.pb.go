// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/compiler.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Adjacency list for the workflow. This is created as part of the compilation process. Every process after the compilation
// step uses this created ConnectionSet
type ConnectionSet struct {
	// A list of all the node ids that are downstream from a given node id
	Downstream map[string]*ConnectionSet_IdList `protobuf:"bytes,7,rep,name=downstream,proto3" json:"downstream,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of all the node ids, that are upstream of this node id
	Upstream             map[string]*ConnectionSet_IdList `protobuf:"bytes,8,rep,name=upstream,proto3" json:"upstream,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ConnectionSet) Reset()         { *m = ConnectionSet{} }
func (m *ConnectionSet) String() string { return proto.CompactTextString(m) }
func (*ConnectionSet) ProtoMessage()    {}
func (*ConnectionSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{0}
}

func (m *ConnectionSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionSet.Unmarshal(m, b)
}
func (m *ConnectionSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionSet.Marshal(b, m, deterministic)
}
func (m *ConnectionSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionSet.Merge(m, src)
}
func (m *ConnectionSet) XXX_Size() int {
	return xxx_messageInfo_ConnectionSet.Size(m)
}
func (m *ConnectionSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionSet.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionSet proto.InternalMessageInfo

func (m *ConnectionSet) GetDownstream() map[string]*ConnectionSet_IdList {
	if m != nil {
		return m.Downstream
	}
	return nil
}

func (m *ConnectionSet) GetUpstream() map[string]*ConnectionSet_IdList {
	if m != nil {
		return m.Upstream
	}
	return nil
}

type ConnectionSet_IdList struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionSet_IdList) Reset()         { *m = ConnectionSet_IdList{} }
func (m *ConnectionSet_IdList) String() string { return proto.CompactTextString(m) }
func (*ConnectionSet_IdList) ProtoMessage()    {}
func (*ConnectionSet_IdList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{0, 0}
}

func (m *ConnectionSet_IdList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionSet_IdList.Unmarshal(m, b)
}
func (m *ConnectionSet_IdList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionSet_IdList.Marshal(b, m, deterministic)
}
func (m *ConnectionSet_IdList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionSet_IdList.Merge(m, src)
}
func (m *ConnectionSet_IdList) XXX_Size() int {
	return xxx_messageInfo_ConnectionSet_IdList.Size(m)
}
func (m *ConnectionSet_IdList) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionSet_IdList.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionSet_IdList proto.InternalMessageInfo

func (m *ConnectionSet_IdList) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

// Output of the compilation Step. This object represents one workflow. We store more metadata at this layer
type CompiledWorkflow struct {
	// Completely contained Workflow Template
	Template *WorkflowTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	// For internal use only! This field is used by the system and must not be filled in. Any values set will be ignored.
	Connections          *ConnectionSet `protobuf:"bytes,2,opt,name=connections,proto3" json:"connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CompiledWorkflow) Reset()         { *m = CompiledWorkflow{} }
func (m *CompiledWorkflow) String() string { return proto.CompactTextString(m) }
func (*CompiledWorkflow) ProtoMessage()    {}
func (*CompiledWorkflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{1}
}

func (m *CompiledWorkflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompiledWorkflow.Unmarshal(m, b)
}
func (m *CompiledWorkflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompiledWorkflow.Marshal(b, m, deterministic)
}
func (m *CompiledWorkflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompiledWorkflow.Merge(m, src)
}
func (m *CompiledWorkflow) XXX_Size() int {
	return xxx_messageInfo_CompiledWorkflow.Size(m)
}
func (m *CompiledWorkflow) XXX_DiscardUnknown() {
	xxx_messageInfo_CompiledWorkflow.DiscardUnknown(m)
}

var xxx_messageInfo_CompiledWorkflow proto.InternalMessageInfo

func (m *CompiledWorkflow) GetTemplate() *WorkflowTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *CompiledWorkflow) GetConnections() *ConnectionSet {
	if m != nil {
		return m.Connections
	}
	return nil
}

// Output of the Compilation step. This object represent one Task. We store more metadata at this layer
type CompiledTask struct {
	// Completely contained TaskTemplate
	Template             *TaskTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CompiledTask) Reset()         { *m = CompiledTask{} }
func (m *CompiledTask) String() string { return proto.CompactTextString(m) }
func (*CompiledTask) ProtoMessage()    {}
func (*CompiledTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{2}
}

func (m *CompiledTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompiledTask.Unmarshal(m, b)
}
func (m *CompiledTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompiledTask.Marshal(b, m, deterministic)
}
func (m *CompiledTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompiledTask.Merge(m, src)
}
func (m *CompiledTask) XXX_Size() int {
	return xxx_messageInfo_CompiledTask.Size(m)
}
func (m *CompiledTask) XXX_DiscardUnknown() {
	xxx_messageInfo_CompiledTask.DiscardUnknown(m)
}

var xxx_messageInfo_CompiledTask proto.InternalMessageInfo

func (m *CompiledTask) GetTemplate() *TaskTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

// A Compiled Workflow Closure contains all the information required to start a new execution, or to visualize a workflow
// and its details. The CompiledWorkflowClosure should always contain a primary workflow, that is the main workflow that
// will being the execution. All subworkflows are denormalized. WorkflowNodes refer to the workflow identifiers of
// compiled subworkflows.
type CompiledWorkflowClosure struct {
	//+required
	Primary *CompiledWorkflow `protobuf:"bytes,1,opt,name=primary,proto3" json:"primary,omitempty"`
	// Guaranteed that there will only exist one and only one workflow with a given id, i.e., every sub workflow has a
	// unique identifier. Also every enclosed subworkflow is used either by a primary workflow or by a subworkflow
	// as an inlined workflow
	//+optional
	SubWorkflows []*CompiledWorkflow `protobuf:"bytes,2,rep,name=sub_workflows,json=subWorkflows,proto3" json:"sub_workflows,omitempty"`
	// Guaranteed that there will only exist one and only one task with a given id, i.e., every task has a unique id
	//+required (at least 1)
	Tasks                []*CompiledTask `protobuf:"bytes,3,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CompiledWorkflowClosure) Reset()         { *m = CompiledWorkflowClosure{} }
func (m *CompiledWorkflowClosure) String() string { return proto.CompactTextString(m) }
func (*CompiledWorkflowClosure) ProtoMessage()    {}
func (*CompiledWorkflowClosure) Descriptor() ([]byte, []int) {
	return fileDescriptor_d310dba155e4f065, []int{3}
}

func (m *CompiledWorkflowClosure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompiledWorkflowClosure.Unmarshal(m, b)
}
func (m *CompiledWorkflowClosure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompiledWorkflowClosure.Marshal(b, m, deterministic)
}
func (m *CompiledWorkflowClosure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompiledWorkflowClosure.Merge(m, src)
}
func (m *CompiledWorkflowClosure) XXX_Size() int {
	return xxx_messageInfo_CompiledWorkflowClosure.Size(m)
}
func (m *CompiledWorkflowClosure) XXX_DiscardUnknown() {
	xxx_messageInfo_CompiledWorkflowClosure.DiscardUnknown(m)
}

var xxx_messageInfo_CompiledWorkflowClosure proto.InternalMessageInfo

func (m *CompiledWorkflowClosure) GetPrimary() *CompiledWorkflow {
	if m != nil {
		return m.Primary
	}
	return nil
}

func (m *CompiledWorkflowClosure) GetSubWorkflows() []*CompiledWorkflow {
	if m != nil {
		return m.SubWorkflows
	}
	return nil
}

func (m *CompiledWorkflowClosure) GetTasks() []*CompiledTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectionSet)(nil), "flyteidl.core.ConnectionSet")
	proto.RegisterMapType((map[string]*ConnectionSet_IdList)(nil), "flyteidl.core.ConnectionSet.DownstreamEntry")
	proto.RegisterMapType((map[string]*ConnectionSet_IdList)(nil), "flyteidl.core.ConnectionSet.UpstreamEntry")
	proto.RegisterType((*ConnectionSet_IdList)(nil), "flyteidl.core.ConnectionSet.IdList")
	proto.RegisterType((*CompiledWorkflow)(nil), "flyteidl.core.CompiledWorkflow")
	proto.RegisterType((*CompiledTask)(nil), "flyteidl.core.CompiledTask")
	proto.RegisterType((*CompiledWorkflowClosure)(nil), "flyteidl.core.CompiledWorkflowClosure")
}

func init() { proto.RegisterFile("flyteidl/core/compiler.proto", fileDescriptor_d310dba155e4f065) }

var fileDescriptor_d310dba155e4f065 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcb, 0xca, 0xd3, 0x40,
	0x14, 0xc7, 0xc9, 0x17, 0x7a, 0xf1, 0xa4, 0xc1, 0x32, 0x1b, 0x63, 0x5a, 0x30, 0xc4, 0x4d, 0x11,
	0x4d, 0x68, 0xbb, 0xd0, 0x2a, 0xb8, 0xb0, 0x55, 0x11, 0xba, 0x8a, 0x15, 0xc1, 0x8d, 0xe6, 0x32,
	0x8d, 0x21, 0x97, 0x09, 0x33, 0x13, 0x4b, 0x9e, 0xc0, 0xa5, 0x4f, 0xe6, 0x3b, 0x49, 0xae, 0x34,
	0x29, 0x2d, 0x2e, 0xbe, 0xdd, 0x84, 0xf3, 0xfb, 0x5f, 0x0e, 0x93, 0x81, 0xf9, 0x31, 0xca, 0x39,
	0x0e, 0xbc, 0xc8, 0x74, 0x09, 0xc5, 0xa6, 0x4b, 0xe2, 0x34, 0x88, 0x30, 0x35, 0x52, 0x4a, 0x38,
	0x41, 0x72, 0x33, 0x35, 0x8a, 0xa9, 0xda, 0x83, 0x4f, 0x84, 0x86, 0xc7, 0x88, 0x9c, 0x2a, 0x58,
	0x7d, 0xdc, 0x9d, 0x72, 0x9b, 0x85, 0xac, 0x1a, 0xe9, 0xbf, 0x45, 0x90, 0xb7, 0x24, 0x49, 0xb0,
	0xcb, 0x03, 0x92, 0x7c, 0xc6, 0x1c, 0xed, 0x01, 0x3c, 0x72, 0x4a, 0x18, 0xa7, 0xd8, 0x8e, 0x95,
	0x91, 0x26, 0x2e, 0xa4, 0xd5, 0x73, 0xa3, 0x13, 0x67, 0x74, 0x14, 0xc6, 0xae, 0xc5, 0xdf, 0x27,
	0x9c, 0xe6, 0xd6, 0x99, 0x1e, 0x7d, 0x80, 0x71, 0x96, 0xd6, 0x5e, 0xe3, 0xd2, 0xeb, 0xd9, 0x4d,
	0xaf, 0x2f, 0xe9, 0xb9, 0x53, 0xab, 0x55, 0x55, 0x18, 0x7e, 0xf2, 0xf6, 0x01, 0xe3, 0x68, 0x0a,
	0x62, 0xe0, 0x31, 0x45, 0xd0, 0xc4, 0xc5, 0x03, 0xab, 0x38, 0xaa, 0x0e, 0x3c, 0xec, 0x55, 0x28,
	0xa0, 0x10, 0xe7, 0x8a, 0xa0, 0x09, 0x05, 0x14, 0xe2, 0x1c, 0x6d, 0x60, 0xf0, 0xcb, 0x8e, 0x32,
	0xac, 0xdc, 0x69, 0xc2, 0x42, 0x5a, 0x3d, 0xbd, 0xd9, 0xa2, 0x8a, 0xb2, 0x2a, 0xc5, 0xeb, 0xbb,
	0x57, 0x82, 0xfa, 0x03, 0xe4, 0x4e, 0xb5, 0x7b, 0x4f, 0xd0, 0xff, 0x08, 0x30, 0xdd, 0x56, 0x97,
	0xec, 0x7d, 0xad, 0xef, 0x0f, 0xbd, 0x81, 0x31, 0xc7, 0x71, 0x1a, 0xd9, 0x1c, 0x97, 0x51, 0xd2,
	0xea, 0x49, 0xcf, 0xb6, 0x41, 0x0f, 0x35, 0x66, 0xb5, 0x02, 0xf4, 0x16, 0x24, 0xb7, 0x0d, 0x65,
	0x75, 0xad, 0xf9, 0xad, 0x5a, 0xd6, 0xb9, 0x40, 0xff, 0x08, 0x93, 0xa6, 0xd0, 0xc1, 0x66, 0x21,
	0x7a, 0x79, 0x51, 0x66, 0xd6, 0x33, 0x2b, 0xb0, 0xcb, 0x22, 0xfa, 0x5f, 0x01, 0x1e, 0xf5, 0x57,
	0xdb, 0x46, 0x84, 0x65, 0x14, 0xa3, 0x0d, 0x8c, 0x52, 0x1a, 0xc4, 0x36, 0xcd, 0xaf, 0x2c, 0xd8,
	0x17, 0x5a, 0x0d, 0x8f, 0x76, 0x20, 0xb3, 0xcc, 0xf9, 0xde, 0xfc, 0xec, 0xc5, 0x86, 0xe2, 0xff,
	0x18, 0x4c, 0x58, 0xe6, 0x34, 0x1f, 0x0c, 0x2d, 0x61, 0x50, 0x3e, 0x08, 0x45, 0x2c, 0xd5, 0xb3,
	0x2b, 0xea, 0x62, 0x35, 0xab, 0x22, 0xdf, 0xad, 0xbf, 0x2d, 0xfd, 0x80, 0xff, 0xcc, 0x1c, 0xc3,
	0x25, 0xb1, 0x59, 0xf2, 0x84, 0xfa, 0xd5, 0xc1, 0xf4, 0x71, 0x62, 0xa6, 0xce, 0x0b, 0x9f, 0x98,
	0x9d, 0x57, 0xe7, 0x0c, 0xcb, 0x07, 0xb7, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x12, 0xd3, 0x6d,
	0xb8, 0xd8, 0x03, 0x00, 0x00,
}
