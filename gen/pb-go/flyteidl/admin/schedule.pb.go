// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: flyteidl/admin/schedule.proto

package admin

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

// Represents a frequency at which to run a schedule.
type FixedRateUnit int32

const (
	FixedRateUnit_MINUTE FixedRateUnit = 0
	FixedRateUnit_HOUR   FixedRateUnit = 1
	FixedRateUnit_DAY    FixedRateUnit = 2
)

// Enum value maps for FixedRateUnit.
var (
	FixedRateUnit_name = map[int32]string{
		0: "MINUTE",
		1: "HOUR",
		2: "DAY",
	}
	FixedRateUnit_value = map[string]int32{
		"MINUTE": 0,
		"HOUR":   1,
		"DAY":    2,
	}
)

func (x FixedRateUnit) Enum() *FixedRateUnit {
	p := new(FixedRateUnit)
	*p = x
	return p
}

func (x FixedRateUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FixedRateUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_flyteidl_admin_schedule_proto_enumTypes[0].Descriptor()
}

func (FixedRateUnit) Type() protoreflect.EnumType {
	return &file_flyteidl_admin_schedule_proto_enumTypes[0]
}

func (x FixedRateUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FixedRateUnit.Descriptor instead.
func (FixedRateUnit) EnumDescriptor() ([]byte, []int) {
	return file_flyteidl_admin_schedule_proto_rawDescGZIP(), []int{0}
}

// Option for schedules run at a certain frequency e.g. every 2 minutes.
type FixedRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32        `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit  FixedRateUnit `protobuf:"varint,2,opt,name=unit,proto3,enum=flyteidl.admin.FixedRateUnit" json:"unit,omitempty"`
}

func (x *FixedRate) Reset() {
	*x = FixedRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FixedRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixedRate) ProtoMessage() {}

func (x *FixedRate) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixedRate.ProtoReflect.Descriptor instead.
func (*FixedRate) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *FixedRate) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *FixedRate) GetUnit() FixedRateUnit {
	if x != nil {
		return x.Unit
	}
	return FixedRateUnit_MINUTE
}

// Options for schedules to run according to a cron expression.
type CronSchedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Standard/default cron implementation as described by https://en.wikipedia.org/wiki/Cron#CRON_expression;
	// Also supports nonstandard predefined scheduling definitions
	// as described by https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions
	// except @reboot
	Schedule string `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// ISO 8601 duration as described by https://en.wikipedia.org/wiki/ISO_8601#Durations
	Offset string `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *CronSchedule) Reset() {
	*x = CronSchedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronSchedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronSchedule) ProtoMessage() {}

func (x *CronSchedule) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronSchedule.ProtoReflect.Descriptor instead.
func (*CronSchedule) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *CronSchedule) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *CronSchedule) GetOffset() string {
	if x != nil {
		return x.Offset
	}
	return ""
}

// Defines complete set of information required to trigger an execution on a schedule.
type Schedule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ScheduleExpression:
	//	*Schedule_CronExpression
	//	*Schedule_Rate
	//	*Schedule_CronSchedule
	ScheduleExpression isSchedule_ScheduleExpression `protobuf_oneof:"ScheduleExpression"`
	// Name of the input variable that the kickoff time will be supplied to when the workflow is kicked off.
	KickoffTimeInputArg string `protobuf:"bytes,3,opt,name=kickoff_time_input_arg,json=kickoffTimeInputArg,proto3" json:"kickoff_time_input_arg,omitempty"`
}

func (x *Schedule) Reset() {
	*x = Schedule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_schedule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schedule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schedule) ProtoMessage() {}

func (x *Schedule) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_schedule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schedule.ProtoReflect.Descriptor instead.
func (*Schedule) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_schedule_proto_rawDescGZIP(), []int{2}
}

func (m *Schedule) GetScheduleExpression() isSchedule_ScheduleExpression {
	if m != nil {
		return m.ScheduleExpression
	}
	return nil
}

// Deprecated: Do not use.
func (x *Schedule) GetCronExpression() string {
	if x, ok := x.GetScheduleExpression().(*Schedule_CronExpression); ok {
		return x.CronExpression
	}
	return ""
}

func (x *Schedule) GetRate() *FixedRate {
	if x, ok := x.GetScheduleExpression().(*Schedule_Rate); ok {
		return x.Rate
	}
	return nil
}

func (x *Schedule) GetCronSchedule() *CronSchedule {
	if x, ok := x.GetScheduleExpression().(*Schedule_CronSchedule); ok {
		return x.CronSchedule
	}
	return nil
}

func (x *Schedule) GetKickoffTimeInputArg() string {
	if x != nil {
		return x.KickoffTimeInputArg
	}
	return ""
}

type isSchedule_ScheduleExpression interface {
	isSchedule_ScheduleExpression()
}

type Schedule_CronExpression struct {
	// Uses AWS syntax: Minutes Hours Day-of-month Month Day-of-week Year
	// e.g. for a schedule that runs every 15 minutes: 0/15 * * * ? *
	//
	// Deprecated: Do not use.
	CronExpression string `protobuf:"bytes,1,opt,name=cron_expression,json=cronExpression,proto3,oneof"`
}

type Schedule_Rate struct {
	Rate *FixedRate `protobuf:"bytes,2,opt,name=rate,proto3,oneof"`
}

type Schedule_CronSchedule struct {
	CronSchedule *CronSchedule `protobuf:"bytes,4,opt,name=cron_schedule,json=cronSchedule,proto3,oneof"`
}

func (*Schedule_CronExpression) isSchedule_ScheduleExpression() {}

func (*Schedule_Rate) isSchedule_ScheduleExpression() {}

func (*Schedule_CronSchedule) isSchedule_ScheduleExpression() {}

var File_flyteidl_admin_schedule_proto protoreflect.FileDescriptor

var file_flyteidl_admin_schedule_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22,
	0x54, 0x0a, 0x09, 0x46, 0x69, 0x78, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x42, 0x0a, 0x0c, 0x43, 0x72, 0x6f, 0x6e, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xfa, 0x01, 0x0a, 0x08, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x0f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x72, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43,
	0x72, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x63,
	0x72, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6b,
	0x69, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x61, 0x72, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6b, 0x69, 0x63,
	0x6b, 0x6f, 0x66, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x72, 0x67,
	0x42, 0x14, 0x0a, 0x12, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x2e, 0x0a, 0x0d, 0x46, 0x69, 0x78, 0x65, 0x64, 0x52,
	0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x55, 0x54,
	0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x44, 0x41, 0x59, 0x10, 0x02, 0x42, 0xb2, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x0d, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x03, 0x46, 0x41, 0x58, 0xaa, 0x02, 0x0e, 0x46, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0e, 0x46, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x1a, 0x46,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x46, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_admin_schedule_proto_rawDescOnce sync.Once
	file_flyteidl_admin_schedule_proto_rawDescData = file_flyteidl_admin_schedule_proto_rawDesc
)

func file_flyteidl_admin_schedule_proto_rawDescGZIP() []byte {
	file_flyteidl_admin_schedule_proto_rawDescOnce.Do(func() {
		file_flyteidl_admin_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_admin_schedule_proto_rawDescData)
	})
	return file_flyteidl_admin_schedule_proto_rawDescData
}

var file_flyteidl_admin_schedule_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flyteidl_admin_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_flyteidl_admin_schedule_proto_goTypes = []interface{}{
	(FixedRateUnit)(0),   // 0: flyteidl.admin.FixedRateUnit
	(*FixedRate)(nil),    // 1: flyteidl.admin.FixedRate
	(*CronSchedule)(nil), // 2: flyteidl.admin.CronSchedule
	(*Schedule)(nil),     // 3: flyteidl.admin.Schedule
}
var file_flyteidl_admin_schedule_proto_depIdxs = []int32{
	0, // 0: flyteidl.admin.FixedRate.unit:type_name -> flyteidl.admin.FixedRateUnit
	1, // 1: flyteidl.admin.Schedule.rate:type_name -> flyteidl.admin.FixedRate
	2, // 2: flyteidl.admin.Schedule.cron_schedule:type_name -> flyteidl.admin.CronSchedule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_flyteidl_admin_schedule_proto_init() }
func file_flyteidl_admin_schedule_proto_init() {
	if File_flyteidl_admin_schedule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_admin_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FixedRate); i {
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
		file_flyteidl_admin_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronSchedule); i {
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
		file_flyteidl_admin_schedule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schedule); i {
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
	file_flyteidl_admin_schedule_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Schedule_CronExpression)(nil),
		(*Schedule_Rate)(nil),
		(*Schedule_CronSchedule)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_admin_schedule_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_admin_schedule_proto_goTypes,
		DependencyIndexes: file_flyteidl_admin_schedule_proto_depIdxs,
		EnumInfos:         file_flyteidl_admin_schedule_proto_enumTypes,
		MessageInfos:      file_flyteidl_admin_schedule_proto_msgTypes,
	}.Build()
	File_flyteidl_admin_schedule_proto = out.File
	file_flyteidl_admin_schedule_proto_rawDesc = nil
	file_flyteidl_admin_schedule_proto_goTypes = nil
	file_flyteidl_admin_schedule_proto_depIdxs = nil
}
