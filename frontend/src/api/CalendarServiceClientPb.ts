/**
 * @fileoverview gRPC-Web generated client stub for api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as calendar_pb from './calendar_pb';


export class CalendarServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorCreateEvent = new grpcWeb.MethodDescriptor(
    '/api.CalendarService/CreateEvent',
    grpcWeb.MethodType.UNARY,
    calendar_pb.CreateEventRequest,
    calendar_pb.CreateEventReply,
    (request: calendar_pb.CreateEventRequest) => {
      return request.serializeBinary();
    },
    calendar_pb.CreateEventReply.deserializeBinary
  );

  createEvent(
    request: calendar_pb.CreateEventRequest,
    metadata: grpcWeb.Metadata | null): Promise<calendar_pb.CreateEventReply>;

  createEvent(
    request: calendar_pb.CreateEventRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: calendar_pb.CreateEventReply) => void): grpcWeb.ClientReadableStream<calendar_pb.CreateEventReply>;

  createEvent(
    request: calendar_pb.CreateEventRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: calendar_pb.CreateEventReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.CalendarService/CreateEvent',
        request,
        metadata || {},
        this.methodDescriptorCreateEvent,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.CalendarService/CreateEvent',
    request,
    metadata || {},
    this.methodDescriptorCreateEvent);
  }

  methodDescriptorUpdateEvent = new grpcWeb.MethodDescriptor(
    '/api.CalendarService/UpdateEvent',
    grpcWeb.MethodType.UNARY,
    calendar_pb.UpdateEventRequest,
    calendar_pb.UpdateEventReply,
    (request: calendar_pb.UpdateEventRequest) => {
      return request.serializeBinary();
    },
    calendar_pb.UpdateEventReply.deserializeBinary
  );

  updateEvent(
    request: calendar_pb.UpdateEventRequest,
    metadata: grpcWeb.Metadata | null): Promise<calendar_pb.UpdateEventReply>;

  updateEvent(
    request: calendar_pb.UpdateEventRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: calendar_pb.UpdateEventReply) => void): grpcWeb.ClientReadableStream<calendar_pb.UpdateEventReply>;

  updateEvent(
    request: calendar_pb.UpdateEventRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: calendar_pb.UpdateEventReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.CalendarService/UpdateEvent',
        request,
        metadata || {},
        this.methodDescriptorUpdateEvent,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.CalendarService/UpdateEvent',
    request,
    metadata || {},
    this.methodDescriptorUpdateEvent);
  }

  methodDescriptorDeleteEvent = new grpcWeb.MethodDescriptor(
    '/api.CalendarService/DeleteEvent',
    grpcWeb.MethodType.UNARY,
    calendar_pb.DeleteEventRequest,
    calendar_pb.DeleteEventReply,
    (request: calendar_pb.DeleteEventRequest) => {
      return request.serializeBinary();
    },
    calendar_pb.DeleteEventReply.deserializeBinary
  );

  deleteEvent(
    request: calendar_pb.DeleteEventRequest,
    metadata: grpcWeb.Metadata | null): Promise<calendar_pb.DeleteEventReply>;

  deleteEvent(
    request: calendar_pb.DeleteEventRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: calendar_pb.DeleteEventReply) => void): grpcWeb.ClientReadableStream<calendar_pb.DeleteEventReply>;

  deleteEvent(
    request: calendar_pb.DeleteEventRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: calendar_pb.DeleteEventReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.CalendarService/DeleteEvent',
        request,
        metadata || {},
        this.methodDescriptorDeleteEvent,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.CalendarService/DeleteEvent',
    request,
    metadata || {},
    this.methodDescriptorDeleteEvent);
  }

  methodDescriptorListEvents = new grpcWeb.MethodDescriptor(
    '/api.CalendarService/ListEvents',
    grpcWeb.MethodType.UNARY,
    calendar_pb.ListEventsRequest,
    calendar_pb.ListEventsReply,
    (request: calendar_pb.ListEventsRequest) => {
      return request.serializeBinary();
    },
    calendar_pb.ListEventsReply.deserializeBinary
  );

  listEvents(
    request: calendar_pb.ListEventsRequest,
    metadata: grpcWeb.Metadata | null): Promise<calendar_pb.ListEventsReply>;

  listEvents(
    request: calendar_pb.ListEventsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: calendar_pb.ListEventsReply) => void): grpcWeb.ClientReadableStream<calendar_pb.ListEventsReply>;

  listEvents(
    request: calendar_pb.ListEventsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: calendar_pb.ListEventsReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.CalendarService/ListEvents',
        request,
        metadata || {},
        this.methodDescriptorListEvents,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.CalendarService/ListEvents',
    request,
    metadata || {},
    this.methodDescriptorListEvents);
  }

  methodDescriptorGetRemind = new grpcWeb.MethodDescriptor(
    '/api.CalendarService/GetRemind',
    grpcWeb.MethodType.UNARY,
    calendar_pb.GetRemindRequest,
    calendar_pb.GetRemindReply,
    (request: calendar_pb.GetRemindRequest) => {
      return request.serializeBinary();
    },
    calendar_pb.GetRemindReply.deserializeBinary
  );

  getRemind(
    request: calendar_pb.GetRemindRequest,
    metadata: grpcWeb.Metadata | null): Promise<calendar_pb.GetRemindReply>;

  getRemind(
    request: calendar_pb.GetRemindRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: calendar_pb.GetRemindReply) => void): grpcWeb.ClientReadableStream<calendar_pb.GetRemindReply>;

  getRemind(
    request: calendar_pb.GetRemindRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: calendar_pb.GetRemindReply) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/api.CalendarService/GetRemind',
        request,
        metadata || {},
        this.methodDescriptorGetRemind,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/api.CalendarService/GetRemind',
    request,
    metadata || {},
    this.methodDescriptorGetRemind);
  }

}

