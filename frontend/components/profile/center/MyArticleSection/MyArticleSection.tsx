import ArticleList from "@/components/articles/ArticleList/ArticleList";

interface Props {
  userId: number;
}

const MyArticleSection: React.FC<Props> = ({ userId }) => {
  return (
    <>
      <div className="article-section">
        <ArticleList
          userId={userId}
          type="profile"
          noServerInit
          marginLeft="30px"
          data={[]}
        />
      </div>
      <style jsx>{`
        .article-section {
        }
      `}</style>
    </>
  );
};

export default MyArticleSection;
