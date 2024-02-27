interface Props {}

import { useRouter } from "next/router";
import React, { useState } from "react";
import Tab, { TabData } from "@/components/profile/center/SideBar/Tab";

const tabData: TabData[] = [
  { name: "最新文章", icon: "none", id: "newest" },
  { name: "熱門文章", icon: "none", id: "hot" },
  { name: "最多觀看", icon: "none", id: "view" },
];

interface Props {
  tab: string;
}

const TabList: React.FC<Props> = ({ tab }) => {
  const [activeTab, setActiveTab] = useState(tab);
  const router = useRouter();

  const onClickHandler = (id: string) => {
    setActiveTab(id);
    router.push(`/?type=${id}`);
  };

  return (
    <>
      <div className="list">
        {tabData.map((d) => (
          <Tab
            key={d.id}
            tabData={d}
            activeTab={activeTab}
            onClick={() => onClickHandler(d.id)}
          />
        ))}
      </div>
      <style jsx>{`
        .list {
          margin-top: 6rem;
          width: 100%;
          font-family: "Poppins";
          font-weight: 600;
          color: white;
          @media only screen and (max-width: 650px) {
            margin-top: 1.5rem;
          }
        }
      `}</style>
    </>
  );
};

export default TabList;
