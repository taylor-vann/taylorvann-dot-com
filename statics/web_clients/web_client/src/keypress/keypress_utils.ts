import { KeyPressStateMap, KeyPressHistoryEntry } from "./keypress_types";

function removeKeyFromKeypressStateMap(
  stateMap: KeyPressStateMap,
  targetKey: string,
): KeyPressStateMap {
  const modifiedStateMap: KeyPressStateMap = {};
  for (let key in stateMap) {
    if (key === targetKey) {
      continue;
    }

    modifiedStateMap[key] = stateMap[key];
  }

  return modifiedStateMap;
}

export { removeKeyFromKeypressStateMap };
