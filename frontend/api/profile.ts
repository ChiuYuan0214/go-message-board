import { GENERAL_IP, SECURITY_IP } from "@/constants/env";
import api from "./utils";

export const getProfile = async (userId = 0) => {
  return await api.get(
    `${GENERAL_IP}/profile${userId ? `?userId=${userId}` : ""}`
  );
};

export const updatePassword = async (body: {
  oldPassword: string;
  newPassword: string;
}) => {
  return await api.put(`${SECURITY_IP}/updatePassword`, body);
};

export const updateProfile = async (body: {
  username: string;
  phone: string;
  job: string;
  address: string;
}) => {
  return await api.post(`${SECURITY_IP}/updateProfile`, body);
};

export const updateImage = async (file: File) => {
  const formData = new FormData();
  formData.append("file", file, file.name);
  formData.append("desc", file.name);

  return api.postForm(`${SECURITY_IP}/uploadImage`, formData);
};
