import useRunner from '@/hooks/useRunner';
import React from 'react';

interface Props {
  num: number;
  color: string;
}

const RunnerBar: React.FC<Props> = ({ num, color }) => {
  const { curNum } = useRunner(num, 1500, 1500);
  return (
    <>
      <div className="progress-bar">
        <div className="progress-inner"></div>
      </div>
      <style jsx>{`
        .progress-bar {
          margin: 0.5rem 0;
          width: 100%;
          height: 0.7rem;
          border-radius: 10px;
          background-color: #d9d9d9;
          box-shadow: inset 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
          .progress-inner {
            height: 100%;
            width: ${`${curNum}%`};
            border-radius: 10px;
            background-color: ${color};
            box-shadow: inset 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
          }
        }
      `}</style>
    </>
  );
};

export default RunnerBar;
