import { IMAGE_PATH } from "@/constants/env";
import { convertDate } from "@/utils/date";
import { useRouter } from "next/router";

interface Props {
  authorId: number;
  author: string;
  authorImage: string;
  publishTime: string;
}

const ArticleInfo: React.FC<Props> = ({
  authorId,
  author,
  authorImage,
  publishTime,
}) => {
  const router = useRouter();
  const datetime = convertDate(publishTime);

  return (
    <>
      <div className="info">
        <div
          className="author"
          onClick={() => router.push(`/profile?userId=${authorId}`)}
        >
          <div className="image">
            <img src={IMAGE_PATH + authorImage} alt={authorImage} />
          </div>
          <p>作者：{author}</p>
        </div>
        <p>發布日期：{datetime}</p>
      </div>
      <style jsx>{`
        .info {
          display: flex;
          align-items: flex-end;
          justify-content: space-between;
          margin-top: 0.5rem;
          > p {
            font-size: 0.7rem;
          }
        }
        .author {
          cursor: pointer;
          > p {
            margin: 0.5rem 0;
            font-size: 0.8rem;
            font-weight: bold;
            color: black;
          }
        }
        .image {
          background-color: #000000;
          border-radius: 50%;
          overflow: hidden;
          width: 45px;
          height: 45px;
          display: flex;
          justify-content: center;
          align-items: center;
          > img {
            display: block;
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
      `}</style>
    </>
  );
};

export default ArticleInfo;
