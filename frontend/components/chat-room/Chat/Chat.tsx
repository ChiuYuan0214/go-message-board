import React, { useState } from "react";
import ChatContent from "./ChatContent";

interface Props {
  name: string;
  targetId: number;
  targetImage: string;
  closeChatHandler: () => void;
}

const Chat: React.FC<Props> = ({
  name,
  targetId,
  targetImage,
  closeChatHandler,
}) => {
  const [isOpen, setIsOpen] = useState(true);

  return (
    <>
      <li>
        <div className="header">
          <p onClick={() => setIsOpen((prev) => !prev)}>{name}</p>
          <p onClick={() => setIsOpen((prev) => !prev)}>
            {isOpen ? "一" : "□"}
          </p>
          <p onClick={closeChatHandler}>X</p>
        </div>
        <ChatContent
          targetId={targetId}
          targetImage={targetImage}
          isOpen={isOpen}
        />
      </li>
      <style jsx>{`
        li {
          width: 300px;
          height: ${isOpen ? "330px" : "30px"};
          padding: 0;
          border: 0px solid transparent;
          border-radius: 5px 5px 0 0;
          background-color: #5d5d5d;
          box-shadow: inset 0px 0px 1px 1px rgba(0, 0, 0, 0.25);
          transition: height 0.2s;
          .header {
            width: 100%;
            height: 30px;
            padding: 0;
            display: flex;
            justify-content: flex-end;
            align-items: center;
            p {
              cursor: pointer;
              margin: 0 0.7rem 0 0;
              font-size: 0.7rem;
              &:first-of-type {
                margin: auto;
              }
            }
          }
        }
      `}</style>
    </>
  );
};

export default Chat;
