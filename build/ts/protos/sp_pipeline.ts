// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "sp_pipeline.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { CustomStep } from "./steps/sp_steps_custom.js";
import { DecodeStep } from "./steps/sp_steps_decode.js";
import { EncodeStep } from "./steps/sp_steps_encode.js";
import { TransformStep } from "./steps/sp_steps_transform.js";
import { DetectiveStep } from "./steps/sp_steps_detective.js";
/**
 * Pipeline is a structure that holds one or more pipeline steps. This structure
 * is intended to be immutable; clients are expected to generate WASMRequest's
 * that contain a pipeline step.
 *
 * @generated from protobuf message protos.Pipeline
 */
export interface Pipeline {
    /**
     * ID should NOT be set by external gRPC client on CreatePipelineRequest - it
     * will be ignored; it _does_ need to be set on UpdatePipelineRequest.
     *
     * @generated from protobuf field: string id = 1;
     */
    id: string;
    /**
     * Friendly name for the pipeline
     *
     * @generated from protobuf field: string name = 2;
     */
    name: string;
    /**
     * One or more steps to execute
     *
     * @generated from protobuf field: repeated protos.PipelineStep steps = 3;
     */
    steps: PipelineStep[];
}
/**
 * A pipeline step is a single step in a pipeline.
 *
 * @generated from protobuf message protos.PipelineStep
 */
export interface PipelineStep {
    /**
     * Friendly name for the step
     *
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * What should SDK do if step succeeds?
     *
     *
     * @generated from protobuf field: repeated protos.PipelineStepCondition on_success = 2;
     */
    onSuccess: PipelineStepCondition[];
    /**
     * What should SDK do if step fails?
     *
     * @generated from protobuf field: repeated protos.PipelineStepCondition on_failure = 3;
     */
    onFailure: PipelineStepCondition[];
    /**
     * @generated from protobuf oneof: step
     */
    step: {
        oneofKind: "detective";
        /**
         * @generated from protobuf field: protos.steps.DetectiveStep detective = 1000;
         */
        detective: DetectiveStep;
    } | {
        oneofKind: "transform";
        /**
         * @generated from protobuf field: protos.steps.TransformStep transform = 1001;
         */
        transform: TransformStep;
    } | {
        oneofKind: "encode";
        /**
         * @generated from protobuf field: protos.steps.EncodeStep encode = 1002;
         */
        encode: EncodeStep;
    } | {
        oneofKind: "decode";
        /**
         * @generated from protobuf field: protos.steps.DecodeStep decode = 1003;
         */
        decode: DecodeStep;
    } | {
        oneofKind: "custom";
        /**
         * @generated from protobuf field: protos.steps.CustomStep custom = 1004;
         */
        custom: CustomStep;
    } | {
        oneofKind: undefined;
    };
    /**
     * ID is a uuid(sha256(_wasm_bytes)) that is set by snitch-server
     *
     * @generated from protobuf field: optional string _wasm_id = 10000;
     */
    WasmId?: string; // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
    /**
     * WASM module bytes (set by snitch-server)
     *
     * @generated from protobuf field: optional bytes _wasm_bytes = 10001;
     */
    WasmBytes?: Uint8Array; // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
    /**
     * WASM function name to execute (set by snitch-server)
     *
     * @generated from protobuf field: optional string _wasm_function = 10002;
     */
    WasmFunction?: string; // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}
/**
 * A condition defines how the SDK should handle a step response -- should it
 * continue executing the pipeline, should it abort, should it notify the server?
 * Each step can have multiple conditions.
 *
 * @generated from protobuf enum protos.PipelineStepCondition
 */
export enum PipelineStepCondition {
    /**
     * @generated from protobuf enum value: PIPELINE_STEP_CONDITION_UNSET = 0;
     */
    UNSET = 0,
    /**
     * @generated from protobuf enum value: PIPELINE_STEP_CONDITION_ABORT = 1;
     */
    ABORT = 1,
    /**
     * @generated from protobuf enum value: PIPELINE_STEP_CONDITION_NOTIFY = 2;
     */
    NOTIFY = 2
}
// @generated message type with reflection information, may provide speed optimized methods
class Pipeline$Type extends MessageType<Pipeline> {
    constructor() {
        super("protos.Pipeline", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "steps", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => PipelineStep }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.Pipeline
 */
export const Pipeline = new Pipeline$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PipelineStep$Type extends MessageType<PipelineStep> {
    constructor() {
        super("protos.PipelineStep", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "on_success", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["protos.PipelineStepCondition", PipelineStepCondition, "PIPELINE_STEP_CONDITION_"] },
            { no: 3, name: "on_failure", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["protos.PipelineStepCondition", PipelineStepCondition, "PIPELINE_STEP_CONDITION_"] },
            { no: 1000, name: "detective", kind: "message", oneof: "step", T: () => DetectiveStep },
            { no: 1001, name: "transform", kind: "message", oneof: "step", T: () => TransformStep },
            { no: 1002, name: "encode", kind: "message", oneof: "step", T: () => EncodeStep },
            { no: 1003, name: "decode", kind: "message", oneof: "step", T: () => DecodeStep },
            { no: 1004, name: "custom", kind: "message", oneof: "step", T: () => CustomStep },
            { no: 10000, name: "_wasm_id", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 10001, name: "_wasm_bytes", kind: "scalar", opt: true, T: 12 /*ScalarType.BYTES*/ },
            { no: 10002, name: "_wasm_function", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.PipelineStep
 */
export const PipelineStep = new PipelineStep$Type();