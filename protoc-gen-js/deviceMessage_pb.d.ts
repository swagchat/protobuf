// package: swagchat.protobuf
// file: deviceMessage.proto

import * as jspb from "google-protobuf";
import * as gogoproto_gogo_pb from "./gogoproto/gogo_pb";

export class Device extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): void;

  getPlatform(): number;
  setPlatform(value: number): void;

  getToken(): string;
  setToken(value: string): void;

  getNotificationDeviceId(): string;
  setNotificationDeviceId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Device.AsObject;
  static toObject(includeInstance: boolean, msg: Device): Device.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Device, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Device;
  static deserializeBinaryFromReader(message: Device, reader: jspb.BinaryReader): Device;
}

export namespace Device {
  export type AsObject = {
    userId: string,
    platform: number,
    token: string,
    notificationDeviceId: string,
  }
}

