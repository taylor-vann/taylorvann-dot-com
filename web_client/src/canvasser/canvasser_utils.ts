// Brian Taylor Vann
// taylorvann.com

import {
  CanvasserChannelActionTypes,
  CanvasserInterfaceType,
  CanvasserStateType,
  CanvasserSubscriptionActionTypes,
  ChannelMapType,
  SubscriptionChannelMapType,
  SubscriptionsMapType,
} from "./canvasser_utils.types";

import { createSubscription } from "./canvasser_actions";

import { canvasserSubscriptionsReducer } from "./canvasser_subscriptions_reducer";
import { canvasserChannelReducer } from "./canvasser_channel_reducer";

function dispatchToAllChannelSubscriptions<T>(
  subscriptionChannels: SubscriptionChannelMapType<T>,
  action: CanvasserChannelActionTypes<T>,
): void {
  const { channelId } = action.payload;
  const subscriptionsMap: SubscriptionsMapType<T> | undefined =
    subscriptionChannels[channelId];
  if (subscriptionsMap == undefined) {
    return;
  }
  for (const subscriptionsStub in subscriptionsMap) {
    const subsciptionRequest = subscriptionsMap[subscriptionsStub];
    subsciptionRequest.callback(action);
  }
}

function canvasserDispatch<T>(
  channels: ChannelMapType<T>,
  subscriptionChannels: SubscriptionChannelMapType<T>,
  action: CanvasserChannelActionTypes<T>,
): ChannelMapType<T> {
  const modifiedChannels = canvasserChannelReducer(channels, action);
  dispatchToAllChannelSubscriptions(subscriptionChannels, action);

  return modifiedChannels;
}

function createCanvasserService<T>(): CanvasserInterfaceType<T> {
  let channels: ChannelMapType<T> = {};
  let subscriptionChannels = {};
  let subscriptionStub = 0;

  function dispatch(action: CanvasserChannelActionTypes<T>): void {
    channels = canvasserDispatch(channels, subscriptionChannels, action);
  }

  function getState(): CanvasserStateType<T> {
    return { channels, subscriptionChannels };
  }

  // we've subscribed an agnostic callback
  function subscribe(
    channel: keyof T,
    callback: (action: CanvasserChannelActionTypes<T>) => void,
  ): void {
    subscriptionChannels = canvasserSubscriptionsReducer({
      subscriptionChannels,
      action: createSubscription({ channelId: channel, callback }),
      subscriptionStub,
    });

    subscriptionStub += 1;
  }

  return {
    dispatch,
    getState,
    subscribe,
  };
}

export { createCanvasserService };
