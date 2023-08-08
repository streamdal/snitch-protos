// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "external.proto" (package "protos", syntax proto3)
// tslint:disable
import { StandardResponse } from "./common.js";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { NotificationConfig } from "./notify.js";
import { Pipeline } from "./pipeline.js";
import { PipelineInfo } from "./info.js";
import { Audience } from "./common.js";
import { LiveInfo } from "./info.js";
/**
 * Don't think there is anything to pass in (yet)?
 *
 * @generated from protobuf message protos.GetAllRequest
 */
export interface GetAllRequest {
}
/**
 * @generated from protobuf message protos.GetAllResponse
 */
export interface GetAllResponse {
    /**
     * Clients currently connected to the server
     *
     * @generated from protobuf field: repeated protos.LiveInfo live = 1;
     */
    live: LiveInfo[];
    /**
     * All of the audiences that are known to the server
     *
     * @generated from protobuf field: repeated protos.Audience audiences = 2;
     */
    audiences: Audience[];
    /**
     * All of the pipelines known to the server + pipeline <-> audience mappings
     * key == pipeline_id; if "Audience" is not filled out - pipeline is not attached
     * to any audience.
     *
     * @generated from protobuf field: map<string, protos.PipelineInfo> pipelines = 3;
     */
    pipelines: {
        [key: string]: PipelineInfo;
    };
}
/**
 * Don't think we need anything here
 *
 * @generated from protobuf message protos.GetPipelinesRequest
 */
export interface GetPipelinesRequest {
}
/**
 * @generated from protobuf message protos.GetPipelinesResponse
 */
export interface GetPipelinesResponse {
    /**
     * @generated from protobuf field: repeated protos.Pipeline pipelines = 1;
     */
    pipelines: Pipeline[];
}
/**
 * @generated from protobuf message protos.GetPipelineRequest
 */
export interface GetPipelineRequest {
    /**
     * @generated from protobuf field: string pipeline_id = 1;
     */
    pipelineId: string;
}
/**
 * @generated from protobuf message protos.GetPipelineResponse
 */
export interface GetPipelineResponse {
    /**
     * @generated from protobuf field: protos.Pipeline pipeline = 1;
     */
    pipeline?: Pipeline;
}
/**
 * @generated from protobuf message protos.CreatePipelineRequest
 */
export interface CreatePipelineRequest {
    /**
     * @generated from protobuf field: protos.Pipeline pipeline = 1;
     */
    pipeline?: Pipeline;
}
/**
 * @generated from protobuf message protos.UpdatePipelineRequest
 */
export interface UpdatePipelineRequest {
    /**
     * @generated from protobuf field: protos.Pipeline pipeline = 1;
     */
    pipeline?: Pipeline;
}
/**
 * @generated from protobuf message protos.DeletePipelineRequest
 */
export interface DeletePipelineRequest {
    /**
     * @generated from protobuf field: string pipeline_id = 1;
     */
    pipelineId: string;
}
/**
 * @generated from protobuf message protos.AttachPipelineRequest
 */
export interface AttachPipelineRequest {
    /**
     * @generated from protobuf field: string pipeline_id = 1;
     */
    pipelineId: string;
    /**
     * @generated from protobuf field: protos.Audience audience = 2;
     */
    audience?: Audience;
}
/**
 * @generated from protobuf message protos.DetachPipelineRequest
 */
export interface DetachPipelineRequest {
    /**
     * @generated from protobuf field: string pipeline_id = 1;
     */
    pipelineId: string;
    /**
     * @generated from protobuf field: protos.Audience audience = 2;
     */
    audience?: Audience;
}
/**
 * @generated from protobuf message protos.PausePipelineRequest
 */
export interface PausePipelineRequest {
    /**
     * @generated from protobuf field: string pipeline_id = 1;
     */
    pipelineId: string;
    /**
     * @generated from protobuf field: protos.Audience audience = 2;
     */
    audience?: Audience;
}
/**
 * @generated from protobuf message protos.ResumePipelineRequest
 */
export interface ResumePipelineRequest {
    /**
     * @generated from protobuf field: string pipeline_id = 1;
     */
    pipelineId: string;
    /**
     * @generated from protobuf field: protos.Audience audience = 2;
     */
    audience?: Audience;
}
/**
 * Notifications
 *
 * @generated from protobuf message protos.CreateNotificationRequest
 */
export interface CreateNotificationRequest {
    /**
     * @generated from protobuf field: protos.NotificationConfig notification = 1;
     */
    notification?: NotificationConfig;
}
/**
 * @generated from protobuf message protos.UpdateNotificationRequest
 */
export interface UpdateNotificationRequest {
    /**
     * @generated from protobuf field: protos.NotificationConfig notification = 1;
     */
    notification?: NotificationConfig;
}
/**
 * @generated from protobuf message protos.DeleteNotificationRequest
 */
export interface DeleteNotificationRequest {
    /**
     * @generated from protobuf field: string notification_id = 1;
     */
    notificationId: string;
}
/**
 * Don't think we need anything here
 *
 * @generated from protobuf message protos.GetNotificationsRequest
 */
export interface GetNotificationsRequest {
}
/**
 * @generated from protobuf message protos.GetNotificationsResponse
 */
export interface GetNotificationsResponse {
    /**
     * @generated from protobuf field: map<string, protos.NotificationConfig> notifications = 1;
     */
    notifications: {
        [key: string]: NotificationConfig;
    };
}
/**
 * @generated from protobuf message protos.GetNotificationRequest
 */
export interface GetNotificationRequest {
    /**
     * @generated from protobuf field: string notification_id = 1;
     */
    notificationId: string;
}
/**
 * @generated from protobuf message protos.GetNotificationResponse
 */
export interface GetNotificationResponse {
    /**
     * @generated from protobuf field: protos.NotificationConfig notification = 1;
     */
    notification?: NotificationConfig;
}
/**
 * @generated from protobuf message protos.AttachNotificationRequest
 */
export interface AttachNotificationRequest {
    /**
     * @generated from protobuf field: string notification_id = 1;
     */
    notificationId: string;
    /**
     * @generated from protobuf field: string pipeline_id = 2;
     */
    pipelineId: string;
}
/**
 * @generated from protobuf message protos.TestRequest
 */
export interface TestRequest {
    /**
     * @generated from protobuf field: string input = 1;
     */
    input: string;
}
/**
 * @generated from protobuf message protos.TestResponse
 */
export interface TestResponse {
    /**
     * @generated from protobuf field: string output = 2;
     */
    output: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class GetAllRequest$Type extends MessageType<GetAllRequest> {
    constructor() {
        super("protos.GetAllRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetAllRequest
 */
export const GetAllRequest = new GetAllRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetAllResponse$Type extends MessageType<GetAllResponse> {
    constructor() {
        super("protos.GetAllResponse", [
            { no: 1, name: "live", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => LiveInfo },
            { no: 2, name: "audiences", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Audience },
            { no: 3, name: "pipelines", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => PipelineInfo } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetAllResponse
 */
export const GetAllResponse = new GetAllResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetPipelinesRequest$Type extends MessageType<GetPipelinesRequest> {
    constructor() {
        super("protos.GetPipelinesRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetPipelinesRequest
 */
export const GetPipelinesRequest = new GetPipelinesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetPipelinesResponse$Type extends MessageType<GetPipelinesResponse> {
    constructor() {
        super("protos.GetPipelinesResponse", [
            { no: 1, name: "pipelines", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Pipeline }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetPipelinesResponse
 */
export const GetPipelinesResponse = new GetPipelinesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetPipelineRequest$Type extends MessageType<GetPipelineRequest> {
    constructor() {
        super("protos.GetPipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetPipelineRequest
 */
export const GetPipelineRequest = new GetPipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetPipelineResponse$Type extends MessageType<GetPipelineResponse> {
    constructor() {
        super("protos.GetPipelineResponse", [
            { no: 1, name: "pipeline", kind: "message", T: () => Pipeline }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetPipelineResponse
 */
export const GetPipelineResponse = new GetPipelineResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreatePipelineRequest$Type extends MessageType<CreatePipelineRequest> {
    constructor() {
        super("protos.CreatePipelineRequest", [
            { no: 1, name: "pipeline", kind: "message", T: () => Pipeline }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.CreatePipelineRequest
 */
export const CreatePipelineRequest = new CreatePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdatePipelineRequest$Type extends MessageType<UpdatePipelineRequest> {
    constructor() {
        super("protos.UpdatePipelineRequest", [
            { no: 1, name: "pipeline", kind: "message", T: () => Pipeline }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.UpdatePipelineRequest
 */
export const UpdatePipelineRequest = new UpdatePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeletePipelineRequest$Type extends MessageType<DeletePipelineRequest> {
    constructor() {
        super("protos.DeletePipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.DeletePipelineRequest
 */
export const DeletePipelineRequest = new DeletePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AttachPipelineRequest$Type extends MessageType<AttachPipelineRequest> {
    constructor() {
        super("protos.AttachPipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.AttachPipelineRequest
 */
export const AttachPipelineRequest = new AttachPipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DetachPipelineRequest$Type extends MessageType<DetachPipelineRequest> {
    constructor() {
        super("protos.DetachPipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.DetachPipelineRequest
 */
export const DetachPipelineRequest = new DetachPipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PausePipelineRequest$Type extends MessageType<PausePipelineRequest> {
    constructor() {
        super("protos.PausePipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.PausePipelineRequest
 */
export const PausePipelineRequest = new PausePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ResumePipelineRequest$Type extends MessageType<ResumePipelineRequest> {
    constructor() {
        super("protos.ResumePipelineRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.ResumePipelineRequest
 */
export const ResumePipelineRequest = new ResumePipelineRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateNotificationRequest$Type extends MessageType<CreateNotificationRequest> {
    constructor() {
        super("protos.CreateNotificationRequest", [
            { no: 1, name: "notification", kind: "message", T: () => NotificationConfig }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.CreateNotificationRequest
 */
export const CreateNotificationRequest = new CreateNotificationRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateNotificationRequest$Type extends MessageType<UpdateNotificationRequest> {
    constructor() {
        super("protos.UpdateNotificationRequest", [
            { no: 1, name: "notification", kind: "message", T: () => NotificationConfig }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.UpdateNotificationRequest
 */
export const UpdateNotificationRequest = new UpdateNotificationRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteNotificationRequest$Type extends MessageType<DeleteNotificationRequest> {
    constructor() {
        super("protos.DeleteNotificationRequest", [
            { no: 1, name: "notification_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.DeleteNotificationRequest
 */
export const DeleteNotificationRequest = new DeleteNotificationRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetNotificationsRequest$Type extends MessageType<GetNotificationsRequest> {
    constructor() {
        super("protos.GetNotificationsRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetNotificationsRequest
 */
export const GetNotificationsRequest = new GetNotificationsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetNotificationsResponse$Type extends MessageType<GetNotificationsResponse> {
    constructor() {
        super("protos.GetNotificationsResponse", [
            { no: 1, name: "notifications", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => NotificationConfig } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetNotificationsResponse
 */
export const GetNotificationsResponse = new GetNotificationsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetNotificationRequest$Type extends MessageType<GetNotificationRequest> {
    constructor() {
        super("protos.GetNotificationRequest", [
            { no: 1, name: "notification_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetNotificationRequest
 */
export const GetNotificationRequest = new GetNotificationRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetNotificationResponse$Type extends MessageType<GetNotificationResponse> {
    constructor() {
        super("protos.GetNotificationResponse", [
            { no: 1, name: "notification", kind: "message", T: () => NotificationConfig }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetNotificationResponse
 */
export const GetNotificationResponse = new GetNotificationResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class AttachNotificationRequest$Type extends MessageType<AttachNotificationRequest> {
    constructor() {
        super("protos.AttachNotificationRequest", [
            { no: 1, name: "notification_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.AttachNotificationRequest
 */
export const AttachNotificationRequest = new AttachNotificationRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TestRequest$Type extends MessageType<TestRequest> {
    constructor() {
        super("protos.TestRequest", [
            { no: 1, name: "input", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.TestRequest
 */
export const TestRequest = new TestRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TestResponse$Type extends MessageType<TestResponse> {
    constructor() {
        super("protos.TestResponse", [
            { no: 2, name: "output", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.TestResponse
 */
export const TestResponse = new TestResponse$Type();
/**
 * @generated ServiceType for protobuf service protos.External
 */
export const External = new ServiceType("protos.External", [
    { name: "GetAll", options: {}, I: GetAllRequest, O: GetAllResponse },
    { name: "GetPipelines", options: {}, I: GetPipelinesRequest, O: GetPipelinesResponse },
    { name: "GetPipeline", options: {}, I: GetPipelineRequest, O: GetPipelineResponse },
    { name: "CreatePipeline", options: {}, I: CreatePipelineRequest, O: StandardResponse },
    { name: "UpdatePipeline", options: {}, I: UpdatePipelineRequest, O: StandardResponse },
    { name: "DeletePipeline", options: {}, I: DeletePipelineRequest, O: StandardResponse },
    { name: "AttachPipeline", options: {}, I: AttachPipelineRequest, O: StandardResponse },
    { name: "DetachPipeline", options: {}, I: DetachPipelineRequest, O: StandardResponse },
    { name: "PausePipeline", options: {}, I: PausePipelineRequest, O: StandardResponse },
    { name: "ResumePipeline", options: {}, I: ResumePipelineRequest, O: StandardResponse },
    { name: "CreateNotification", options: {}, I: CreateNotificationRequest, O: StandardResponse },
    { name: "UpdateNotification", options: {}, I: UpdateNotificationRequest, O: StandardResponse },
    { name: "DeleteNotification", options: {}, I: DeleteNotificationRequest, O: StandardResponse },
    { name: "GetNotifications", options: {}, I: GetNotificationsRequest, O: GetNotificationsResponse },
    { name: "GetNotification", options: {}, I: GetNotificationRequest, O: GetNotificationResponse },
    { name: "AttachNotification", options: {}, I: AttachNotificationRequest, O: StandardResponse },
    { name: "Test", options: {}, I: TestRequest, O: TestResponse }
]);
