import { SSR_GENERAL_IP, GENERAL_IP } from "@/constants/env";
import api from "./utils";

export const vote = async (body: {
  sourceId: number;
  score: number;
  voteType: string;
}) => {
  return await api.post(`${GENERAL_IP}/vote`, body);
};
