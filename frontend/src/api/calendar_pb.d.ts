import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class RangeTime extends jspb.Message {
  getStart(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setStart(value?: google_protobuf_timestamp_pb.Timestamp): RangeTime;
  hasStart(): boolean;
  clearStart(): RangeTime;

  getEnd(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setEnd(value?: google_protobuf_timestamp_pb.Timestamp): RangeTime;
  hasEnd(): boolean;
  clearEnd(): RangeTime;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RangeTime.AsObject;
  static toObject(includeInstance: boolean, msg: RangeTime): RangeTime.AsObject;
  static serializeBinaryToWriter(message: RangeTime, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RangeTime;
  static deserializeBinaryFromReader(message: RangeTime, reader: jspb.BinaryReader): RangeTime;
}

export namespace RangeTime {
  export type AsObject = {
    start?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    end?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Event extends jspb.Message {
  getId(): number;
  setId(value: number): Event;

  getTitle(): string;
  setTitle(value: string): Event;

  getRangeTime(): RangeTime | undefined;
  setRangeTime(value?: RangeTime): Event;
  hasRangeTime(): boolean;
  clearRangeTime(): Event;

  getLocation(): Location | undefined;
  setLocation(value?: Location): Event;
  hasLocation(): boolean;
  clearLocation(): Event;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Event.AsObject;
  static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
  static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Event;
  static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
}

export namespace Event {
  export type AsObject = {
    id: number,
    title: string,
    rangeTime?: RangeTime.AsObject,
    location?: Location.AsObject,
  }
}

export class ListEvents extends jspb.Message {
  getEventList(): Array<Event>;
  setEventList(value: Array<Event>): ListEvents;
  clearEventList(): ListEvents;
  addEvent(value?: Event, index?: number): Event;

  getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTime(value?: google_protobuf_timestamp_pb.Timestamp): ListEvents;
  hasTime(): boolean;
  clearTime(): ListEvents;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListEvents.AsObject;
  static toObject(includeInstance: boolean, msg: ListEvents): ListEvents.AsObject;
  static serializeBinaryToWriter(message: ListEvents, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListEvents;
  static deserializeBinaryFromReader(message: ListEvents, reader: jspb.BinaryReader): ListEvents;
}

export namespace ListEvents {
  export type AsObject = {
    eventList: Array<Event.AsObject>,
    time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Location extends jspb.Message {
  getX(): number;
  setX(value: number): Location;

  getY(): number;
  setY(value: number): Location;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Location.AsObject;
  static toObject(includeInstance: boolean, msg: Location): Location.AsObject;
  static serializeBinaryToWriter(message: Location, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Location;
  static deserializeBinaryFromReader(message: Location, reader: jspb.BinaryReader): Location;
}

export namespace Location {
  export type AsObject = {
    x: number,
    y: number,
  }
}

export class CreateEventRequest extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): CreateEventRequest;

  getRangeTime(): RangeTime | undefined;
  setRangeTime(value?: RangeTime): CreateEventRequest;
  hasRangeTime(): boolean;
  clearRangeTime(): CreateEventRequest;

  getLocation(): Location | undefined;
  setLocation(value?: Location): CreateEventRequest;
  hasLocation(): boolean;
  clearLocation(): CreateEventRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateEventRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateEventRequest): CreateEventRequest.AsObject;
  static serializeBinaryToWriter(message: CreateEventRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateEventRequest;
  static deserializeBinaryFromReader(message: CreateEventRequest, reader: jspb.BinaryReader): CreateEventRequest;
}

export namespace CreateEventRequest {
  export type AsObject = {
    title: string,
    rangeTime?: RangeTime.AsObject,
    location?: Location.AsObject,
  }
}

export class CreateEventReply extends jspb.Message {
  getId(): number;
  setId(value: number): CreateEventReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateEventReply.AsObject;
  static toObject(includeInstance: boolean, msg: CreateEventReply): CreateEventReply.AsObject;
  static serializeBinaryToWriter(message: CreateEventReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateEventReply;
  static deserializeBinaryFromReader(message: CreateEventReply, reader: jspb.BinaryReader): CreateEventReply;
}

export namespace CreateEventReply {
  export type AsObject = {
    id: number,
  }
}

export class ListEventsRequest extends jspb.Message {
  getRangeTime(): RangeTime | undefined;
  setRangeTime(value?: RangeTime): ListEventsRequest;
  hasRangeTime(): boolean;
  clearRangeTime(): ListEventsRequest;

  getDateType(): ListEventsRequest.DATE_TYPE;
  setDateType(value: ListEventsRequest.DATE_TYPE): ListEventsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListEventsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListEventsRequest): ListEventsRequest.AsObject;
  static serializeBinaryToWriter(message: ListEventsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListEventsRequest;
  static deserializeBinaryFromReader(message: ListEventsRequest, reader: jspb.BinaryReader): ListEventsRequest;
}

export namespace ListEventsRequest {
  export type AsObject = {
    rangeTime?: RangeTime.AsObject,
    dateType: ListEventsRequest.DATE_TYPE,
  }

  export enum DATE_TYPE { 
    DAY = 0,
    MONTH = 1,
    YEAR = 2,
  }
}

export class ListEventsReply extends jspb.Message {
  getListEventsList(): Array<ListEvents>;
  setListEventsList(value: Array<ListEvents>): ListEventsReply;
  clearListEventsList(): ListEventsReply;
  addListEvents(value?: ListEvents, index?: number): ListEvents;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListEventsReply.AsObject;
  static toObject(includeInstance: boolean, msg: ListEventsReply): ListEventsReply.AsObject;
  static serializeBinaryToWriter(message: ListEventsReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListEventsReply;
  static deserializeBinaryFromReader(message: ListEventsReply, reader: jspb.BinaryReader): ListEventsReply;
}

export namespace ListEventsReply {
  export type AsObject = {
    listEventsList: Array<ListEvents.AsObject>,
  }
}

export class GetRemindRequest extends jspb.Message {
  getLocation(): Location | undefined;
  setLocation(value?: Location): GetRemindRequest;
  hasLocation(): boolean;
  clearLocation(): GetRemindRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRemindRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRemindRequest): GetRemindRequest.AsObject;
  static serializeBinaryToWriter(message: GetRemindRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRemindRequest;
  static deserializeBinaryFromReader(message: GetRemindRequest, reader: jspb.BinaryReader): GetRemindRequest;
}

export namespace GetRemindRequest {
  export type AsObject = {
    location?: Location.AsObject,
  }
}

export class GetRemindReply extends jspb.Message {
  getDuration(): number;
  setDuration(value: number): GetRemindReply;

  getOkay(): boolean;
  setOkay(value: boolean): GetRemindReply;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRemindReply.AsObject;
  static toObject(includeInstance: boolean, msg: GetRemindReply): GetRemindReply.AsObject;
  static serializeBinaryToWriter(message: GetRemindReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRemindReply;
  static deserializeBinaryFromReader(message: GetRemindReply, reader: jspb.BinaryReader): GetRemindReply;
}

export namespace GetRemindReply {
  export type AsObject = {
    duration: number,
    okay: boolean,
  }
}

export class UpdateEventRequest extends jspb.Message {
  getEvent(): Event | undefined;
  setEvent(value?: Event): UpdateEventRequest;
  hasEvent(): boolean;
  clearEvent(): UpdateEventRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateEventRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateEventRequest): UpdateEventRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateEventRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateEventRequest;
  static deserializeBinaryFromReader(message: UpdateEventRequest, reader: jspb.BinaryReader): UpdateEventRequest;
}

export namespace UpdateEventRequest {
  export type AsObject = {
    event?: Event.AsObject,
  }
}

export class UpdateEventReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateEventReply.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateEventReply): UpdateEventReply.AsObject;
  static serializeBinaryToWriter(message: UpdateEventReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateEventReply;
  static deserializeBinaryFromReader(message: UpdateEventReply, reader: jspb.BinaryReader): UpdateEventReply;
}

export namespace UpdateEventReply {
  export type AsObject = {
  }
}

export class DeleteEventRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteEventRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteEventRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteEventRequest): DeleteEventRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteEventRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteEventRequest;
  static deserializeBinaryFromReader(message: DeleteEventRequest, reader: jspb.BinaryReader): DeleteEventRequest;
}

export namespace DeleteEventRequest {
  export type AsObject = {
    id: number,
  }
}

export class DeleteEventReply extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteEventReply.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteEventReply): DeleteEventReply.AsObject;
  static serializeBinaryToWriter(message: DeleteEventReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteEventReply;
  static deserializeBinaryFromReader(message: DeleteEventReply, reader: jspb.BinaryReader): DeleteEventReply;
}

export namespace DeleteEventReply {
  export type AsObject = {
  }
}

