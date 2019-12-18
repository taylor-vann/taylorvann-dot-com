// Brian Taylor Vann
// taylorvann dot com

// We can extend the global interface in typescript. This is a rare occurance
// where "any" can be used for broser detection through feature detection.

// We will never use these features directly and we will never want to
// upkeep with broser specific window properties. We only need to know
// if they exist.

declare global {
  interface Window {
    opr?: {
      addons: any;
      [key: string]: any;
    };
    opera?: any;
    InstallTrigger?: any;
    safari?: {
      pushNotification: any;
      [key: string]: any;
    };
    chrome?: {
      webstore: any;
      runtime: any;
      [key: string]: any;
    };
  }
  interface Document {
    documentMode?: any;
  }
}

import {
  BrowsersType,
  BrowserVerifyMapType,
  CHROME,
  FIREFOX,
  IE,
  OPERA,
  SAFARI,
} from "./device_details.types";

// Opera
var isOpera =
  (window.opr != null && window.opr.addons != null) ||
  window.opera != null ||
  navigator.userAgent.indexOf(" OPR/") >= 0;

// Firefox
var isFirefox = typeof window.InstallTrigger !== "undefined";

// Safari
const isWindowForSafari =
  window.safari != null && window.safari.pushNotification != null
    ? window.safari.pushNotification.toString() ===
      "[object SafariRemoteNotification]"
    : false;

//"[object HTMLElementConstructor]"
const isHTMLElementForSafari = /constructor/i.test(
  window.HTMLElement.toString(),
);

var isSafari = isHTMLElementForSafari || isWindowForSafari;

// Internet Explorer
var isIE = /*@cc_on!@*/ false || document["documentMode"] != null;

// Edge
var isEdge = isIE === false && window.StyleMedia != null;

// Chrome
var isChrome =
  window.chrome != null &&
  (window.chrome.webstore != null || window.chrome.runtime != null);

// Blink
var isBlink = (isChrome || isOpera) && window.CSS != null;

// Get the current browser as a string
let currentBrowser: BrowsersType = FIREFOX;
if (isFirefox) {
  currentBrowser = FIREFOX;
}
if (isChrome) {
  currentBrowser = CHROME;
}
if (isSafari) {
  currentBrowser = SAFARI;
}
if (isOpera) {
  currentBrowser = OPERA;
}
if (isIE) {
  currentBrowser = IE;
}

const browserMap: BrowserVerifyMapType = {
  CHROME: isChrome,
  FIREFOX: isFirefox,
  SAFARI: isSafari,
  OPERA: isOpera,
  EDGE: isEdge,
  IE: isIE,
  BLINK: isBlink,
};

export {
  isFirefox,
  isChrome,
  isSafari,
  isEdge,
  isIE,
  isOpera,
  isBlink,
  currentBrowser,
  browserMap,
};
