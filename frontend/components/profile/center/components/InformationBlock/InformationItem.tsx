import { mainColor } from "@/constants/colors";
import React from "react";

export interface InformData {
  title: string;
  informTime: number;
  group: string;
  isRead: boolean;
}

interface Props {
  data: InformData;
}

const InformationItem: React.FC<Props> = ({ data }) => {
  const { title, informTime, group, isRead } = data;

  const displayedTime = informTime + "小時前";

  return (
    <>
      <div className={`item ${isRead ? "isRead" : ""}`}>
        <div className="circle"></div>
        <p className="title">{title}</p>
        <p className="sub-info">
          {displayedTime}・{group}
        </p>
      </div>
      <style jsx>{`
        .item {
          display: flex;
          align-items: center;
          font-family: "Noto Sans";
          margin: 0.8rem 1.2rem;
          cursor: pointer;
          .circle {
            border-radius: 50%;
            width: 0.6rem;
            height: 0.6rem;
            background-color: ${mainColor};
            margin: 0 0.8rem;
          }
          .title {
            font-weight: 600;
            font-size: 14px;
            color: #212121;
          }
          .sub-info {
            font-weight: 400;
            font-size: 12px;
            color: #787878;
          }
          &.isRead {
            .circle {
              background-color: #c4c4c4;
            }
            .title {
              font-weight: 400;
            }
          }
          @media only screen and (max-width: 1200px) {
            flex-direction: column;
            align-items: flex-start;
            .circle {
              display: none;
            }
          }
          @media only screen and (max-width: 750px) {
            margin: 0.8rem;
          }
        }
      `}</style>
    </>
  );
};

export default InformationItem;
