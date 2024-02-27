import Head from "next/head";
import { useState } from "react";
import { ModeType } from "./types";
import Tabs from "./Tabs/Tabs";
import InfoSection from "./Sections/InfoSection";
import PasswordSection from "./Sections/PasswordSection";
import ImageSection from "./Sections/ImageSection";

interface Props {}

const EditProfileSection: React.FC<Props> = ({}) => {
  const [mode, setMode] = useState<ModeType>("info");

  return (
    <>
      <Head>
        <title>登入會員</title>
      </Head>
      <section className="edit-section">
        <Tabs mode={mode} setMode={setMode} />
        <InfoSection open={mode === "info"} />
        <PasswordSection open={mode === "password"} />
        <ImageSection open={mode === "image"} />
      </section>
      <style jsx>{`
        .edit-section {
          min-height: 100vh;
          margin: 3rem 0 0 3rem;
          display: flex;
          flex-direction: column;
        }
      `}</style>
    </>
  );
};

export default EditProfileSection;
