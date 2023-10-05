# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: staff_svc/staff_records.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dstaff_svc/staff_records.proto\x12\x0cStaffRecords\"g\n\x0bStaffRecord\x12\x0f\n\x07staffID\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x10\n\x08jobTitle\x18\x03 \x01(\t\x12\x12\n\ndepartment\x18\x04 \x01(\t\x12\x13\n\x0bisAvailable\x18\x05 \x01(\x08\"7\n\x0b\x43reateStaff\x12(\n\x05staff\x18\x01 \x01(\x0b\x32\x19.StaffRecords.StaffRecord\"A\n\x0e\x43reateResponse\x12\x0f\n\x07staffID\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\r\n\x05\x65rror\x18\x03 \x01(\t\"\x11\n\x0fGetStaffRecords\"M\n\x0bGetResponse\x12/\n\x0cstaffRecords\x18\x01 \x03(\x0b\x32\x19.StaffRecords.StaffRecord\x12\r\n\x05\x65rror\x18\x02 \x01(\t\"\'\n\x14GetStaffAvailability\x12\x0f\n\x07staffID\x18\x01 \x01(\t\"N\n\x17GetAvailabilityResponse\x12\x0f\n\x07staffID\x18\x01 \x01(\t\x12\x13\n\x0bisAvailable\x18\x02 \x01(\x08\x12\r\n\x05\x65rror\x18\x03 \x01(\t\"=\n\x0bUpdateStaff\x12.\n\x0bstaffRecord\x18\x01 \x01(\x0b\x32\x19.StaffRecords.StaffRecord\"A\n\x0eUpdateResponse\x12\x0f\n\x07staffID\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\r\n\x05\x65rror\x18\x03 \x01(\t\"\x1e\n\x0b\x44\x65leteStaff\x12\x0f\n\x07staffID\x18\x01 \x01(\t\"A\n\x0e\x44\x65leteResponse\x12\x0f\n\x07staffID\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\r\n\x05\x65rror\x18\x03 \x01(\t\"%\n\x12HealthCheckRequest\x12\x0f\n\x07service\x18\x01 \x01(\t\"\x92\x01\n\x13HealthCheckResponse\x12?\n\x06status\x18\x01 \x01(\x0e\x32/.StaffRecords.HealthCheckResponse.ServingStatus\":\n\rServingStatus\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SERVING\x10\x01\x12\x0f\n\x0bNOT_SERVING\x10\x02\x32\xc4\x03\n\x0cStaffRecords\x12\x41\n\x06\x43reate\x12\x19.StaffRecords.CreateStaff\x1a\x1c.StaffRecords.CreateResponse\x12?\n\x03Get\x12\x1d.StaffRecords.GetStaffRecords\x1a\x19.StaffRecords.GetResponse\x12\\\n\x0fGetAvailability\x12\".StaffRecords.GetStaffAvailability\x1a%.StaffRecords.GetAvailabilityResponse\x12\x41\n\x06Update\x12\x19.StaffRecords.UpdateStaff\x1a\x1c.StaffRecords.UpdateResponse\x12\x41\n\x06\x44\x65lete\x12\x19.StaffRecords.DeleteStaff\x1a\x1c.StaffRecords.DeleteResponse\x12L\n\x05\x43heck\x12 .StaffRecords.HealthCheckRequest\x1a!.StaffRecords.HealthCheckResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'staff_svc.staff_records_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_STAFFRECORD']._serialized_start=47
  _globals['_STAFFRECORD']._serialized_end=150
  _globals['_CREATESTAFF']._serialized_start=152
  _globals['_CREATESTAFF']._serialized_end=207
  _globals['_CREATERESPONSE']._serialized_start=209
  _globals['_CREATERESPONSE']._serialized_end=274
  _globals['_GETSTAFFRECORDS']._serialized_start=276
  _globals['_GETSTAFFRECORDS']._serialized_end=293
  _globals['_GETRESPONSE']._serialized_start=295
  _globals['_GETRESPONSE']._serialized_end=372
  _globals['_GETSTAFFAVAILABILITY']._serialized_start=374
  _globals['_GETSTAFFAVAILABILITY']._serialized_end=413
  _globals['_GETAVAILABILITYRESPONSE']._serialized_start=415
  _globals['_GETAVAILABILITYRESPONSE']._serialized_end=493
  _globals['_UPDATESTAFF']._serialized_start=495
  _globals['_UPDATESTAFF']._serialized_end=556
  _globals['_UPDATERESPONSE']._serialized_start=558
  _globals['_UPDATERESPONSE']._serialized_end=623
  _globals['_DELETESTAFF']._serialized_start=625
  _globals['_DELETESTAFF']._serialized_end=655
  _globals['_DELETERESPONSE']._serialized_start=657
  _globals['_DELETERESPONSE']._serialized_end=722
  _globals['_HEALTHCHECKREQUEST']._serialized_start=724
  _globals['_HEALTHCHECKREQUEST']._serialized_end=761
  _globals['_HEALTHCHECKRESPONSE']._serialized_start=764
  _globals['_HEALTHCHECKRESPONSE']._serialized_end=910
  _globals['_HEALTHCHECKRESPONSE_SERVINGSTATUS']._serialized_start=852
  _globals['_HEALTHCHECKRESPONSE_SERVINGSTATUS']._serialized_end=910
  _globals['_STAFFRECORDS']._serialized_start=913
  _globals['_STAFFRECORDS']._serialized_end=1365
# @@protoc_insertion_point(module_scope)
