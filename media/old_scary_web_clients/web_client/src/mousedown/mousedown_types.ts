// Brian Taylor Vann
// taylorvann dot com

export type Vector2d = {
  x: number;
  y: number;
};

export type Radians3d = {
  alpha: number;
  beta: number;
};

export type PointerStateActions = "down" | "up" | "hover" | "drag";
export type PointerStateBasicEntry = {
  button: number;
  buttons: number;
  id: number;
  movement: Vector2d;
  pagePosition: Vector2d;
  pointerType: string;
  position: Vector2d;
  pressure: number;
  screenPosition: Vector2d;
  tiltRadians: Radians3d;
  timestamp: number;
};

type Buttons = number;
export type PointerStatePressedButtons = {
  [button: number]: Buttons;
};

export type PointerStateBasicEntryMap = {
  [timestamp: number]: PointerStateBasicEntry;
};

export type PointerStateEntry = {
  currentEntry: PointerStateBasicEntry;
  gesturedEntries?: {
    [timestamp: number]: PointerStateBasicEntry;
  };
};

export type PointerStateMap = {
  [id: number]: PointerStateEntry;
};
