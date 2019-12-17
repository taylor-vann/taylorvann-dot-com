//Brian Taylor Vann
// taylorvann dot com

// mousedown
// keep track of pointer (mouse, touch, pen, tablet...) state

"strict";

import { PointerStateMap } from "./mousedown_types";

import {
  createPointerStateEntry,
  removePointerStateEntry,
  updatePointerStateEntry,
  getMouseDownAction,
} from "./mousedown_utils";

export interface MouseDownClass {}

export type MouseDownInterfaceType = {
  move: (e: PointerEvent) => void;
  down: (e: PointerEvent) => void;
  up: (e: PointerEvent) => void;
  remove: (e: PointerEvent) => void;
  removeAllPointers: (e: PointerEvent) => void;
  getMouseDownState: () => PointerStateMap;
};

function createMouseDown() {
  const pointerStateMap: PointerStateMap = {};

  return {
    move: (e: PointerEvent): void => {
      getMouseDownAction(e, this._pointerStateMap);
      this._pointerStateMap = updatePointerStateEntry(e, this._pointerStateMap);
    },

    down: (e: PointerEvent): void => {
      getMouseDownAction(e, this._pointerStateMap);
      this._pointerStateMap = createPointerStateEntry(e, this._pointerStateMap);
    },

    up: (e: PointerEvent): void => {
      getMouseDownAction(e, this._pointerStateMap);
      this._pointerStateMap = updatePointerStateEntry(e, this._pointerStateMap);
    },

    remove: (e: PointerEvent): void => {
      getMouseDownAction(e, this._pointerStateMap);
      this._pointerStateMap = removePointerStateEntry(e, this._pointerStateMap);
    },

    removeAllPointers: (): void => {
      this._pointerStateMap = {};
    },

    getMouseDownState: (): Readonly<PointerStateMap> => {
      return Object.freeze({ ...this._pointerStateMap });
    },
  };
}

class MouseDown implements MouseDownClass {
  private _pointerStateMap: PointerStateMap = {};

  move(e: PointerEvent): void {
    getMouseDownAction(e, this._pointerStateMap);
    this._pointerStateMap = updatePointerStateEntry(e, this._pointerStateMap);
  }

  down(e: PointerEvent): void {
    getMouseDownAction(e, this._pointerStateMap);
    this._pointerStateMap = createPointerStateEntry(e, this._pointerStateMap);
  }

  up(e: PointerEvent): void {
    getMouseDownAction(e, this._pointerStateMap);
    this._pointerStateMap = updatePointerStateEntry(e, this._pointerStateMap);
  }

  remove(e: PointerEvent): void {
    getMouseDownAction(e, this._pointerStateMap);
    this._pointerStateMap = removePointerStateEntry(e, this._pointerStateMap);
  }

  removeAllPointers(): void {
    this._pointerStateMap = {};
  }

  getMouseDownState(): Readonly<PointerStateMap> {
    return Object.freeze({ ...this._pointerStateMap });
  }
}

export { createMouseDown };
