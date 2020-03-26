// Brian Taylor Vann
// taylorvann dot com

// Actions - Start, Stop, Step
// Callback - a singular callback to handle message publications

// It does not make any calls to any interfaces directly, interfaces should
// subscribe.

import {
  HALT_LOOP,
  LooperActionTypes,
  LooperCallbackPayloadType,
  LooperCallbackType,
  LooperStateType,
  LooperSubscriptionsMapType,
  RESET_LOOP,
  START_LOOP,
  LooperInterfaceType,
} from "./looper_utils.types";

import { looperReducer } from "./looper_service_reducer";

function subscribeToLooper(
  subscriptions: LooperSubscriptionsMapType,
  callback: LooperCallbackType,
  stub: number,
): LooperSubscriptionsMapType {
  return {
    ...subscriptions,
    ...{ [stub]: callback },
  };
}

function unscubscribeToLooper(
  subscriptions: LooperSubscriptionsMapType,
  stub: number,
): LooperSubscriptionsMapType {
  const modifedSubscriptions: LooperSubscriptionsMapType = {};
  const stubAsString: string = stub.toString();
  for (let subscriptionId in subscriptions) {
    if (subscriptionId === stubAsString) {
      continue;
    }
    modifedSubscriptions[subscriptionId] = subscriptions[subscriptionId];
  }

  return modifedSubscriptions;
}

function dispatchToAllSubcriptions(
  subscriptions: LooperSubscriptionsMapType,
  payload: Readonly<LooperCallbackPayloadType>,
) {
  for (const subscriptionId in subscriptions) {
    const subscriptionCallback = subscriptions[subscriptionId];
    subscriptionCallback(payload);
  }
}

function updateLooperStateByFrame(state: LooperStateType, stub: number) {
  let { previousTime, currentTime } = state.timeDetails;
  state.timeDetails.previousTime = currentTime;
  state.timeDetails.currentTime = window.performance.now();
  state.timeDetails.deltaTime = currentTime - previousTime;
  state.requestAnimationFrameStub = stub;
}

function cloneState(state: LooperStateType): Readonly<LooperStateType> {
  const stateCopy = {
    ...state,
    ...{
      timeDetails: {
        ...state.timeDetails,
      },
    },
  };
  return stateCopy;
}

function createLooperInstance(): LooperInterfaceType {
  let subscriptionStub = -1;
  let subscriptions: LooperSubscriptionsMapType = {};

  let looperState: LooperStateType = looperReducer(null, { type: RESET_LOOP });

  // This is the function called every event animation frame request.
  function looperCallback() {
    // request a new animation frame for this function and its context
    const stub = window.requestAnimationFrame(looperCallback);

    // side effects
    updateLooperStateByFrame(looperState, stub);

    // send a copy of time details to subscriptions
    dispatchToAllSubcriptions(subscriptions, { ...looperState.timeDetails });
  }

  function reduceAndDispatch(action: LooperActionTypes) {
    looperReducer(looperState, action);
    if (action.type === START_LOOP) {
      looperCallback();
    }
    if (action.type === HALT_LOOP) {
      if (looperState.requestAnimationFrameStub != null) {
        window.cancelAnimationFrame(looperState.requestAnimationFrameStub);
      }
    }
  }

  return {
    subscribe: (callback: LooperCallbackType) => {
      subscriptionStub += 1;
      subscriptions = subscribeToLooper(
        subscriptions,
        callback,
        subscriptionStub,
      );
      return () => {
        subscriptions = unscubscribeToLooper(subscriptions, subscriptionStub);
      };
    },
    dispatch: (action: LooperActionTypes) => {
      reduceAndDispatch(action);
    },
    getState: () => {
      return cloneState(looperState);
    },
  };
}

export { createLooperInstance };
