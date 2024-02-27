type Props = {};

const Loading: React.FC<Props> = () => {
  return (
    <>
      <div className="loader-ellips" id="infinite-loader-animation">
        <span className="loader-ellips__dot" />
        <span className="loader-ellips__dot" />
        <span className="loader-ellips__dot" />
        <span className="loader-ellips__dot" />
      </div>
      <style jsx>{`
        @keyframes reveal {
          from {
            transform: scale(0.001);
          }

          to {
            transform: scale(1);
          }
        }
        @keyframes slide {
          to {
            transform: translateX(1.5em);
          }
        }

        .loader-ellips {
          font-size: 20px;
          position: relative;
          width: 4em;
          height: 1em;
          margin: 25px auto;
          .loader-ellips__dot {
            display: block;
            width: 0.5em;
            height: 0.5em;
            border-radius: 0.5em;
            background: ${"grey"};
            position: absolute;
            animation-duration: 0.5s;
            animation-timing-function: ease;
            animation-iteration-count: infinite;
            &:nth-child(1) {
              animation-name: reveal;
              left: 0;
            }
            &:nth-child(2) {
              animation-name: slide;
              left: 0;
            }
            &:nth-child(3) {
              animation-name: slide;
              left: 1.5em;
            }
            &:nth-child(4) {
              animation-name: reveal;
              animation-direction: reverse;
              left: 3em;
            }
          }
        }
      `}</style>
    </>
  );
};

export default Loading;
