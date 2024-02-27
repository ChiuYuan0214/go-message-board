import { getArticle, viewArticle } from "@/api/article";
import { Article } from "@/types/article";
import { GetServerSideProps, NextPage } from "next";
import { authCtx } from "@/context/auth";
import React, { useCallback, useContext, useEffect, useState } from "react";
import ActionButtonList from "@/components/profile/center/components/AddNewContentBlock/ActionButtonList";
import { EditType } from "@/components/article/ContentEditor/EditContentBlock";
import { getComments } from "@/api/comment";
import { Comment } from "@/types/comment";
import CommentList from "@/components/article/CommentList";
import { parseCookies } from "@/utils/cookie";
import Votes from "@/components/article/Votes/Votes";
import ArticleTitle from "@/components/article/ArticleTitle/ArticleTitle";
import ArticleInfo from "@/components/article/ArticleInfo/ArticleInfo";
import ContentEditor, {
  DeleteType,
} from "@/components/article/ContentEditor/ContentEditor";

interface Props {
  data: Article;
  comments: Comment[];
}

const ArticleDetail: NextPage<Props> = (props) => {
  const [data, setData] = useState(props.data);
  const { articleId, userId: authorId, content } = data;
  const [comments, setComments] = useState(props.comments);
  const {
    userInfo: { userId },
  } = useContext(authCtx);

  const [type, setType] = useState<EditType>(""); // for edit block
  const [deleteType, setDeleteType] = useState<DeleteType>(""); // for dialog
  const [selectedComment, setSelectedComment] = useState(0);

  const refreshArticle = useCallback(async () => {
    const { status, data } = await getArticle(userId, articleId);
    if (status === "success") {
      setData(data);
    }
  }, [userId, articleId]);

  const refreshComments = useCallback(async () => {
    const { status, list } = await getComments(userId, articleId);
    if (status === "success") {
      setComments(list);
    }
  }, [userId, articleId]);

  useEffect(() => {
    viewArticle(articleId);
  }, [articleId]);

  const buttonList = [
    {
      text: userId === authorId ? "編輯文章" : "",
      onClick: () => setType("article"),
    },
    {
      text: userId === authorId ? "刪除文章" : "",
      onClick: () => setDeleteType("article"),
    },
    { text: userId ? "新增評論" : "", onClick: () => setType("comment") },
  ];

  return (
    <>
      <section>
        <div className="content-block" onClick={() => setType("")}>
          <div className="article">
            <ArticleTitle
              userId={userId}
              articleId={articleId}
              title={props.data.title}
              hasCollec={props.data.hasCollec}
            />
            <p className="content">{content}</p>
            <Votes
              articleId={articleId}
              voteUp={props.data.voteUp}
              voteDown={props.data.voteDown}
              myScore={props.data.myScore}
            />
            <ArticleInfo
              authorId={props.data.userId}
              author={props.data.author}
              authorImage={props.data.authorImage}
              publishTime={props.data.publishTime}
            />
          </div>
          <CommentList
            data={comments}
            onDelete={(commentId: number) => {
              setDeleteType("comment");
              setSelectedComment(commentId);
            }}
          />
        </div>
        <ActionButtonList list={buttonList} />
      </section>
      <ContentEditor
        articleId={articleId}
        data={props.data}
        editType={type}
        selectedComment={selectedComment}
        deleteType={deleteType}
        cancelEditHandler={() => setType("")}
        cancelDeletionHandler={() => setDeleteType("")}
        refreshArticle={refreshArticle}
        refreshComments={refreshComments}
      />
      <style jsx>{`
        section {
          width: 100%;
          min-height: 100vh;
          display: flex;
          justify-content: center;
          padding-top: 7rem;
          background-color: #dadada;
          .content-block {
            margin-right: 1rem;
            .article {
              width: 600px;
              padding: 1rem;
              background-color: white;
              color: black;
              border-radius: 5px;
              box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
              overflow: hidden;
              margin-bottom: 1rem;
              .content {
                cursor: pointer;
                overflow-wrap: break-word;
                font-size: 0.85rem;
              }
            }
          }
        }
      `}</style>
    </>
  );
};

export const getServerSideProps: GetServerSideProps = async ({
  req,
  query,
}) => {
  const articleId = Array.isArray(query.id) || !query.id ? 0 : +query.id;
  const cookies = req.headers.cookie;
  const userId = cookies ? +parseCookies(cookies)["userId"] : 0;
  const { status: articleStatus, data } = await getArticle(
    userId,
    articleId,
    true
  );
  const { status: commentStatus, list } = await getComments(
    userId,
    articleId,
    true
  );
  return {
    props: {
      articleId,
      data: articleStatus === "success" ? data : {},
      comments: commentStatus === "success" ? list : [],
    },
  };
};

export default ArticleDetail;
