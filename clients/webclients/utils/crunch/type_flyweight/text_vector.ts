// brian taylor vann
// text vector

interface Position {
  arrayIndex: number;
  stringIndex: number;
}

interface Vector {
  origin: Position;
  target: Position;
}

export { Position, Vector };
