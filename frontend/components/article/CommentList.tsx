import { useContext } from "react";
import Comment from "./Comment";
import { authCtx } from "@/context/auth";

interface Props {
  data: Comment[];
  onDelete: (id: number) => void;
}

const CommentList: React.FC<Props> = ({ data, onDelete }) => {
  const {
    userInfo: { userId },
  } = useContext(authCtx);

  return (
    <>
      <h3>大家的評論</h3>
      {data.length < 1 ? (
        <p>目前還沒有評論...</p>
      ) : (
        <ul>
          {data.map((d, i) => (
            <Comment
              key={d.commentId}
              data={d}
              hasBorder={i > 0}
              onDelete={() => onDelete(d.commentId)}
              userId={userId}
            />
          ))}
        </ul>
      )}
      <style jsx>{`
        h3 {
          margin-top: 3rem;
          color: black;
        }
        p {
          margin-top: 3rem;
          color: #6d6d6d;
          font-size: 1rem;
        }
        ul {
          padding: 1rem;
          border-radius: 5px;
          overflow: hidden;
          background-color: white;
          box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
        }
      `}</style>
    </>
  );
};

export default CommentList;
