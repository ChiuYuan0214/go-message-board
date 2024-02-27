import type { GetServerSideProps, NextPage } from "next";
import Head from "next/head";
import styles from "../styles/Home.module.css";
import ArticleList, {
  getType,
} from "@/components/articles/ArticleList/ArticleList";
import SideBar from "@/components/articles/SideBar/SideBar";
import FadeIn from "@/components/animation/FadeIn";
import AddNewContentBlock from "@/components/profile/center/components/AddNewContentBlock/AddNewContentBlock";
import AddArticleBlock from "@/components/articles/AddArticleBlock/AddArticleBlock";
import { useContext, useState } from "react";
import { parseCookies } from "@/utils/cookie";
import { useRouter } from "next/router";
import { authCtx } from "@/context/auth";
import { ArticleListData } from "@/models/article";
import { getData } from "@/components/articles/ArticleList/utils";

interface Props {
  articleList: ArticleListData[];
}

const Home: NextPage<Props> = (props) => {
  const [isEdit, setIsEdit] = useState(false);
  const {
    userInfo: { userId },
  } = useContext(authCtx);
  const router = useRouter();
  const type = router.query.type;

  return (
    <div className={styles.container}>
      <Head>
        <title>討論版</title>
        <meta
          name="description"
          content="message board created with react + golang"
        />
      </Head>
      <main className={styles.main} onClick={() => setIsEdit(false)}>
        <SideBar />
        <ArticleList
          userId={userId}
          type={type as string}
          data={props.articleList}
        />
        <div style={{ flexGrow: 1, height: "100vh" }}>
          <FadeIn duration={1}>
            <div className="add-article-block">
              <AddNewContentBlock
                toggle={(e: React.MouseEvent<HTMLDivElement>) => {
                  e.stopPropagation();
                  setIsEdit((prev) => (userId ? !prev : false));
                }}
              />
            </div>
            <AddArticleBlock isEdit={isEdit} onClose={() => setIsEdit(false)} />
          </FadeIn>
        </div>
      </main>
      <style jsx>{`
        .add-article-block {
          position: fixed;
          top: 3.6rem;
          transform: translateX(50%);
        }
      `}</style>
    </div>
  );
};

export const getServerSideProps: GetServerSideProps = async ({
  query,
  req,
}) => {
  const cookies = req.headers.cookie;
  const map = cookies ? parseCookies(cookies) : {};
  const userId = map["userId"] ? +map["userId"] : 0;
  const data = await getData(
    getType(query.type),
    isNaN(userId) ? 0 : userId,
    1,
    true
  );
  return { props: { articleList: data.list } };
};

export default Home;
