import React from "react";
import BlockTitle from "../../UI/BlockTitle";
import { mainColor, subColor } from "@/constants/colors";

interface Props {
  // data: { [key in Content]: number };
}

const AwaitingContentBlock: React.FC<Props> = ({}) => {
  // let count = Object.entries(data).reduce((sum, pair) => sum + pair[1], 0);
  return (
    <>
      <div className="block">
        <BlockTitle iconType="ReviewList" text="待審查" innerFill={subColor} />
        <div className="awaiting-info">
          <header>您有{0}則訊息</header>
          <div className="title">
            <h4>待審查的文章</h4>
          </div>
          <div className="list">
            {/* {itemList.map((name, index) => (
              <AwaitingItem
                key={name}
                name={itemMap[name]}
                count={data[name]}
                iconType={iconMap[name]}
              />
            ))} */}
          </div>
        </div>
      </div>
      <style jsx>{`
        .block {
          .awaiting-info {
            margin-top: 1.3rem;
            display: flex;
            flex-wrap: wrap;
            box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
            border-radius: 10px;
            font-family: "Noto Sans";
            overflow: hidden;
            header {
              width: 100%;
              padding: 1rem 1.5rem;
              background-color: ${mainColor};
              color: white;
              font-weight: 700;
              font-size: 14px;
            }
            .title,
            .list {
              margin: 2rem 0;
              height: 2.5rem;
              display: flex;
              align-items: center;
            }
            .title {
              width: 30%;
              border-right: 1px solid #c4c4c4;
              justify-content: center;
            }
            .list {
              display: flex;
              align-items: center;
              width: 70%;
              border-left: 1px solid #c4c4c4;
              justify-content: space-evenly;
            }
            @media only screen and (max-width: 650px) {
              .title,
              .list {
                height: 4rem;
              }
              .list {
                padding-left: 1rem;
                flex-wrap: wrap;
                justify-content: flex-start;
              }
            }
          }
        }
      `}</style>
    </>
  );
};

export default AwaitingContentBlock;
