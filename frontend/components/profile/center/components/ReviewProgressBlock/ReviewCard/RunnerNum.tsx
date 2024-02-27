import useRunner from "@/hooks/useRunner";
import React from "react";

interface Props {
  num: number;
}

const RunnerNum: React.FC<Props> = ({ num }) => {
  const { curNum } = useRunner(num, 1500, 1500);

  return <p className="num">{curNum}</p>;
};

export default RunnerNum;
