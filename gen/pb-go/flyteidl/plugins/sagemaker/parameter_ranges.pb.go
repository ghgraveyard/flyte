// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/sagemaker/parameter_ranges.proto

package plugins

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

type HyperparameterScalingType_Value int32

const (
	HyperparameterScalingType_AUTO               HyperparameterScalingType_Value = 0
	HyperparameterScalingType_LINEAR             HyperparameterScalingType_Value = 1
	HyperparameterScalingType_LOGARITHMIC        HyperparameterScalingType_Value = 2
	HyperparameterScalingType_REVERSELOGARITHMIC HyperparameterScalingType_Value = 3
)

var HyperparameterScalingType_Value_name = map[int32]string{
	0: "AUTO",
	1: "LINEAR",
	2: "LOGARITHMIC",
	3: "REVERSELOGARITHMIC",
}

var HyperparameterScalingType_Value_value = map[string]int32{
	"AUTO":               0,
	"LINEAR":             1,
	"LOGARITHMIC":        2,
	"REVERSELOGARITHMIC": 3,
}

func (x HyperparameterScalingType_Value) String() string {
	return proto.EnumName(HyperparameterScalingType_Value_name, int32(x))
}

func (HyperparameterScalingType_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5f31fcc87eba0a70, []int{0, 0}
}

// HyperparameterScalingType defines the way to increase or decrease the value of the hyperparameter
// For details, refer to: https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-ranges.html
// See examples of these scaling type, refer to: https://aws.amazon.com/blogs/machine-learning/amazon-sagemaker-automatic-model-tuning-now-supports-random-search-and-hyperparameter-scaling/
type HyperparameterScalingType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HyperparameterScalingType) Reset()         { *m = HyperparameterScalingType{} }
func (m *HyperparameterScalingType) String() string { return proto.CompactTextString(m) }
func (*HyperparameterScalingType) ProtoMessage()    {}
func (*HyperparameterScalingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f31fcc87eba0a70, []int{0}
}

func (m *HyperparameterScalingType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HyperparameterScalingType.Unmarshal(m, b)
}
func (m *HyperparameterScalingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HyperparameterScalingType.Marshal(b, m, deterministic)
}
func (m *HyperparameterScalingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HyperparameterScalingType.Merge(m, src)
}
func (m *HyperparameterScalingType) XXX_Size() int {
	return xxx_messageInfo_HyperparameterScalingType.Size(m)
}
func (m *HyperparameterScalingType) XXX_DiscardUnknown() {
	xxx_messageInfo_HyperparameterScalingType.DiscardUnknown(m)
}

var xxx_messageInfo_HyperparameterScalingType proto.InternalMessageInfo

// ContinuousParameterRange refers to a continuous range of hyperparameter values, allowing
// users to specify the search space of a floating-point hyperparameter
// https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-ranges.html
type ContinuousParameterRange struct {
	MaxValue             float64                         `protobuf:"fixed64,1,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	MinValue             float64                         `protobuf:"fixed64,2,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	ScalingType          HyperparameterScalingType_Value `protobuf:"varint,3,opt,name=scaling_type,json=scalingType,proto3,enum=flyteidl.plugins.sagemaker.HyperparameterScalingType_Value" json:"scaling_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ContinuousParameterRange) Reset()         { *m = ContinuousParameterRange{} }
func (m *ContinuousParameterRange) String() string { return proto.CompactTextString(m) }
func (*ContinuousParameterRange) ProtoMessage()    {}
func (*ContinuousParameterRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f31fcc87eba0a70, []int{1}
}

func (m *ContinuousParameterRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContinuousParameterRange.Unmarshal(m, b)
}
func (m *ContinuousParameterRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContinuousParameterRange.Marshal(b, m, deterministic)
}
func (m *ContinuousParameterRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousParameterRange.Merge(m, src)
}
func (m *ContinuousParameterRange) XXX_Size() int {
	return xxx_messageInfo_ContinuousParameterRange.Size(m)
}
func (m *ContinuousParameterRange) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousParameterRange.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousParameterRange proto.InternalMessageInfo

func (m *ContinuousParameterRange) GetMaxValue() float64 {
	if m != nil {
		return m.MaxValue
	}
	return 0
}

func (m *ContinuousParameterRange) GetMinValue() float64 {
	if m != nil {
		return m.MinValue
	}
	return 0
}

func (m *ContinuousParameterRange) GetScalingType() HyperparameterScalingType_Value {
	if m != nil {
		return m.ScalingType
	}
	return HyperparameterScalingType_AUTO
}

// IntegerParameterRange refers to a discrete range of hyperparameter values, allowing
// users to specify the search space of an integer hyperparameter
// https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-ranges.html
type IntegerParameterRange struct {
	MaxValue             int64                           `protobuf:"varint,1,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	MinValue             int64                           `protobuf:"varint,2,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	ScalingType          HyperparameterScalingType_Value `protobuf:"varint,3,opt,name=scaling_type,json=scalingType,proto3,enum=flyteidl.plugins.sagemaker.HyperparameterScalingType_Value" json:"scaling_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *IntegerParameterRange) Reset()         { *m = IntegerParameterRange{} }
func (m *IntegerParameterRange) String() string { return proto.CompactTextString(m) }
func (*IntegerParameterRange) ProtoMessage()    {}
func (*IntegerParameterRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f31fcc87eba0a70, []int{2}
}

func (m *IntegerParameterRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegerParameterRange.Unmarshal(m, b)
}
func (m *IntegerParameterRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegerParameterRange.Marshal(b, m, deterministic)
}
func (m *IntegerParameterRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegerParameterRange.Merge(m, src)
}
func (m *IntegerParameterRange) XXX_Size() int {
	return xxx_messageInfo_IntegerParameterRange.Size(m)
}
func (m *IntegerParameterRange) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegerParameterRange.DiscardUnknown(m)
}

var xxx_messageInfo_IntegerParameterRange proto.InternalMessageInfo

func (m *IntegerParameterRange) GetMaxValue() int64 {
	if m != nil {
		return m.MaxValue
	}
	return 0
}

func (m *IntegerParameterRange) GetMinValue() int64 {
	if m != nil {
		return m.MinValue
	}
	return 0
}

func (m *IntegerParameterRange) GetScalingType() HyperparameterScalingType_Value {
	if m != nil {
		return m.ScalingType
	}
	return HyperparameterScalingType_AUTO
}

// ContinuousParameterRange refers to a continuous range of hyperparameter values, allowing
// users to specify the search space of a floating-point hyperparameter
// https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-ranges.html
type CategoricalParameterRange struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoricalParameterRange) Reset()         { *m = CategoricalParameterRange{} }
func (m *CategoricalParameterRange) String() string { return proto.CompactTextString(m) }
func (*CategoricalParameterRange) ProtoMessage()    {}
func (*CategoricalParameterRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f31fcc87eba0a70, []int{3}
}

func (m *CategoricalParameterRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoricalParameterRange.Unmarshal(m, b)
}
func (m *CategoricalParameterRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoricalParameterRange.Marshal(b, m, deterministic)
}
func (m *CategoricalParameterRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoricalParameterRange.Merge(m, src)
}
func (m *CategoricalParameterRange) XXX_Size() int {
	return xxx_messageInfo_CategoricalParameterRange.Size(m)
}
func (m *CategoricalParameterRange) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoricalParameterRange.DiscardUnknown(m)
}

var xxx_messageInfo_CategoricalParameterRange proto.InternalMessageInfo

func (m *CategoricalParameterRange) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// ParameterRangeOneOf describes a single ParameterRange, which is a one-of structure that can be one of
// the three possible types: ContinuousParameterRange, IntegerParameterRange, and CategoricalParameterRange.
// This one-of structure in Flyte enables specifying a Parameter in a type-safe manner
// See: https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-ranges.html
type ParameterRangeOneOf struct {
	// Types that are valid to be assigned to ParameterRangeType:
	//	*ParameterRangeOneOf_ContinuousParameterRange
	//	*ParameterRangeOneOf_IntegerParameterRange
	//	*ParameterRangeOneOf_CategoricalParameterRange
	ParameterRangeType   isParameterRangeOneOf_ParameterRangeType `protobuf_oneof:"parameter_range_type"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ParameterRangeOneOf) Reset()         { *m = ParameterRangeOneOf{} }
func (m *ParameterRangeOneOf) String() string { return proto.CompactTextString(m) }
func (*ParameterRangeOneOf) ProtoMessage()    {}
func (*ParameterRangeOneOf) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f31fcc87eba0a70, []int{4}
}

func (m *ParameterRangeOneOf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParameterRangeOneOf.Unmarshal(m, b)
}
func (m *ParameterRangeOneOf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParameterRangeOneOf.Marshal(b, m, deterministic)
}
func (m *ParameterRangeOneOf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterRangeOneOf.Merge(m, src)
}
func (m *ParameterRangeOneOf) XXX_Size() int {
	return xxx_messageInfo_ParameterRangeOneOf.Size(m)
}
func (m *ParameterRangeOneOf) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterRangeOneOf.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterRangeOneOf proto.InternalMessageInfo

type isParameterRangeOneOf_ParameterRangeType interface {
	isParameterRangeOneOf_ParameterRangeType()
}

type ParameterRangeOneOf_ContinuousParameterRange struct {
	ContinuousParameterRange *ContinuousParameterRange `protobuf:"bytes,1,opt,name=continuous_parameter_range,json=continuousParameterRange,proto3,oneof"`
}

type ParameterRangeOneOf_IntegerParameterRange struct {
	IntegerParameterRange *IntegerParameterRange `protobuf:"bytes,2,opt,name=integer_parameter_range,json=integerParameterRange,proto3,oneof"`
}

type ParameterRangeOneOf_CategoricalParameterRange struct {
	CategoricalParameterRange *CategoricalParameterRange `protobuf:"bytes,3,opt,name=categorical_parameter_range,json=categoricalParameterRange,proto3,oneof"`
}

func (*ParameterRangeOneOf_ContinuousParameterRange) isParameterRangeOneOf_ParameterRangeType() {}

func (*ParameterRangeOneOf_IntegerParameterRange) isParameterRangeOneOf_ParameterRangeType() {}

func (*ParameterRangeOneOf_CategoricalParameterRange) isParameterRangeOneOf_ParameterRangeType() {}

func (m *ParameterRangeOneOf) GetParameterRangeType() isParameterRangeOneOf_ParameterRangeType {
	if m != nil {
		return m.ParameterRangeType
	}
	return nil
}

func (m *ParameterRangeOneOf) GetContinuousParameterRange() *ContinuousParameterRange {
	if x, ok := m.GetParameterRangeType().(*ParameterRangeOneOf_ContinuousParameterRange); ok {
		return x.ContinuousParameterRange
	}
	return nil
}

func (m *ParameterRangeOneOf) GetIntegerParameterRange() *IntegerParameterRange {
	if x, ok := m.GetParameterRangeType().(*ParameterRangeOneOf_IntegerParameterRange); ok {
		return x.IntegerParameterRange
	}
	return nil
}

func (m *ParameterRangeOneOf) GetCategoricalParameterRange() *CategoricalParameterRange {
	if x, ok := m.GetParameterRangeType().(*ParameterRangeOneOf_CategoricalParameterRange); ok {
		return x.CategoricalParameterRange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ParameterRangeOneOf) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ParameterRangeOneOf_ContinuousParameterRange)(nil),
		(*ParameterRangeOneOf_IntegerParameterRange)(nil),
		(*ParameterRangeOneOf_CategoricalParameterRange)(nil),
	}
}

// ParameterRanges is a map that maps hyperparameter name to the corresponding hyperparameter range
// https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-ranges.html
type ParameterRanges struct {
	ParameterRangeMap    map[string]*ParameterRangeOneOf `protobuf:"bytes,1,rep,name=parameter_range_map,json=parameterRangeMap,proto3" json:"parameter_range_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ParameterRanges) Reset()         { *m = ParameterRanges{} }
func (m *ParameterRanges) String() string { return proto.CompactTextString(m) }
func (*ParameterRanges) ProtoMessage()    {}
func (*ParameterRanges) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f31fcc87eba0a70, []int{5}
}

func (m *ParameterRanges) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParameterRanges.Unmarshal(m, b)
}
func (m *ParameterRanges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParameterRanges.Marshal(b, m, deterministic)
}
func (m *ParameterRanges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterRanges.Merge(m, src)
}
func (m *ParameterRanges) XXX_Size() int {
	return xxx_messageInfo_ParameterRanges.Size(m)
}
func (m *ParameterRanges) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterRanges.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterRanges proto.InternalMessageInfo

func (m *ParameterRanges) GetParameterRangeMap() map[string]*ParameterRangeOneOf {
	if m != nil {
		return m.ParameterRangeMap
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.sagemaker.HyperparameterScalingType_Value", HyperparameterScalingType_Value_name, HyperparameterScalingType_Value_value)
	proto.RegisterType((*HyperparameterScalingType)(nil), "flyteidl.plugins.sagemaker.HyperparameterScalingType")
	proto.RegisterType((*ContinuousParameterRange)(nil), "flyteidl.plugins.sagemaker.ContinuousParameterRange")
	proto.RegisterType((*IntegerParameterRange)(nil), "flyteidl.plugins.sagemaker.IntegerParameterRange")
	proto.RegisterType((*CategoricalParameterRange)(nil), "flyteidl.plugins.sagemaker.CategoricalParameterRange")
	proto.RegisterType((*ParameterRangeOneOf)(nil), "flyteidl.plugins.sagemaker.ParameterRangeOneOf")
	proto.RegisterType((*ParameterRanges)(nil), "flyteidl.plugins.sagemaker.ParameterRanges")
	proto.RegisterMapType((map[string]*ParameterRangeOneOf)(nil), "flyteidl.plugins.sagemaker.ParameterRanges.ParameterRangeMapEntry")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/sagemaker/parameter_ranges.proto", fileDescriptor_5f31fcc87eba0a70)
}

var fileDescriptor_5f31fcc87eba0a70 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0x1a, 0x56, 0xad, 0xaf, 0x88, 0x15, 0x8f, 0x95, 0xb6, 0xbb, 0x54, 0x3d, 0xf5, 0x42,
	0xa2, 0x75, 0x03, 0x21, 0x38, 0xb5, 0x55, 0xa0, 0x95, 0x36, 0x8a, 0xbc, 0xb2, 0x03, 0x07, 0x2a,
	0x37, 0x78, 0xc1, 0x6a, 0xe2, 0x58, 0x8e, 0x03, 0xcb, 0x9f, 0xc0, 0x9f, 0x83, 0xc4, 0x9f, 0xc7,
	0x01, 0xd5, 0xe9, 0xba, 0x2d, 0x6b, 0x02, 0xa7, 0xdd, 0x6c, 0xbf, 0x1f, 0xdf, 0xf7, 0xde, 0xf7,
	0xc9, 0x70, 0x74, 0xe9, 0x27, 0x8a, 0xb2, 0xaf, 0xbe, 0x2d, 0xfc, 0xd8, 0x63, 0x3c, 0xb2, 0x23,
	0xe2, 0xd1, 0x80, 0x2c, 0xa9, 0xb4, 0x05, 0x91, 0x24, 0xa0, 0x8a, 0xca, 0xb9, 0x24, 0xdc, 0xa3,
	0x91, 0x25, 0x64, 0xa8, 0x42, 0xd4, 0xbe, 0x2e, 0xb1, 0xd6, 0x25, 0xd6, 0xa6, 0xa4, 0xeb, 0x42,
	0x6b, 0x9c, 0x08, 0x2a, 0x37, 0xa5, 0xe7, 0x2e, 0xf1, 0x19, 0xf7, 0x66, 0x89, 0xa0, 0xdd, 0x77,
	0xb0, 0x73, 0x41, 0xfc, 0x98, 0xa2, 0x5d, 0x78, 0x34, 0xf8, 0x34, 0x9b, 0xd6, 0x4b, 0x08, 0xa0,
	0x72, 0x3a, 0xf9, 0xe0, 0x0c, 0x70, 0xdd, 0x40, 0x7b, 0x50, 0x3b, 0x9d, 0xbe, 0x1f, 0xe0, 0xc9,
	0x6c, 0x7c, 0x36, 0x19, 0xd5, 0xcb, 0xa8, 0x01, 0x08, 0x3b, 0x17, 0x0e, 0x3e, 0x77, 0x6e, 0xbf,
	0x9b, 0xdd, 0xdf, 0x06, 0x34, 0x47, 0x21, 0x57, 0x8c, 0xc7, 0x61, 0x1c, 0x7d, 0xbc, 0x86, 0xc2,
	0x2b, 0x92, 0xe8, 0x10, 0xaa, 0x01, 0xb9, 0x9a, 0x7f, 0x5f, 0x01, 0x35, 0x8d, 0x8e, 0xd1, 0x33,
	0xf0, 0x6e, 0x40, 0xae, 0x52, 0xe0, 0x55, 0x90, 0xf1, 0x75, 0xb0, 0xbc, 0x0e, 0x32, 0x9e, 0x06,
	0xbf, 0xc0, 0xe3, 0x28, 0x65, 0x3b, 0x57, 0x89, 0xa0, 0x4d, 0xb3, 0x63, 0xf4, 0x9e, 0xf4, 0xdf,
	0x5a, 0xf9, 0xe3, 0x5a, 0xb9, 0xb3, 0x5a, 0xba, 0x25, 0xae, 0x45, 0xb7, 0xc6, 0xff, 0x65, 0xc0,
	0xc1, 0x84, 0x2b, 0xea, 0x51, 0xf9, 0x2f, 0xce, 0x66, 0x11, 0x67, 0xf3, 0x01, 0x39, 0x1f, 0x43,
	0x6b, 0x44, 0x14, 0xf5, 0x42, 0xc9, 0x5c, 0xe2, 0x67, 0x68, 0x37, 0xa0, 0xa2, 0x59, 0x45, 0x4d,
	0xa3, 0x63, 0xf6, 0xaa, 0x78, 0x7d, 0xeb, 0xfe, 0x34, 0x61, 0xff, 0x6e, 0xea, 0x94, 0xd3, 0xe9,
	0x25, 0x52, 0xd0, 0x76, 0x37, 0xb2, 0xcd, 0x33, 0xee, 0xd2, 0x73, 0xd7, 0xfa, 0x27, 0x45, 0xd4,
	0xf3, 0x44, 0x1f, 0x97, 0x70, 0xd3, 0xcd, 0x33, 0xc4, 0x12, 0x9e, 0xb3, 0x74, 0xeb, 0xf7, 0x20,
	0xcb, 0x1a, 0xf2, 0xa8, 0x08, 0x72, 0xab, 0x60, 0xe3, 0x12, 0x3e, 0x60, 0x5b, 0x95, 0xfc, 0x01,
	0x87, 0xee, 0xcd, 0xbe, 0xee, 0x01, 0x9a, 0x1a, 0xf0, 0x65, 0xe1, 0x8c, 0x79, 0xeb, 0x1e, 0x97,
	0x70, 0xcb, 0xcd, 0x0b, 0x0e, 0x1b, 0xf0, 0x2c, 0x03, 0xa6, 0x0d, 0xd1, 0xfd, 0x63, 0xc0, 0xde,
	0xdd, 0xd4, 0x08, 0x49, 0xd8, 0xcf, 0xe6, 0x06, 0x44, 0x68, 0x11, 0x6b, 0xfd, 0x61, 0x11, 0xb9,
	0x4c, 0xa7, 0xcc, 0xfd, 0x8c, 0x08, 0x87, 0x2b, 0x99, 0xe0, 0xa7, 0x22, 0xfb, 0xde, 0x8e, 0xa1,
	0xb1, 0x3d, 0x19, 0xd5, 0xc1, 0x5c, 0xd2, 0x44, 0xcb, 0x5f, 0xc5, 0xab, 0x23, 0x72, 0x60, 0xe7,
	0xc6, 0xed, 0xb5, 0xbe, 0xfd, 0xff, 0x8c, 0xb4, 0xcf, 0x70, 0x5a, 0xfd, 0xa6, 0xfc, 0xda, 0x18,
	0xbe, 0xfa, 0x7c, 0xe2, 0x31, 0xf5, 0x2d, 0x5e, 0x58, 0x6e, 0x18, 0xd8, 0xba, 0x4f, 0x28, 0xbd,
	0xf4, 0x60, 0x7b, 0x94, 0xdb, 0x62, 0xf1, 0xc2, 0x0b, 0xed, 0xec, 0x27, 0xb8, 0xa8, 0xe8, 0xaf,
	0xee, 0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x56, 0x2d, 0x28, 0x1f, 0x05, 0x00, 0x00,
}
