// @generated by protoc-gen-es v1.0.0
// @generated from file zitadel/auth_n_key.proto (package zitadel.authn.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";
import { ObjectDetails } from "./object_pb.js";

/**
 * @generated from enum zitadel.authn.v1.KeyType
 */
export const KeyType = proto3.makeEnum(
  "zitadel.authn.v1.KeyType",
  [
    {no: 0, name: "KEY_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "KEY_TYPE_JSON", localName: "JSON"},
  ],
);

/**
 * @generated from message zitadel.authn.v1.Key
 */
export const Key = proto3.makeMessageType(
  "zitadel.authn.v1.Key",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "details", kind: "message", T: ObjectDetails },
    { no: 3, name: "type", kind: "enum", T: proto3.getEnumType(KeyType) },
    { no: 4, name: "expiration_date", kind: "message", T: Timestamp },
  ],
);

