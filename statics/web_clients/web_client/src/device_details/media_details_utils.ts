// Brian Taylor Vann
// taylorvann dot com

// We're gathering details about media capabilities on a device.
// Media Device queries are asynchronous.

import {
  GetMediaDevicesRawCallbackType,
  MediaDevicesMapByGroupIdType,
  MediaDevicesMapByKindType,
  MediaDevicesMapType,
  MediaSubscriptionsMapType,
} from "./device_details.types";

const getMediaDevicesRaw = async (
  responseCallback: GetMediaDevicesRawCallbackType,
) => {
  const mediaDevicesList = await navigator.mediaDevices.enumerateDevices();

  const mediaDevicesById: MediaDevicesMapType = {};
  for (const device of mediaDevicesList) {
    mediaDevicesById[device.deviceId] = device;
  }

  // build map by id
  const mediaDevicesByGroupId: MediaDevicesMapByGroupIdType = {};
  for (const device of mediaDevicesList) {
    const { groupId, deviceId, kind } = device;
    if (mediaDevicesByGroupId[groupId] == null) {
      mediaDevicesByGroupId[groupId] = {};
    }
    if (mediaDevicesByGroupId[groupId][kind] == null) {
      mediaDevicesByGroupId[groupId][kind] = {};
    }

    mediaDevicesByGroupId[groupId] = {
      ...mediaDevicesByGroupId[groupId],
      [kind]: {
        ...mediaDevicesByGroupId[groupId][kind],
        [deviceId]: device,
      },
    };
  }

  // build map by kind
  const mediaDevicesByKind: MediaDevicesMapByKindType = {};
  for (const device of mediaDevicesList) {
    const { kind, deviceId } = device;
    mediaDevicesByKind[kind] = {
      ...mediaDevicesByKind[deviceId],
      [deviceId]: device,
    };
  }

  responseCallback({
    mediaDevicesByGroupId,
    mediaDevicesById,
    mediaDevicesByKind,
  });
};

const createMediaDeviceDetailsInstance = () => {
  let subscriptionStub: number = -1;
  let subscriptionsMap: MediaSubscriptionsMapType = {};

  let mediaDevicesById = {};
  let mediaDevicesByGroupId = {};
  let mediaDevicesByKind = {};

  const getCurrentDeviceMaps = () => ({
    mediaDevicesById,
    mediaDevicesByGroupId,
    mediaDevicesByKind,
  });

  const subsribeCallback = (
    callback: GetMediaDevicesRawCallbackType,
  ): number => {
    subscriptionStub += 1;
    subscriptionsMap = {
      ...subscriptionsMap,
      ...{ [subscriptionStub]: callback },
    };
    return subscriptionStub;
  };

  const publishToAllSubscriptions: GetMediaDevicesRawCallbackType = payload => {
    for (const subscriptionId in subscriptionsMap) {
      const subscription = subscriptionsMap[subscriptionId];
      subscription(payload);
    }
  };

  const handleRequest: GetMediaDevicesRawCallbackType = payload => {
    // set local device state
    console.log(payload);
    mediaDevicesById = payload.mediaDevicesById;
    mediaDevicesByGroupId = payload.mediaDevicesByGroupId;
    mediaDevicesByKind = payload.mediaDevicesByKind;

    // dispatch all callbacks here.
    publishToAllSubscriptions(payload);
  };

  // on load, get media
  getMediaDevicesRaw(handleRequest);

  return {
    updateMediaDevices: () => {
      getMediaDevicesRaw(handleRequest);
    },
    getMediaDeviceMaps: () => {
      return getCurrentDeviceMaps();
    },
    subsribe: (callback: GetMediaDevicesRawCallbackType) => {
      return subsribeCallback(callback);
    },
  };
};

export { createMediaDeviceDetailsInstance };
