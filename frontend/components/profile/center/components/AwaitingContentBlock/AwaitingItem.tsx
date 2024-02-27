import { subColor } from "@/constants/colors";
import React from "react";

interface Props {
  name: string;
  count: number;
  iconType: string;
}

const AwaitingItem: React.FC<Props> = ({ name, count, iconType }) => {
  return (
    <>
      <div className="review">
        <div className="icon"></div>
        <p>{`${name} (${count})`}</p>
      </div>
      <style jsx>{`
        .review {
          display: flex;
          .icon {
            display: flex;
            align-items: center;
            justify-content: center;
          }
          p {
            margin-left: 0.2rem;
            color: ${subColor};
            font-family: "Poppins";
            font-weight: 500;
            font-size: 14px;
          }
          @media only screen and (max-width: 650px) {
            width: 50%;
            margin-bottom: 0.25rem;
          }
        }
      `}</style>
    </>
  );
};

export default AwaitingItem;
