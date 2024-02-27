import { addArticle } from "@/api/article";
import FadeIn from "@/components/animation/FadeIn";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

interface Props {
  isEdit: boolean;
  onClose: () => void;
}

const AddArticleBlock: React.FC<Props> = ({ isEdit, onClose }) => {
  const [isOpen, setIsOpen] = useState(false);
  const [style, setStyle] = useState("");
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("請輸入文章內容...");
  const [publishTime, setPublishTime] = useState("");
  const [error, setError] = useState("");
  const router = useRouter();

  useEffect(() => {
    if (isEdit) {
      setIsOpen(true);
      setStyle("in");
    } else {
      setStyle("out");
      setTimeout(() => {
        setIsOpen(false);
        setStyle("");
      }, 300);
    }
  }, [isEdit]);

  const submitHandler = async () => {
    if (!title.trim() || !content.trim()) {
      return setError("標題和內容不可為空");
    }
    if (new Date(publishTime).getTime() < new Date().getTime() - 60 * 1000) {
      return setError("發布時間不可為過去時間");
    }
    const { status, id, message } = await addArticle({
      title,
      content,
      publishTime,
    });
    if (status !== "success") {
      return setError(message || "伺服器異常");
    }
    setTitle("");
    setContent("");
    setPublishTime("");
    setError("");
    router.push(`/articles/${id}`);
  };

  return (
    <>
      <div
        className={`container ${style}`}
        onClick={(e) => e.stopPropagation()}
      >
        {isOpen && (
          <>
            <FadeIn duration={0.15} delayTime={0.3}>
              <div className="control">
                <label htmlFor="title">文章標題</label>
                <input
                  id="title"
                  type="text"
                  placeholder="請輸入文章標題"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                />
              </div>
            </FadeIn>
            <FadeIn duration={0.15} delayTime={0.45}>
              <div className="control">
                <label htmlFor="content">文章內容</label>
                <textarea
                  id="content"
                  rows={8}
                  cols={40}
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                ></textarea>
              </div>
            </FadeIn>
            <FadeIn duration={0.15} delayTime={0.6}>
              <div className="control">
                <label htmlFor="time">發布時間</label>
                <input
                  id="time"
                  type="datetime-local"
                  value={publishTime}
                  onChange={(e) => setPublishTime(e.target.value)}
                />
              </div>
            </FadeIn>
            <p style={{ color: "red", margin: "1rem 1rem 0 3rem" }}>{error}</p>
            <FadeIn duration={0.15} delayTime={0.75}>
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

export default AddArticleBlock;
