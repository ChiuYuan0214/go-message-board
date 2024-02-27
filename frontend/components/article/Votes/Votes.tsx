import { vote } from "@/api/vote";
import { useState } from "react";
import UpVote from "./UpVote";
import DownVote from "./DownVote";

interface Props {
  articleId: number;
  voteUp: number;
  voteDown: number;
  myScore: number;
}

const Votes: React.FC<Props> = (props) => {
  const [upVote, setUpVote] = useState(props.voteUp);
  const [downVote, setDownVote] = useState(props.voteDown);
  const [myScore, setMyScore] = useState(props.myScore);

  const doVote = async (newScore: 1 | -1) => {
    vote({ sourceId: props.articleId, score: newScore, voteType: "article" });
    setMyScore((prev) => {
      if (prev === 0) {
        if (newScore === 1) {
          setUpVote((prev) => prev + 1);
        } else {
          setDownVote((prev) => prev - 1);
        }
      }
      if (prev === 1) {
        if (newScore === 1) {
          setUpVote((prev) => prev - 1);
        } else {
          setUpVote((prev) => prev - 1);
          setDownVote((prev) => prev + 1);
        }
      }
      if (prev === -1) {
        if (newScore === 1) {
          setUpVote((prev) => prev + 1);
          setDownVote((prev) => prev - 1);
        } else {
          setDownVote((prev) => prev - 1);
        }
      }
      return newScore;
    });
  };

  return (
    <>
      <div className="votes">
        <UpVote voteUp={upVote} myScore={myScore} doVote={() => doVote(1)} />
        <DownVote
          downVote={downVote}
          myScore={myScore}
          doVote={() => doVote(-1)}
        />
      </div>
      <style jsx>{`
        div {
          display: flex;
          justify-content: flex-end;
          align-items: center;
          position: relative;
          width: 150px;
          margin-left: auto;
          bottom: -35px;
        }
      `}</style>
    </>
  );
};

export default Votes;
