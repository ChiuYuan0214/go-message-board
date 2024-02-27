import React from "react";
import Header from "./Header/Header";
import ChatRoom from "@/components/chat-room/ChatRoom";

interface Props {
  children: React.ReactNode;
}

const Layout: React.FC<Props> = ({ children }) => {
  return (
    <>
      <Header />
      {children}
      <ChatRoom />
    </>
  );
};

export default Layout;
