// Brian Taylor Vann
// taylorvann dot com

// This is a rare occurance where xtend the global interface and "any" can be
// used for browser detection through feature detection.

// We will never use these features directly and we will never want to
// upkeep browser specific window properties. We only need to know
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
  CHROME,
  FIREFOX,
  IE,
  OPERA,
  SAFARI,
  BLINK,
  EDGE,
} from "./device_details.types";

const safariRemoteNotification = "[object SafariRemoteNotification]";
const operaTag = " OPR/";

// Opera
var isOpera: boolean =
  (window.opr != null && window.opr.addons != null) ||
  window.opera != null ||
  navigator.userAgent.indexOf(operaTag) >= 0;

// Firefox
// Value is never used, but we assume Browser is Firefox in a cascading
// if statement later
var isFirefox: boolean = window.InstallTrigger != null;

// Safari
const isWindowFromSafari: boolean =
  window.safari != null && window.safari.pushNotification != null
    ? window.safari.pushNotification.toString() === safariRemoteNotification
    : false;

// test for [object HTMLElementConstructor]
const isHTMLElementFromSafari: boolean = /constructor/i.test(
  window.HTMLElement.toString(),
);

var isSafari: boolean = isHTMLElementFromSafari || isWindowFromSafari;

// Internet Explorer
var isIE: boolean = /*@cc_on!@*/ false || document.documentMode != null;

// Edge
var isEdge: boolean = isIE === false && window.StyleMedia != null;

// Chrome
var isChrome: boolean =
  window.chrome != null &&
  (window.chrome.webstore != null || window.chrome.runtime != null);

// Blink
var isBlink = (isChrome || isOpera) && window.CSS != null;

// Reduce to a single reference to Browser
// [TODO]: Find a way to map consts as key strings
let currentBrowser: BrowsersType = FIREFOX;
if (isChrome) {
  currentBrowser = CHROME;
}
if (isSafari) {
  currentBrowser = SAFARI;
}
if (isEdge) {
  currentBrowser = EDGE;
}
if (isOpera) {
  currentBrowser = OPERA;
}
if (isIE) {
  currentBrowser = IE;
}
if (isBlink) {
  currentBrowser = BLINK;
}

const platform: string = window.navigator.platform;
const oscpu: string = window.navigator.oscpu;
const languages: readonly string[] = window.navigator.languages;

export { currentBrowser, platform, oscpu, languages };
