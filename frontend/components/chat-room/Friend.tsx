import { IMAGE_PATH } from "@/constants/env";
import { Friend as FType } from "@/models/friend";
import React from "react";

interface Props {
  data: FType;
  onClick: () => void;
  isOther?: boolean;
  isSearch?: boolean;
}

const Friend: React.FC<Props> = ({
  data: { username, userImage, isOnline },
  onClick,
  isOther,
  isSearch,
}) => {
  return (
    <>
      <li onClick={onClick}>
        <div className="image">
          <img src={IMAGE_PATH + userImage} alt={userImage} />
        </div>
        <p className="name">{username}</p>
        {!isOther && !isSearch && (
          <div className="info">
            <p>{isOnline ? "已上線" : "未上線"}</p>
            <div className="mark"></div>
          </div>
        )}
      </li>
      <style jsx>{`
        li {
          cursor: pointer;
          color: black;
          height: 50px;
          margin: 0 0.5rem;
          padding: 0 0.5rem;
          display: flex;
          align-items: center;
          border-bottom: 0.5px dotted black;
          min-width: 200px;
          .image {
            background-color: #000000;
            width: 35px;
            height: 35px;
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
          .name {
            margin: 0 auto 0 0.5rem;
          }
          .info {
            display: flex;
            align-items: center;
            p {
              font-size: 0.8rem;
              opacity: ${isOnline ? "1" : "0.5"};
            }
            .mark {
              margin-left: 1rem;
              border-radius: 50%;
              width: 15px;
              height: 15px;
              background-color: ${isOnline ? "#04d404" : "#8f8f8f"};
            }
          }
        }
      `}</style>
    </>
  );
};

export default Friend;
