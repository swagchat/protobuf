// package: swagchat.protobuf
// file: userMessage.proto

import * as jspb from "google-protobuf";
import * as gogoproto_gogo_pb from "./gogoproto/gogo_pb";
import * as deviceMessage_pb from "./deviceMessage_pb";
import * as roomMessage_pb from "./roomMessage_pb";
import * as commonMessage_pb from "./commonMessage_pb";

export class User extends jspb.Message {
  hasId(): boolean;
  clearId(): void;
  getId(): number | undefined;
  setId(value: number): void;

  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasPictureUrl(): boolean;
  clearPictureUrl(): void;
  getPictureUrl(): string | undefined;
  setPictureUrl(value: string): void;

  hasInformationUrl(): boolean;
  clearInformationUrl(): void;
  getInformationUrl(): string | undefined;
  setInformationUrl(value: string): void;

  hasUnreadCount(): boolean;
  clearUnreadCount(): void;
  getUnreadCount(): number | undefined;
  setUnreadCount(value: number): void;

  hasMetaData(): boolean;
  clearMetaData(): void;
  getMetaData(): Uint8Array | string;
  getMetaData_asU8(): Uint8Array;
  getMetaData_asB64(): string;
  setMetaData(value: Uint8Array | string): void;

  hasPublic(): boolean;
  clearPublic(): void;
  getPublic(): boolean | undefined;
  setPublic(value: boolean): void;

  hasCanBlock(): boolean;
  clearCanBlock(): void;
  getCanBlock(): boolean | undefined;
  setCanBlock(value: boolean): void;

  hasLang(): boolean;
  clearLang(): void;
  getLang(): string | undefined;
  setLang(value: string): void;

  hasAccessToken(): boolean;
  clearAccessToken(): void;
  getAccessToken(): string | undefined;
  setAccessToken(value: string): void;

  hasLastAccessRoomId(): boolean;
  clearLastAccessRoomId(): void;
  getLastAccessRoomId(): string | undefined;
  setLastAccessRoomId(value: string): void;

  hasLastAccessed(): boolean;
  clearLastAccessed(): void;
  getLastAccessed(): number | undefined;
  setLastAccessed(value: number): void;

  hasCreated(): boolean;
  clearCreated(): void;
  getCreated(): number | undefined;
  setCreated(value: number): void;

  hasModified(): boolean;
  clearModified(): void;
  getModified(): number | undefined;
  setModified(value: number): void;

  hasDeleted(): boolean;
  clearDeleted(): void;
  getDeleted(): number | undefined;
  setDeleted(value: number): void;

  clearRolesList(): void;
  getRolesList(): Array<number>;
  setRolesList(value: Array<number>): void;
  addRoles(value: number, index?: number): number;

  clearRoomsList(): void;
  getRoomsList(): Array<RoomForUser>;
  setRoomsList(value: Array<RoomForUser>): void;
  addRooms(value?: RoomForUser, index?: number): RoomForUser;

  clearDevicesList(): void;
  getDevicesList(): Array<deviceMessage_pb.Device>;
  setDevicesList(value: Array<deviceMessage_pb.Device>): void;
  addDevices(value?: deviceMessage_pb.Device, index?: number): deviceMessage_pb.Device;

  clearBlocksList(): void;
  getBlocksList(): Array<string>;
  setBlocksList(value: Array<string>): void;
  addBlocks(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id?: number,
    userId?: string,
    name?: string,
    pictureUrl?: string,
    informationUrl?: string,
    unreadCount?: number,
    metaData: Uint8Array | string,
    pb_public?: boolean,
    canBlock?: boolean,
    lang?: string,
    accessToken?: string,
    lastAccessRoomId?: string,
    lastAccessed?: number,
    created?: number,
    modified?: number,
    deleted?: number,
    rolesList: Array<number>,
    roomsList: Array<RoomForUser.AsObject>,
    devicesList: Array<deviceMessage_pb.Device.AsObject>,
    blocksList: Array<string>,
  }
}

export class RoomForUser extends jspb.Message {
  hasRoomId(): boolean;
  clearRoomId(): void;
  getRoomId(): string | undefined;
  setRoomId(value: string): void;

  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasPictureUrl(): boolean;
  clearPictureUrl(): void;
  getPictureUrl(): string | undefined;
  setPictureUrl(value: string): void;

  hasInformationUrl(): boolean;
  clearInformationUrl(): void;
  getInformationUrl(): string | undefined;
  setInformationUrl(value: string): void;

  hasMetaData(): boolean;
  clearMetaData(): void;
  getMetaData(): Uint8Array | string;
  getMetaData_asU8(): Uint8Array;
  getMetaData_asB64(): string;
  setMetaData(value: Uint8Array | string): void;

  hasType(): boolean;
  clearType(): void;
  getType(): roomMessage_pb.RoomType | undefined;
  setType(value: roomMessage_pb.RoomType): void;

  hasLastMessage(): boolean;
  clearLastMessage(): void;
  getLastMessage(): string | undefined;
  setLastMessage(value: string): void;

  hasLastMessageUpdated(): boolean;
  clearLastMessageUpdated(): void;
  getLastMessageUpdated(): number | undefined;
  setLastMessageUpdated(value: number): void;

  hasCanLeft(): boolean;
  clearCanLeft(): void;
  getCanLeft(): boolean | undefined;
  setCanLeft(value: boolean): void;

  hasCreated(): boolean;
  clearCreated(): void;
  getCreated(): number | undefined;
  setCreated(value: number): void;

  hasModified(): boolean;
  clearModified(): void;
  getModified(): number | undefined;
  setModified(value: number): void;

  clearUsersList(): void;
  getUsersList(): Array<roomMessage_pb.UserForRoom>;
  setUsersList(value: Array<roomMessage_pb.UserForRoom>): void;
  addUsers(value?: roomMessage_pb.UserForRoom, index?: number): roomMessage_pb.UserForRoom;

  hasRuUnreadCount(): boolean;
  clearRuUnreadCount(): void;
  getRuUnreadCount(): number | undefined;
  setRuUnreadCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomForUser.AsObject;
  static toObject(includeInstance: boolean, msg: RoomForUser): RoomForUser.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RoomForUser, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomForUser;
  static deserializeBinaryFromReader(message: RoomForUser, reader: jspb.BinaryReader): RoomForUser;
}

export namespace RoomForUser {
  export type AsObject = {
    roomId?: string,
    userId?: string,
    name?: string,
    pictureUrl?: string,
    informationUrl?: string,
    metaData: Uint8Array | string,
    type?: roomMessage_pb.RoomType,
    lastMessage?: string,
    lastMessageUpdated?: number,
    canLeft?: boolean,
    created?: number,
    modified?: number,
    usersList: Array<roomMessage_pb.UserForRoom.AsObject>,
    ruUnreadCount?: number,
  }
}

export class CreateUserRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasPictureUrl(): boolean;
  clearPictureUrl(): void;
  getPictureUrl(): string | undefined;
  setPictureUrl(value: string): void;

  hasInformationUrl(): boolean;
  clearInformationUrl(): void;
  getInformationUrl(): string | undefined;
  setInformationUrl(value: string): void;

  hasMetaData(): boolean;
  clearMetaData(): void;
  getMetaData(): Uint8Array | string;
  getMetaData_asU8(): Uint8Array;
  getMetaData_asB64(): string;
  setMetaData(value: Uint8Array | string): void;

  hasPublic(): boolean;
  clearPublic(): void;
  getPublic(): boolean | undefined;
  setPublic(value: boolean): void;

  hasCanBlock(): boolean;
  clearCanBlock(): void;
  getCanBlock(): boolean | undefined;
  setCanBlock(value: boolean): void;

  hasLang(): boolean;
  clearLang(): void;
  getLang(): string | undefined;
  setLang(value: string): void;

  clearRoleIdsList(): void;
  getRoleIdsList(): Array<number>;
  setRoleIdsList(value: Array<number>): void;
  addRoleIds(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateUserRequest): CreateUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateUserRequest;
  static deserializeBinaryFromReader(message: CreateUserRequest, reader: jspb.BinaryReader): CreateUserRequest;
}

export namespace CreateUserRequest {
  export type AsObject = {
    userId?: string,
    name?: string,
    pictureUrl?: string,
    informationUrl?: string,
    metaData: Uint8Array | string,
    pb_public?: boolean,
    canBlock?: boolean,
    lang?: string,
    roleIdsList: Array<number>,
  }
}

export class GetUsersRequest extends jspb.Message {
  hasLimit(): boolean;
  clearLimit(): void;
  getLimit(): number | undefined;
  setLimit(value: number): void;

  hasOffset(): boolean;
  clearOffset(): void;
  getOffset(): number | undefined;
  setOffset(value: number): void;

  clearOrdersList(): void;
  getOrdersList(): Array<commonMessage_pb.OrderInfo>;
  setOrdersList(value: Array<commonMessage_pb.OrderInfo>): void;
  addOrders(value?: commonMessage_pb.OrderInfo, index?: number): commonMessage_pb.OrderInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUsersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUsersRequest): GetUsersRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUsersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUsersRequest;
  static deserializeBinaryFromReader(message: GetUsersRequest, reader: jspb.BinaryReader): GetUsersRequest;
}

export namespace GetUsersRequest {
  export type AsObject = {
    limit?: number,
    offset?: number,
    ordersList: Array<commonMessage_pb.OrderInfo.AsObject>,
  }
}

export class UsersResponse extends jspb.Message {
  clearUsersList(): void;
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): void;
  addUsers(value?: User, index?: number): User;

  hasAllcount(): boolean;
  clearAllcount(): void;
  getAllcount(): number | undefined;
  setAllcount(value: number): void;

  hasLimit(): boolean;
  clearLimit(): void;
  getLimit(): number | undefined;
  setLimit(value: number): void;

  hasOffset(): boolean;
  clearOffset(): void;
  getOffset(): number | undefined;
  setOffset(value: number): void;

  clearOrdersList(): void;
  getOrdersList(): Array<commonMessage_pb.OrderInfo>;
  setOrdersList(value: Array<commonMessage_pb.OrderInfo>): void;
  addOrders(value?: commonMessage_pb.OrderInfo, index?: number): commonMessage_pb.OrderInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UsersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UsersResponse): UsersResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UsersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UsersResponse;
  static deserializeBinaryFromReader(message: UsersResponse, reader: jspb.BinaryReader): UsersResponse;
}

export namespace UsersResponse {
  export type AsObject = {
    usersList: Array<User.AsObject>,
    allcount?: number,
    limit?: number,
    offset?: number,
    ordersList: Array<commonMessage_pb.OrderInfo.AsObject>,
  }
}

export class GetUserRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRequest): GetUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRequest;
  static deserializeBinaryFromReader(message: GetUserRequest, reader: jspb.BinaryReader): GetUserRequest;
}

export namespace GetUserRequest {
  export type AsObject = {
    userId?: string,
  }
}

export class UpdateUserRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasPictureUrl(): boolean;
  clearPictureUrl(): void;
  getPictureUrl(): string | undefined;
  setPictureUrl(value: string): void;

  hasInformationUrl(): boolean;
  clearInformationUrl(): void;
  getInformationUrl(): string | undefined;
  setInformationUrl(value: string): void;

  hasMetaData(): boolean;
  clearMetaData(): void;
  getMetaData(): Uint8Array | string;
  getMetaData_asU8(): Uint8Array;
  getMetaData_asB64(): string;
  setMetaData(value: Uint8Array | string): void;

  hasPublic(): boolean;
  clearPublic(): void;
  getPublic(): boolean | undefined;
  setPublic(value: boolean): void;

  hasCanBlock(): boolean;
  clearCanBlock(): void;
  getCanBlock(): boolean | undefined;
  setCanBlock(value: boolean): void;

  hasLang(): boolean;
  clearLang(): void;
  getLang(): string | undefined;
  setLang(value: string): void;

  clearRoleIdsList(): void;
  getRoleIdsList(): Array<number>;
  setRoleIdsList(value: Array<number>): void;
  addRoleIds(value: number, index?: number): number;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateUserRequest): UpdateUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateUserRequest;
  static deserializeBinaryFromReader(message: UpdateUserRequest, reader: jspb.BinaryReader): UpdateUserRequest;
}

export namespace UpdateUserRequest {
  export type AsObject = {
    userId?: string,
    name?: string,
    pictureUrl?: string,
    informationUrl?: string,
    metaData: Uint8Array | string,
    pb_public?: boolean,
    canBlock?: boolean,
    lang?: string,
    roleIdsList: Array<number>,
  }
}

export class DeleteUserRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteUserRequest): DeleteUserRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteUserRequest;
  static deserializeBinaryFromReader(message: DeleteUserRequest, reader: jspb.BinaryReader): DeleteUserRequest;
}

export namespace DeleteUserRequest {
  export type AsObject = {
    userId?: string,
  }
}

export class GetContactsRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  hasLimit(): boolean;
  clearLimit(): void;
  getLimit(): number | undefined;
  setLimit(value: number): void;

  hasOffset(): boolean;
  clearOffset(): void;
  getOffset(): number | undefined;
  setOffset(value: number): void;

  clearOrdersList(): void;
  getOrdersList(): Array<commonMessage_pb.OrderInfo>;
  setOrdersList(value: Array<commonMessage_pb.OrderInfo>): void;
  addOrders(value?: commonMessage_pb.OrderInfo, index?: number): commonMessage_pb.OrderInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetContactsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetContactsRequest): GetContactsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetContactsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetContactsRequest;
  static deserializeBinaryFromReader(message: GetContactsRequest, reader: jspb.BinaryReader): GetContactsRequest;
}

export namespace GetContactsRequest {
  export type AsObject = {
    userId?: string,
    limit?: number,
    offset?: number,
    ordersList: Array<commonMessage_pb.OrderInfo.AsObject>,
  }
}

export class GetProfileRequest extends jspb.Message {
  hasUserId(): boolean;
  clearUserId(): void;
  getUserId(): string | undefined;
  setUserId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProfileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetProfileRequest): GetProfileRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProfileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProfileRequest;
  static deserializeBinaryFromReader(message: GetProfileRequest, reader: jspb.BinaryReader): GetProfileRequest;
}

export namespace GetProfileRequest {
  export type AsObject = {
    userId?: string,
  }
}

