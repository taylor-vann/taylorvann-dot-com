// Brian Taylor Vann
// taylorvann dot com

import { Renderer } from "three";
import {
  RendererChannelActionTypes,
  RendererInterfaceType,
  RendererStateType,
  ChannelMapType,
} from "./renderer_utils.types";

export const PIXEL_RATIO = window.devicePixelRatio;

import { rendererChannelReducer } from "./renderer_channel_reducer";

function updateRendererDimensions<T>(
  canvas: HTMLCanvasElement,
  renderer: Renderer,
) {
  const { clientWidth, clientHeight } = canvas;

  const width = Math.floor(clientWidth * PIXEL_RATIO);
  const height = Math.floor(clientHeight * PIXEL_RATIO);

  canvas.width = width;
  canvas.height = height;
  renderer.setSize(width, height, false);
}

function updateChannel<T>(
  channels: ChannelMapType<T>,
  action: RendererChannelActionTypes<T>,
) {
  const { channelId } = action.payload;

  const channel = channels[channelId];
  if (channel == null) {
    return;
  }

  const canvas = channel.canvas;
  const renderer = channel.renderer;
  if (canvas == null || renderer == null) {
    return;
  }

  // update canvas size and stuff
  updateRendererDimensions(canvas, renderer);
}

function createRendererService<T>(): RendererInterfaceType<T> {
  let channels: ChannelMapType<T> = {};

  function dispatch(action: RendererChannelActionTypes<T>): void {
    if (action.type === "RENDER_CHANNEL") {
      // render channel
      return;
    }
    if (action.type === "UPDATE_CHANNEL") {
      // update dimensions
      updateChannel(channels, action);
      return;
    }
    channels = rendererChannelReducer(channels, action);
  }

  function getState(): RendererStateType<T> {
    return { channels };
  }

  return {
    dispatch,
    getState,
  };
}

export { createRendererService, updateRendererDimensions };
