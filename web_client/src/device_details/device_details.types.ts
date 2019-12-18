export const FIREFOX = "FIREFOX";
export const CHROME = "CHROME";
export const SAFARI = "SAFARI";
export const IE = "IE";
export const OPERA = "OPERA";
export const EDGE = "EDGE";
export const BLINK = "BLINK";

export type BrowserMapType = {
  [FIREFOX]: typeof FIREFOX;
  [CHROME]: typeof CHROME;
  [SAFARI]: typeof SAFARI;
  [IE]: typeof IE;
  [OPERA]: typeof OPERA;
  [EDGE]: typeof EDGE;
  [BLINK]: typeof BLINK;
};

export type BrowsersType =
  | typeof FIREFOX
  | typeof CHROME
  | typeof SAFARI
  | typeof IE
  | typeof OPERA
  | typeof BLINK;

export type BrowserVerifyMapType = {
  [K in keyof BrowserMapType]: boolean;
};
