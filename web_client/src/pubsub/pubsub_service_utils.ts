// Brian Taylor Vann
// taylorvann dot com

// PubSub Service Utils
// Create a message hub and keep parts of an application separated
// through messaging.
// Provide a map of {[channels]: actionTypes}.

import {
  AddSubToChannelArgsType,
  PublishToAllSubsArgsType,
  PubSubChannelMapType,
  PubSubInterfaceType,
  PubSubMapType,
  RemoveSubToChannelArgsType,
  UnsubscribeType,
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

  const modifiedChannelSubs: PubSubChannelMapType<T[keyof T]> = {};
  const subStubStr = subscriptionStub.toString();
  for (let stub in channelSubs) {
    if (subStubStr === stub) {
      continue;
    }
    modifiedChannelSubs[stub] = channelSubs[stub];
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
  let subscriptionStub: number = -1;
  let pubsubs = {};

  return {
    subscribe: (channel, callback): UnsubscribeType => {
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
    dispatch: (channel, action) => {
      publishToAllSubscriptions({ pubsubs, channel, action });
    },
    getState: () => {
      return pubsubs;
    },
  };
}

export { createPubSubService };
