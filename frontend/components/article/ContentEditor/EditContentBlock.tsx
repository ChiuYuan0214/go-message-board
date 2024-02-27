import { updateArticle } from "@/api/article";
import { addComment } from "@/api/comment";
import FadeIn from "@/components/animation/FadeIn";
import { Article } from "@/types/article";
import { useEffect, useState } from "react";

export type EditType = "article" | "comment" | "";

interface Props {
  editType: EditType;
  articleId: number;
  onClose: () => void;
  articleData: Article;
  refreshArticle: () => void;
  refreshComments: () => void;
}

const EditContentBlock: React.FC<Props> = ({
  editType,
  articleId,
  onClose,
  articleData,
  refreshArticle,
  refreshComments,
}) => {
  const [isOpen, setIsOpen] = useState(false);
  const [style, setStyle] = useState("");
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("請輸入評論內容...");
  const [error, setError] = useState("");

  useEffect(() => {
    if (editType) {
      setIsOpen(true);
      setStyle("in");
      if (editType === "article") {
        setTitle(articleData.title);
        setContent(articleData.content);
      } else {
        setTitle("");
        setContent("請輸入評論內容...");
      }
    } else {
      setStyle("out");
      setTimeout(() => {
        setIsOpen(false);
        setStyle("");
      }, 300);
    }
    setError("");
  }, [editType, articleData.title, articleData.content]);

  const submitHandler = async () => {
    if (!title.trim() || !content.trim()) {
      return setError("標題和內容不可為空");
    }
    if (editType === "article") {
      const { status, id, message } = await updateArticle(articleId, {
        title,
        content,
      });
      if (status !== "success") {
        return setError(message || "伺服器異常");
      }
      refreshArticle();
      onClose();
    } else {
      const { status, id, message } = await addComment({
        articleId,
        title,
        content,
      });
      if (status !== "success") {
        return setError(message || "伺服器異常");
      }
      refreshComments();
      onClose();
    }
    setTitle("");
    setContent("");
    setError("");
  };

  return (
    <>
      <div
        className={`container ${style}`}
        onClick={(e) => e.stopPropagation()}
      >
        {!isOpen ? null : (
          <>
            <FadeIn duration={0.15} delayTime={0.3}>
              <div className="control">
                <label htmlFor="title">
                  {editType === "article" ? "文章" : "評論"}標題
                </label>
                <input
                  id="title"
                  type="text"
                  placeholder={`請輸入${
                    editType === "article" ? "文章" : "評論"
                  }標題`}
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                />
              </div>
            </FadeIn>
            <FadeIn duration={0.15} delayTime={0.45}>
              <div className="control">
                <label htmlFor="content">
                  {editType === "article" ? "文章" : "評論"}內容
                </label>
                <textarea
                  id="content"
                  rows={8}
                  cols={40}
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                ></textarea>
              </div>
            </FadeIn>
            <p style={{ color: "red", margin: "1rem 1rem 0 3rem" }}>{error}</p>
            <FadeIn duration={0.15} delayTime={0.6}>
              <div className="actions">
                <button onClick={submitHandler}>確認</button>
                <button onClick={onClose}>取消</button>
              </div>
            </FadeIn>
          </>
        )}
      </div>
      <style jsx>{`
        .container {
          visibility: ${isOpen ? "visible" : "hidden"};
          position: fixed;
          z-index: 900;
          top: -100%;
          right: 0;
          height: 100vh;
          width: 500px;
          background: white;
          transition: top 0.3s;
          display: flex;
          flex-direction: column;
          justify-content: center;
        }
        .in {
          top: 0;
        }
        .out {
          top: 100%;
        }
        .control {
          display: flex;
          flex-direction: column;
          margin-left: 3rem;
          margin-bottom: 1rem;
          align-items: flex-start;
          color: black;
          label {
            width: 100px;
            margin-bottom: 0.3rem;
          }
          input,
          textarea {
            color: black;
            background-color: #dadada;
            outline: none;
            border: 0px solid transparent;
            padding: 0.3rem 0.5rem;
            border-radius: 5px;
          }
          input::placeholder {
            color: black;
          }
        }
        .actions {
          display: flex;
          margin-top: 2rem;
          padding-left: 3rem;
          > button {
            display: block;
            padding: 0.2rem 1.5rem;
            border: 0px solid transparent;
            border-radius: 5px;
            cursor: pointer;
            &:first-of-type {
              margin-right: 0.5rem;
            }
            &:last-of-type {
              background-color: #740000;
            }
          }
        }
      `}</style>
    </>
  );
};

export default EditContentBlock;
