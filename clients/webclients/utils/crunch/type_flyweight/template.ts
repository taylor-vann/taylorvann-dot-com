// brian taylor vann
// template

// N Node
// A Attributables

import { AttributeValue } from "./hooks";

interface Template<N, A> {
  templateArray: TemplateStringsArray;
  injections: AttributeValue<N, A>[];
}

export { Template };
