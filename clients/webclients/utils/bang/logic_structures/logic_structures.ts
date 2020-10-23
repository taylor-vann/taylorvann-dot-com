// brian taylor vann

type LogicStructure<C, K> = <P>(params: P, ctx: K) => C;
type LogicStructureMap<C, K> = Record<number, LogicStructure<C, K>>;

type CreateLogicStructure = <C, K>(
  structure: LogicStructure<C, K>
) => LogicStructure<C, K>;

const CreateLogicStructure: CreateLogicStructure = <C, K>(structureFunc) => {
  // register structure

  // now we have functionID

  const curry = <P>(params: P) => {
    // get context (simple stub)
    // set context as current

    // that way when structureFunc is called, context is a number available
    const ctx = {
      bang: () => {},
      onDisconnected: () => {},
    };
    // the structure is reused call a render
    return structureFunc(params, ctx);
  };

  return curry;
};

export { LogicStructure, LogicStructureMap };
