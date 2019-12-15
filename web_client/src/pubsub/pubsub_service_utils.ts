import {
  AddSubToChannelArgsType,
  RemoveSubToChannelArgsType,
  UnsubscribeType,
  PubSubInterfaceType,
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
  let channelSubs = pubsubs[channel];

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
function publishToAllSubscriptions<T, R>({
  pubsubs,
  channel,
  action,
}: PublishToAllSubsArgsType<T>): void {
  console.log("publish to all subscriptions");
  const serviceSubs: PubSubChannelMapType<T[keyof T]> | undefined =
    pubsubs[channel];

  if (serviceSubs == null) {
    return;
  }
  console.log("channel subs:", serviceSubs);
  for (let stub in serviceSubs) {
    const subCallback = serviceSubs[stub];
    subCallback(action);
  }
}

// create a PubSub Service
function createPubSubService<T>(): PubSubInterfaceType<T> {
  console.log("created pub sub service");
  let subscriptionStub: number = -1;
  let pubsubs: PubSubMapType<T> = {};

  return {
    subscribe: (
      channel: keyof T,
      callback: (action: T[keyof T]) => void,
    ): UnsubscribeType => {
      console.log("pubsub subscribe called");
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

    // this isn't a callback, it's an action
    // we have a map of callback types,
    // we need a map of potential callback actions
    dispatch: (channel: keyof T, action: T[keyof T]) => {
      console.log("dispatched:", channel, action);
      publishToAllSubscriptions({ pubsubs, channel, action });
    },
    getState: () => {
      return pubsubs;
    },
  };
}

export { createPubSubService };
