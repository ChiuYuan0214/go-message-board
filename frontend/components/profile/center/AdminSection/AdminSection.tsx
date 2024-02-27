import React from "react";
import FadeIn from "@/components/animation/FadeIn";
import AwaitingContentBlock from "../components/AwaitingContentBlock/AwaitingContentBlock";
import InformationBlock from "../components/InformationBlock/InformationBlock";

interface Props {}

const AdminSection: React.FC<Props> = ({}) => {
  return (
    <>
      <div className="admin-section">
        <FadeIn duration={1}>
          <AwaitingContentBlock />
        </FadeIn>
        <FadeIn duration={1} delayTime={0.3}>
          <InformationBlock list={[]} onClick={() => {}} />
        </FadeIn>
      </div>
      <style jsx>{`
        .admin-section {
          margin: 3rem 0 0 3rem;
          @media only screen and (max-width: 650px) {
            margin: 2rem 0;
          }
        }
      `}</style>
    </>
  );
};

export default AdminSection;
