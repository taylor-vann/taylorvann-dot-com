// brian taylor vann

// N Node
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
  kind: "OPEN_NODE" | "INDEPENDENT_NODE";
  attributes: Attributes<A>;
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
