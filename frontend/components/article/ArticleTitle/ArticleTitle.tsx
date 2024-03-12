import { addCollection, removeCollection } from "@/api/collections";
import BookMarkIcon from "@/components/svgs/book-mark";
import { useEffect, useState } from "react";

interface Props {
  userId: number;
  articleId: number;
  title: string;
  hasCollec: boolean;
}

const ArticleTitle: React.FC<Props> = ({
  userId,
  articleId,
  title,
  hasCollec,
}) => {
  const [hasMark, setHasMark] = useState(false);

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
    setHasMark(hasCollec);
  }, [hasCollec]);

  return (
    <>
      <div className="title">
        <h3>{title}</h3>
        <BookMarkIcon color="#434343" isMark={hasMark} onClick={doMark} />
      </div>
      <style jsx>{`
        .title {
          display: flex;
          align-items: center;
          h3 {
            margin-right: auto;
            cursor: pointer;
            overflow-wrap: break-word;
            max-width: 80%;
          }
          > svg {
            cursor: pointer;
          }
        }
      `}</style>
    </>
  );
};

export default ArticleTitle;
