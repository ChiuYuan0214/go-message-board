import React from 'react';

interface Props {
  duration: number;
  delayTime?: number;
  children: React.ReactNode;
}

const FadeIn: React.FC<Props> = ({ duration, delayTime, children }) => {
  return (
    <>
      <div className="wrapper">{children}</div>
      <style jsx>{`
        .wrapper {
          opacity: 0;
          animation-name: fade-in;
          animation-timing-function: ease-in-out;
          animation-duration: ${duration}s;
          animation-fill-mode: forwards;
          animation-delay: ${delayTime || 0}s;
        }
        @keyframes fade-in {
          0% {
            opacity: 0;
          }
          100% {
            opacity: 1;
          }
        }
      `}</style>
    </>
  );
};

export default FadeIn;
