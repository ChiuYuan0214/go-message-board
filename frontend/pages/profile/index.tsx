import React, { useContext, useEffect, useState } from "react";
import { GetServerSideProps, GetServerSidePropsContext, NextPage } from "next";
import SideBar from "../../components/profile/center/SideBar/SideBar";
import GeneralSection from "../../components/profile/center/GeneralSection/GeneralSection";
import Head from "next/head";
import { authCtx } from "@/context/auth";
import { useRouter } from "next/router";
import FollowSection from "@/components/profile/center/FollowSection/FollowSection";
import MyArticleSection from "@/components/profile/center/MyArticleSection/MyArticleSection";
import MyCollectionSection from "@/components/profile/center/MyCollectionSection/MyCollectionSection";
import EditProfileSection from "@/components/profile/center/EditProfileSection/EditProfileSection";
import { getParam } from "@/utils/query";
import { getProfile } from "@/api/profile";
import { UserInfo, Count } from "@/models/auth";

interface Props {
  mode: string;
  userId: number;
}

const Profile: NextPage<Props> = ({ mode, userId: paramUserId }) => {
  const {
    userInfo: { userId, username },
    isAuthInit,
  } = useContext(authCtx);
  const router = useRouter();
  const isSelf = paramUserId === userId || !paramUserId;
  const [otherUserInfo, setOtherUserInfo] = useState<UserInfo>(new UserInfo());
  const [otherCount, setOtherCount] = useState<Count>(new Count());

  useEffect(() => {
    if (!isAuthInit) return;
    (async () => {
      if (!isSelf) {
        const { status, data } = await getProfile(paramUserId);
        if (status === "success") {
          setOtherUserInfo((prev) => ({ ...prev, ...data }));
          setOtherCount({
            article: data.articleCount,
            comment: data.commentCount,
            upVote: data.upVoteCount,
          });
        }
      } else {
        setOtherUserInfo(new UserInfo());
        setOtherCount(new Count());
      }
    })();
  }, [isAuthInit, isSelf, paramUserId]);

  useEffect(() => {
    if (!isAuthInit) return;
    if ((!userId || !username) && !paramUserId) {
      router.push("/login");
    }
  }, [userId, username, isAuthInit, router, paramUserId]);

  if (!isAuthInit || (!userId && isSelf) || (!isSelf && !otherUserInfo.userId))
    return null;

  return (
    <>
      <div className="container">
        <Head>
          <title>會員中心</title>
        </Head>
        <SideBar userInfo={otherUserInfo} />
        <section className="main-block">
          {mode === "general" && <GeneralSection otherCount={otherCount} />}
          {mode === "article" && (
            <MyArticleSection userId={otherUserInfo.userId || userId} />
          )}
          {mode === "fans" && (
            <FollowSection otherUserId={otherUserInfo.userId} />
          )}
          {mode === "collection" && isSelf && <MyCollectionSection />}
          {mode === "setting" && isSelf && <EditProfileSection />}
        </section>
      </div>
      <style jsx>{`
        .container {
          display: flex;
          background-color: #e3e2e2;
          .main-block {
            flex: 1;
            margin-top: 30px;
          }
          @media only screen and (max-width: 650px) {
            flex-direction: column;
          }
        }
      `}</style>
    </>
  );
};

export const getServerSideProps: GetServerSideProps = async ({
  req,
  res,
  query,
  locale,
}: GetServerSidePropsContext) => {
  const mode = getParam(query.mode, "general");
  const userId = getParam(query.userId, "0");
  return { props: { mode, userId: +userId } };
};

export default Profile;
