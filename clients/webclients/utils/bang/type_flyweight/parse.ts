// brian taylor vann
// parse

// A Attributables

interface AttributeBase<A> {
  name: string;
  value: A;
}
type Attributes<A> = Record<string, AttributeBase<A>>;

interface ContentParseResults {
  kind: "CONTENT_NODE";
  content: string;
}
interface OpenParseResults<A> {
  tag: string;
  kind: "OPEN_NODE" | "SELF_CLOSING_NODE_CONFIRMED";
  attributes?: Attributes<A>;
}
interface CloseParseResults {
  tag: string;
  kind: "CLOSE_NODE";
}
type ParseResults<A> =
  | ContentParseResults
  | OpenParseResults<A>
  | CloseParseResults;

export {
  AttributeBase,
  Attributes,
  CloseParseResults,
  ContentParseResults,
  OpenParseResults,
  ParseResults,
};
