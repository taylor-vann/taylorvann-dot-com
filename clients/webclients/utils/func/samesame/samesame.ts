// brian taylor vann
// samesame

// deep compare obejcts and arrays and collections of objects and arrays
// does not support DateTime obejcts

type SameSame = <T>(source: T, comparator: T) => boolean;

const samesame: SameSame = (source, comparator) => {
  if (source instanceof Object === false) {
    return source === comparator;
  }

  if (comparator instanceof Object === false) {
    return false;
  }

  for (const sourceKey in source) {
    const nextSource = source[sourceKey];
    const nextComparator = comparator[sourceKey];

    const result = samesame(nextSource, nextComparator);
    if (!result) {
      return result;
    }
  }

  return true;
};

export { samesame };
