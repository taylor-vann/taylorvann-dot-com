declare module "worker_bridge.worker" {
  class WebpackWorker extends Worker {
    constructor();
  }

  export default WebpackWorker;
}
