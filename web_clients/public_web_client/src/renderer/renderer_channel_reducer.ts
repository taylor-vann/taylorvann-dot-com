import {
  ChannelMapType,
  ChannelType,
  CREATE_CHANNEL,
  RendererChannelActionTypes,
  REMOVE_CHANNEL,
} from "./renderer_utils.types";

import * as THREE from "three";

import { updateRendererDimensions, PIXEL_RATIO } from "./renderer_utils";

function rendererChannelReducer<T>(
  channels: ChannelMapType<T>,
  action: RendererChannelActionTypes<T>,
): ChannelMapType<T> {
  switch (action.type) {
    case CREATE_CHANNEL: {
      const { canvas } = action.payload;
      const channelId = action.payload.channelId;
      if (channels[channelId] != null) {
        return channels;
      }

      canvas.width = Math.floor(canvas.clientWidth * PIXEL_RATIO);
      canvas.height = Math.floor(canvas.clientHeight * PIXEL_RATIO);

      // instantiate renderer
      const renderer = new THREE.WebGLRenderer({
        canvas: action.payload.canvas,
      });

      // update dimensions
      updateRendererDimensions(canvas, renderer);

      console.log("created element");
      console.log(canvas);
      console.log(renderer);
      const newChannel: ChannelType<keyof T> = {
        channelId,
        canvas,
        renderer,
      };

      return {
        ...channels,
        ...{ [channelId]: newChannel },
      };
    }
    case REMOVE_CHANNEL: {
      const { channelId } = action.payload;
      if (channels[channelId] != null) {
        return channels;
      }
      const modifiedChannels: ChannelMapType<T> = {};

      for (const channelKey in channels) {
        if (channelId === channelKey) {
          continue;
        }
        modifiedChannels[channelKey] = channels[channelKey];
      }

      return modifiedChannels;
    }
  }

  return channels;
}

export { rendererChannelReducer };
