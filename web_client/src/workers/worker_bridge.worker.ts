// Brian Taylor Vann

// WorkerBridge is a messaging system between the main thread
// and auxillary threads.

// TODO DANGEROUS As of 2020, webpack asks developers to cast self as any

// eslint-disable-next-line
const ctx: Worker = self as any;

console.log("hello from a webworker");

ctx.addEventListener("message", (e: MessageEvent) => {
  console.log("yo we got your message!");
  console.log(e.data);
});

ctx.postMessage({ whats: "good" });
