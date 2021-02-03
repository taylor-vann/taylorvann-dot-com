// brian taylor vann
// samestuff

// deep compare obejcts and arrays and collections of objects and arrays
// does not support DateTime obejcts

type SameStuff = <T>(source: T, comparator: T) => boolean;

const samestuff: SameStuff = (source, comparator) => {
  if (source instanceof Object === false) {
    return source === comparator;
  }

  if (comparator instanceof Object === false) {
    return false;
  }

  for (const sourceKey in source) {
    // update to iterate over intreis
    if (!Object.hasOwnProperty(sourceKey)) {
      continue;
    }
    const nextSource = source[sourceKey];
    const nextComparator = comparator[sourceKey];

    const result = samestuff(nextSource, nextComparator);
    if (!result) {
      return result;
    }
  }

  return true;
};

export { samestuff };
