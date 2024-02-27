import { Message } from "../models/message";

export type HistroyMap = {
  [targetId: number]: Message[];
};

export type HistoryEndMap = {
  [targetId: number]: boolean;
};
