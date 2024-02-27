import React, { useContext } from "react";
import Chat from "./Chat/Chat";
import { chatCtx } from "@/context/chat";

interface Props {}

const ChatList: React.FC<Props> = ({}) => {
  const { openList: list, setOpenList: setList } = useContext(chatCtx);

  return (
    <>
      <ul>
        {list.map((d) => (
          <Chat
            key={d.userId}
            name={d.name}
            targetId={d.userId}
            targetImage={d.userImage}
            closeChatHandler={() =>
              setList((prev) => prev.filter((u) => u.userId !== d.userId))
            }
          />
        ))}
      </ul>
      <style jsx>{`
        ul {
          text-decoration: none;
          list-style-type: none;
          padding: 0;
          margin: 0;
          display: flex;
          flex-direction: row;
          align-items: flex-end;
        }
      `}</style>
    </>
  );
};

export default ChatList;
