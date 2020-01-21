// Brian Taylor Vann
// taylorvann dot com

// This is a rare occurance where xtend the global interface and "any" can be
// used for browser detection through feature detection.

// We will never use these features directly and we will never want to
// We only need to know if they exist for polyfills and web standards.

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
} from "./browser_details.types";

import {
  getBlinkStatus,
  getChromeStatus,
  getEdgeStatus,
  getInternetExplorerStatus,
  getMozillaStatus,
  getOperaStatus,
  getSafariStatus,
} from "./browser_details_utils";

// Opera
const isOpera: boolean = getOperaStatus();

// Firefox
const isFirefox: boolean = getMozillaStatus();

// Safari
const isSafari: boolean = getSafariStatus();

// Internet Explorer
const isIE: boolean = getInternetExplorerStatus();

// Edge
const isEdge: boolean = getEdgeStatus(isIE);

// Chrome
const isChrome: boolean = getChromeStatus();

// Blink
const isBlink = getBlinkStatus(isChrome, isOpera);

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
