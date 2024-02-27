import { GENERAL_IP, SECURITY_IP } from "@/constants/env";
import api, { getLocalToken } from "./utils";

interface RegisterBody {
  username: string;
  email: string;
  password: string;
  phone: string;
  job: string;
  address: string;
}

interface VerifyCodeBody {
  userId: number;
  code: number;
}

interface ResendCodeBody {
  email: string;
  password: string;
}

interface LoginBody {
  email: string;
  password: string;
}

interface RefreshTokenBody {
  userId: number;
  token: string;
}

export const register = async (body: RegisterBody) => {
  return await api.post(`${SECURITY_IP}/register`, body);
};

export const verifyCode = async (body: VerifyCodeBody) => {
  return await api.post(`${SECURITY_IP}/verifyCode`, body);
};

export const resendVerificationCode = async (body: ResendCodeBody) => {
  return await api.post(`${SECURITY_IP}/resendVerificationCode`, body);
};

export const login = async (body: LoginBody) => {
  return await api.post(`${SECURITY_IP}/login`, body);
};

export const refreshToken = async (userId: number) => {
  return await api.put(`${SECURITY_IP}/login`, {
    userId,
    token: getLocalToken(),
  });
};
