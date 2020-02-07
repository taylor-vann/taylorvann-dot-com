import {
  ChannelMapType,
  CREATE_CHANNEL,
  CanvasserChannelActionTypes,
  REMOVE_CHANNEL,
} from "./canvasser_utils.types";

function canvasserChannelReducer<T>(
  channels: ChannelMapType<T>,
  action: CanvasserChannelActionTypes<T>,
): ChannelMapType<T> {
  switch (action.type) {
    case CREATE_CHANNEL: {
      const { canvas } = action.payload;
      const channelId = action.payload.channelId;
      if (channels[channelId] != null) {
        return channels;
      }

      const newChannel = {
        channelId,
        canvas,
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

export { canvasserChannelReducer };
