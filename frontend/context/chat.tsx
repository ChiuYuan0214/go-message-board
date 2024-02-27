import {
  addDBFollow,
  getFollowers,
  getFollows,
  removeDBFollow,
  removeDBFollower,
} from "@/api/follower";
import React, {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useState,
} from "react";
import { authCtx } from "./auth";
import { getLocalToken } from "@/api/utils";

import { Message } from "@/models/message";
import {
  followLoginHandler,
  followLogoutHandler,
  followerLoginHandler,
  followerLogoutHandler,
  historyHandler,
  onlineFollowListHandler,
  onlineFollowerListHandler,
  receiveMessageHandler,
} from "@/utils/chat";
import { HistoryEndMap, HistroyMap } from "@/types/chat";
import { ChatListData } from "@/models/chat";
import { Friend } from "@/models/friend";
import { Socket } from "@/models/socket";

class ChatContext {
  historyMap: HistroyMap = {};
  historyEndMap: HistoryEndMap = {};
  isChatOpen: boolean = false;
  setIsChatOpen: React.Dispatch<React.SetStateAction<boolean>> = () => {};
  followList: Friend[] = [];
  followerList: Friend[] = [];
  openList: ChatListData[] = [];
  setOpenList: React.Dispatch<React.SetStateAction<ChatListData[]>> = () => {};
  addFollow = (id: number, name: string, image: string) => {};
  removeFollowFromList = (id: number) => {};
  removeFollowerFromList = (id: number) => {};
  sendMessage = (id: number, content: string) => {};
  getHistory = (id: number) => {};
}

export const chatCtx = createContext(new ChatContext());

interface Props {
  children: React.ReactNode;
}

export const ChatProvider: React.FC<Props> = ({ children }) => {
  const {
    userInfo: { userId },
    isAuthInit,
    isChatNeedRefresh,
    setIsChatNeedRefresh,
  } = useContext(authCtx);
  const [socket, setSocket] = useState<Socket | null>(null);
  const [isChatOpen, setIsChatOpen] = useState(false);
  const [followList, setFollowList] = useState<Friend[]>([]);
  const [followerList, setFollowerList] = useState<Friend[]>([]);
  const [openList, setOpenList] = useState<ChatListData[]>([]);
  const [historyMap, setHistoryMap] = useState<HistroyMap>({});
  const [historyEndMap, setHistoryEndMap] = useState<HistoryEndMap>({});

  const addFollow = useCallback(
    async (targetId: number, username: string, userImage: string) => {
      const { status } = await addDBFollow(targetId);
      if (status !== "success") return;
      setFollowList((prev) => [
        ...prev,
        new Friend(targetId, username, userImage, false),
      ]);
      socket?.addFollow(userId, targetId);
    },
    [userId, socket]
  );

  const removeFollowFromList = useCallback(
    async (targetId: number) => {
      const { status } = await removeDBFollow(targetId);
      if (status !== "success") return;
      setFollowList((prev) => prev.filter((f) => f.userId !== targetId));
      socket?.removeFollow(userId, targetId);
    },
    [userId, socket]
  );

  const removeFollowerFromList = useCallback(
    async (targetId: number) => {
      const { status } = await removeDBFollower(targetId);
      if (status !== "success") return;
      setFollowerList((prev) => prev.filter((f) => f.userId !== targetId));
      socket?.removeFollower(userId, targetId);
    },
    [userId, socket]
  );

  const sendMessage = useCallback(
    (receiverId: number, content: string) => {
      socket?.sendMessage(userId, receiverId, content);
      setHistoryMap((prev) => ({
        ...prev,
        [receiverId]: [
          ...(prev[receiverId] || []),
          new Message(userId, receiverId, content, new Date().getTime() * 1e6),
        ],
      }));
    },
    [userId, socket]
  );

  const getHistory = useCallback(
    (receiverId: number) => {
      const msgs = historyMap[receiverId];
      const time = !msgs || msgs.length < 1 ? 0 : msgs[0].time;
      socket?.getHistroy(userId, receiverId, time);
    },
    [historyMap, socket, userId]
  );

  const logout = () => {
    setIsChatOpen(false);
    setFollowList([]);
    setFollowerList([]);
    setOpenList([]);
    setHistoryMap({});
    setHistoryEndMap({});
    socket?.logout();
    setSocket(null);
  };

  useEffect(() => {
    if (!userId || !isAuthInit) {
      return logout();
    }
    (async () => {
      const data = await getFollows(userId);
      if (data.status !== "success") return;
      setFollowList(data.list.map((d: any) => ({ ...d, isOnline: false })));
    })();
    (async () => {
      const data = await getFollowers(userId);
      if (data.status !== "success") return;
      setFollowerList(data.list.map((d: any) => ({ ...d, isOnline: false })));
    })();
  }, [userId, isAuthInit]);

  useEffect(() => {
    if (!userId || !isAuthInit) {
      return;
    }
    const newSocket = new Socket(getLocalToken());
    setSocket(newSocket);
  }, [userId, isAuthInit]);

  useEffect(() => {
    if (!socket) return;
    socket.initHandlers({
      onOpen: (e: Event) => {
        console.log("socket opened. event:", e);
      },
      onMessage: (e: MessageEvent) => {
        const data = JSON.parse(e.data);
        switch (data.event) {
          case "follow-login":
            setFollowList(followLoginHandler(data));
            break;
          case "follower-login":
            setFollowerList(followerLoginHandler(data));
            break;
          case "follow-logout":
            setFollowList(followLogoutHandler(data));
            break;
          case "follower-logout":
            setFollowerList(followerLogoutHandler(data));
            break;
          case "online-follow-list":
            onlineFollowListHandler(userId, data, setFollowList);
            break;
          case "online-follower-list":
            onlineFollowerListHandler(userId, data, setFollowerList);
            break;
          case "history":
            setHistoryMap(
              historyHandler(data, (id: number) =>
                setHistoryEndMap((prev) => ({ ...prev, [id]: true }))
              )
            );
            break;
          case "message":
            setHistoryMap(receiveMessageHandler(data));
            break;
          case "error":
            console.log(data.content);
        }
      },
      onClose: (e: CloseEvent) => {
        console.log("close event:", e);
      },
      onError: (e: Event) => {
        console.log("error event:", e);
      },
    });
  }, [socket]);

  useEffect(() => {
    if (!isChatNeedRefresh) return;
    socket?.refresh(getLocalToken());
    setIsChatNeedRefresh(false);
  }, [socket, setIsChatNeedRefresh, isChatNeedRefresh]);

  return (
    <chatCtx.Provider
      value={{
        historyMap,
        isChatOpen,
        setIsChatOpen,
        followList,
        followerList,
        openList,
        setOpenList,
        addFollow,
        removeFollowFromList,
        removeFollowerFromList,
        sendMessage,
        getHistory,
        historyEndMap,
      }}
    >
      {children}
    </chatCtx.Provider>
  );
};
