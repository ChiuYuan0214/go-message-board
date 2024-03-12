import { IMAGE_PATH } from "@/constants/env";
import { Message as MType } from "@/models/message";

interface Props {
  userId: number;
  imagePath: string;
  targetImage: string;
  msg: MType;
}

const Message: React.FC<Props> = ({ userId, imagePath, targetImage, msg }) => {
  return (
    <>
      <li className={userId === msg.senderId ? "isSelf" : ""}>
        {userId !== msg.senderId && (
          <div className="image">
            <img src={IMAGE_PATH + targetImage} alt={targetImage} />
          </div>
        )}
        <p>{msg.content}</p>
        <div
          className={`arrow ${userId === msg.senderId ? "right" : "left"}`}
        ></div>
        {userId === msg.senderId && (
          <div className="image">
            <img src={IMAGE_PATH + imagePath} alt={imagePath} />
          </div>
        )}
      </li>
      <style jsx>{`
        li {
          position: relative;
          display: flex;
          justify-content: flex-start;
          padding: 0.3rem;
          &.isSelf {
            justify-content: flex-end;
            p {
              background-color: #c1d8dc;
            }
          }
          p {
            color: black;
            display: block;
            max-width: 170px;
            background-color: #d0cece;
            padding: 0.5rem;
            font-size: 0.8rem;
            border-radius: 5px;
            margin: 0.2rem 0.6rem;
          }
          .arrow {
            position: absolute;
            top: 15px;
            border-top: 5px solid transparent;
            border-bottom: 5px solid transparent;
            &.right {
              right: 46px;
              border-left: 10px solid #c1d8dc;
            }
            &.left {
              left: 46px;
              border-right: 10px solid #d0cece;
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
        }
      `}</style>
    </>
  );
};

export default Message;
