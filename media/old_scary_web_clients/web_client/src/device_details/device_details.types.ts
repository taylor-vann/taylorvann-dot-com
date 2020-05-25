// Brian Taylor Vann
// taylorvann dot com

// Media Details
export type MediaDevicesMapType = {
  [deviceId: string]: MediaDeviceInfo;
};

export type MediaDevicesMapByGroupIdType = {
  [groupId: string]: {
    [kind: string]: {
      [deviceId: string]: MediaDeviceInfo;
    };
  };
};

export type MediaDevicesMapByKindType = {
  [kind: string]: {
    [deviceId: string]: MediaDeviceInfo;
  };
};

export type GetMediaDevicesRawArgsType = {
  mediaDevicesById: MediaDevicesMapType;
  mediaDevicesByGroupId: MediaDevicesMapByGroupIdType;
  mediaDevicesByKind: MediaDevicesMapByKindType;
};

export type GetMediaDevicesRawCallbackType = (
  args: GetMediaDevicesRawArgsType,
) => void;

export type MediaSubscriptionsMapType = {
  [key: number]: GetMediaDevicesRawCallbackType;
};
