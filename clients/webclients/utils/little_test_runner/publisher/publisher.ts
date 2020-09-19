const subpub = new SubPub<RunResultsState>();

const subscribe: Subscribe<RunResultsState> = (callback) => {
  const stub = subpub.subscribe(callback);
  return () => {
    subpub.unsubscribe(stub);
  };
};
