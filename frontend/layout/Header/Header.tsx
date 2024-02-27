import { mainColor } from "@/constants/colors";
import { authCtx } from "@/context/auth";
import { chatCtx } from "@/context/chat";
import Link from "next/link";
import React, { useCallback, useContext, useEffect, useState } from "react";

interface Props {}

const Header: React.FC<Props> = () => {
  const [isHidden, setIsHidden] = useState(false);
  const [scrollPos, setScrollPos] = useState(2000);
  const {
    userInfo: { userId },
  } = useContext(authCtx);
  const { setIsChatOpen } = useContext(chatCtx);

  const onScrollHandler = useCallback(() => {
    const newPos = window.scrollY;
    setIsHidden(newPos > scrollPos);
    setScrollPos(newPos);
  }, [scrollPos]);

  useEffect(() => {
    window.onscroll = onScrollHandler;
    return () => {
      window.onscroll = null;
    };
  }, [onScrollHandler]);

  return (
    <>
      <header className={isHidden ? "hidden" : ""}>
        <ul>
          <li>
            <Link href="/">討論版</Link>
          </li>
          <li>
            <Link href="/profile">會員中心</Link>
          </li>
          <li
            onClick={() => {
              setIsChatOpen((prev) => (userId ? !prev : false));
            }}
          >
            聊天室
          </li>
        </ul>
      </header>
      <style jsx>{`
        header {
          position: fixed;
          z-index: 950;
          width: 100vw;
          height: 60px;
          background-color: ${mainColor};
          top: 0;
          transition: top 0.5s;
          &.hidden {
            top: -60px;
          }
          > ul {
            text-decoration: none;
            list-style-type: none;
            display: flex;
            justify-content: flex-end;
            li {
              cursor: pointer;
              margin: 0 1.5rem;
            }
          }
        }
      `}</style>
    </>
  );
};

export default Header;
