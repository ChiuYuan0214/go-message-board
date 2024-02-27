import { useRouter } from "next/router";
import React, { useContext, useEffect, useState } from "react";
import Tab, { TabData } from "./Tab";
import { removeLocalToken } from "@/api/utils";
import { authCtx } from "@/context/auth";
import { chatCtx } from "@/context/chat";
import { getParam } from "@/utils/query";
import Dialog from "@/components/UI/Dialog";
import { UserInfo } from "@/models/auth";

const selfTabDataList: TabData[] = [
  { name: "總覽", icon: "none", id: "general" },
  { name: "我的文章", icon: "none", id: "article" },
  { name: "我的追蹤/粉絲", icon: "none", id: "fans" },
  { name: "我的收藏", icon: "none", id: "collection" },
  { name: "設定", icon: "none", id: "setting" },
  { name: "登出", icon: "none", id: "logout" },
];

const otherTabDataList: TabData[] = [
  { name: "總覽", icon: "none", id: "general" },
  { name: "他的文章", icon: "none", id: "article" },
  { name: "他的追蹤/粉絲", icon: "none", id: "fans" },
  { name: "追蹤", icon: "none", id: "follow" },
];

const getTabData = (isSelf: boolean, hasFollow = false) =>
  isSelf
    ? selfTabDataList
    : hasFollow
    ? otherTabDataList.map((t) => ({
        ...t,
        name: t.id === "follow" ? "取消追蹤" : t.name,
      }))
    : otherTabDataList;

interface Props {
  defaultTab: string;
  otherInfo: UserInfo;
}

const TabList: React.FC<Props> = ({ defaultTab, otherInfo }) => {
  const {
    userInfo: { userId },
    clearUserInfo,
  } = useContext(authCtx);
  const { setIsChatOpen, addFollow, removeFollowFromList, followList } =
    useContext(chatCtx);
  const [activeTab, setActiveTab] = useState(defaultTab);
  const [followDialogTitle, setFollowDialogTitle] = useState("");
  const router = useRouter();
  const paramUserId = +getParam(router.query.userId, "0");
  const mode = getParam(router.query.mode, "general");
  const hasFollow =
    followList.findIndex((f) => f.userId === otherInfo.userId) >= 0;
  const isSelf = userId === paramUserId || !paramUserId;
  const tabListData = getTabData(isSelf, hasFollow);

  const onClickHandler = (id: string) => {
    if (id === "logout") {
      removeLocalToken();
      clearUserInfo();
      setIsChatOpen(false);
      router.push("/");
    } else if (id === "follow") {
      return setFollowDialogTitle(
        hasFollow
          ? `確定要取消追蹤${otherInfo.username}嗎？`
          : `確定要追蹤${otherInfo.username}嗎？`
      );
    } else {
      router.push({
        pathname: "/profile",
        query: { ...router.query, mode: id },
      });
    }
    return setActiveTab(id);
  };

  useEffect(() => {
    setActiveTab(mode);
  }, [mode]);

  return (
    <>
      <div className="list">
        {tabListData.map((d) => (
          <Tab
            key={d.id}
            tabData={d}
            activeTab={activeTab}
            onClick={() => onClickHandler(d.id)}
          />
        ))}
        {!isSelf && (
          <Dialog
            title={followDialogTitle}
            desc=""
            confirmText={hasFollow ? "取消追蹤" : "確認追蹤"}
            cancelText="取消"
            onConfirm={
              hasFollow
                ? () => {
                    removeFollowFromList(otherInfo.userId);
                    setFollowDialogTitle("");
                  }
                : () => {
                    addFollow(
                      otherInfo.userId,
                      otherInfo.username,
                      otherInfo.imagePath
                    );
                    setFollowDialogTitle("");
                  }
            }
            onCancel={() => setFollowDialogTitle("")}
          />
        )}
      </div>
      <style jsx>{`
        .list {
          width: 100%;
          font-family: "Poppins";
          font-weight: 600;
          color: white;
          @media only screen and (max-width: 650px) {
            margin-top: 1.5rem;
          }
        }
      `}</style>
    </>
  );
};

export default TabList;
