import { refreshToken } from "@/api/auth";
import { getProfile } from "@/api/profile";
import {
  getLocalExpireTime,
  getLocalToken,
  getLocalUserId,
  setLocalToken,
  setLocalTokenExpire,
  syncLocalStorage,
} from "@/api/utils";
import { Count, UserInfo } from "@/models/auth";
import { UserInfoArgs } from "@/types/auth";
import { SetState } from "@/types/general";
import { setCookies } from "@/utils/cookie";
import { createContext, useCallback, useEffect, useState } from "react";

class AuthContext {
  userInfo = new UserInfo();
  isAuthInit = false;
  profileCount = new Count();
  setUserInfo = (info: UserInfoArgs) => {};
  getMyProfile = (bool: boolean) => {};
  clearUserInfo = () => {};
  isChatNeedRefresh = false;
  setIsChatNeedRefresh: SetState<boolean> = () => {};
}

export const authCtx = createContext(new AuthContext());

interface Props {
  children: React.ReactNode;
}

export const AuthProvider: React.FC<Props> = ({ children }) => {
  const [isAuthInit, setIsAuthInit] = useState(false);
  const [userId, setUserId] = useState(0);
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [job, setJob] = useState("");
  const [address, setAddress] = useState("");
  const [isActive, setIsActive] = useState(false);
  const [imagePath, setImagePath] = useState("");
  const [profileCount, setProfileCount] = useState(new Count());
  const [isChatNeedRefresh, setIsChatNeedRefresh] = useState(false);

  const clearUserInfo = () => {
    setUserId(0);
    setUsername("");
    setEmail("");
    setPhone("");
    setJob("");
    setAddress("");
    setIsActive(false);
  };

  const setUserInfo = useCallback((info: UserInfoArgs) => {
    setUserId(info.userId);
    setUsername(info.username);
    setEmail(info.email);
    setPhone(info.phone);
    setJob(info.job);
    setAddress(info.address);
    setIsActive(info.isActive);
    setImagePath(info.imagePath);
    setProfileCount(info.count);
  }, []);

  const getMyProfile = useCallback(
    async (isInitAuth = false) => {
      const { status, data } = await getProfile();
      if (status !== "success") {
        return;
      }
      setUserInfo({
        userId: data.userId,
        username: data.username,
        email: data.email,
        phone: data.phone,
        job: data.job,
        address: data.address,
        isActive: data.isActive,
        imagePath: data.imagePath + `?refresh=${new Date().getTime()}`,
        count: {
          article: data.articleCount,
          comment: data.commentCount,
          upVote: data.upVoteCount,
        },
      });
      setCookies(data.userId, data.token);
      if (isInitAuth) setIsAuthInit(true);
    },
    [setUserInfo]
  );

  const refreshTokenHandler = async (userId: number) => {
    if (!getLocalToken()) return false;
    return await refreshToken(userId).then((data) => {
      if (data.status === "success") {
        setLocalToken(data.token);
        setLocalTokenExpire(data.expireTime * 1000);
        return true;
      }
      return false;
    });
  };

  useEffect(() => {
    (async () => {
      syncLocalStorage();
      const localUserId = getLocalUserId();
      const refreshed = await refreshTokenHandler(localUserId);
      if (refreshed) {
        await getMyProfile();
      }
      setIsAuthInit(true);
    })();
  }, [getMyProfile]);

  useEffect(() => {
    if (!userId) return;

    const timer = setInterval(() => {
      const expireTime = getLocalExpireTime();
      if (
        new Date(expireTime).getTime() - new Date().getTime() <
        5 * 60 * 1000
      ) {
        (async () => {
          const refreshed = await refreshTokenHandler(userId);
          if (refreshed) setIsChatNeedRefresh(true);
        })();
      }
    }, 30000);
    return () => clearInterval(timer);
  }, [userId]);

  return (
    <authCtx.Provider
      value={{
        userInfo: {
          userId,
          username,
          email,
          phone,
          job,
          address,
          isActive,
          imagePath,
        },
        profileCount,
        isAuthInit,
        setUserInfo,
        clearUserInfo,
        getMyProfile,
        isChatNeedRefresh,
        setIsChatNeedRefresh,
      }}
    >
      {children}
    </authCtx.Provider>
  );
};
