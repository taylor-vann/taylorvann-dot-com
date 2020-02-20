export interface KeyPressHistoryEntry {
  key: string;
  timestamp: number;
}

export type KeyPressStateMap = {
  [code: string]: KeyPressHistoryEntry;
};
