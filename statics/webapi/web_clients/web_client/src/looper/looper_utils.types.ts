// Brian Taylor Vann

// Looper actions
export const START_LOOP = "START_LOOP";
export const HALT_LOOP = "HALT_LOOP";
export const RESET_LOOP = "RESET_LOOP";

export const UNREQUESTED = "UNREQUESTED";
export const REQUESTING_FRAMES = "REQUESTING_FRAMES";
export const HALTED = "HALTED";

export type LooperFrameRequestStateType =
  | typeof UNREQUESTED
  | typeof REQUESTING_FRAMES
  | typeof HALTED;

export type LooperStartActionType = {
  type: typeof START_LOOP;
};

export type LooperStopActionType = {
  type: typeof HALT_LOOP;
};

export type LooperResetActionType = {
  type: typeof RESET_LOOP;
};

export type LooperActionTypes =
  | LooperStartActionType
  | LooperStopActionType
  | LooperResetActionType;

export type LooperDispatchType = (action: LooperActionTypes) => void;

export type LooperCallbackPayloadType = {
  loopStartTime: number;
  previousTime: number;
  currentTime: number;
  deltaTime: number;
};

export type LooperCallbackType = (payload: LooperCallbackPayloadType) => void;

export type LooperSubscriptionsMapType = {
  [key: number]: LooperCallbackType;
};

export type LooperSubscriptionRequestType = (
  callback: LooperCallbackType,
) => void;

export type LooperStateType = {
  requestState: LooperFrameRequestStateType;
  timeDetails: LooperCallbackPayloadType;
  requestAnimationFrameStub?: number | null;
};

export type LooperInterfaceType = {
  subscribe: LooperSubscriptionRequestType;
  dispatch: LooperDispatchType;
  getState: () => Readonly<LooperStateType>;
};
