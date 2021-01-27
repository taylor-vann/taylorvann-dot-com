// brian taylor vann
// build structure

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
  template: Template<A>;
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

type AppendDescendant = <N, A>(
  rs: RenderStructure<N, A>,
  descendant: N
) => void;

interface CreateContextInjectionParam<N, A> {
  rs: RenderStructure<N, A>;
  injection: A;
  injectionID: number;
}
type CreateContextInjection = <N, A>(
  params: CreateContextInjectionParam<N, A>
) => void;

const appendDescendant: AppendDescendant = (rs, descendant) => {
  // clean up self closing nodes
  let parent = rs.stack[rs.stack.length - 1];
  while (parent.kind === "NODE" && parent.selfClosing === true) {
    rs.stack.pop();
    parent = rs.stack[rs.stack.length - 1];
  }

  const parentNode = parent?.node;
  const leftNode = rs.siblings[rs.siblings.length - 1];

  rs.hooks.appendDescendant({ descendant, parentNode, leftNode });
  if (rs.stack.length === 0) {
    rs.siblings.push(descendant);
  }
};

const createTextNode: RenderTextNode = (rs, integral) => {
  const text = getText(rs.template, integral.textVector);
  if (text === undefined) {
    return;
  }

  const textNode = rs.hooks.createTextNode(text);

  appendDescendant(rs, textNode);
};

const createNode: RenderNode = (rs, integral) => {
  const tagName = getText(rs.template, integral.tagNameVector);
  if (tagName === undefined) {
    return;
  }

  const node = rs.hooks.createNode(tagName);
  const selfClosing = integral.kind === "SELF_CLOSING_NODE";

  rs.stack.push({
    kind: "NODE",
    selfClosing,
    tagName,
    node,
  });

  appendDescendant(rs, node);
};

const closeNode: RenderCloseNode = (rs, integral) => {
  if (rs.stack.length === 0) {
    return;
  }

  const tagName = getText(rs.template, integral.tagNameVector);
  const nodeBit = rs.stack[rs.stack.length - 1];
  if (nodeBit.kind !== "NODE") {
    return;
  }

  if (nodeBit.tagName === tagName) {
    rs.stack.pop();
  }
};

const createContentInjection: CreateContextInjection = ({
  rs,
  injection,
  injectionID,
}) => {
  // attach injection as Context
  const parent = rs.stack[rs.stack.length - 1]?.node;
  const left = rs.siblings[rs.stack.length - 1];

  if (injection instanceof Context) {
    const siblings = injection.getSiblings();
    for (const sibling of siblings) {
      appendDescendant(rs, sibling);
    }

    rs.injections[injectionID] = {
      kind: "CONTEXT",
      params: { siblings, left, parent },
    };
    return;
  }

  // attach injection as content
  const text = String(injection);
  const textNode = rs.hooks.createTextNode(text);

  appendDescendant(rs, textNode);

  rs.injections[injectionID] = {
    kind: "CONTENT",
    params: { textNode, left, parent },
  };
};

const createContent: RenderContentInjection = (rs, integral) => {};

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
      createContent(rs, integral);
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

export { buildRender };
