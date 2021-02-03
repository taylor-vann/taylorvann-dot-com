// brian taylor vann
// build render

import { Hooks } from "../../type_flyweight/hooks";
import { Template } from "../../type_flyweight/template";
import { RenderStructure } from "../../type_flyweight/render";
import { Context } from "../../type_flyweight/context";
import {
  CloseNodeAction,
  NodeAction,
  SelfClosingNodeAction,
  ExplicitAttributeAction,
  InjectedAttributeAction,
  ImplicitAttributeAction,
  Integrals,
  ContextInjectionAction,
  TextAction,
} from "../../type_flyweight/integrals";
import { getText, Vector } from "../../text_vector/text_vector";

interface BuildRenderParams<N, A> {
  hooks: Hooks<N, A>;
  template: Template<N, A>;
  integrals: Integrals;
}

type BuildRender = <N, A>(
  params: BuildRenderParams<N, A>
) => RenderStructure<N, A>;

type RenderNode = <N, A>(
  rs: RenderStructure<N, A>,
  integral: NodeAction | SelfClosingNodeAction
) => void;

type RenderTextNode = <N, A>(
  rs: RenderStructure<N, A>,
  integral: TextAction
) => void;

type RenderCloseNode = <N, A>(
  rs: RenderStructure<N, A>,
  integral: CloseNodeAction
) => void;

type RenderContentInjection = <N, A>(
  rs: RenderStructure<N, A>,
  integral: ContextInjectionAction
) => void;

type RenderAppendExplicitAttribute = <N, A>(
  rs: RenderStructure<N, A>,
  integral: ExplicitAttributeAction
) => void;

type RenderInjectedAttribute = <N, A>(
  rs: RenderStructure<N, A>,
  integral: InjectedAttributeAction
) => void;

type RenderImplicitAttribute = <N, A>(
  rs: RenderStructure<N, A>,
  integral: ImplicitAttributeAction
) => void;

type popSelfClosingNodes = <N, A>(rs: RenderStructure<N, A>) => void;

// add integral on to stack
const popSelfClosingNodes: popSelfClosingNodes = (rs) => {
  let parent = rs.stack[rs.stack.length - 1];
  while (
    parent !== undefined &&
    parent.kind === "NODE" &&
    parent.selfClosing === true
  ) {
    rs.stack.pop();
    parent = rs.stack[rs.stack.length - 1];
  }
};

const createTextNode: RenderTextNode = (rs, integral) => {
  // bounce through stack for self closing nodes
  popSelfClosingNodes(rs);

  const text = getText(rs.template, integral.textVector);
  if (text === undefined) {
    return;
  }

  const node = rs.hooks.createTextNode(text);
  const parentNode = rs.stack[rs.stack.length - 1]?.node;
  rs.hooks.appendDescendant(parentNode, node);

  if (rs.stack.length === 0) {
    rs.siblings.push(node);
  }
};

const createNode: RenderNode = (rs, integral) => {
  popSelfClosingNodes(rs);

  const tagName = getText(rs.template, integral.tagNameVector);
  if (tagName === undefined) {
    return;
  }

  const parent = rs.stack[rs.stack.length - 1];
  const node = rs.hooks.createNode(tagName);

  // get parent node
  const parentNode = parent?.node;
  rs.hooks.appendDescendant(parentNode, node);

  // add to silblings when stack is flat
  if (rs.stack.length === 0) {
    rs.siblings.push(node);
  }

  const selfClosing = integral.kind === "SELF_CLOSING_NODE";
  rs.stack.push({
    kind: "NODE",
    selfClosing,
    tagName,
    node,
  });
};

const closeNode: RenderCloseNode = (rs, integral) => {
  if (rs.stack.length === 0) {
    return;
  }

  popSelfClosingNodes(rs);

  const tagName = getText(rs.template, integral.tagNameVector);
  const nodeBit = rs.stack[rs.stack.length - 1];
  if (nodeBit.kind !== "NODE") {
    return;
  }

  if (nodeBit.tagName === tagName) {
    rs.stack.pop();
  }
};

const createContentInjection: RenderContentInjection = (rs, integral) => {
  popSelfClosingNodes(rs);

  // attach injection as Context
  const parent = rs.stack[rs.stack.length - 1]?.node;
  const left = rs.siblings[rs.stack.length - 1];
  const injection = rs.template.injections[integral.injectionID];
  if (injection === undefined) {
    return;
  }

  if (injection instanceof Context) {
    const siblings = injection.getSiblings();
    for (const sibling of siblings) {
      rs.hooks.appendDescendant(parent, sibling);
      if (rs.stack.length === 0) {
        rs.siblings.push(sibling);
      }
    }

    rs.injections[integral.injectionID] = {
      kind: "CONTEXT",
      params: { siblings, left, parent },
    };
    return;
  }

  // attach injection as content
  const text = String(injection);
  const textNode = rs.hooks.createTextNode(text);

  rs.hooks.appendDescendant(parent, textNode);
  if (rs.stack.length === 0) {
    rs.siblings.push(textNode);
  }

  rs.injections[integral.injectionID] = {
    kind: "CONTENT",
    params: { textNode, left, parent },
  };
};

const appendExplicitAttribute: RenderAppendExplicitAttribute = (
  rs,
  integral
) => {
  const node = rs.stack[rs.stack.length - 1].node;

  const attribute = getText(rs.template, integral.attributeVector);
  if (attribute === undefined) {
    return;
  }

  const value = getText(rs.template, integral.valueVector);
  if (value === undefined) {
    return;
  }

  rs.hooks.setAttribute({ node, attribute, value });
};

const appendImplicitAttribute: RenderImplicitAttribute = (rs, integral) => {
  if (rs.stack.length === 0) {
    return;
  }

  const { node } = rs.stack[rs.stack.length - 1];

  const attribute = getText(rs.template, integral.attributeVector);
  if (attribute === undefined) {
    return;
  }

  rs.hooks.setAttribute({ node, attribute, value: true });
};

const appendInjectedAttribute: RenderInjectedAttribute = (rs, integral) => {
  if (rs.stack.length === 0) {
    return;
  }

  const { node } = rs.stack[rs.stack.length - 1];

  const attribute = getText(rs.template, integral.attributeVector);
  if (attribute === undefined) {
    return;
  }

  const { injectionID } = integral;
  const value = rs.template.injections[injectionID];

  if (value instanceof Context) {
    return;
  }
  // add to injection map
  rs.injections[injectionID] = {
    kind: "ATTRIBUTE",
    params: { node, attribute, value },
  };

  rs.hooks.setAttribute({ node, attribute, value });
};

const buildRender: BuildRender = ({ hooks, template, integrals }) => {
  const rs = {
    hooks,
    template,
    injections: {},
    siblings: [],
    stack: [],
  };

  for (const integral of integrals) {
    if (integral.kind === "NODE") {
      createNode(rs, integral);
    }
    if (integral.kind === "SELF_CLOSING_NODE") {
      createNode(rs, integral);
    }
    if (integral.kind === "CLOSE_NODE") {
      closeNode(rs, integral);
    }
    if (integral.kind === "TEXT") {
      createTextNode(rs, integral);
    }
    if (integral.kind === "CONTEXT_INJECTION") {
      createContentInjection(rs, integral);
    }
    if (integral.kind === "EXPLICIT_ATTRIBUTE") {
      appendExplicitAttribute(rs, integral);
    }
    if (integral.kind === "IMPLICIT_ATTRIBUTE") {
      appendImplicitAttribute(rs, integral);
    }
    if (integral.kind === "INJECTED_ATTRIBUTE") {
      appendInjectedAttribute(rs, integral);
    }
  }

  return rs;
};

export {
  appendExplicitAttribute,
  appendImplicitAttribute,
  appendInjectedAttribute,
  buildRender,
  closeNode,
  createContentInjection,
  createNode,
  createTextNode,
};
