import Loading from "@/components/animation/Loading";
import { authCtx } from "@/context/auth";
import { chatCtx } from "@/context/chat";
import React, {
  useCallback,
  useContext,
  useEffect,
  useRef,
  useState,
} from "react";
import Message from "./Message";
import TextArea from "./TextArea";

interface Props {
  targetId: number;
  targetImage: string;
  isOpen: boolean;
}

const initMap: { [targetId: number]: boolean } = {};

const ChatContent: React.FC<Props> = ({ targetId, targetImage, isOpen }) => {
  const {
    userInfo: { userId, imagePath },
  } = useContext(authCtx);
  const { historyMap, historyEndMap, getHistory } = useContext(chatCtx);
  const [isFetching, setIsFetching] = useState(false);
  const containerRef = useRef<HTMLDivElement>(null);
  const listRef = useRef<HTMLUListElement>(null);
  const messages = historyMap[targetId];
  const isEnd = historyEndMap[targetId];
  const dataExist = messages && messages.length > 0;

  const scrollToBottom = useCallback((behavior: "instant" | "smooth") => {
    listRef.current?.parentElement?.scrollTo({
      top: listRef.current.offsetHeight + 100,
      left: 0,
      behavior,
    });
  }, []);

  const onScroll = () => {
    const listElement = containerRef.current;
    if (listElement) {
      const isAtTop = listElement.scrollTop === 0;
      if (isAtTop && !isFetching && !isEnd) {
        setIsFetching(true);
        setTimeout(() => setIsFetching(false), 1500);
        getHistory(targetId);
      }
    }
  };

  useEffect(() => {
    if (!dataExist) {
      return;
    }
    scrollToBottom("instant");
  }, [dataExist, scrollToBottom]);

  useEffect(() => {
    if (initMap[targetId]) return;
    if (!dataExist) {
      getHistory(targetId);
      initMap[targetId] = true;
    }
  }, [dataExist, getHistory, targetId]);

  return (
    <>
      <div className="content">
        <div className="messages" ref={containerRef}>
          <ul ref={listRef} onWheel={onScroll}>
            {isFetching && (
              <li>
                <Loading />
              </li>
            )}
            {dataExist && (
              <p className="time">
                {new Date(messages[0].time / 1e6).toLocaleString()}
              </p>
            )}
            {dataExist &&
              messages.map((msg) => (
                <Message
                  key={msg.time + "_" + targetId}
                  userId={userId}
                  imagePath={imagePath}
                  targetImage={targetImage}
                  msg={msg}
                />
              ))}
          </ul>
        </div>
        <TextArea targetId={targetId} scroll={scrollToBottom} />
      </div>
      <style jsx>{`
        .content {
          border: 0.5px solid black;
          height: ${isOpen ? "300px" : "0px"};
          transition: height 0.2s;
          background-color: white;
          display: flex;
          flex-direction: column;
          .messages {
            flex-grow: 1;
            overflow: scroll;
            ul {
              padding: 0;
              min-height: 100%;
              .time {
                color: #7a7a7a;
                margin: 1rem auto;
                font-size: 0.7rem;
                text-align: center;
              }
            }
          }
        }
      `}</style>
    </>
  );
};

export default ChatContent;
