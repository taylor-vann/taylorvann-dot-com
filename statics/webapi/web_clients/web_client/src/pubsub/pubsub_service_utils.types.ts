export type PubSubChannelMapType<T> = { [stub: number]: (action: T) => void };

export type PubSubMapType<M> = {
  [P in keyof M]?: PubSubChannelMapType<M[P]>;
};

export type CanvasserDispatchCallback<T> = (action: T) => void;

export type AddSubToChannelArgsType<M> = {
  pubsubs: PubSubMapType<M>;
  channel: keyof M;
  callback: CanvasserDispatchCallback<M[keyof M]>;
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

export type UnsubscribeType = () => void;

export type SubscribeType = <T>(
  channel: keyof T,
  callback: (action: T[keyof T]) => void,
) => UnsubscribeType;

export type DispatchType<T> = (channel: keyof T, action: T[keyof T]) => void;

export type GetStateType<T> = () => PubSubMapType<T>;

export type PubSubInterfaceType<M> = Readonly<{
  subscribe: SubscribeType;
  dispatch: DispatchType<M>;
  getState: GetStateType<M>;
}>;
