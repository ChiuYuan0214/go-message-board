import React from "react";
import InformationItem, { InformData } from "./InformationItem";
import { mainColor } from "@/constants/colors";

interface Props {
  list: InformData[];
  onClick: () => void;
}

const InformationList: React.FC<Props> = ({ list, onClick }) => {
  return (
    <>
      <div className="wrapper">
        <header>
          <p>您有{list.length}則通知</p>
          <button onClick={onClick}>全部顯示已讀</button>
        </header>
        <div className="list">
          {list.map((data) => (
            <InformationItem key={data.title} data={data} />
          ))}
        </div>
      </div>
      <style jsx>{`
        .wrapper {
          overflow: hidden;
          box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
          border-radius: 10px;
          font-family: "Noto Sans";
          header {
            height: 3rem;
            padding: 0 1.5rem;
            background-color: ${mainColor};
            display: flex;
            align-items: center;
            justify-content: space-between;
            color: white;
            font-size: 14px;
            p {
              font-weight: 700;
            }
            button {
              font-weight: 600;
              background-color: transparent;
              border: 0px solid transparent;
              padding: 0 1px;
              color: white;
              text-decoration: underline;
            }
          }
          .list {
            padding: 0.5rem 0;
          }
        }
      `}</style>
    </>
  );
};

export default InformationList;
