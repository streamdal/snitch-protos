// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "steps/encode.proto" (package "protos.steps", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
/**
 * WIP
 *
 * @generated from protobuf message protos.steps.EncodeStep
 */
export interface EncodeStep {
    /**
     * @generated from protobuf field: string id = 1;
     */
    id: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class EncodeStep$Type extends MessageType<EncodeStep> {
    constructor() {
        super("protos.steps.EncodeStep", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.steps.EncodeStep
 */
export const EncodeStep = new EncodeStep$Type();