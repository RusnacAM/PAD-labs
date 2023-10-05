# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: scheduler_svc/scheduler.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dscheduler_svc/scheduler.proto\x12\tscheduler\"n\n\x0b\x41ppointment\x12\x0e\n\x06\x61pptID\x18\x01 \x01(\t\x12\x13\n\x0bpatientName\x18\x02 \x01(\t\x12\x12\n\ndoctorName\x18\x03 \x01(\t\x12\x14\n\x0c\x61pptDateTime\x18\x04 \x01(\t\x12\x10\n\x08\x61pptType\x18\x05 \x01(\t\"@\n\x11\x43reateAppointment\x12+\n\x0b\x61ppointment\x18\x01 \x01(\x0b\x32\x16.scheduler.Appointment\"?\n\x0e\x43reateResponse\x12\x0e\n\x06\x61pptID\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\x04\x12\r\n\x05\x65rror\x18\x03 \x01(\t\"\x11\n\x0fGetAppointments\"Z\n\x0bGetResponse\x12,\n\x0c\x61ppointments\x18\x01 \x03(\x0b\x32\x16.scheduler.Appointment\x12\x0e\n\x06status\x18\x02 \x01(\x04\x12\r\n\x05\x65rror\x18\x03 \x01(\t\"@\n\x11UpdateAppointment\x12+\n\x0b\x61ppointment\x18\x01 \x01(\x0b\x32\x16.scheduler.Appointment\"\\\n\x0eUpdateResponse\x12+\n\x0b\x61ppointment\x18\x01 \x01(\x0b\x32\x16.scheduler.Appointment\x12\x0e\n\x06status\x18\x02 \x01(\x04\x12\r\n\x05\x65rror\x18\x03 \x01(\t\"#\n\x11\x44\x65leteAppointment\x12\x0e\n\x06\x61pptID\x18\x01 \x01(\t\"?\n\x0e\x44\x65leteResponse\x12\x0e\n\x06\x61pptID\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\x04\x12\r\n\x05\x65rror\x18\x03 \x01(\t\"%\n\x12HealthCheckRequest\x12\x0f\n\x07service\x18\x01 \x01(\t\"\x8f\x01\n\x13HealthCheckResponse\x12<\n\x06status\x18\x01 \x01(\x0e\x32,.scheduler.HealthCheckResponse.ServingStatus\":\n\rServingStatus\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SERVING\x10\x01\x12\x0f\n\x0bNOT_SERVING\x10\x02\x32\xe7\x02\n\tScheduler\x12\x45\n\nCreateAppt\x12\x1c.scheduler.CreateAppointment\x1a\x19.scheduler.CreateResponse\x12=\n\x07GetAppt\x12\x1a.scheduler.GetAppointments\x1a\x16.scheduler.GetResponse\x12\x45\n\nUpdateAppt\x12\x1c.scheduler.UpdateAppointment\x1a\x19.scheduler.UpdateResponse\x12\x45\n\nDeleteAppt\x12\x1c.scheduler.DeleteAppointment\x1a\x19.scheduler.DeleteResponse\x12\x46\n\x05\x43heck\x12\x1d.scheduler.HealthCheckRequest\x1a\x1e.scheduler.HealthCheckResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'scheduler_svc.scheduler_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_APPOINTMENT']._serialized_start=44
  _globals['_APPOINTMENT']._serialized_end=154
  _globals['_CREATEAPPOINTMENT']._serialized_start=156
  _globals['_CREATEAPPOINTMENT']._serialized_end=220
  _globals['_CREATERESPONSE']._serialized_start=222
  _globals['_CREATERESPONSE']._serialized_end=285
  _globals['_GETAPPOINTMENTS']._serialized_start=287
  _globals['_GETAPPOINTMENTS']._serialized_end=304
  _globals['_GETRESPONSE']._serialized_start=306
  _globals['_GETRESPONSE']._serialized_end=396
  _globals['_UPDATEAPPOINTMENT']._serialized_start=398
  _globals['_UPDATEAPPOINTMENT']._serialized_end=462
  _globals['_UPDATERESPONSE']._serialized_start=464
  _globals['_UPDATERESPONSE']._serialized_end=556
  _globals['_DELETEAPPOINTMENT']._serialized_start=558
  _globals['_DELETEAPPOINTMENT']._serialized_end=593
  _globals['_DELETERESPONSE']._serialized_start=595
  _globals['_DELETERESPONSE']._serialized_end=658
  _globals['_HEALTHCHECKREQUEST']._serialized_start=660
  _globals['_HEALTHCHECKREQUEST']._serialized_end=697
  _globals['_HEALTHCHECKRESPONSE']._serialized_start=700
  _globals['_HEALTHCHECKRESPONSE']._serialized_end=843
  _globals['_HEALTHCHECKRESPONSE_SERVINGSTATUS']._serialized_start=785
  _globals['_HEALTHCHECKRESPONSE_SERVINGSTATUS']._serialized_end=843
  _globals['_SCHEDULER']._serialized_start=846
  _globals['_SCHEDULER']._serialized_end=1205
# @@protoc_insertion_point(module_scope)
