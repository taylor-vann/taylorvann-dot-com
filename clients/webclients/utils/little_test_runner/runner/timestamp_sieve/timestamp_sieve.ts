// brian taylor vann

type GetTimestamp = () => number;
type UpdateTimestamp = GetTimestamp;

// timestamps
let currentTestTimestamp = performance.now();
const getTimestamp: GetTimestamp = () => {
  return currentTestTimestamp;
};
const updateTimestamp: UpdateTimestamp = () => {
  currentTestTimestamp = performance.now();
  return currentTestTimestamp;
};

export { getTimestamp, updateTimestamp };
