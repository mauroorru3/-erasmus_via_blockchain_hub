/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

export interface Universities {
  universityName: string;
  universitiesKey: string;
}

const baseUniversities: object = { universityName: "", universitiesKey: "" };

export const Universities = {
  encode(message: Universities, writer: Writer = Writer.create()): Writer {
    if (message.universityName !== "") {
      writer.uint32(10).string(message.universityName);
    }
    if (message.universitiesKey !== "") {
      writer.uint32(18).string(message.universitiesKey);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Universities {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUniversities } as Universities;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.universityName = reader.string();
          break;
        case 2:
          message.universitiesKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Universities {
    const message = { ...baseUniversities } as Universities;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = String(object.universityName);
    } else {
      message.universityName = "";
    }
    if (
      object.universitiesKey !== undefined &&
      object.universitiesKey !== null
    ) {
      message.universitiesKey = String(object.universitiesKey);
    } else {
      message.universitiesKey = "";
    }
    return message;
  },

  toJSON(message: Universities): unknown {
    const obj: any = {};
    message.universityName !== undefined &&
      (obj.universityName = message.universityName);
    message.universitiesKey !== undefined &&
      (obj.universitiesKey = message.universitiesKey);
    return obj;
  },

  fromPartial(object: DeepPartial<Universities>): Universities {
    const message = { ...baseUniversities } as Universities;
    if (object.universityName !== undefined && object.universityName !== null) {
      message.universityName = object.universityName;
    } else {
      message.universityName = "";
    }
    if (
      object.universitiesKey !== undefined &&
      object.universitiesKey !== null
    ) {
      message.universitiesKey = object.universitiesKey;
    } else {
      message.universitiesKey = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
