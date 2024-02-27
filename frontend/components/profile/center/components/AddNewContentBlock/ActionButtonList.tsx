import React from "react";
import ActionButton from "./ActionButton";

interface Props {
  list: {
    text: string;
    onClick: (e: React.MouseEvent<HTMLDivElement>) => void;
  }[];
}

const ActionButtonList: React.FC<Props> = ({ list }) => {
  return (
    <>
      <div className="list">
        {list.map((btn, i) => (
          <ActionButton key={i} text={btn.text} onClick={btn.onClick} />
        ))}
      </div>
      <style jsx>{`
        .list {
          display: flex;
          flex-direction: column;
          font-family: "Poppins";
          @media only screen and (max-width: 1100px) {
            flex-wrap: wrap;
          }
        }
      `}</style>
    </>
  );
};

export default ActionButtonList;
