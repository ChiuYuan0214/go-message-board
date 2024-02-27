import DownVoteIcon from "../../svgs/down-vote";

interface Props {
  downVote: number;
  myScore: number;
  doVote: () => void;
}

const DownVote: React.FC<Props> = ({ downVote, myScore, doVote }) => {
  return (
    <>
      <p>{downVote}</p>
      <div
        onClick={(e) => {
          e.stopPropagation();
          doVote();
        }}
        className="icon"
      >
        <DownVoteIcon hasVote={myScore === -1} />
      </div>
      <style jsx>{`
        p {
          margin: 0;
          width: 1rem;
          font-size: 0.8rem;
        }
        .icon {
          margin-right: 1rem;
          width: 20px;
          height: 20px;
          cursor: pointer;
        }
      `}</style>
    </>
  );
};

export default DownVote;
