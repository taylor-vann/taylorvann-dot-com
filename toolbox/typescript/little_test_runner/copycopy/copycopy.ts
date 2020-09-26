// brian taylor vann
// copycopy

// deep clone obejcts and arrays and collections of objects and arrays
// does not support DateTime obejcts

type CopyCopy = <T>(atomToCopy: T) => T;

const copycopy: CopyCopy = (atomToCopy) => {
  if (atomToCopy instanceof Object === false) {
    return atomToCopy;
  }

  const entries = Array.isArray(atomToCopy)
    ? ([...atomToCopy] as typeof atomToCopy)
    : { ...atomToCopy };

  for (const index in entries) {
    const entry = entries[index];
    if (entries instanceof Object) {
      entries[index] = copycopy(entry);
    }
  }

  return entries;
};

export { copycopy };
