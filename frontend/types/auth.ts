import { Count } from "@/models/auth";

export type UserInfoArgs = {
  userId: number;
  username: string;
  email: string;
  phone: string;
  job: string;
  address: string;
  isActive: boolean;
  imagePath: string;
  count: Count;
};
