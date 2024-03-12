import { IMAGE_PATH } from "@/constants/env";
import { Comment } from "@/types/comment";

interface Props {
  data: Comment;
  hasBorder: boolean;
  userId: number;
  onDelete: () => void;
}

const Comment: React.FC<Props> = ({
  data: { userId: commenterId, commenter, commenterImage, title, content },
  hasBorder,
  userId,
  onDelete,
}) => {
  return (
    <>
      <li className="comment-block">
        <div className="comment-info">
          <div className="image">
            <img src={IMAGE_PATH + commenterImage} alt={commenterImage} />
          </div>
          <p>{commenter}</p>
        </div>
        <div className="content">
          <h4>{title}</h4>
          <p>{content}</p>
        </div>
        <div className="delete-btn" onClick={onDelete}>
          X
        </div>
      </li>
      <style jsx>{`
        .comment-block {
          position: relative;
          display: flex;
          align-items: flex-end;
          padding: 1rem 0;
          border-top: ${hasBorder ? "1px dotted black" : ""};
          background-color: white;
          color: black;
          /* &:first-of-type {
            border-top-left-radius: 5px;
            border-top-right-radius: 5px;
          }
          &:last-of-type {
            border-bottom-right-radius: 5px;
            border-bottom-left-radius: 5px;
          } */
          .comment-info {
            display: flex;
            flex-direction: column;
            align-items: center;
            margin-top: 1rem;
            margin-right: 1rem;
            > p {
              margin: 0.5rem 0;
              font-size: 1rem;
              font-weight: bold;
              color: black;
            }
          }
          .image {
            display: flex;
            justify-content: center;
            align-items: center;
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
          .content {
            max-width: 80%;
            margin-left: 1rem;
            > p {
              font-size: 0.8rem;
            }
          }
          .delete-btn {
            display: none;
            position: absolute;
            top: 0.8rem;
            right: 0.3rem;
            font-size: 1.2rem;
            font-weight: 600;
            cursor: pointer;
          }
          &:hover {
            .delete-btn {
              display: ${userId === commenterId ? "block" : "none"};
            }
          }
        }
      `}</style>
    </>
  );
};

export default Comment;
