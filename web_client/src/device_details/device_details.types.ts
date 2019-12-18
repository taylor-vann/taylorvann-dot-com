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
