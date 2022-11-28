# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/launch_plan.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import execution_pb2 as flyteidl_dot_core_dot_execution__pb2
from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2
from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2
from flyteidl.core import interface_pb2 as flyteidl_dot_core_dot_interface__pb2
from flyteidl.core import security_pb2 as flyteidl_dot_core_dot_security__pb2
from flyteidl.admin import schedule_pb2 as flyteidl_dot_admin_dot_schedule__pb2
from flyteidl.admin import common_pb2 as flyteidl_dot_admin_dot_common__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/admin/launch_plan.proto',
  package='flyteidl.admin',
  syntax='proto3',
  serialized_options=_b('Z2github.com/flyteorg/flyte/gen/pb-go/flyteidl/admin'),
  serialized_pb=_b('\n flyteidl/admin/launch_plan.proto\x12\x0e\x66lyteidl.admin\x1a\x1d\x66lyteidl/core/execution.proto\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x1e\x66lyteidl/core/identifier.proto\x1a\x1d\x66lyteidl/core/interface.proto\x1a\x1c\x66lyteidl/core/security.proto\x1a\x1d\x66lyteidl/admin/schedule.proto\x1a\x1b\x66lyteidl/admin/common.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"n\n\x17LaunchPlanCreateRequest\x12%\n\x02id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12,\n\x04spec\x18\x02 \x01(\x0b\x32\x1e.flyteidl.admin.LaunchPlanSpec\"\x1a\n\x18LaunchPlanCreateResponse\"\x95\x01\n\nLaunchPlan\x12%\n\x02id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12,\n\x04spec\x18\x02 \x01(\x0b\x32\x1e.flyteidl.admin.LaunchPlanSpec\x12\x32\n\x07\x63losure\x18\x03 \x01(\x0b\x32!.flyteidl.admin.LaunchPlanClosure\"Q\n\x0eLaunchPlanList\x12\x30\n\x0claunch_plans\x18\x01 \x03(\x0b\x32\x1a.flyteidl.admin.LaunchPlan\x12\r\n\x05token\x18\x02 \x01(\t\"J\n\x04\x41uth\x12\x1a\n\x12\x61ssumable_iam_role\x18\x01 \x01(\t\x12\"\n\x1akubernetes_service_account\x18\x02 \x01(\t:\x02\x18\x01\"\xb0\x05\n\x0eLaunchPlanSpec\x12.\n\x0bworkflow_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12;\n\x0f\x65ntity_metadata\x18\x02 \x01(\x0b\x32\".flyteidl.admin.LaunchPlanMetadata\x12\x33\n\x0e\x64\x65\x66\x61ult_inputs\x18\x03 \x01(\x0b\x32\x1b.flyteidl.core.ParameterMap\x12/\n\x0c\x66ixed_inputs\x18\x04 \x01(\x0b\x32\x19.flyteidl.core.LiteralMap\x12\x10\n\x04role\x18\x05 \x01(\tB\x02\x18\x01\x12&\n\x06labels\x18\x06 \x01(\x0b\x32\x16.flyteidl.admin.Labels\x12\x30\n\x0b\x61nnotations\x18\x07 \x01(\x0b\x32\x1b.flyteidl.admin.Annotations\x12&\n\x04\x61uth\x18\x08 \x01(\x0b\x32\x14.flyteidl.admin.AuthB\x02\x18\x01\x12/\n\tauth_role\x18\t \x01(\x0b\x32\x18.flyteidl.admin.AuthRoleB\x02\x18\x01\x12\x38\n\x10security_context\x18\n \x01(\x0b\x32\x1e.flyteidl.core.SecurityContext\x12;\n\x12quality_of_service\x18\x10 \x01(\x0b\x32\x1f.flyteidl.core.QualityOfService\x12\x43\n\x16raw_output_data_config\x18\x11 \x01(\x0b\x32#.flyteidl.admin.RawOutputDataConfig\x12\x17\n\x0fmax_parallelism\x18\x12 \x01(\x05\x12\x31\n\rinterruptible\x18\x13 \x01(\x0b\x32\x1a.google.protobuf.BoolValue\"\x8f\x02\n\x11LaunchPlanClosure\x12.\n\x05state\x18\x01 \x01(\x0e\x32\x1f.flyteidl.admin.LaunchPlanState\x12\x34\n\x0f\x65xpected_inputs\x18\x02 \x01(\x0b\x32\x1b.flyteidl.core.ParameterMap\x12\x34\n\x10\x65xpected_outputs\x18\x03 \x01(\x0b\x32\x1a.flyteidl.core.VariableMap\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"u\n\x12LaunchPlanMetadata\x12*\n\x08schedule\x18\x01 \x01(\x0b\x32\x18.flyteidl.admin.Schedule\x12\x33\n\rnotifications\x18\x02 \x03(\x0b\x32\x1c.flyteidl.admin.Notification\"p\n\x17LaunchPlanUpdateRequest\x12%\n\x02id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12.\n\x05state\x18\x02 \x01(\x0e\x32\x1f.flyteidl.admin.LaunchPlanState\"\x1a\n\x18LaunchPlanUpdateResponse\"L\n\x17\x41\x63tiveLaunchPlanRequest\x12\x31\n\x02id\x18\x01 \x01(\x0b\x32%.flyteidl.admin.NamedEntityIdentifier\"\x83\x01\n\x1b\x41\x63tiveLaunchPlanListRequest\x12\x0f\n\x07project\x18\x01 \x01(\t\x12\x0e\n\x06\x64omain\x18\x02 \x01(\t\x12\r\n\x05limit\x18\x03 \x01(\r\x12\r\n\x05token\x18\x04 \x01(\t\x12%\n\x07sort_by\x18\x05 \x01(\x0b\x32\x14.flyteidl.admin.Sort*+\n\x0fLaunchPlanState\x12\x0c\n\x08INACTIVE\x10\x00\x12\n\n\x06\x41\x43TIVE\x10\x01\x42\x34Z2github.com/flyteorg/flyte/gen/pb-go/flyteidl/adminb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_execution__pb2.DESCRIPTOR,flyteidl_dot_core_dot_literals__pb2.DESCRIPTOR,flyteidl_dot_core_dot_identifier__pb2.DESCRIPTOR,flyteidl_dot_core_dot_interface__pb2.DESCRIPTOR,flyteidl_dot_core_dot_security__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_schedule__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_common__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_wrappers__pb2.DESCRIPTOR,])

_LAUNCHPLANSTATE = _descriptor.EnumDescriptor(
  name='LaunchPlanState',
  full_name='flyteidl.admin.LaunchPlanState',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='INACTIVE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ACTIVE', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2220,
  serialized_end=2263,
)
_sym_db.RegisterEnumDescriptor(_LAUNCHPLANSTATE)

LaunchPlanState = enum_type_wrapper.EnumTypeWrapper(_LAUNCHPLANSTATE)
INACTIVE = 0
ACTIVE = 1



_LAUNCHPLANCREATEREQUEST = _descriptor.Descriptor(
  name='LaunchPlanCreateRequest',
  full_name='flyteidl.admin.LaunchPlanCreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.LaunchPlanCreateRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spec', full_name='flyteidl.admin.LaunchPlanCreateRequest.spec', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=331,
  serialized_end=441,
)


_LAUNCHPLANCREATERESPONSE = _descriptor.Descriptor(
  name='LaunchPlanCreateResponse',
  full_name='flyteidl.admin.LaunchPlanCreateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=443,
  serialized_end=469,
)


_LAUNCHPLAN = _descriptor.Descriptor(
  name='LaunchPlan',
  full_name='flyteidl.admin.LaunchPlan',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.LaunchPlan.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spec', full_name='flyteidl.admin.LaunchPlan.spec', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='closure', full_name='flyteidl.admin.LaunchPlan.closure', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=472,
  serialized_end=621,
)


_LAUNCHPLANLIST = _descriptor.Descriptor(
  name='LaunchPlanList',
  full_name='flyteidl.admin.LaunchPlanList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='launch_plans', full_name='flyteidl.admin.LaunchPlanList.launch_plans', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='token', full_name='flyteidl.admin.LaunchPlanList.token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=623,
  serialized_end=704,
)


_AUTH = _descriptor.Descriptor(
  name='Auth',
  full_name='flyteidl.admin.Auth',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='assumable_iam_role', full_name='flyteidl.admin.Auth.assumable_iam_role', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='kubernetes_service_account', full_name='flyteidl.admin.Auth.kubernetes_service_account', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('\030\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=706,
  serialized_end=780,
)


_LAUNCHPLANSPEC = _descriptor.Descriptor(
  name='LaunchPlanSpec',
  full_name='flyteidl.admin.LaunchPlanSpec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='workflow_id', full_name='flyteidl.admin.LaunchPlanSpec.workflow_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='entity_metadata', full_name='flyteidl.admin.LaunchPlanSpec.entity_metadata', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_inputs', full_name='flyteidl.admin.LaunchPlanSpec.default_inputs', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fixed_inputs', full_name='flyteidl.admin.LaunchPlanSpec.fixed_inputs', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='role', full_name='flyteidl.admin.LaunchPlanSpec.role', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\030\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='labels', full_name='flyteidl.admin.LaunchPlanSpec.labels', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='annotations', full_name='flyteidl.admin.LaunchPlanSpec.annotations', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auth', full_name='flyteidl.admin.LaunchPlanSpec.auth', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\030\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auth_role', full_name='flyteidl.admin.LaunchPlanSpec.auth_role', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\030\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='security_context', full_name='flyteidl.admin.LaunchPlanSpec.security_context', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='quality_of_service', full_name='flyteidl.admin.LaunchPlanSpec.quality_of_service', index=10,
      number=16, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='raw_output_data_config', full_name='flyteidl.admin.LaunchPlanSpec.raw_output_data_config', index=11,
      number=17, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max_parallelism', full_name='flyteidl.admin.LaunchPlanSpec.max_parallelism', index=12,
      number=18, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='interruptible', full_name='flyteidl.admin.LaunchPlanSpec.interruptible', index=13,
      number=19, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=783,
  serialized_end=1471,
)


_LAUNCHPLANCLOSURE = _descriptor.Descriptor(
  name='LaunchPlanClosure',
  full_name='flyteidl.admin.LaunchPlanClosure',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='state', full_name='flyteidl.admin.LaunchPlanClosure.state', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='expected_inputs', full_name='flyteidl.admin.LaunchPlanClosure.expected_inputs', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='expected_outputs', full_name='flyteidl.admin.LaunchPlanClosure.expected_outputs', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='flyteidl.admin.LaunchPlanClosure.created_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='flyteidl.admin.LaunchPlanClosure.updated_at', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1474,
  serialized_end=1745,
)


_LAUNCHPLANMETADATA = _descriptor.Descriptor(
  name='LaunchPlanMetadata',
  full_name='flyteidl.admin.LaunchPlanMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='schedule', full_name='flyteidl.admin.LaunchPlanMetadata.schedule', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='notifications', full_name='flyteidl.admin.LaunchPlanMetadata.notifications', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1747,
  serialized_end=1864,
)


_LAUNCHPLANUPDATEREQUEST = _descriptor.Descriptor(
  name='LaunchPlanUpdateRequest',
  full_name='flyteidl.admin.LaunchPlanUpdateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.LaunchPlanUpdateRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='state', full_name='flyteidl.admin.LaunchPlanUpdateRequest.state', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1866,
  serialized_end=1978,
)


_LAUNCHPLANUPDATERESPONSE = _descriptor.Descriptor(
  name='LaunchPlanUpdateResponse',
  full_name='flyteidl.admin.LaunchPlanUpdateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1980,
  serialized_end=2006,
)


_ACTIVELAUNCHPLANREQUEST = _descriptor.Descriptor(
  name='ActiveLaunchPlanRequest',
  full_name='flyteidl.admin.ActiveLaunchPlanRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.ActiveLaunchPlanRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=2008,
  serialized_end=2084,
)


_ACTIVELAUNCHPLANLISTREQUEST = _descriptor.Descriptor(
  name='ActiveLaunchPlanListRequest',
  full_name='flyteidl.admin.ActiveLaunchPlanListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='flyteidl.admin.ActiveLaunchPlanListRequest.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='domain', full_name='flyteidl.admin.ActiveLaunchPlanListRequest.domain', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='flyteidl.admin.ActiveLaunchPlanListRequest.limit', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='token', full_name='flyteidl.admin.ActiveLaunchPlanListRequest.token', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sort_by', full_name='flyteidl.admin.ActiveLaunchPlanListRequest.sort_by', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=2087,
  serialized_end=2218,
)

_LAUNCHPLANCREATEREQUEST.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_LAUNCHPLANCREATEREQUEST.fields_by_name['spec'].message_type = _LAUNCHPLANSPEC
_LAUNCHPLAN.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_LAUNCHPLAN.fields_by_name['spec'].message_type = _LAUNCHPLANSPEC
_LAUNCHPLAN.fields_by_name['closure'].message_type = _LAUNCHPLANCLOSURE
_LAUNCHPLANLIST.fields_by_name['launch_plans'].message_type = _LAUNCHPLAN
_LAUNCHPLANSPEC.fields_by_name['workflow_id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_LAUNCHPLANSPEC.fields_by_name['entity_metadata'].message_type = _LAUNCHPLANMETADATA
_LAUNCHPLANSPEC.fields_by_name['default_inputs'].message_type = flyteidl_dot_core_dot_interface__pb2._PARAMETERMAP
_LAUNCHPLANSPEC.fields_by_name['fixed_inputs'].message_type = flyteidl_dot_core_dot_literals__pb2._LITERALMAP
_LAUNCHPLANSPEC.fields_by_name['labels'].message_type = flyteidl_dot_admin_dot_common__pb2._LABELS
_LAUNCHPLANSPEC.fields_by_name['annotations'].message_type = flyteidl_dot_admin_dot_common__pb2._ANNOTATIONS
_LAUNCHPLANSPEC.fields_by_name['auth'].message_type = _AUTH
_LAUNCHPLANSPEC.fields_by_name['auth_role'].message_type = flyteidl_dot_admin_dot_common__pb2._AUTHROLE
_LAUNCHPLANSPEC.fields_by_name['security_context'].message_type = flyteidl_dot_core_dot_security__pb2._SECURITYCONTEXT
_LAUNCHPLANSPEC.fields_by_name['quality_of_service'].message_type = flyteidl_dot_core_dot_execution__pb2._QUALITYOFSERVICE
_LAUNCHPLANSPEC.fields_by_name['raw_output_data_config'].message_type = flyteidl_dot_admin_dot_common__pb2._RAWOUTPUTDATACONFIG
_LAUNCHPLANSPEC.fields_by_name['interruptible'].message_type = google_dot_protobuf_dot_wrappers__pb2._BOOLVALUE
_LAUNCHPLANCLOSURE.fields_by_name['state'].enum_type = _LAUNCHPLANSTATE
_LAUNCHPLANCLOSURE.fields_by_name['expected_inputs'].message_type = flyteidl_dot_core_dot_interface__pb2._PARAMETERMAP
_LAUNCHPLANCLOSURE.fields_by_name['expected_outputs'].message_type = flyteidl_dot_core_dot_interface__pb2._VARIABLEMAP
_LAUNCHPLANCLOSURE.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_LAUNCHPLANCLOSURE.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_LAUNCHPLANMETADATA.fields_by_name['schedule'].message_type = flyteidl_dot_admin_dot_schedule__pb2._SCHEDULE
_LAUNCHPLANMETADATA.fields_by_name['notifications'].message_type = flyteidl_dot_admin_dot_common__pb2._NOTIFICATION
_LAUNCHPLANUPDATEREQUEST.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_LAUNCHPLANUPDATEREQUEST.fields_by_name['state'].enum_type = _LAUNCHPLANSTATE
_ACTIVELAUNCHPLANREQUEST.fields_by_name['id'].message_type = flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIER
_ACTIVELAUNCHPLANLISTREQUEST.fields_by_name['sort_by'].message_type = flyteidl_dot_admin_dot_common__pb2._SORT
DESCRIPTOR.message_types_by_name['LaunchPlanCreateRequest'] = _LAUNCHPLANCREATEREQUEST
DESCRIPTOR.message_types_by_name['LaunchPlanCreateResponse'] = _LAUNCHPLANCREATERESPONSE
DESCRIPTOR.message_types_by_name['LaunchPlan'] = _LAUNCHPLAN
DESCRIPTOR.message_types_by_name['LaunchPlanList'] = _LAUNCHPLANLIST
DESCRIPTOR.message_types_by_name['Auth'] = _AUTH
DESCRIPTOR.message_types_by_name['LaunchPlanSpec'] = _LAUNCHPLANSPEC
DESCRIPTOR.message_types_by_name['LaunchPlanClosure'] = _LAUNCHPLANCLOSURE
DESCRIPTOR.message_types_by_name['LaunchPlanMetadata'] = _LAUNCHPLANMETADATA
DESCRIPTOR.message_types_by_name['LaunchPlanUpdateRequest'] = _LAUNCHPLANUPDATEREQUEST
DESCRIPTOR.message_types_by_name['LaunchPlanUpdateResponse'] = _LAUNCHPLANUPDATERESPONSE
DESCRIPTOR.message_types_by_name['ActiveLaunchPlanRequest'] = _ACTIVELAUNCHPLANREQUEST
DESCRIPTOR.message_types_by_name['ActiveLaunchPlanListRequest'] = _ACTIVELAUNCHPLANLISTREQUEST
DESCRIPTOR.enum_types_by_name['LaunchPlanState'] = _LAUNCHPLANSTATE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

LaunchPlanCreateRequest = _reflection.GeneratedProtocolMessageType('LaunchPlanCreateRequest', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANCREATEREQUEST,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.LaunchPlanCreateRequest)
  ))
_sym_db.RegisterMessage(LaunchPlanCreateRequest)

LaunchPlanCreateResponse = _reflection.GeneratedProtocolMessageType('LaunchPlanCreateResponse', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANCREATERESPONSE,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.LaunchPlanCreateResponse)
  ))
_sym_db.RegisterMessage(LaunchPlanCreateResponse)

LaunchPlan = _reflection.GeneratedProtocolMessageType('LaunchPlan', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLAN,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.LaunchPlan)
  ))
_sym_db.RegisterMessage(LaunchPlan)

LaunchPlanList = _reflection.GeneratedProtocolMessageType('LaunchPlanList', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANLIST,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.LaunchPlanList)
  ))
_sym_db.RegisterMessage(LaunchPlanList)

Auth = _reflection.GeneratedProtocolMessageType('Auth', (_message.Message,), dict(
  DESCRIPTOR = _AUTH,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.Auth)
  ))
_sym_db.RegisterMessage(Auth)

LaunchPlanSpec = _reflection.GeneratedProtocolMessageType('LaunchPlanSpec', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANSPEC,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.LaunchPlanSpec)
  ))
_sym_db.RegisterMessage(LaunchPlanSpec)

LaunchPlanClosure = _reflection.GeneratedProtocolMessageType('LaunchPlanClosure', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANCLOSURE,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.LaunchPlanClosure)
  ))
_sym_db.RegisterMessage(LaunchPlanClosure)

LaunchPlanMetadata = _reflection.GeneratedProtocolMessageType('LaunchPlanMetadata', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANMETADATA,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.LaunchPlanMetadata)
  ))
_sym_db.RegisterMessage(LaunchPlanMetadata)

LaunchPlanUpdateRequest = _reflection.GeneratedProtocolMessageType('LaunchPlanUpdateRequest', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANUPDATEREQUEST,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.LaunchPlanUpdateRequest)
  ))
_sym_db.RegisterMessage(LaunchPlanUpdateRequest)

LaunchPlanUpdateResponse = _reflection.GeneratedProtocolMessageType('LaunchPlanUpdateResponse', (_message.Message,), dict(
  DESCRIPTOR = _LAUNCHPLANUPDATERESPONSE,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.LaunchPlanUpdateResponse)
  ))
_sym_db.RegisterMessage(LaunchPlanUpdateResponse)

ActiveLaunchPlanRequest = _reflection.GeneratedProtocolMessageType('ActiveLaunchPlanRequest', (_message.Message,), dict(
  DESCRIPTOR = _ACTIVELAUNCHPLANREQUEST,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ActiveLaunchPlanRequest)
  ))
_sym_db.RegisterMessage(ActiveLaunchPlanRequest)

ActiveLaunchPlanListRequest = _reflection.GeneratedProtocolMessageType('ActiveLaunchPlanListRequest', (_message.Message,), dict(
  DESCRIPTOR = _ACTIVELAUNCHPLANLISTREQUEST,
  __module__ = 'flyteidl.admin.launch_plan_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.ActiveLaunchPlanListRequest)
  ))
_sym_db.RegisterMessage(ActiveLaunchPlanListRequest)


DESCRIPTOR._options = None
_AUTH._options = None
_LAUNCHPLANSPEC.fields_by_name['role']._options = None
_LAUNCHPLANSPEC.fields_by_name['auth']._options = None
_LAUNCHPLANSPEC.fields_by_name['auth_role']._options = None
# @@protoc_insertion_point(module_scope)
