import FadeIn from "@/components/animation/FadeIn";
import React from "react";
import ReviewProgressBlock from "../components/ReviewProgressBlock/ReviewProgressBlock";
import { Count } from "@/models/auth";

interface Props {
  otherCount: Count;
}

const GeneralSection: React.FC<Props> = ({ otherCount }) => {
  return (
    <>
      <div className="general-section">
        <FadeIn duration={1} delayTime={0}>
          <ReviewProgressBlock otherCount={otherCount} />
        </FadeIn>
        {/* <FadeIn duration={1} delayTime={0.9}>
          <InformationBlock list={[]} onClick={() => {}} />
        </FadeIn> */}
      </div>
      <style jsx>{`
        .general-section {
          margin: 3rem 0 0 3rem;
          @media only screen and (max-width: 650px) {
            margin: 2rem 0;
          }
        }
      `}</style>
    </>
  );
};

export default GeneralSection;
