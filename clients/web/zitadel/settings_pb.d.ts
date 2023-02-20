// @generated by protoc-gen-es v1.0.0
// @generated from file zitadel/settings.proto (package zitadel.settings.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { ObjectDetails } from "./object_pb.js";

/**
 * @generated from enum zitadel.settings.v1.SecretGeneratorType
 */
export declare enum SecretGeneratorType {
  /**
   * @generated from enum value: SECRET_GENERATOR_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SECRET_GENERATOR_TYPE_INIT_CODE = 1;
   */
  INIT_CODE = 1,

  /**
   * @generated from enum value: SECRET_GENERATOR_TYPE_VERIFY_EMAIL_CODE = 2;
   */
  VERIFY_EMAIL_CODE = 2,

  /**
   * @generated from enum value: SECRET_GENERATOR_TYPE_VERIFY_PHONE_CODE = 3;
   */
  VERIFY_PHONE_CODE = 3,

  /**
   * @generated from enum value: SECRET_GENERATOR_TYPE_PASSWORD_RESET_CODE = 4;
   */
  PASSWORD_RESET_CODE = 4,

  /**
   * @generated from enum value: SECRET_GENERATOR_TYPE_PASSWORDLESS_INIT_CODE = 5;
   */
  PASSWORDLESS_INIT_CODE = 5,

  /**
   * @generated from enum value: SECRET_GENERATOR_TYPE_APP_SECRET = 6;
   */
  APP_SECRET = 6,
}

/**
 * @generated from enum zitadel.settings.v1.SMSProviderConfigState
 */
export declare enum SMSProviderConfigState {
  /**
   * @generated from enum value: SMS_PROVIDER_CONFIG_STATE_UNSPECIFIED = 0;
   */
  SMS_PROVIDER_CONFIG_STATE_UNSPECIFIED = 0,

  /**
   * @generated from enum value: SMS_PROVIDER_CONFIG_ACTIVE = 1;
   */
  SMS_PROVIDER_CONFIG_ACTIVE = 1,

  /**
   * @generated from enum value: SMS_PROVIDER_CONFIG_INACTIVE = 2;
   */
  SMS_PROVIDER_CONFIG_INACTIVE = 2,
}

/**
 * @generated from message zitadel.settings.v1.SecretGenerator
 */
export declare class SecretGenerator extends Message<SecretGenerator> {
  /**
   * @generated from field: zitadel.settings.v1.SecretGeneratorType generator_type = 1;
   */
  generatorType: SecretGeneratorType;

  /**
   * @generated from field: zitadel.v1.ObjectDetails details = 2;
   */
  details?: ObjectDetails;

  /**
   * @generated from field: uint32 length = 3;
   */
  length: number;

  /**
   * @generated from field: google.protobuf.Duration expiry = 4;
   */
  expiry?: Duration;

  /**
   * @generated from field: bool include_lower_letters = 5;
   */
  includeLowerLetters: boolean;

  /**
   * @generated from field: bool include_upper_letters = 6;
   */
  includeUpperLetters: boolean;

  /**
   * @generated from field: bool include_digits = 7;
   */
  includeDigits: boolean;

  /**
   * @generated from field: bool include_symbols = 8;
   */
  includeSymbols: boolean;

  constructor(data?: PartialMessage<SecretGenerator>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zitadel.settings.v1.SecretGenerator";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SecretGenerator;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SecretGenerator;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SecretGenerator;

  static equals(a: SecretGenerator | PlainMessage<SecretGenerator> | undefined, b: SecretGenerator | PlainMessage<SecretGenerator> | undefined): boolean;
}

/**
 * @generated from message zitadel.settings.v1.SecretGeneratorQuery
 */
export declare class SecretGeneratorQuery extends Message<SecretGeneratorQuery> {
  /**
   * @generated from oneof zitadel.settings.v1.SecretGeneratorQuery.query
   */
  query: {
    /**
     * @generated from field: zitadel.settings.v1.SecretGeneratorTypeQuery type_query = 1;
     */
    value: SecretGeneratorTypeQuery;
    case: "typeQuery";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<SecretGeneratorQuery>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zitadel.settings.v1.SecretGeneratorQuery";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SecretGeneratorQuery;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SecretGeneratorQuery;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SecretGeneratorQuery;

  static equals(a: SecretGeneratorQuery | PlainMessage<SecretGeneratorQuery> | undefined, b: SecretGeneratorQuery | PlainMessage<SecretGeneratorQuery> | undefined): boolean;
}

/**
 * @generated from message zitadel.settings.v1.SecretGeneratorTypeQuery
 */
export declare class SecretGeneratorTypeQuery extends Message<SecretGeneratorTypeQuery> {
  /**
   * @generated from field: zitadel.settings.v1.SecretGeneratorType generator_type = 1;
   */
  generatorType: SecretGeneratorType;

  constructor(data?: PartialMessage<SecretGeneratorTypeQuery>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zitadel.settings.v1.SecretGeneratorTypeQuery";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SecretGeneratorTypeQuery;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SecretGeneratorTypeQuery;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SecretGeneratorTypeQuery;

  static equals(a: SecretGeneratorTypeQuery | PlainMessage<SecretGeneratorTypeQuery> | undefined, b: SecretGeneratorTypeQuery | PlainMessage<SecretGeneratorTypeQuery> | undefined): boolean;
}

/**
 * @generated from message zitadel.settings.v1.SMTPConfig
 */
export declare class SMTPConfig extends Message<SMTPConfig> {
  /**
   * @generated from field: zitadel.v1.ObjectDetails details = 1;
   */
  details?: ObjectDetails;

  /**
   * @generated from field: string sender_address = 2;
   */
  senderAddress: string;

  /**
   * @generated from field: string sender_name = 3;
   */
  senderName: string;

  /**
   * @generated from field: bool tls = 4;
   */
  tls: boolean;

  /**
   * @generated from field: string host = 5;
   */
  host: string;

  /**
   * @generated from field: string user = 6;
   */
  user: string;

  constructor(data?: PartialMessage<SMTPConfig>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zitadel.settings.v1.SMTPConfig";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SMTPConfig;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SMTPConfig;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SMTPConfig;

  static equals(a: SMTPConfig | PlainMessage<SMTPConfig> | undefined, b: SMTPConfig | PlainMessage<SMTPConfig> | undefined): boolean;
}

/**
 * @generated from message zitadel.settings.v1.SMSProvider
 */
export declare class SMSProvider extends Message<SMSProvider> {
  /**
   * @generated from field: zitadel.v1.ObjectDetails details = 1;
   */
  details?: ObjectDetails;

  /**
   * @generated from field: string id = 2;
   */
  id: string;

  /**
   * @generated from field: zitadel.settings.v1.SMSProviderConfigState state = 3;
   */
  state: SMSProviderConfigState;

  /**
   * @generated from oneof zitadel.settings.v1.SMSProvider.config
   */
  config: {
    /**
     * @generated from field: zitadel.settings.v1.TwilioConfig twilio = 4;
     */
    value: TwilioConfig;
    case: "twilio";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<SMSProvider>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zitadel.settings.v1.SMSProvider";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SMSProvider;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SMSProvider;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SMSProvider;

  static equals(a: SMSProvider | PlainMessage<SMSProvider> | undefined, b: SMSProvider | PlainMessage<SMSProvider> | undefined): boolean;
}

/**
 * @generated from message zitadel.settings.v1.TwilioConfig
 */
export declare class TwilioConfig extends Message<TwilioConfig> {
  /**
   * @generated from field: string sid = 1;
   */
  sid: string;

  /**
   * @generated from field: string sender_number = 2;
   */
  senderNumber: string;

  constructor(data?: PartialMessage<TwilioConfig>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zitadel.settings.v1.TwilioConfig";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TwilioConfig;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TwilioConfig;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TwilioConfig;

  static equals(a: TwilioConfig | PlainMessage<TwilioConfig> | undefined, b: TwilioConfig | PlainMessage<TwilioConfig> | undefined): boolean;
}

/**
 * @generated from message zitadel.settings.v1.DebugNotificationProvider
 */
export declare class DebugNotificationProvider extends Message<DebugNotificationProvider> {
  /**
   * @generated from field: zitadel.v1.ObjectDetails details = 1;
   */
  details?: ObjectDetails;

  /**
   * @generated from field: bool compact = 2;
   */
  compact: boolean;

  constructor(data?: PartialMessage<DebugNotificationProvider>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zitadel.settings.v1.DebugNotificationProvider";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DebugNotificationProvider;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DebugNotificationProvider;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DebugNotificationProvider;

  static equals(a: DebugNotificationProvider | PlainMessage<DebugNotificationProvider> | undefined, b: DebugNotificationProvider | PlainMessage<DebugNotificationProvider> | undefined): boolean;
}

/**
 * @generated from message zitadel.settings.v1.OIDCSettings
 */
export declare class OIDCSettings extends Message<OIDCSettings> {
  /**
   * @generated from field: zitadel.v1.ObjectDetails details = 1;
   */
  details?: ObjectDetails;

  /**
   * @generated from field: google.protobuf.Duration access_token_lifetime = 2;
   */
  accessTokenLifetime?: Duration;

  /**
   * @generated from field: google.protobuf.Duration id_token_lifetime = 3;
   */
  idTokenLifetime?: Duration;

  /**
   * @generated from field: google.protobuf.Duration refresh_token_idle_expiration = 4;
   */
  refreshTokenIdleExpiration?: Duration;

  /**
   * @generated from field: google.protobuf.Duration refresh_token_expiration = 5;
   */
  refreshTokenExpiration?: Duration;

  constructor(data?: PartialMessage<OIDCSettings>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zitadel.settings.v1.OIDCSettings";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OIDCSettings;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OIDCSettings;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OIDCSettings;

  static equals(a: OIDCSettings | PlainMessage<OIDCSettings> | undefined, b: OIDCSettings | PlainMessage<OIDCSettings> | undefined): boolean;
}

/**
 * @generated from message zitadel.settings.v1.SecurityPolicy
 */
export declare class SecurityPolicy extends Message<SecurityPolicy> {
  /**
   * @generated from field: zitadel.v1.ObjectDetails details = 1;
   */
  details?: ObjectDetails;

  /**
   * states if iframe embedding is enabled or disabled
   *
   * @generated from field: bool enable_iframe_embedding = 2;
   */
  enableIframeEmbedding: boolean;

  /**
   * origins allowed to load ZITADEL in an iframe if enable_iframe_embedding is true
   *
   * @generated from field: repeated string allowed_origins = 3;
   */
  allowedOrigins: string[];

  constructor(data?: PartialMessage<SecurityPolicy>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zitadel.settings.v1.SecurityPolicy";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SecurityPolicy;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SecurityPolicy;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SecurityPolicy;

  static equals(a: SecurityPolicy | PlainMessage<SecurityPolicy> | undefined, b: SecurityPolicy | PlainMessage<SecurityPolicy> | undefined): boolean;
}

