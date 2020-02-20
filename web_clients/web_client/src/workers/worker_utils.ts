// Brian Taylor Vann

// DANGEROUS: Ignore typsecript warnings about not being a module.

// The WorkerBridge import is a symbolic link to an actual worker
// created through Webpack as a blob.

// the next two warngins disable typescript and eslint ignores

// eslint-disable-next-line
// @ts-ignore
import * as WorkerBridge from "./worker_bridge.worker";

const createWorkerInterfaceInstance = (): { getState: () => void } => {
  console.log("attempting to create worker");
  console.log(WorkerBridge);

  // @ts-ignore
  const worker: Worker = new WorkerBridge();
  worker.addEventListener("message", (e: MessageEvent) => {
    console.log("received message:");
    console.log(e.data);
  });

  worker.postMessage({ hello: "hello, world" });

  return {
    getState: (): void => {
      console.log("yoooo its a worker");
    },
  };
};

export { createWorkerInterfaceInstance };
