import { mainColor } from "@/constants/colors";
import React from "react";

interface Props {
  iconType: string;
  text: string;
  innerFill?: string;
}

const BlockTitle: React.FC<Props> = ({ iconType, text, innerFill }) => {
  return (
    <>
      <div className="title">
        <h3>{text}</h3>
      </div>
      <style jsx>{`
        .title {
          display: flex;
          align-items: center;
          font-family: "Poppins";
          h3 {
            margin-left: 0.5rem;
            color: ${mainColor};
            font-weight: 600;
            font-size: 18px;
          }
        }
      `}</style>
    </>
  );
};

export default BlockTitle;
