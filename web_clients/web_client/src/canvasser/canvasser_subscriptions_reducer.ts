import {
  CanvasserSubscriptionActionTypes,
  SubscriptionChannelMapType,
  CREATE_SUBSCRIPTION,
  REMOVE_SUBSCRIPTION,
  SubscriptionsMapType,
} from "./canvasser_utils.types";

type CanvasserSubscriptionReducerArgs<T> = {
  subscriptionChannels: SubscriptionChannelMapType<T>;
  action: CanvasserSubscriptionActionTypes<T>;
  subscriptionStub: number;
};

// update correct subscription format
function canvasserSubscriptionsReducer<T>({
  subscriptionChannels,
  action,
  subscriptionStub,
}: CanvasserSubscriptionReducerArgs<T>): SubscriptionChannelMapType<T> {
  switch (action.type) {
    case CREATE_SUBSCRIPTION: {
      const { channelId } = action.payload;
      if (subscriptionChannels[channelId] != null) {
        return subscriptionChannels;
      }

      const channelSubscriptions = subscriptionChannels[channelId];

      const modifiedSubscriptions = {
        ...subscriptionChannels,
        ...{
          [channelId]: {
            ...channelSubscriptions,
            ...{ [subscriptionStub]: action.payload },
          },
        },
      };

      return modifiedSubscriptions;
    }
    case REMOVE_SUBSCRIPTION: {
      const { channelId, stub } = action.payload;
      if (subscriptionChannels[channelId] == null) {
        return subscriptionChannels;
      }

      const subscriptions = subscriptionChannels[channelId];

      const modifiedSubscriptions: SubscriptionsMapType<T> = {};
      const numStub = stub.toString();
      for (const subStub in subscriptions) {
        if (subStub === numStub) {
          continue;
        }
        subscriptions[stub] = subscriptions[stub];
      }

      return {
        ...subscriptionChannels,
        ...{ [channelId]: modifiedSubscriptions },
      };
    }
  }
  return subscriptionChannels;
}

export { canvasserSubscriptionsReducer };
