import { GENERAL_IP } from "@/constants/env";
import api from "./utils";

export const getCollections = async (page = 1) => {
  return await api.get(`${GENERAL_IP}/collections?page=${page}&size=10`);
};

export const addCollection = async (articleId: number) => {
  return await api.post(`${GENERAL_IP}/collections`, { articleId });
};

export const removeCollection = async (articleId: number) => {
  return await api.delete(`${GENERAL_IP}/collections`, { articleId });
};
