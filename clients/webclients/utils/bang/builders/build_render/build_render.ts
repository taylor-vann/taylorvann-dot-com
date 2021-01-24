// brian taylor vann
// build structure

import { Hooks } from "../../type_flyweight/hooks";
import { Template } from "../../type_flyweight/template";
import {
  RenderStructure,
  Injection,
  TextNode,
  ElementNode,
} from "../../type_flyweight/render";
import { Context } from "../../type_flyweight/context";
import {
  CloseNodeAction,
  CreateContentAction,
  CreateNodeAction,
  CreateSelfClosingNode,
  ExplicitAttributeAction,
  InjectedAttributeAction,
  ImplicitAttributeAction,
  Integrals,
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

type RenderCreateNode = <N, A>(
  rs: RenderStructure<N, A>,
  integral: CreateNodeAction | CreateSelfClosingNode
) => void;

type RenderCreateTextNode = <N, A>(
  rs: RenderStructure<N, A>,
  textVector: Vector
) => void;

type RenderCloseNode = <N, A>(
  rs: RenderStructure<N, A>,
  integral: CloseNodeAction
) => void;

type RenderCreateContent = <N, A>(
  rs: RenderStructure<N, A>,
  integral: CreateContentAction
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

const createTextNode: RenderCreateTextNode = (rs, vector) => {
  const text = getText(rs.template, vector);
  if (text === undefined) {
    return;
  }

  const textNode = rs.hooks.createTextNode(text);

  appendDescendant(rs, textNode);
};

const createNode: RenderCreateNode = (rs, integral) => {
  const tagName = getText(rs.template, integral.params.tagNameVector);
  if (tagName === undefined) {
    return;
  }

  const node = rs.hooks.createNode(tagName);
  const selfClosing = integral.action === "CREATE_SELF_CLOSING_NODE";

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

  const tagName = getText(rs.template, integral.params.tagNameVector);
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

const createContent: RenderCreateContent = (rs, integral) => {
  const { origin, target } = integral.params.contentVector;

  // does injection come first?
  if (origin.stringIndex === 0) {
    const injection = rs.template.injections[origin.arrayIndex - 1];
    if (injection !== undefined) {
      createContentInjection({
        rs,
        injection,
        injectionID: origin.arrayIndex - 1,
      });
    }
  }

  // what if content has no injections?
  if (origin.arrayIndex === target.arrayIndex) {
    createTextNode(rs, integral.params.contentVector);
    return;
  }

  // get that beginning stuff
  const stringIndex = rs.template.templateArray[origin.arrayIndex].length - 1;
  const startVector = {
    origin,
    target: {
      arrayIndex: origin.arrayIndex,
      stringIndex,
    },
  };
  createTextNode(rs, startVector);

  const injection = rs.template.injections[origin.arrayIndex];
  createContentInjection({
    rs,
    injection,
    injectionID: origin.arrayIndex,
  });

  // get that middle stuff
  let innerIndex = origin.arrayIndex + 1;
  while (innerIndex < target.arrayIndex) {
    const stringIndex = rs.template.templateArray[innerIndex].length - 1;
    const startVector = {
      origin: {
        arrayIndex: innerIndex,
        stringIndex: 0,
      },
      target: {
        arrayIndex: innerIndex,
        stringIndex,
      },
    };
    createTextNode(rs, startVector);

    // attach injections
    const startInjection = rs.template.injections[innerIndex];
    createContentInjection({
      injection: startInjection,
      injectionID: innerIndex,
      rs,
    });

    innerIndex += 1;
  }

  // get that end stuff
  const endVector = {
    origin: {
      arrayIndex: target.arrayIndex,
      stringIndex: 0,
    },
    target,
  };
  createTextNode(rs, endVector);

  // if text node ends before an injection
  const endStringIndex =
    rs.template.templateArray[target.arrayIndex].length - 1;
  if (endStringIndex === target.stringIndex) {
    const injection = rs.template.injections[target.arrayIndex];
    createContentInjection({
      injectionID: target.arrayIndex,
      injection,
      rs,
    });
  }
};

const appendExplicitAttribute: RenderAppendExplicitAttribute = (
  rs,
  integral
) => {
  const node = rs.stack[rs.stack.length - 1].node;

  const attribute = getText(rs.template, integral.params.attributeVector);
  if (attribute === undefined) {
    return;
  }

  const value = getText(rs.template, integral.params.valueVector);
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

  const attribute = getText(rs.template, integral.params.attributeVector);
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

  const attribute = getText(rs.template, integral.params.attributeVector);
  if (attribute === undefined) {
    return;
  }

  const { injectionID } = integral.params;
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
    if (integral.action === "CREATE_NODE") {
      createNode(rs, integral);
    }
    if (integral.action === "CREATE_SELF_CLOSING_NODE") {
      createNode(rs, integral);
    }
    if (integral.action === "CLOSE_NODE") {
      closeNode(rs, integral);
    }
    if (integral.action === "CREATE_CONTENT") {
      createContent(rs, integral);
    }
    if (integral.action === "APPEND_EXPLICIT_ATTRIBUTE") {
      appendExplicitAttribute(rs, integral);
    }
    if (integral.action === "APPEND_IMPLICIT_ATTRIBUTE") {
      appendImplicitAttribute(rs, integral);
    }
    if (integral.action === "APPEND_INJECTED_ATTRIBUTE") {
      appendInjectedAttribute(rs, integral);
    }
  }

  return rs;
};

export { buildRender };
