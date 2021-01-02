// brian taylor vann
// build integrals

// start super loose
// we are taking a text_vector and returning the modification

import { StructureRender } from "../../../type_flyweight/structure";
import { Vector } from "../../../type_flyweight/text_vector";

import {
  decrementTarget,
  incrementTarget,
  hasOriginEclipsedTaraget,
} from "../../../text_vector/text_vector";
import { getCharAtPosition } from "../../../text_position/text_position";

type Routes = Record<string, string>;
type Routers = Partial<Record<string, Routes>>;

const INITAL_TAG_NAME = "INITAL_TAG_NAME";
const TAG_NAME_CONFIRMED = "TAG_NAME_CONFIRMED";

const ROUTERS: Routers = {
  [INITAL_TAG_NAME]: {
    " ": TAG_NAME_CONFIRMED,
  },
};

const crawlForTagName = <A>(
  template: StructureRender<A>,
  textVector: Vector
) => {
  // first character must not be a space
  let positionChar = getCharAtPosition(template, textVector.target);
  if (positionChar === undefined || positionChar === " ") {
    return;
  }

  let tagNameCrawlState =
    ROUTERS?.[INITAL_TAG_NAME]?.[positionChar] ?? INITAL_TAG_NAME;

  while (
    !hasOriginEclipsedTaraget(textVector) &&
    tagNameCrawlState === INITAL_TAG_NAME
  ) {
    if (incrementTarget(template, textVector) === undefined) {
      // this is a bad nil return
      return;
    }

    positionChar = getCharAtPosition(template, textVector.target);
    if (positionChar === undefined) {
      // this is also a bad nil return
      return;
    }

    tagNameCrawlState =
      ROUTERS?.[tagNameCrawlState]?.[positionChar] ?? INITAL_TAG_NAME;
  }

  // only return valid tags
  if (tagNameCrawlState === TAG_NAME_CONFIRMED) {
    decrementTarget(template, textVector);
  }

  return textVector;
};

export { crawlForTagName };
