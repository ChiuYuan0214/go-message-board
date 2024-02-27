import AddIcon from "@/components/svgs/add";
import { mainColor } from "@/constants/colors";
import React from "react";

interface Props {
  text: string;
  onClick: (e: React.MouseEvent<HTMLDivElement>) => void;
}

const ActionButton: React.FC<Props> = ({ text, onClick }) => {
  if (!text) return null;
  return (
    <>
      <div className="btn" onClick={onClick}>
        <div className="add-icon">
          <div>
            <AddIcon />
          </div>
        </div>
        <h4>{text}</h4>
      </div>
      <style jsx>{`
        .btn {
          cursor: pointer;
          display: flex;
          align-items: center;
          max-height: 80px;
          max-width: 200px;
          box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
          border-radius: 10px;
          padding: 1rem 1.8rem;
          margin-right: 1.5rem;
          margin-bottom: 1.5rem;
          opacity: 0.85;
          transition: opacity 0.5s ease-in-out;
          min-width: 250px;
          background-color: #f7fbff;
          .add-icon {
            width: 2rem;
            height: 2rem;
            border-radius: 50%;
            box-shadow: 2px 2px 2px rgba(0, 0, 0, 0.25);
            overflow: hidden;
            background-color: ${mainColor};
            display: flex;
            align-items: center;
            justify-content: center;
            > div {
              width: 70%;
              height: 70%;
              display: flex;
              justify-content: center;
              align-items: center;
              transition: transform 0.5s ease-in-out;
            }
          }
          h4 {
            margin-left: 0.5rem;
            font-weight: 600;
            font-size: 20px;
            color: ${mainColor};
          }
          &:hover {
            opacity: 1;
            .add-icon > div {
              transform: rotate(180deg);
            }
          }
          @media only screen and (max-width: 1100px) {
            width: 45%;
            margin-bottom: 1rem;
          }
          @media only screen and (max-width: 830px) {
            width: 100%;
            margin-right: 0;
          }
        }
      `}</style>
    </>
  );
};

export default ActionButton;
