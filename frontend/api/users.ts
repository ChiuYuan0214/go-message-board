import { SECURITY_IP } from "@/constants/env";
import api from "./utils";

export const getUsers = async (name: string) => {
  return await api.get(`${SECURITY_IP}/users?name=${name}`);
};
