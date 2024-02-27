import { GENERAL_IP } from "@/constants/env";
import api from "./utils";

export const getFollows = async (userId: number) => {
  return await api.get(`${GENERAL_IP}/follows?userId=${userId}`);
};

export const getFollowers = async (userId: number) => {
  return await api.get(`${GENERAL_IP}/follower?userId=${userId}`);
};

export const addDBFollow = async (followee: number) => {
  return await api.post(`${GENERAL_IP}/follow`, { followee });
};

export const removeDBFollow = async (followee: number) => {
  return await api.delete(`${GENERAL_IP}/follow`, { followee });
};

export const removeDBFollower = async (follower: number) => {
  return await api.delete(`${GENERAL_IP}/follower`, { follower });
};
