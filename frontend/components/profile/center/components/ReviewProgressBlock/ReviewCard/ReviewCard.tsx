import React from "react";
import { mainColor, subColor } from "@/constants/colors";
import RunnerNum from "./RunnerNum";

const shallowBlue = "#C8E5FF";

const getContent = (count: number, type: string) => {
  switch (type) {
    case "article":
      return count === 0
        ? "你還沒有任何文章唷！寫點什麼吧！"
        : count < 5
        ? "你的文章還不多唷！需要多加油"
        : count < 10
        ? "你有一些文章了，但可以再多一點！"
        : "你是文章好手！";
    case "comment":
      return count === 0
        ? "你還沒有任何留言唷！跟大家多點互動吧！"
        : count < 3
        ? "你有一些留言了！可以多寫一點"
        : count < 10
        ? "你是個活潑的人！"
        : "你是專業評論家！";
    case "upVote":
      return count === 0
        ? "你還沒有任何讚唷！跟大家多點互動吧！"
        : count < 10
        ? "你有一些讚了！再接再厲"
        : count < 30
        ? "你被不少人喜歡！"
        : "你是人氣王！";
    default:
      return "";
  }
};

const getTag = (type: string, count: number) => {
  switch (type) {
    case "article":
      return count === 0
        ? "新手"
        : count < 5
        ? "微新手"
        : count < 10
        ? "文章老手"
        : "專業寫手";
    case "comment":
      return count === 0
        ? "新手"
        : count < 3
        ? "微新手"
        : count < 10
        ? "留言老手"
        : "專業評論家";
    case "upVote":
      return count === 0
        ? "默默無名"
        : count < 10
        ? "還算討喜"
        : count < 30
        ? "八面玲瓏"
        : "眾星捧月";
    default:
      return "";
  }
};

const getColor = (type: string, count: number) => {
  switch (type) {
    case "article":
      return count === 0
        ? "#f5568e"
        : count < 5
        ? "#ed56f5"
        : count < 10
        ? shallowBlue
        : "#eeb230";
    case "comment":
      return count === 0
        ? "#f5568e"
        : count < 3
        ? "#ed56f5"
        : count < 10
        ? shallowBlue
        : "#eeb230";
    case "upVote":
      return count === 0
        ? "#f5568e"
        : count < 10
        ? "#ed56f5"
        : count < 30
        ? shallowBlue
        : "#eeb230";
    default:
      return "";
  }
};
export interface ReviewData {
  title: string;
  count: number;
  type: string;
}

interface Props {
  isSelf: boolean;
  data: ReviewData;
}

const ReviewCard: React.FC<Props> = ({ isSelf, data }) => {
  const { title, count, type } = data;

  const content = getContent(count, type);

  return (
    <>
      <div className="card">
        <div className="info-btn">
          <div>
            <div></div>
          </div>
          <div>
            <div></div>
          </div>
          <div>
            <div></div>
          </div>
        </div>
        <h4>{title}</h4>
        <div className="status-section">
          <p className="status">數量</p>
          <div style={{ color: "#ea5e42", fontSize: "1.5rem" }}>
            <RunnerNum num={count} />
          </div>
        </div>
        {isSelf && <p className="content">{content}</p>}
        <div className="date-section">
          <div className="left-day">{getTag(type, count)}</div>
        </div>
      </div>
      <style jsx>{`
        .card {
          font-family: "Noto Sans";
          margin-right: 1rem;
          padding: 1.2rem 1.2rem;
          font-family: "Poppins";
          color: #212121;
          width: 33%;
          background-color: #f7fbff;
          box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
          border-radius: 10px;
          .info-btn {
            width: 10px;
            margin-left: auto;
            > div {
              width: 100%;
              > div {
                border-radius: 50%;
                background-color: ${mainColor};
                width: 2px;
                height: 2px;
                margin-left: auto;
                margin-bottom: 2px;
              }
              &:hover > div {
                transform: scale(2);
              }
            }
          }
          h4 {
            font-family: "Poppins";
            font-style: normal;
            font-weight: 600;
            font-size: 16px;
          }
          .content {
            margin: 0.7rem 0;
            font-style: normal;
            font-weight: 400;
            font-size: 14px;
            border-bottom: 2px solid #dadada;
            padding-bottom: 1.5rem;
          }
          .progress-bar {
            margin: 0.5rem 0;
            width: 100%;
            height: 0.7rem;
            border-radius: 10px;
            background-color: #d9d9d9;
            box-shadow: inset 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
            .progress-inner {
              height: 100%;
              border-radius: 10px;
              background-color: ${mainColor};
              box-shadow: inset 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
            }
          }
          .status-section {
            display: flex;
            justify-content: space-between;
            align-items: center;
            font-style: normal;
            font-weight: 600;
            font-size: 14px;
            color: ${subColor};
            .status {
            }
          }
          .date-section {
            margin-top: 1rem;
            display: flex;
            justify-content: space-between;
            align-items: center;
            .upload-date {
              font-style: normal;
              font-weight: 400;
              font-size: 14px;
              color: #787878;
            }
            .left-day {
              background-color: ${getColor(type, count)};
              padding: 0.4rem 0.8rem;
              border-radius: 25px;
              font-family: "Helvetica";
              font-style: normal;
              font-weight: 700;
              font-size: 12px;
              color: ${subColor};
            }
          }
          @media only screen and (max-width: 950px) {
            width: 100%;
            margin-bottom: 1rem;
          }
        }
      `}</style>
    </>
  );
};

export default ReviewCard;
