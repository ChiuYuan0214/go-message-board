import { GENERAL_IP, SSR_GENERAL_IP } from "@/constants/env";
import api from "./utils";

export const addArticle = async (body: {
  title: string;
  content: string;
  publishTime: string;
}) => {
  return await api.post(`${GENERAL_IP}/article`, { tags: [], ...body });
};

export const updateArticle = async (
  articleId: number,
  body: {
    title: string;
    content: string;
  }
) => {
  return await api.put(`${GENERAL_IP}/article?articleId=${articleId}`, {
    tags: [],
    ...body,
  });
};

export const deleteArticle = async (articleId: number) => {
  return await api.delete(`${GENERAL_IP}/article?articleId=${articleId}`, {});
};

export const getArticle = async (
  userId: number,
  articleId: number,
  isSSR = false
) => {
  return await api.get(
    `${
      isSSR ? SSR_GENERAL_IP : GENERAL_IP
    }/article?userId=${userId}&articleId=${articleId}`
  );
};

export const viewArticle = async (articleId: number) => {
  return await api.put(`${GENERAL_IP}/view?articleId=${articleId}`, {});
};
