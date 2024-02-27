import ArticleList from "@/components/articles/ArticleList/ArticleList";

interface Props {}

const MyCollectionSection: React.FC<Props> = ({}) => {
  return (
    <>
      <div className="collection-section">
        <ArticleList
          userId={0}
          type="collection"
          noServerInit
          marginLeft="30px"
          data={[]}
        />
      </div>
      <style jsx>{`
        .collection-section {
        }
      `}</style>
    </>
  );
};

export default MyCollectionSection;
