# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/plugins/pytorch.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/plugins/pytorch.proto',
  package='flyteidl.plugins',
  syntax='proto3',
  serialized_options=_b('Z4github.com/flyteorg/flyte/gen/pb-go/flyteidl/plugins'),
  serialized_pb=_b('\n\x1e\x66lyteidl/plugins/pytorch.proto\x12\x10\x66lyteidl.plugins\"1\n\x1e\x44istributedPyTorchTrainingTask\x12\x0f\n\x07workers\x18\x01 \x01(\x05\x42\x36Z4github.com/flyteorg/flyte/gen/pb-go/flyteidl/pluginsb\x06proto3')
)




_DISTRIBUTEDPYTORCHTRAININGTASK = _descriptor.Descriptor(
  name='DistributedPyTorchTrainingTask',
  full_name='flyteidl.plugins.DistributedPyTorchTrainingTask',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='workers', full_name='flyteidl.plugins.DistributedPyTorchTrainingTask.workers', index=0,
      number=1, type=5, cpp_type=1, label=1,
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
  serialized_start=52,
  serialized_end=101,
)

DESCRIPTOR.message_types_by_name['DistributedPyTorchTrainingTask'] = _DISTRIBUTEDPYTORCHTRAININGTASK
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DistributedPyTorchTrainingTask = _reflection.GeneratedProtocolMessageType('DistributedPyTorchTrainingTask', (_message.Message,), dict(
  DESCRIPTOR = _DISTRIBUTEDPYTORCHTRAININGTASK,
  __module__ = 'flyteidl.plugins.pytorch_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.plugins.DistributedPyTorchTrainingTask)
  ))
_sym_db.RegisterMessage(DistributedPyTorchTrainingTask)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
