import { KeyPressStateMap, KeyPressHistoryEntry } from "./keypress_types";
import { removeKeyFromKeypressStateMap } from "./keypress_utils";

declare class KeyPressClass {
  onKeyUp(e: KeyboardEvent): void;
  onKeyDown(e: KeyboardEvent): void;
  getState(): string;
}

class KeyPress implements KeyPressClass {
  private _keyStateMap: KeyPressStateMap;

  constructor() {
    this._keyStateMap = {};
  }

  onKeyUp(e: KeyboardEvent): void {
    e.stopPropagation();
    const code = e.key.toLowerCase();

    if (code === "meta" && e.metaKey === false) {
      this._keyStateMap = removeKeyFromKeypressStateMap(
        this._keyStateMap,
        "alt",
      );
      return;
    }
    this._keyStateMap = removeKeyFromKeypressStateMap(this._keyStateMap, code);
  }

  onKeyDown(e: KeyboardEvent): void {
    e.stopPropagation();
    const code = e.key.toLowerCase();
    if (code === "meta" && e.metaKey === false) {
      this._keyStateMap["alt"] = {
        key: "alt",
        timestamp: Date.now(),
      };
      return;
    }

    this._keyStateMap[code] = {
      key: code,
      timestamp: Date.now(),
    };
  }

  getState(): string {
    const trues = [];
    for (let key in this._keyStateMap) {
      if (this._keyStateMap[key] != null) {
        trues.push(key);
      }
    }

    const stateString = trues.sort().join("");
    return stateString;
  }

  getHistoryEntries(): Array<KeyPressHistoryEntry> {
    const trues = [];
    for (let key in this._keyStateMap) {
      if (this._keyStateMap[key] != null) {
        trues.push(this._keyStateMap[key]);
      }
    }

    const sortedTrues = trues.sort((a, b) => {
      if (a.timestamp < b.timestamp) {
        return -1;
      }
      return 1;
    });

    return sortedTrues;
  }

  getStateByTime(): string {
    return this.getHistoryEntries()
      .map(historyEntry => historyEntry.key)
      .join("");
  }

  clearState(): void {
    this._keyStateMap = {};
  }
}

export { KeyPress };
