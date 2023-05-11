/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "hub.hub";

export interface ChainInfo {
  chainKey: string;
  chainAdministratorKey: string;
  startFlag: boolean;
}

const baseChainInfo: object = {
  chainKey: "",
  chainAdministratorKey: "",
  startFlag: false,
};

export const ChainInfo = {
  encode(message: ChainInfo, writer: Writer = Writer.create()): Writer {
    if (message.chainKey !== "") {
      writer.uint32(10).string(message.chainKey);
    }
    if (message.chainAdministratorKey !== "") {
      writer.uint32(18).string(message.chainAdministratorKey);
    }
    if (message.startFlag === true) {
      writer.uint32(24).bool(message.startFlag);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ChainInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseChainInfo } as ChainInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.chainKey = reader.string();
          break;
        case 2:
          message.chainAdministratorKey = reader.string();
          break;
        case 3:
          message.startFlag = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChainInfo {
    const message = { ...baseChainInfo } as ChainInfo;
    if (object.chainKey !== undefined && object.chainKey !== null) {
      message.chainKey = String(object.chainKey);
    } else {
      message.chainKey = "";
    }
    if (
      object.chainAdministratorKey !== undefined &&
      object.chainAdministratorKey !== null
    ) {
      message.chainAdministratorKey = String(object.chainAdministratorKey);
    } else {
      message.chainAdministratorKey = "";
    }
    if (object.startFlag !== undefined && object.startFlag !== null) {
      message.startFlag = Boolean(object.startFlag);
    } else {
      message.startFlag = false;
    }
    return message;
  },

  toJSON(message: ChainInfo): unknown {
    const obj: any = {};
    message.chainKey !== undefined && (obj.chainKey = message.chainKey);
    message.chainAdministratorKey !== undefined &&
      (obj.chainAdministratorKey = message.chainAdministratorKey);
    message.startFlag !== undefined && (obj.startFlag = message.startFlag);
    return obj;
  },

  fromPartial(object: DeepPartial<ChainInfo>): ChainInfo {
    const message = { ...baseChainInfo } as ChainInfo;
    if (object.chainKey !== undefined && object.chainKey !== null) {
      message.chainKey = object.chainKey;
    } else {
      message.chainKey = "";
    }
    if (
      object.chainAdministratorKey !== undefined &&
      object.chainAdministratorKey !== null
    ) {
      message.chainAdministratorKey = object.chainAdministratorKey;
    } else {
      message.chainAdministratorKey = "";
    }
    if (object.startFlag !== undefined && object.startFlag !== null) {
      message.startFlag = object.startFlag;
    } else {
      message.startFlag = false;
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
