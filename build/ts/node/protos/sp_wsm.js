// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "sp_wsm.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { PipelineStep } from "./sp_pipeline.js";
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
export var WASMExitCode;
(function (WASMExitCode) {
    /**
     * @generated from protobuf enum value: WASM_EXIT_CODE_UNSET = 0;
     */
    WASMExitCode[WASMExitCode["WASM_EXIT_CODE_UNSET"] = 0] = "WASM_EXIT_CODE_UNSET";
    /**
     * @generated from protobuf enum value: WASM_EXIT_CODE_SUCCESS = 1;
     */
    WASMExitCode[WASMExitCode["WASM_EXIT_CODE_SUCCESS"] = 1] = "WASM_EXIT_CODE_SUCCESS";
    /**
     * Probably need better names for these as FAILURE is too harsh
     *
     * @generated from protobuf enum value: WASM_EXIT_CODE_FAILURE = 2;
     */
    WASMExitCode[WASMExitCode["WASM_EXIT_CODE_FAILURE"] = 2] = "WASM_EXIT_CODE_FAILURE";
    /**
     * @generated from protobuf enum value: WASM_EXIT_CODE_INTERNAL_ERROR = 3;
     */
    WASMExitCode[WASMExitCode["WASM_EXIT_CODE_INTERNAL_ERROR"] = 3] = "WASM_EXIT_CODE_INTERNAL_ERROR";
})(WASMExitCode || (WASMExitCode = {}));
// @generated message type with reflection information, may provide speed optimized methods
class WASMRequest$Type extends MessageType {
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
class WASMResponse$Type extends MessageType {
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
//# sourceMappingURL=sp_wsm.js.map