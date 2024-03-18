import React from "react";
import { mainColor } from "@/constants/colors";
import TabList from "./TabList";
import { useRouter } from "next/router";

interface Props {}

const SideBar: React.FC<Props> = ({}) => {
  const router = useRouter();
  const type = router.query?.type;
  const tab = Array.isArray(type) || !type ? "newest" : type;

  return (
    <>
      <div className="side-bar">
        <TabList tab={tab} />
      </div>
      <style jsx>{`
        .side-bar {
          position: fixed;
          top: 0;
          left: 0;
          width: 30%;
          max-width: 270px;
          height: 100vh;
          background-color: ${mainColor};
          @media only screen and (max-width: 650px) {
            display: none;
            max-width: 650px;
            width: 100%;
            min-height: 300px;
          }
        }
      `}</style>
    </>
  );
};

export default SideBar;
