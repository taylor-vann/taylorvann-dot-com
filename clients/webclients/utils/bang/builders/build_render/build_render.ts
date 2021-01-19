// brian taylor vann
// build structure

import { Hooks } from "../../type_flyweight/hooks";
import { Template } from "../../type_flyweight/template";
import { Render, RenderStructure } from "../../type_flyweight/render";
import {
  AttributeAction,
  CloseNodeAction,
  CreateContentAction,
  CreateNodeAction,
  CreateSelfClosingNode,
  InjectedAttributeAction,
  ImplicitAttributeAction,
  Integrals,
} from "../../type_flyweight/integrals";

interface BuildRenderParams<N, A> {
  hooks: Hooks<N, A>;
  template: Template<A>;
  integrals: Integrals;
}

type BuildRender = <N, A>(params: BuildRenderParams<N, A>) => Render<N, A>;

// create node
type RenderCreateNode = <N, A>(
  ref: RenderStructure<N, A>,
  action: CreateNodeAction
) => void;
type RenderSelfClosingNode = <N, A>(
  ref: RenderStructure<N, A>,
  action: CreateSelfClosingNode
) => void;
type RenderCloseNode = <N, A>(
  ref: RenderStructure<N, A>,
  action: CloseNodeAction
) => void;
type RenderCreateContent = <N, A>(
  ref: RenderStructure<N, A>,
  action: CreateContentAction
) => void;
type RenderAppendAttribute = <N, A>(
  ref: RenderStructure<N, A>,
  action: AttributeAction
) => void;
type RenderInjectedAttribute = <N, A>(
  ref: RenderStructure<N, A>,
  action: InjectedAttributeAction
) => void;
type RenderImplicitAttribute = <N, A>(
  ref: RenderStructure<N, A>,
  action: ImplicitAttributeAction
) => void;

const createNode: RenderCreateNode = (ref, action) => {
  // create node
};
const createSelfClosingNode: RenderSelfClosingNode = (ref, action) => {
  // create node
  // pop from stack
};
const closeNode: RenderCloseNode = (ref, action) => {
  // if stack length < 1 and stack[stack.length - 1] === close node tag
  //   pop from stack
};
const createContent: RenderCreateContent = (ref, action) => {
  // create content
  //  if array distance === 0
  //    create text node
  //  else
  //    create injection
};
const appendAttribute: RenderAppendAttribute = (ref, action) => {};
const appendImplicitAttribute: RenderImplicitAttribute = (ref, action) => {};

const injectedAttribute: RenderInjectedAttribute = (ref, action) => {};

// close node

// add attribute

const buildRender: BuildRender = ({ hooks, template, integrals }) => {
  const ref = {
    hooks,
    template,
    render: {
      injections: {},
      siblings: [],
    },
    siblings: [],
    stack: [],
  };

  for (const integral of integrals) {
    if (integral.action === "CREATE_NODE") {
      createNode(ref, integral);
    }
    if (integral.action === "CREATE_SELF_CLOSING_NODE") {
      createSelfClosingNode(ref, integral);
    }
    if (integral.action === "CLOSE_NODE") {
      closeNode(ref, integral);
    }
    if (integral.action === "CREATE_CONTENT") {
      createContent(ref, integral);
    }
    if (integral.action === "APPEND_IMPLICIT_ATTRIBUTE") {
      appendImplicitAttribute(ref, integral);
    }
    if (integral.action === "APPEND_EXPLICIT_ATTRIBUTE") {
      appendAttribute(ref, integral);
    }
    if (integral.action === "APPEND_INJECTED_ATTRIBUTE") {
      injectedAttribute(ref, integral);
    }
  }
  // if stack.length === 0
  //
  // if create node and siblins
  //

  // if create node
  //   create node
  // if stack length === 0
  //   add to siblings

  // if content node, add content injection
  //   if string index === 0, we should add injection first
  //   than strings

  // if close node and last element is === close tag
  //   push off stack
  //
  // see if there's a starting content node
  return {
    injections: {},
    siblings: [],
  };
};

export { buildRender };
