import {
  CanvasserCreateChannelActionArgsType,
  CanvasserDefaultActionArgsType,
  CanvasserRemoveSubscriptionActionArgsType,
  CreateChannelActionType,
  CreateSubscriptionActionType,
  RemoveChannelActionType,
  RemoveSubscriptionActionType,
  SubscriptionRequestType,
} from "./canvasser_utils.types";

function createChannel<T>(
  payload: CanvasserCreateChannelActionArgsType<T>,
): CreateChannelActionType<T> {
  return {
    type: "CREATE_CHANNEL",
    payload,
  };
}

function removeChannel<T>(
  payload: CanvasserDefaultActionArgsType<T>,
): RemoveChannelActionType<T> {
  return {
    type: "REMOVE_CHANNEL",
    payload,
  };
}

function createSubscription<T>(
  payload: SubscriptionRequestType<T>,
): CreateSubscriptionActionType<T> {
  return {
    type: "CREATE_SUBSCRIPTION",
    payload,
  };
}

function removeSubscription<T>(
  payload: CanvasserRemoveSubscriptionActionArgsType<T>,
): RemoveSubscriptionActionType<T> {
  return {
    type: "REMOVE_SUBSCRIPTION",
    payload,
  };
}

export { createChannel, removeChannel, createSubscription, removeSubscription };
