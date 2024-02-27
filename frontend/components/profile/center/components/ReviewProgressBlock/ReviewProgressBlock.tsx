import React, { useContext } from "react";
import BlockTitle from "../../UI/BlockTitle";
import ReviewCard from "./ReviewCard/ReviewCard";
import { authCtx } from "@/context/auth";
import { getParam } from "@/utils/query";
import { useRouter } from "next/router";
import { Count } from "@/models/auth";

const getNameList = (
  isSelf: boolean
): { title: string; type: keyof Count }[] => [
  {
    title: `${isSelf ? "我" : "他"}的文章`,
    type: "article",
  },
  {
    title: `${isSelf ? "我" : "他"}的留言`,
    type: "comment",
  },
  {
    title: `${isSelf ? "我" : "他"}收到的讚`,
    type: "upVote",
  },
];

interface Props {
  otherCount: Count;
}

const ReviewProgressBlock: React.FC<Props> = (props) => {
  const {
    userInfo: { userId },
    profileCount,
  } = useContext(authCtx);
  const paramUserId = +getParam(useRouter().query.userId, "0");
  const isSelf = !paramUserId || userId === paramUserId;

  return (
    <>
      <div className="block">
        <BlockTitle iconType="Pie" text={`${isSelf ? "我" : "他"}的內容`} />
        <div className="content">
          {getNameList(isSelf).map((data) => (
            <ReviewCard
              key={data.type}
              isSelf={isSelf}
              data={{
                title: data.title,
                count: (isSelf ? profileCount : props.otherCount)[
                  data.type as keyof Count
                ],
                type: data.type,
              }}
            />
          ))}
        </div>
      </div>
      <style jsx>{`
        .block {
          margin: 3rem 0;
          .content {
            margin-top: 1.3rem;
            width: 100%;
            display: flex;
            @media only screen and (max-width: 950px) {
              flex-wrap: wrap;
            }
          }
        }
      `}</style>
    </>
  );
};

export default ReviewProgressBlock;
