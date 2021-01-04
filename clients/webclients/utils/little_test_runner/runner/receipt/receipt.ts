// brian taylor vann

type GetStub = () => number;
type UpdateStub = GetStub;

let stub = 0;

const getStub: GetStub = () => {
  return stub;
};

const updateStub: UpdateStub = () => {
  stub += 1;

  return stub;
};

export { getStub, updateStub };
