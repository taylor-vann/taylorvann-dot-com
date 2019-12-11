// Brian Taylor Vann
// taylorvann dot com

// PubSub Utility Types

// Map Types
export type PubSubChannelMapType<T> = { [stub: number]: (action: T) => void };

export type PubSubMapType<M> = {
  [P in keyof M]?: PubSubChannelMapType<M[P]>;
};

// PubSubInterface Types
export type UnsubscribeType = () => void;

export type SubscribeType<M> = (
  channel: keyof M,
  callback: (action: M[keyof M]) => void,
) => UnsubscribeType;

export type DispatchType<T> = (channel: keyof T, action: T[keyof T]) => void;

export type GetStateType<T> = () => PubSubMapType<T>;

export type SubPubInterfaceType<M> = Readonly<{
  subscribe: SubscribeType<M>;
  dispatch: DispatchType<M>;
  getState: GetStateType<M>;
}>;

// Utility Function Argument Types
export type AddSubToChannelArgsType<M> = {
  pubsubs: PubSubMapType<M>;
  channel: keyof M;
  callback: (action: M[keyof M]) => void;
  subscriptionStub: number;
};

export type RemoveSubToChannelArgsType<M> = {
  pubsubs: PubSubMapType<M>;
  channel: keyof M;
  subscriptionStub: number;
};

export type PublishToAllSubsArgsType<M> = {
  pubsubs: PubSubMapType<M>;
  channel: keyof M;
  action: M[keyof M];
};
