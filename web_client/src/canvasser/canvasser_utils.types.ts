// Brian Taylor Vann

// canvasser constants
export const CREATE_SUBSCRIPTION = "CREATE_SUBSCRIPTION";
export const REMOVE_SUBSCRIPTION = "REMOVE_SUBSCRIPTION";
export const CREATE_CHANNEL = "CREATE_CHANNEL";
export const REMOVE_CHANNEL = "REMOVE_CHANNEL";

// canvasser types
export type ChannelType<T> = {
  channelId: keyof T;
  canvas: HTMLCanvasElement;
};

export type ChannelMapType<T> = {
  [K in keyof T]?: ChannelType<T>;
};

// canvasser channel action creator types
export type CreateChannelActionType<T> = {
  type: typeof CREATE_CHANNEL;
  payload: {
    channelId: keyof T;
    canvas: HTMLCanvasElement;
  };
};

export type RemoveChannelActionType<T> = {
  type: typeof REMOVE_CHANNEL;
  payload: {
    channelId: keyof T;
  };
};

export type CanvasserChannelActionTypes<T> =
  | CreateChannelActionType<T>
  | RemoveChannelActionType<T>;

// canvasser subscription action creator types
export type SubscriptionRequestType<T> = {
  channelId: keyof T;
  callback: (action: CanvasserChannelActionTypes<T>) => void;
};

export type CreateSubscriptionActionType<T> = {
  type: typeof CREATE_SUBSCRIPTION;
  payload: SubscriptionRequestType<T>;
};

export type RemoveSubscriptionActionType<T> = {
  type: typeof REMOVE_SUBSCRIPTION;
  payload: {
    channelId: keyof T;
    stub: number;
  };
};

export type CanvasserSubscriptionActionTypes<T> =
  | CreateSubscriptionActionType<T>
  | RemoveSubscriptionActionType<T>;

// canvasser reducer types
export type SubscriptionsMapType<T> = {
  [key: number]: SubscriptionRequestType<T>;
};

export type SubscriptionChannelMapType<T> = {
  [K in keyof T]?: SubscriptionsMapType<T>;
};

// canvasser state
export type CanvasserStateType<T> = Readonly<{
  channels: ChannelMapType<T>;
  subscriptionChannels: SubscriptionChannelMapType<T>;
}>;

// canvasser interface
export type CanvasserDispatchType<T> = (
  action: CanvasserChannelActionTypes<T>,
) => void;
export type CanvasserGetStateType<T> = () => CanvasserStateType<T>;
export type CanvasserSubscriptionType<T> = (
  channelId: keyof T,
  callback: (action: CanvasserChannelActionTypes<T>) => void,
) => void;

export type CanvasserInterfaceType<T> = {
  dispatch: CanvasserDispatchType<T>;
  getState: CanvasserGetStateType<T>;
  subscribe: CanvasserSubscriptionType<T>;
};

// action args
export type CanvasserDefaultActionArgsType<T> = { channelId: keyof T };
export type CanvasserCreateChannelActionArgsType<T> = {
  channelId: keyof T;
  canvas: HTMLCanvasElement;
};
export type CanvasserRemoveSubscriptionActionArgsType<T> = {
  channelId: keyof T;
  stub: number;
};
