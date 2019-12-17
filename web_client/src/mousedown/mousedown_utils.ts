//  Brian Taylor Vann
// taylorvann dot com

"strict";

import {
  PointerStateBasicEntry,
  PointerStateEntry,
  PointerStateMap,
  Radians3d,
  Vector2d,
} from "./mousedown_types";

function degreesToRadians(degrees: number): number {
  return (degrees * Math.PI) / 180;
}

function getVector2d(x: number, y: number): Vector2d {
  return Object.freeze({
    x,
    y,
  });
}

function getRadians3d(alpha: number, beta: number): Radians3d {
  return Object.freeze({
    alpha,
    beta,
  });
}

function getPointerStateBasicEntry(
  e: PointerEvent,
): Readonly<PointerStateBasicEntry> {
  const movement = getVector2d(e.movementX, e.movementY);
  const nowUTCmS = Date.now();
  const pagePosition = getVector2d(e.pageX, e.pageY);
  const position = getVector2d(e.x, e.y);
  const screenPosition = getVector2d(e.screenX, e.screenY);
  const tiltAlpha = degreesToRadians(e.tiltX);
  const tiltBeta = degreesToRadians(e.tiltY);
  const tiltRadians = getRadians3d(tiltAlpha, tiltBeta);

  return Object.freeze({
    button: e.button,
    buttons: e.buttons,
    id: e.pointerId,
    movement,
    pagePosition,
    pointerType: e.pointerType,
    position,
    pressure: e.pressure,
    screenPosition,
    tiltRadians: tiltRadians,
    timestamp: nowUTCmS,
  });
}

function cascadeNewEntryIntoPreviousPointerStateEntry(
  pointerStateEntry: PointerStateEntry,
  eventEntry: PointerStateBasicEntry,
): PointerStateEntry {
  const { currentEntry, gesturedEntries } = pointerStateEntry;
  if (currentEntry.buttons === 0) {
    return Object.freeze({
      currentEntry: eventEntry,
      gesturedEntries,
    });
  }

  return Object.freeze({
    currentEntry: eventEntry,
    gesturedEntries: Object.freeze({
      ...gesturedEntries,
      ...{ [currentEntry.timestamp]: currentEntry },
    }),
  });
}

function hasCurrentEntryChanged(
  e: PointerEvent,
  currentEntry: PointerStateBasicEntry,
): boolean {
  return currentEntry != null && currentEntry.buttons != e.buttons;
}

function getMouseDownAction(
  e: PointerEvent,
  pointerStateMap: PointerStateMap,
): void {
  const pointerStateEntry = pointerStateMap[e.pointerId];
  if (pointerStateEntry == null) {
    console.log("initial pointer interaction");
    return;
  }

  const { currentEntry } = pointerStateEntry;

  if (e.button === -1) {
    if (e.buttons !== 0) {
      if (hasCurrentEntryChanged(e, currentEntry)) {
        console.log("button changed");
        console.log("previous:", currentEntry.buttons);
        console.log("current:", e.buttons);
      }
      console.log("drag", e.buttons);

      // drag events
    } else {
      console.log("hover", e.buttons);
      // hover events
    }
    // drag and hover will be here
    return;
  }

  if (hasCurrentEntryChanged(e, currentEntry)) {
    console.log("button change");
    const buttonDelta = e.buttons - currentEntry.buttons;
    if (buttonDelta > 0) {
      console.log("down:", e.buttons);
      // down events
    } else {
      console.log("up:", e.buttons);
      // down events
    }
  }
}

function createPointerStateEntry(
  e: PointerEvent,
  pointerStateMap: PointerStateMap,
): Readonly<PointerStateMap> {
  const eventEntry = getPointerStateBasicEntry(e);

  const pointerStateEntry: PointerStateEntry = {
    currentEntry: eventEntry,
  };

  return Object.freeze({
    ...pointerStateMap,
    ...{ [eventEntry.id]: pointerStateEntry },
  });
}

function updatePointerStateEntry(
  e: PointerEvent,
  pointerStateMap: PointerStateMap,
): Readonly<PointerStateMap> {
  const eventEntry = getPointerStateBasicEntry(e);
  const pointerStateEntry = pointerStateMap[eventEntry.id];
  if (pointerStateEntry == null) {
    console.log("new entry on update");
    return createPointerStateEntry(e, pointerStateMap);
  }

  // if currentEntry has any button pressed, add to gesturedEntries
  const modifiedPointerStateEntry = cascadeNewEntryIntoPreviousPointerStateEntry(
    pointerStateEntry,
    eventEntry,
  );

  return Object.freeze({
    ...pointerStateMap,
    ...{
      [eventEntry.id]: Object.freeze({
        ...modifiedPointerStateEntry,
      }),
    },
  });
}

function removePointerStateEntry(
  e: PointerEvent,
  pointerStateMap: PointerStateMap,
): Readonly<PointerStateMap> {
  const reducedPointerStateMap: PointerStateMap = {};
  const eventPointerId = e.pointerId.toString();

  for (let pointerId in pointerStateMap) {
    if (pointerId === eventPointerId) {
      continue;
    }

    reducedPointerStateMap[pointerId] = pointerStateMap[pointerId];
  }

  return Object.freeze(reducedPointerStateMap);
}

export {
  getPointerStateBasicEntry,
  createPointerStateEntry,
  removePointerStateEntry,
  updatePointerStateEntry,
  getMouseDownAction,
};
