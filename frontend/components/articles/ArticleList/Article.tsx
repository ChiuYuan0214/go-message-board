import { addCollection, removeCollection } from "@/api/collections";
import { vote } from "@/api/vote";
import BookMarkIcon from "@/components/svgs/book-mark";
import DownVote from "@/components/svgs/down-vote";
import UpVote from "@/components/svgs/up-vote";
import { convertDate } from "@/utils/date";
import { IMAGE_PATH } from "@/constants/env";
import { authCtx } from "@/context/auth";
import { ArticleListData } from "@/models/article";
import { useRouter } from "next/router";
import React, { useContext, useEffect, useState } from "react";

interface Props {
  data: ArticleListData;
  measureRef: any;
}

const Article: React.FC<Props> = ({
  measureRef,
  data: {
    articleId,
    userId: authorId,
    title,
    content,
    voteUp: uv,
    voteDown: dv,
    myScore,
    hasCollec,
    author,
    authorImage,
    commentTitle,
    commentContent,
    commentUser,
    commentUserImage,
    publishTime,
  },
}) => {
  const {
    userInfo: { userId },
  } = useContext(authCtx);
  const router = useRouter();
  const [upVote, setUpVote] = useState(uv);
  const [downVote, setDownVote] = useState(dv);
  const [hasUpVote, setHasUpVote] = useState(false);
  const [hasDownVote, setHasDownVote] = useState(false);
  const [hasMark, setHasMark] = useState(false);
  const datetime = convertDate(publishTime);

  const onClickArticle = () => router.push(`/articles/${articleId}`);

  const doVote = async (score: 1 | -1) => {
    vote({ sourceId: articleId, score, voteType: "article" });
  };

  const doMark = async () => {
    if (!userId) return;
    if (hasMark) {
      const { status } = await removeCollection(articleId);
      if (status !== "success") return;
    } else {
      const { status } = await addCollection(articleId);
      if (status !== "success") return;
    }
    setHasMark((prev) => !prev);
  };

  useEffect(() => {
    setHasUpVote(myScore === 1);
    setHasDownVote(myScore === -1);
    setHasMark(hasCollec);
  }, [myScore, hasCollec]);

  return (
    <>
      <div className="article" ref={measureRef || null}>
        <div className="title">
          <h3 onClick={onClickArticle}>{title}</h3>
          <BookMarkIcon color="#434343" isMark={hasMark} onClick={doMark} />
        </div>
        <p className="content" onClick={onClickArticle}>
          {content}
        </p>
        <div className="votes">
          <p>{upVote}</p>
          <div
            onClick={(e) => {
              e.stopPropagation();
              if (!userId) return;
              doVote(1);
              setHasUpVote((prev) => {
                if (!prev) {
                  setHasDownVote((hasDown) => {
                    if (hasDown) {
                      setDownVote((prev) => prev - 1);
                    }
                    return false;
                  });
                }
                setUpVote((prevCount) => prevCount + (prev ? -1 : 1));
                return !prev;
              });
            }}
            className="icon"
          >
            <UpVote hasVote={hasUpVote} />
          </div>
          <p>{downVote}</p>
          <div
            onClick={(e) => {
              e.stopPropagation();
              if (!userId) return;
              doVote(-1);
              setHasDownVote((prev) => {
                if (!prev) {
                  setHasUpVote((hasUp) => {
                    if (hasUp) {
                      setUpVote((prev) => prev - 1);
                    }
                    return false;
                  });
                }
                setDownVote((prevCount) => prevCount + (prev ? -1 : 1));
                return !prev;
              });
            }}
            className="icon"
          >
            <DownVote hasVote={hasDownVote} />
          </div>
        </div>
        <div className="info">
          <div
            className="author"
            onClick={() => router.push(`/profile?userId=${authorId}`)}
          >
            <div className="image">
              <img src={IMAGE_PATH + authorImage} alt={authorImage} />
            </div>
            <p>作者：{author}</p>
          </div>
          <p>發布日期：{datetime}</p>
        </div>
        {commentTitle && (
          <div className="comment-block">
            <div className="comment-info">
              <div className="image">
                <img
                  src={IMAGE_PATH + commentUserImage}
                  alt={commentUserImage}
                />
              </div>
              <p>{commentUser}</p>
            </div>
            <div className="content">
              <h4>{commentTitle}</h4>
              <p>{commentContent}</p>
            </div>
          </div>
        )}
      </div>
      <style jsx>{`
        .article {
          padding: 1rem;
          background-color: white;
          color: black;
          border-radius: 5px;
          box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
          overflow: hidden;
          margin-bottom: 1rem;
          .title {
            display: flex;
            align-items: center;
            h3 {
              margin-right: auto;
              max-width: 80%;
            }
            > svg {
              cursor: pointer;
            }
          }
          h3,
          .content {
            cursor: pointer;
            overflow-wrap: break-word;
          }
          .content {
            font-size: 0.85rem;
            max-width: 90%;
          }
          .votes {
            display: flex;
            justify-content: flex-end;
            align-items: center;
            position: relative;
            width: 150px;
            margin-left: auto;
            bottom: -35px;
            > p {
              margin: 0;
              width: 1rem;
              font-size: 0.8rem;
            }
            .icon {
              margin-right: 1rem;
              width: 20px;
              height: 20px;
              cursor: pointer;
            }
          }
          .info {
            display: flex;
            align-items: flex-end;
            justify-content: space-between;
            margin-top: 0.5rem;
            > p {
              font-size: 0.7rem;
            }
          }
          .author {
            cursor: pointer;
            > p {
              margin: 0.5rem 0;
              font-size: 0.8rem;
              font-weight: bold;
              color: black;
            }
          }
        }
        .comment-block {
          display: flex;
          border-top: 1px dotted grey;
          .comment-info {
            display: flex;
            flex-direction: column;
            align-items: center;
            margin-top: 1rem;
            margin-right: 1rem;
            .image {
              background-color: #000000;
              width: 40px;
              height: 40px;
              border-radius: 50%;
              overflow: hidden;
            }
            > p {
              margin: 0.5rem 0;
              font-size: 0.8rem;
              font-weight: bold;
              color: black;
            }
          }
        }
        .image {
          background-color: #000000;
          width: 40px;
          height: 40px;
          border-radius: 50%;
          overflow: hidden;
          > img {
            display: block;
            max-width: 100%;
            max-height: 100%;
            object-fit: cover;
            @media (min-aspect-ratio: 1/1) {
              width: 100%;
              height: auto;
            }
            @media (max-aspect-ratio: 1/1) {
              width: auto;
              height: 100%;
            }
          }
        }
      `}</style>
    </>
  );
};

export default Article;
