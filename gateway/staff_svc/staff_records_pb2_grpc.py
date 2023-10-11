# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from staff_svc import staff_records_pb2 as staff__svc_dot_staff__records__pb2


class StaffRecordsStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Create = channel.unary_unary(
                '/StaffRecords.StaffRecords/Create',
                request_serializer=staff__svc_dot_staff__records__pb2.CreateStaff.SerializeToString,
                response_deserializer=staff__svc_dot_staff__records__pb2.CreateResponse.FromString,
                )
        self.Get = channel.unary_unary(
                '/StaffRecords.StaffRecords/Get',
                request_serializer=staff__svc_dot_staff__records__pb2.GetStaffRecords.SerializeToString,
                response_deserializer=staff__svc_dot_staff__records__pb2.GetResponse.FromString,
                )
        self.GetAvailability = channel.unary_unary(
                '/StaffRecords.StaffRecords/GetAvailability',
                request_serializer=staff__svc_dot_staff__records__pb2.GetStaffAvailability.SerializeToString,
                response_deserializer=staff__svc_dot_staff__records__pb2.GetAvailabilityResponse.FromString,
                )
        self.Update = channel.unary_unary(
                '/StaffRecords.StaffRecords/Update',
                request_serializer=staff__svc_dot_staff__records__pb2.UpdateStaff.SerializeToString,
                response_deserializer=staff__svc_dot_staff__records__pb2.UpdateResponse.FromString,
                )
        self.Delete = channel.unary_unary(
                '/StaffRecords.StaffRecords/Delete',
                request_serializer=staff__svc_dot_staff__records__pb2.DeleteStaff.SerializeToString,
                response_deserializer=staff__svc_dot_staff__records__pb2.DeleteResponse.FromString,
                )
        self.Check = channel.unary_unary(
                '/StaffRecords.StaffRecords/Check',
                request_serializer=staff__svc_dot_staff__records__pb2.HealthCheckRequest.SerializeToString,
                response_deserializer=staff__svc_dot_staff__records__pb2.HealthCheckResponse.FromString,
                )


class StaffRecordsServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Create(self, request, context):
        """Service Methods
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAvailability(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Check(self, request, context):
        """Health Check Methods
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_StaffRecordsServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=staff__svc_dot_staff__records__pb2.CreateStaff.FromString,
                    response_serializer=staff__svc_dot_staff__records__pb2.CreateResponse.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=staff__svc_dot_staff__records__pb2.GetStaffRecords.FromString,
                    response_serializer=staff__svc_dot_staff__records__pb2.GetResponse.SerializeToString,
            ),
            'GetAvailability': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAvailability,
                    request_deserializer=staff__svc_dot_staff__records__pb2.GetStaffAvailability.FromString,
                    response_serializer=staff__svc_dot_staff__records__pb2.GetAvailabilityResponse.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=staff__svc_dot_staff__records__pb2.UpdateStaff.FromString,
                    response_serializer=staff__svc_dot_staff__records__pb2.UpdateResponse.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=staff__svc_dot_staff__records__pb2.DeleteStaff.FromString,
                    response_serializer=staff__svc_dot_staff__records__pb2.DeleteResponse.SerializeToString,
            ),
            'Check': grpc.unary_unary_rpc_method_handler(
                    servicer.Check,
                    request_deserializer=staff__svc_dot_staff__records__pb2.HealthCheckRequest.FromString,
                    response_serializer=staff__svc_dot_staff__records__pb2.HealthCheckResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'StaffRecords.StaffRecords', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class StaffRecords(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StaffRecords.StaffRecords/Create',
            staff__svc_dot_staff__records__pb2.CreateStaff.SerializeToString,
            staff__svc_dot_staff__records__pb2.CreateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StaffRecords.StaffRecords/Get',
            staff__svc_dot_staff__records__pb2.GetStaffRecords.SerializeToString,
            staff__svc_dot_staff__records__pb2.GetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetAvailability(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StaffRecords.StaffRecords/GetAvailability',
            staff__svc_dot_staff__records__pb2.GetStaffAvailability.SerializeToString,
            staff__svc_dot_staff__records__pb2.GetAvailabilityResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StaffRecords.StaffRecords/Update',
            staff__svc_dot_staff__records__pb2.UpdateStaff.SerializeToString,
            staff__svc_dot_staff__records__pb2.UpdateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StaffRecords.StaffRecords/Delete',
            staff__svc_dot_staff__records__pb2.DeleteStaff.SerializeToString,
            staff__svc_dot_staff__records__pb2.DeleteResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Check(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/StaffRecords.StaffRecords/Check',
            staff__svc_dot_staff__records__pb2.HealthCheckRequest.SerializeToString,
            staff__svc_dot_staff__records__pb2.HealthCheckResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)