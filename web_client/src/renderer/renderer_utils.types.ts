// Brian Taylor Vann

import { Renderer } from "three";

// Renderer constants
export const CREATE_SUBSCRIPTION = "CREATE_SUBSCRIPTION";
export const REMOVE_SUBSCRIPTION = "REMOVE_SUBSCRIPTION";
export const CREATE_CHANNEL = "CREATE_CHANNEL";
export const UPDATE_CHANNEL = "UPDATE_CHANNEL";
export const REMOVE_CHANNEL = "REMOVE_CHANNEL";
export const RENDER_CHANNEL = "RENDER_CHANNEL";

// Renderer types
export type ChannelType<K> = {
  channelId: K;
  canvas: HTMLCanvasElement;
  renderer: THREE.Renderer;
};

export type ChannelMapType<T> = {
  [K in keyof T]?: ChannelType<K>;
};

// Renderer channel action creator types
export type CreateChannelActionType<T> = {
  type: typeof CREATE_CHANNEL;
  payload: {
    channelId: keyof T;
    canvas: HTMLCanvasElement;
  };
};

export type UpdateChannelActionType<T> = {
  type: typeof UPDATE_CHANNEL;
  payload: {
    channelId: keyof T;
  };
};

export type RemoveChannelActionType<T> = {
  type: typeof REMOVE_CHANNEL;
  payload: {
    channelId: keyof T;
  };
};

export type RenderChannelActionType<T> = {
  type: typeof RENDER_CHANNEL;
  payload: {
    channelId: keyof T;
  };
};

export type RendererChannelActionTypes<T> =
  | CreateChannelActionType<T>
  | UpdateChannelActionType<T>
  | RemoveChannelActionType<T>
  | RenderChannelActionType<T>;

// Renderer subscription action creator types

// Renderer state
export type RendererStateType<T> = Readonly<{
  channels: ChannelMapType<T>;
}>;

// Renderer interface
export type RendererDispatchType<T> = (
  action: RendererChannelActionTypes<T>,
) => void;
export type RendererGetStateType<T> = () => RendererStateType<T>;

export type RendererInterfaceType<T> = {
  dispatch: RendererDispatchType<T>;
  getState: RendererGetStateType<T>;
  // subscribe: RendererSubscriptionType<T>;
};

// action args
export type RendererDefaultActionArgsType<T> = { channelId: keyof T };
export type RendererCreateChannelActionArgsType<T> = {
  channelId: keyof T;
  canvas: HTMLCanvasElement;
};
