// Brian Taylor Vann
// taylorvann dot com

import {
  HALT_LOOP,
  HALTED,
  LooperActionTypes,
  REQUESTING_FRAMES,
  START_LOOP,
  LooperStateType,
  UNREQUESTED,
  RESET_LOOP,
  LooperCallbackPayloadType,
} from "./looper_utils.types";

const getDefaultTimeValues = (): LooperCallbackPayloadType => {
  const now = window.performance.now();

  return {
    loopStartTime: now,
    previousTime: now,
    currentTime: now,
    deltaTime: 0,
  };
};

const getDefaultStateValues = (): LooperStateType => {
  return {
    requestState: UNREQUESTED,
    requestAnimationFrameStub: null,
    timeDetails: getDefaultTimeValues(),
  };
};

const looperReducer = (
  state: LooperStateType | null,
  action: LooperActionTypes,
) => {
  if (state == null) {
    return getDefaultStateValues();
  }
  switch (action.type) {
    case START_LOOP:
      state.requestState = REQUESTING_FRAMES;
      return state;
    case HALT_LOOP:
      state.requestState = HALTED;
      return state;
    case RESET_LOOP:
      state.requestState = UNREQUESTED;
      state.requestAnimationFrameStub = null;
      state.timeDetails = getDefaultTimeValues();
    default:
      return state;
  }
};

export { looperReducer };
