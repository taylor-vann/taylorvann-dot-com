import {
  RendererCreateChannelActionArgsType,
  RendererDefaultActionArgsType,
  CreateChannelActionType,
  RemoveChannelActionType,
  UpdateChannelActionType,
} from "./renderer_utils.types";

function createChannel<T>(
  payload: RendererCreateChannelActionArgsType<T>,
): CreateChannelActionType<T> {
  return {
    type: "CREATE_CHANNEL",
    payload,
  };
}

function updateChannel<T>(
  payload: RendererDefaultActionArgsType<T>,
): UpdateChannelActionType<T> {
  return {
    type: "UPDATE_CHANNEL",
    payload,
  };
}

function removeChannel<T>(
  payload: RendererDefaultActionArgsType<T>,
): RemoveChannelActionType<T> {
  return {
    type: "REMOVE_CHANNEL",
    payload,
  };
}

export { createChannel, updateChannel, removeChannel };
