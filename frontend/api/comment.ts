import { GENERAL_IP, SSR_GENERAL_IP } from "@/constants/env";
import api from "./utils";

export const addComment = async (body: {
  articleId: number;
  title: string;
  content: string;
}) => {
  return await api.post(`${GENERAL_IP}/comment`, body);
};

export const getComments = async (
  userId: number,
  articleId: number,
  isSSR = false
) => {
  return await api.get(
    `${
      isSSR ? SSR_GENERAL_IP : GENERAL_IP
    }/comments?userId=${userId}&articleId=${articleId}`
  );
};

export const deleteComment = async (commentId: number) => {
  return await api.delete(`${GENERAL_IP}/comment?commentId=${commentId}`, {});
};
