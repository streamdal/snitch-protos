// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "bus.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { ResumePipelineRequest } from "./external.js";
import { PausePipelineRequest } from "./external.js";
import { DetachPipelineRequest } from "./external.js";
import { AttachPipelineRequest } from "./external.js";
import { UpdatePipelineRequest } from "./external.js";
import { DeletePipelineRequest } from "./external.js";
import { CreatePipelineRequest } from "./external.js";
import { HeartbeatRequest } from "./internal.js";
import { DeregisterRequest } from "./internal.js";
import { RegisterRequest } from "./internal.js";
/**
 * Type used by `snitch-server` for broadcasting events to other snitch nodes
 *
 * @generated from protobuf message protos.BusEvent
 */
export interface BusEvent {
    /**
     * @generated from protobuf field: string source = 1;
     */
    source: string;
    /**
     * @generated from protobuf oneof: event
     */
    event: {
        oneofKind: "registerRequest";
        /**
         * @generated from protobuf field: protos.RegisterRequest register_request = 100;
         */
        registerRequest: RegisterRequest;
    } | {
        oneofKind: "deregisterRequest";
        /**
         * @generated from protobuf field: protos.DeregisterRequest deregister_request = 101;
         */
        deregisterRequest: DeregisterRequest;
    } | {
        oneofKind: "heartbeatRequest";
        /**
         * @generated from protobuf field: protos.HeartbeatRequest heartbeat_request = 102;
         */
        heartbeatRequest: HeartbeatRequest;
    } | {
        oneofKind: "createPipelineRequest";
        /**
         * @generated from protobuf field: protos.CreatePipelineRequest create_pipeline_request = 103;
         */
        createPipelineRequest: CreatePipelineRequest;
    } | {
        oneofKind: "deletePipelineRequest";
        /**
         * @generated from protobuf field: protos.DeletePipelineRequest delete_pipeline_request = 104;
         */
        deletePipelineRequest: DeletePipelineRequest;
    } | {
        oneofKind: "updatePipelineRequest";
        /**
         * @generated from protobuf field: protos.UpdatePipelineRequest update_pipeline_request = 105;
         */
        updatePipelineRequest: UpdatePipelineRequest;
    } | {
        oneofKind: "attachPipelineRequest";
        /**
         * @generated from protobuf field: protos.AttachPipelineRequest attach_pipeline_request = 106;
         */
        attachPipelineRequest: AttachPipelineRequest;
    } | {
        oneofKind: "detachPipelineRequest";
        /**
         * @generated from protobuf field: protos.DetachPipelineRequest detach_pipeline_request = 107;
         */
        detachPipelineRequest: DetachPipelineRequest;
    } | {
        oneofKind: "pausePipelineRequest";
        /**
         * @generated from protobuf field: protos.PausePipelineRequest pause_pipeline_request = 108;
         */
        pausePipelineRequest: PausePipelineRequest;
    } | {
        oneofKind: "resumePipelineRequest";
        /**
         * @generated from protobuf field: protos.ResumePipelineRequest resume_pipeline_request = 109;
         */
        resumePipelineRequest: ResumePipelineRequest;
    } | {
        oneofKind: undefined;
    };
    /**
     * All gRPC metadata is stored in ctx; when request goes outside of gRPC
     * bounds, we will translate ctx metadata into this field.
     *
     * Example:
     * 1. Request comes into snitch-server via external gRPC to set new pipeline
     * 2. snitch-server has to send SetPipeline cmd to SDK via gRPC - it passes
     *    on original metadata in request.
     * 3. snitch-server has to broadcast SetPipeline cmd to other services via bus
     * 4. Since this is not a gRPC call, snitch-server translates ctx metadata to
     *    this field and includes it in the bus event.
     *
     * @generated from protobuf field: map<string, string> _metadata = 1000;
     */
    Metadata: {
        [key: string]: string;
    }; // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}
// @generated message type with reflection information, may provide speed optimized methods
class BusEvent$Type extends MessageType<BusEvent> {
    constructor() {
        super("protos.BusEvent", [
            { no: 1, name: "source", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 100, name: "register_request", kind: "message", oneof: "event", T: () => RegisterRequest },
            { no: 101, name: "deregister_request", kind: "message", oneof: "event", T: () => DeregisterRequest },
            { no: 102, name: "heartbeat_request", kind: "message", oneof: "event", T: () => HeartbeatRequest },
            { no: 103, name: "create_pipeline_request", kind: "message", oneof: "event", T: () => CreatePipelineRequest },
            { no: 104, name: "delete_pipeline_request", kind: "message", oneof: "event", T: () => DeletePipelineRequest },
            { no: 105, name: "update_pipeline_request", kind: "message", oneof: "event", T: () => UpdatePipelineRequest },
            { no: 106, name: "attach_pipeline_request", kind: "message", oneof: "event", T: () => AttachPipelineRequest },
            { no: 107, name: "detach_pipeline_request", kind: "message", oneof: "event", T: () => DetachPipelineRequest },
            { no: 108, name: "pause_pipeline_request", kind: "message", oneof: "event", T: () => PausePipelineRequest },
            { no: 109, name: "resume_pipeline_request", kind: "message", oneof: "event", T: () => ResumePipelineRequest },
            { no: 1000, name: "_metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.BusEvent
 */
export const BusEvent = new BusEvent$Type();
