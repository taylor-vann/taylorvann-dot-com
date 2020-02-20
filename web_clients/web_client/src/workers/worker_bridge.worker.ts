// Brian Taylor Vann

// WorkerBridge is a messaging system between the main thread
// and auxillary threads.

// eslint-disable-next-line
import * as Three from "three";

const lolz = new Three.Scene();

const ctx: Worker = self as any;

console.log("hello from a webworker");

ctx.addEventListener("message", (e: MessageEvent) => {
  console.log("yo we got your message!");
  console.log(e.data);
});

ctx.postMessage({ whats: lolz });
