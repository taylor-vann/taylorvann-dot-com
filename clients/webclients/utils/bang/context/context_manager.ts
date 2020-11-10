// brian taylor vann
// context manager

// handles stateful updates to contexts

// U Unique Tag names
// N Node
// A Attributables
// P Params

import { ContextInterface, InterfaceBase } from "../references/context";

interface ConnectAction<N, A, P, R> {
  xcontext: ContextInterface<N, A, P, R>;
  params: P;
}
type UpdateAction<N, A, P, R> = ConnectAction<N, A, P, R>;
interface DisconnectAction<N, A, P, R> {
  xcontext: ContextInterface<N, A, P, R>;
  params: R;
}
interface ContextManagerBase {
  connect<N, A, P, R>(action: ConnectAction<N, A, P, R>): void;
  update<N, A, P, R>(action: UpdateAction<N, A, P, R>): void;
  disconnect<N, A, P, R>(action: DisconnectAction<N, A, P, R>): void;
}

class ContextManager implements ContextManagerBase {
  connect<N, A, P, R>(action: ConnectAction<N, A, P, R>) {
    // connect stuff
    // call onConnected
    // call update below()
    // call render with params
    // put structure render
  }

  update<N, A, P, R>(params: UpdateAction<N, A, P, R>) {
    // update stuff
    // call render with params
    // compare render results
  }

  disconnect<N, A, P, R>(params: DisconnectAction<N, A, P, R>) {
    // disconnect stuff
  }
}

const contextManager = new ContextManager();

export { contextManager };
