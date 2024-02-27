import { subColor } from "@/constants/colors";
import { chatCtx } from "@/context/chat";
import React, { useContext } from "react";
import FriendCmp from "./Friend";
import { Friend } from "@/models/friend";
import { ChatListData } from "@/models/chat";

interface Props {}

const FriendList: React.FC<Props> = ({}) => {
  const {
    followList,
    openList: list,
    setOpenList: setList,
  } = useContext(chatCtx);

  const onClickHandler = (data: Friend) => {
    setList((prev) => {
      const newList = [...prev];
      const index = prev.findIndex((u) => u.userId === data.userId);
      if (index >= 0) {
        const target = prev[index];
        newList.splice(index, 1);
        newList.push(target);
      } else {
        if (newList.length >= 3) {
          newList.shift();
        }
        const f = followList.find((f) => f.userId === data.userId);
        newList.push(
          new ChatListData(data.username, data.userId, f?.userImage || "")
        );
      }
      return newList;
    });
  };

  return (
    <>
      <div>
        <h3>追蹤清單</h3>
        <ul>
          {followList.map((f) => (
            <FriendCmp
              key={f.userId}
              data={f}
              onClick={() => onClickHandler(f)}
            />
          ))}
        </ul>
      </div>
      <style jsx>{`
        div {
          width: 100%;
          height: 100%;
          padding: 1rem;
          border-radius: 5px 5px 0 0;
          background-color: ${subColor};
          box-shadow: inset 0px 0px 1px 1px rgba(0, 0, 0, 0.25);
          overflow: hidden;
          white-space: nowrap;
        }
        h3 {
          margin: 0;
          text-align: center;
        }
        ul {
          text-decoration: none;
          list-style-type: none;
          padding: 0;
          background-color: white;
          border-radius: 5px;
          overflow: hidden;
        }
      `}</style>
    </>
  );
};

export default FriendList;
