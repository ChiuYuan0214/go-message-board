import { GENERAL_IP, SSR_GENERAL_IP } from "@/constants/env";
import api from "./utils";

export const getArticles = async (userId = 0, page = 1, isSSR = false) => {
  return await api.get(
    `${
      isSSR ? SSR_GENERAL_IP : GENERAL_IP
    }/articles?page=${page}&userId=${userId}`
  );
};

export const getViewArticles = async (userId = 0, page = 1, isSSR = false) => {
  return await api.get(
    `${
      isSSR ? SSR_GENERAL_IP : GENERAL_IP
    }/articles?type=view&page=${page}&userId=${userId}`
  );
};

export const getHotArticles = async (userId = 0, page = 1, isSSR = false) => {
  return await api.get(
    `${
      isSSR ? SSR_GENERAL_IP : GENERAL_IP
    }/articles?type=hot&page=${page}&userId=${userId}`
  );
};

export const getProfileArticles = async (userId = 0, page = 1) => {
  return await api.get(
    `${GENERAL_IP}/articles?type=profile&page=${page}&userId=${userId}`
  );
};
