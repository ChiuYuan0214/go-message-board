import UpVoteIcon from "../../svgs/up-vote";

interface Props {
  voteUp: number;
  myScore: number;
  doVote: () => void;
}

const UpVote: React.FC<Props> = ({ voteUp, doVote, myScore }) => {
  return (
    <>
      <p>{voteUp}</p>
      <div
        onClick={(e) => {
          e.stopPropagation();
          doVote();
        }}
        className="icon"
      >
        <UpVoteIcon hasVote={myScore === 1} />
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

export default UpVote;
