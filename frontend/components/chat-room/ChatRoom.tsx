import { chatCtx } from "@/context/chat";
import React, { useContext, useEffect, useState } from "react";
import ChatList from "./ChatList";
import FriendList from "./FriendList";

interface Props {}

const ChatRoom: React.FC<Props> = ({}) => {
  const { isChatOpen } = useContext(chatCtx);
  const [isOpen, setIsOpen] = useState(true);

  useEffect(() => {
    let timer: ReturnType<typeof setTimeout>;
    if (isChatOpen) {
      setIsOpen(true);
    } else {
      timer = setTimeout(() => setIsOpen(false), 500);
    }

    return () => clearTimeout(timer);
  }, [isChatOpen]);

  return (
    <>
      <section className={isChatOpen && isOpen ? "isOpen" : ""}>
        <FriendList />
      </section>
      <div className={`chat-list ${isChatOpen && isOpen ? "isOpen" : ""}`}>
        <ChatList />
      </div>
      <style jsx>{`
        section {
          position: fixed;
          right: 0;
          bottom: 0;
          height: 80vh;
          width: 0px;
          transition: width 0.5s ease-in-out;
          &.isOpen {
            width: 300px;
          }
        }
        .chat-list {
          position: fixed;
          right: 0;
          bottom: 0;
          transition: right 0.5s ease-in-out;
          &.isOpen {
            right: 300px;
          }
        }
      `}</style>
    </>
  );
};

export default ChatRoom;
