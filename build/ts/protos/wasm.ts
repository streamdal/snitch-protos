// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "wasm.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { PipelineStep } from "./pipeline.js";
/**
 * SDK generates a WASM request and passes this to the WASM func
 *
 * @generated from protobuf message protos.WASMRequest
 */
export interface WASMRequest {
    /**
     * @generated from protobuf field: protos.PipelineStep step = 1;
     */
    step?: PipelineStep;
    /**
     * @generated from protobuf field: bytes input = 2;
     */
    input: Uint8Array;
}
/**
 * Returned by all WASM functions
 *
 * @generated from protobuf message protos.WASMResponse
 */
export interface WASMResponse {
    /**
     * @generated from protobuf field: bytes output = 1;
     */
    output: Uint8Array;
    /**
     * @generated from protobuf field: protos.WASMExitCode exit_code = 2;
     */
    exitCode: WASMExitCode;
    /**
     * @generated from protobuf field: string exit_msg = 3;
     */
    exitMsg: string;
}
/**
 * Included in WASM response; SDK is responsible for interpreting the response
 * status and how it relates to the step condition.
 * ie. WASM func returns WASM_EXIT_CODE_INTERNAL_ERROR lookup ON_ERROR
 * conditions to determine what to do next.
 * ie. WASM func returns WASM_EXIT_CODE_SUCCESS lookup ON_MATCH conditions
 * to determine what to do next;
 *
 * protolint:disable:this ENUM_FIELD_NAMES_PREFIX
 *
 * @generated from protobuf enum protos.WASMExitCode
 */
export enum WASMExitCode {
    /**
     * @generated from protobuf enum value: WASM_EXIT_CODE_UNSET = 0;
     */
    WASM_EXIT_CODE_UNSET = 0,
    /**
     * @generated from protobuf enum value: WASM_EXIT_CODE_SUCCESS = 1;
     */
    WASM_EXIT_CODE_SUCCESS = 1,
    /**
     * Probably need better names for these as FAILURE is too harsh
     *
     * @generated from protobuf enum value: WASM_EXIT_CODE_FAILURE = 2;
     */
    WASM_EXIT_CODE_FAILURE = 2,
    /**
     * @generated from protobuf enum value: WASM_EXIT_CODE_INTERNAL_ERROR = 3;
     */
    WASM_EXIT_CODE_INTERNAL_ERROR = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class WASMRequest$Type extends MessageType<WASMRequest> {
    constructor() {
        super("protos.WASMRequest", [
            { no: 1, name: "step", kind: "message", T: () => PipelineStep },
            { no: 2, name: "input", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.WASMRequest
 */
export const WASMRequest = new WASMRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class WASMResponse$Type extends MessageType<WASMResponse> {
    constructor() {
        super("protos.WASMResponse", [
            { no: 1, name: "output", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 2, name: "exit_code", kind: "enum", T: () => ["protos.WASMExitCode", WASMExitCode] },
            { no: 3, name: "exit_msg", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.WASMResponse
 */
export const WASMResponse = new WASMResponse$Type();
