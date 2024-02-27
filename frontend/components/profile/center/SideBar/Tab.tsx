import { subColor } from "@/constants/colors";
import React from "react";

export interface TabData {
  name: string;
  icon: string;
  id: string;
}

interface Props {
  tabData: TabData;
  activeTab: string;
  onClick: () => void;
}

const Tab: React.FC<Props> = ({ tabData, activeTab, onClick }) => {
  const { name, id } = tabData;
  const isActive = id === activeTab;

  return (
    <>
      <div className={`tab ${isActive ? "active" : ""}`} onClick={onClick}>
        <h4>{name}</h4>
      </div>
      <style jsx>{`
        .tab {
          display: flex;
          padding: 0.8rem 2rem;
          cursor: pointer;
          h4 {
            margin: 1rem;
            font-size: 16px;
          }
          &.active {
            background-color: ${subColor};
            box-shadow: inset 0px 4px 4px rgba(0, 0, 0, 0.25);
          }
          @media only screen and (max-width: 850px) {
            padding: 0.8rem 1rem;
          }
          @media only screen and (max-width: 650px) {
          }
        }
      `}</style>
    </>
  );
};

export default Tab;
