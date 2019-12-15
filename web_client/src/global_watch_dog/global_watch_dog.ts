// Brian Taylor Vann
// taylorvann dot com

// Global Watch Dog
//
// There may be times when globals are required.
// In production, global usage must be restricted to the unavoidable.
// In development, usage outside of production should be restricted to debug.

declare global {
  interface Window {
    doggo: {};
  }
}

let createConsoleLog: () => void = () => {
  return;
};

if (process.env.NODE_ENV === "development") {
  createConsoleLog = (): void => {
    window.doggo = {};
  };
}

export { createConsoleLog };
