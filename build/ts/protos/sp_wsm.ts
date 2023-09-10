// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "sp_wsm.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { PipelineStep } from "./sp_pipeline.js";
/**
 * SDK generates a WASM request and passes this to the WASM func
 *
 * @generated from protobuf message protos.WASMRequest
 */
export interface WASMRequest {
    /**
     * The actual step that the WASM func will operate on. This is the same step
     * that is declared in protos.Pipeline.
     *
     * @generated from protobuf field: protos.PipelineStep step = 1;
     */
    step?: PipelineStep;
    /**
     * Payload data that WASM func will operate on
     *
     * @generated from protobuf field: bytes input_payload = 2;
     */
    inputPayload: Uint8Array;
    /**
     * Potentially filled out result from previous step. If this is first step in
     * the pipeline, it will be empty.
     *
     * @generated from protobuf field: optional bytes input_step = 3;
     */
    inputStep?: Uint8Array;
}
/**
 * Returned by all WASM functions
 *
 * @generated from protobuf message protos.WASMResponse
 */
export interface WASMResponse {
    /**
     * Potentially modified input payload. Concept: All WASM funcs accept an
     * input_payload in WASMRequest, WASM func reads input payload, modifies it
     * and writes the modified output to output_payload.
     *
     * @generated from protobuf field: bytes output_payload = 1;
     */
    outputPayload: Uint8Array;
    /**
     * Exit code that the WASM func exited with; more info in WASMExitCode's comment
     *
     * @generated from protobuf field: protos.WASMExitCode exit_code = 2;
     */
    exitCode: WASMExitCode;
    /**
     * Additional info about the reason a specific exit code was returned
     *
     * @generated from protobuf field: string exit_msg = 3;
     */
    exitMsg: string;
    /**
     * Potential additional step output - ie. if a WASM func is an HTTPGet,
     * output_step would contain the HTTP response body; if the WASM func is a
     * KVGet, the output_step would be the value of the fetched key.
     *
     * @generated from protobuf field: optional bytes output_step = 4;
     */
    outputStep?: Uint8Array;
}
/**
 * Included in WASM response; the SDK should use the WASMExitCode to determine
 * what to do next - should it execute next step, should it notify or should it
 * stop executing/abort the rest of the steps in the pipeline.
 *
 * Example:
 *
 * a. WASM func returns WASM_EXIT_CODE_FAILURE - read PipelineStep.on_failure
 * conditions to determine what to do next.
 *
 * b. WASM func returns WASM_EXIT_CODE_SUCCESS - read PipelineStep.on_success
 * conditions to determine what to do next.
 *
 * .. and so on.
 * protolint:disable:next ENUM_FIELD_NAMES_PREFIX
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
            { no: 2, name: "input_payload", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 3, name: "input_step", kind: "scalar", opt: true, T: 12 /*ScalarType.BYTES*/ }
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
            { no: 1, name: "output_payload", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 2, name: "exit_code", kind: "enum", T: () => ["protos.WASMExitCode", WASMExitCode] },
            { no: 3, name: "exit_msg", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "output_step", kind: "scalar", opt: true, T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.WASMResponse
 */
export const WASMResponse = new WASMResponse$Type();
