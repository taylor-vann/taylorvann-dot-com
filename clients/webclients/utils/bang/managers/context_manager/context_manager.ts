// brian talor vann

// Bang context manager
// logic_structures return contexts

// we don't want direct access to contexts
// key value storage meets graph storage

// we can simplify things, it's a tree. no cycles.

// N Node
// A Attributables

// P Params

import { Context } from "../../references/context";

interface ContextManagerBase {
  addContext: (structureID: number) => number;
  updateContext: <P>(contextID: number, params?: P) => {};
  removeContext: (stubID: number) => void;
  getSiblings: <A, P, R>(stubID: number) => Context<A, P, R>;
}

// we can control max render
// helps track current amount of nodes (we should track number of nodes)
// recycle stubIDs, object never grows larger than N defined references

// create class for this

export { ContextManagerBase };
