// brian talor vann

// Bang context manager
// logic_structures return contexts

// we don't want direct access to contexts
// key value storage meets graph storage

// we can simplify things, it's a tree. no cycles.

// N Node
// A Attributables

// P Params

// updateContext returns siblings
interface ContextManagerBase<N> {
  addContext: (structureID: number) => number;
  updateContext: <P>(contextID: number, params?: P) => N[];
  removeContext: (stubID: number) => void;
}

// we can control max render
// helps track current amount of nodes (we should track number of nodes)
// recycle stubIDs, object never grows larger than N defined references

// create class for this

// addContext does what you think, references strucutre, return context ID

// updateContext starts a render chain hopefully
//   give contextID and params, returns structurerender, compare previous.
//   if previous strings are different, remove all and render new siblings
//   if injections are not the same,
//     update through previous injection map
//     depending on attribute
//   return previous renderResults

// removeContext
//
//   recursively get render chain, call "onDisconnected" on every substructure

export { ContextManagerBase };
