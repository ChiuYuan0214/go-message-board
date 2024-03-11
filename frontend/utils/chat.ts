import { mergeSortedList } from "./sort";
import { Message } from "@/models/message";
import { getFollowers, getFollows } from "@/api/follower";
import { Friend } from "@/models/friend";

export const followLoginHandler = (data: any) => (prev: Friend[]) => {
  const list = [...prev];
  const index = list.findIndex((f) => f.userId === data.userId);
  if (index >= 0 && list[index]) {
    list[index].isOnline = true;
  }
  return list;
};

export const followerLoginHandler = (data: any) => (prev: Friend[]) => {
  const list = [...prev];
  const index = list.findIndex((f) => f.userId === data.userId);
  if (index >= 0 && list[index]) {
    list[index].isOnline = true;
  }
  return list;
};

export const followLogoutHandler = (data: any) => (prev: Friend[]) => {
  const list = [...prev];
  const index = list.findIndex((f) => f.userId === data.userId);
  if (index >= 0 && list[index]) {
    list[index].isOnline = false;
  }
  return list;
};

export const followerLogoutHandler = (data: any) => (prev: Friend[]) => {
  const list = [...prev];
  const index = list.findIndex((f) => f.userId === data.userId);
  if (index >= 0 && list[index]) {
    list[index].isOnline = false;
  }
  return list;
};

export const onlineFollowListHandler = async (
  userId: number,
  data: any,
  setter: (f: Friend[]) => void
) => {
  const res = await getFollows(userId);
  if (res.status !== "success") return;
  const fList: Friend[] = res.list.map((d: any) => ({
    ...d,
    isOnline: false,
  }));

  data.list.forEach((id: number) => {
    const index = fList.findIndex((f) => f.userId === id);
    if (index >= 0) {
      fList[index].isOnline = true;
    }
  });
  setter(fList);
};

export const onlineFollowerListHandler = async (
  userId: number,
  data: any,
  setter: (f: Friend[]) => void
) => {
  const res = await getFollowers(userId);
  if (res.status !== "success") return;
  const fList: Friend[] = res.list.map((d: any) => ({
    ...d,
    isOnline: false,
  }));
  data.list.forEach((id: number) => {
    const index = fList.findIndex((f) => f.userId === id);
    if (index >= 0) {
      fList[index].isOnline = true;
    }
  });
  setter(fList);
};
const getListFromHsitory = (data: any) => {
  const userList = data.userHistory.map((d: any) => ({
    ...d,
    time: Math.floor(d.time / 1e6),
  }));
  const targetList = data.targetHistory.map((d: any) => ({
    ...d,
    time: Math.floor(d.time / 1e6),
  }));
  userList.reverse();
  targetList.reverse();
  return mergeSortedList(
    userList.map(
      (d: any) => new Message(d.userId, d.targetUserId, d.content, d.time)
    ),
    targetList.map(
      (d: any) => new Message(d.userId, d.targetUserId, d.content, d.time)
    ),
    (d) => d.time
  );
};

export const historyHandler = (
  data: any,
  endNotifier: (id: number) => void
) => {
  const receiverId = data.targetUserId;
  if (!receiverId) return (prev: { [target: number]: Message[] }) => prev;
  const list = getListFromHsitory(data);
  if (!list.length) {
    endNotifier(receiverId);
  }

  return (prev: { [target: number]: Message[] }) => {
    return {
      ...prev,
      [receiverId]: mergeSortedList(
        prev[receiverId] || [],
        list,
        (d) => d.time
      ),
    };
  };
};

export const receiveMessageHandler =
  (data: any) => (prev: { [targetId: number]: Message[] }) => ({
    ...prev,
    [data.userId]: [
      ...(prev[data.userId] || []),
      new Message(
        data.userId,
        data.targetUserId,
        data.content,
        Math.floor(data.time / 1e6)
      ),
    ],
  });
