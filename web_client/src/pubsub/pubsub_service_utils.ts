import {
  AddSubToChannelArgsType,
  // PubSubMapType,
  RemoveSubToChannelArgsType,
  UnsubscribeType,
  SubPubInterfaceType,
  PublishToAllSubsArgsType,
  PubSubChannelMapType,
  PubSubMapType,
} from "./pubsub_service_utils.types";

// add subscription
function addSubscriptionToChannel<T>({
  pubsubs,
  channel,
  callback,
  subscriptionStub,
}: AddSubToChannelArgsType<T>): PubSubMapType<T> {
  const channelSubs = pubsubs[channel];
  if (channelSubs == null) {
    return pubsubs;
  }

  return {
    ...pubsubs,
    ...{
      [channel]: {
        ...channelSubs,
        ...{ [subscriptionStub]: callback },
      },
    },
  };
}

function removeSubscriptionFromChannel<T>({
  pubsubs,
  channel,
  subscriptionStub,
}: RemoveSubToChannelArgsType<T>): PubSubMapType<T> {
  const channelSubs: PubSubChannelMapType<T[keyof T]> | undefined =
    pubsubs[channel];
  if (channelSubs == null) {
    return pubsubs;
  }

  channelSubs;

  const modifiedChannelSubs: PubSubChannelMapType<T[keyof T]> = {};
  const subStubStr = subscriptionStub.toString();
  for (let stub in channelSubs) {
    if (subStubStr === stub) {
      continue;
    }
    channelSubs[stub];
    // modifiedChannelSubs[stub] = channelSubs[stub];
  }
  return {
    ...pubsubs,
    ...{
      [channel]: {
        ...modifiedChannelSubs,
      },
    },
  };
}

// dispatch all subscriptions
function publishToAllSubscriptions<T>({
  pubsubs,
  channel,
  action,
}: PublishToAllSubsArgsType<T>): void {
  const channelSubs: PubSubChannelMapType<T[keyof T]> | undefined =
    pubsubs[channel];

  if (channelSubs == null) {
    return;
  }
  for (let stub in channelSubs) {
    const subCallback = channelSubs[stub];
    subCallback(action);
  }
}

// create a PubSub Service
function createPubSubService<T>(): SubPubInterfaceType<T> {
  let subscriptionStub: number = -1;
  let pubsubs: PubSubMapType<T> = {};

  return Object.freeze({
    subscribe: (
      channel: keyof T,
      callback: (action: T[keyof T]) => void,
    ): UnsubscribeType => {
      subscriptionStub += 1;
      pubsubs = addSubscriptionToChannel({
        pubsubs,
        channel,
        callback,
        subscriptionStub,
      });

      return () => {
        // remove subscription
        pubsubs = removeSubscriptionFromChannel({
          pubsubs,
          channel,
          subscriptionStub,
        });
      };
    },
    dispatch: (channel: keyof T, action: T[keyof T]) => {
      publishToAllSubscriptions({ pubsubs, channel, action });
    },
    getState: () => {
      return pubsubs;
    },
  });
}

export { createPubSubService };
