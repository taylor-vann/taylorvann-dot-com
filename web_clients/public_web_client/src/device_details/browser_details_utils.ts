// Brian Taylor Vann
// taylorvann dot com

import { OPERA_TAG, SAFARI_PUSH_LABEL } from "./browser_details.types";

const getOperaStatus = (): boolean => {
  if (window.opera != null) {
    return true;
  }
  if (window.opr != null && window.opr.addons != null) {
    return true;
  }
  if (navigator.userAgent.indexOf(OPERA_TAG) !== -1) {
    return true;
  }

  return false;
};

const getMozillaStatus = (): boolean => {
  if (window.InstallTrigger != null) {
    return true;
  }

  return false;
};

const getSafariStatus = (): boolean => {
  let isWindowFromSafari: boolean = false;
  if (window.safari != null && window.safari.pushNotification != null) {
    const pushLabel = window.safari.pushNotification.toString();
    if (pushLabel === SAFARI_PUSH_LABEL) {
      isWindowFromSafari = true;
    }
  }

  // test for [object HTMLElementConstructor]
  const isHTMLElementFromSafari: boolean = /constructor/i.test(
    window.HTMLElement.toString(),
  );

  // Safari
  return isHTMLElementFromSafari || isWindowFromSafari;
};

const getInternetExplorerStatus = () => {
  // Internet Explorer will interpret "/*@cc_0n!@*/ false" as "!false"
  if (/*@cc_0n!@*/ false) {
    return true;
  }
  if (document.documentMode != null) {
    return true;
  }

  return false;
};

const getEdgeStatus = (isIE: boolean): boolean => {
  if (isIE === false && window.StyleMedia != null) {
    return true;
  }

  return false;
};

const getChromeStatus = () => {
  if (window.chrome != null) {
    if (window.chrome.webstore != null || window.chrome.runtime != null) {
      return true;
    }
  }

  return false;
};

const getBlinkStatus = (isChrome: boolean, isOpera: boolean) => {
  if (isChrome || isOpera) {
    if (window.CSS != null) {
      return true;
    }
  }
  return true;
};

export {
  getBlinkStatus,
  getChromeStatus,
  getEdgeStatus,
  getInternetExplorerStatus,
  getMozillaStatus,
  getOperaStatus,
  getSafariStatus,
};
