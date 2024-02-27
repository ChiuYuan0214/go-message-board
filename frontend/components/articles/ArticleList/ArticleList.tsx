import React, { useCallback, useEffect, useRef, useState } from "react";
import Article from "./Article";
import { subColor } from "@/constants/colors";
import useOnScreen from "@/hooks/useOnScreen";
import { ArticleListData } from "@/models/article";
import { getData } from "./utils";

export const getType = (type: string | string[] | undefined) =>
  Array.isArray(type) || !type ? "newest" : type;

interface Props {
  userId: number;
  type: string;
  data: ArticleListData[];
  noServerInit?: boolean;
  marginLeft?: string;
}

const ArticleList: React.FC<Props> = (props) => {
  const { userId } = props;
  const [data, setData] = useState(props.data);
  const [page, setPage] = useState(1);
  const [isLoading, setIsLoading] = useState(false);
  const [isEnd, setIsEnd] = useState(false);
  const { measureRef, isIntersecting, observer } = useOnScreen();
  const initRef = useRef(false);

  useEffect(() => {
    if (!initRef.current || props.noServerInit) {
      initRef.current = true;
      return;
    }
    (async () => {
      setPage(1);
      const { list: newData } = await getData(
        getType(props.type),
        userId,
        1,
        false
      );
      setData(newData);
      setIsLoading(false);
    })();
  }, [props.type, userId, props.noServerInit]);

  useEffect(() => {
    if (page === 1 && !props.noServerInit) return;
    (async () => {
      const { list: newData } = await getData(
        getType(props.type),
        userId,
        page,
        false
      );
      setData((prev) => (page === 1 ? newData : [...prev, ...newData]));
      setIsEnd(newData.length <= 0);
      setIsLoading(false);
    })();
  }, [page, userId, props.type, props.noServerInit]);

  const loadMore = useCallback(() => {
    if (isLoading) return;
    setPage((page) => page + 1);
    setIsLoading(true);
  }, [isLoading]);

  useEffect(() => {
    if (isIntersecting && !isEnd) {
      loadMore();
      observer?.disconnect();
    }
  }, [isIntersecting, isEnd, loadMore]);

  return (
    <>
      <div className="container">
        <ul>
          {data.map((d, i) => (
            <Article
              measureRef={i === data.length - 1 ? measureRef : null}
              key={d.articleId}
              data={d}
            />
          ))}
        </ul>
        {!isEnd && (
          <div className={`loader ${isLoading ? "visible" : ""}`}></div>
        )}
        {(isEnd || !data.length) && (
          <p style={{ color: subColor, margin: "2rem 0 3rem" }}>
            已經沒有更多內容
          </p>
        )}
      </div>
      <style jsx>{`
        .container {
          display: flex;
          flex-direction: column;
          align-items: center;
          margin-left: ${props.marginLeft || "330px"};
          width: 600px;
          ul {
            margin-top: 5rem;
            padding: 0;
            width: 600px;
          }
          .loader {
            border: 8px solid #f3f3f3;
            border-top: 8px solid ${subColor};
            border-radius: 50%;
            width: 50px;
            height: 50px;
            animation: spin 1s linear infinite;
            margin: 20px auto;
            visibility: hidden;
          }
          .visible {
            display: visible;
          }
        }
        @keyframes spin {
          0% {
            transform: rotate(0deg);
          }
          100% {
            transform: rotate(360deg);
          }
        }
      `}</style>
    </>
  );
};

export default ArticleList;
