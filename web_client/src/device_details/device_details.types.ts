// Browser Details

export const FIREFOX = "FIREFOX";
export const CHROME = "CHROME";
export const SAFARI = "SAFARI";
export const IE = "IE";
export const OPERA = "OPERA";
export const EDGE = "EDGE";
export const BLINK = "BLINK";

export type BrowsersType =
  | typeof FIREFOX
  | typeof CHROME
  | typeof SAFARI
  | typeof EDGE
  | typeof IE
  | typeof OPERA
  | typeof BLINK;

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
