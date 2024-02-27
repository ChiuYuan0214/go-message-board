import React, { useContext } from "react";
import TabList from "./TabList";
import { mainColor, subColor } from "@/constants/colors";
import { IMAGE_PATH } from "@/constants/env";
import { authCtx } from "@/context/auth";
import { UserInfo } from "@/models/auth";

interface Props {
  userInfo: UserInfo;
}

const SideBar: React.FC<Props> = ({ userInfo: otherInfo }) => {
  const {
    userInfo: { username, imagePath },
  } = useContext(authCtx);
  const isOther = otherInfo.userId;
  return (
    <>
      <div className="side-bar">
        <div className="user-profile-card">
          <div className="user-icon">
            <img
              src={IMAGE_PATH + (isOther ? otherInfo.imagePath : imagePath)}
              alt="profile-image"
            />
          </div>
          <div className="info">
            <h3>{isOther ? otherInfo.username : username}</h3>
          </div>
        </div>
        <TabList defaultTab="general" otherInfo={otherInfo} />
      </div>
      <div style={{ height: "100vh", marginRight: "270px" }}></div>
      <style jsx>{`
        .side-bar {
          position: fixed;
          z-index: 950;
          width: 30%;
          max-width: 270px;
          height: 100vh;
          background-color: ${mainColor};
          .user-profile-card {
            height: 300px;
            padding-top: 30px;
            background-color: ${subColor};
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
            font-family: "Roboto";
            color: white;
            h3 {
              font-weight: 600;
              font-size: 28px;
              margin: 0.3rem 0;
            }
            p {
              font-weight: 400;
              font-size: 16px;
            }
            .user-icon {
              background-color: black;
              border-radius: 50%;
              overflow: hidden;
              width: 45%;
              height: 45%;
              display: flex;
              justify-content: center;
              align-items: center;
              img {
                display: ${isOther && !otherInfo.imagePath ? "none" : "block"};
                max-width: 100%;
                max-height: 100%;
                object-fit: cover;
                @media (min-aspect-ratio: 1/1) {
                  width: 100%;
                  height: auto;
                }
                @media (max-aspect-ratio: 1/1) {
                  width: auto;
                  height: 100%;
                }
              }
            }
          }
          @media only screen and (max-width: 650px) {
            max-width: 650px;
            width: 100%;
            min-height: 300px;
            .user-profile-card {
              height: 200px;
              flex-direction: row;
              .user-icon {
                border-radius: 50%;
                overflow: hidden;
                width: 30%;
              }
              .info {
                margin-left: 1.5rem;
              }
            }
          }
        }
      `}</style>
    </>
  );
};

export default SideBar;
