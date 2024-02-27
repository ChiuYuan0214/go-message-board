import React from 'react';
import BlockTitle from '../../UI/BlockTitle';
import { InformData } from './InformationItem';
import InformationList from './InformationList';

interface Props {
  list: InformData[];
  onClick: () => void;
}

const InformationBlock: React.FC<Props> = ({ list, onClick }) => {
  return (
    <>
      <div className="block">
        <BlockTitle iconType="Bell" text="通知" />
        <div className="content">
          <InformationList list={list} onClick={onClick} />
        </div>
      </div>
      <style jsx>{`
        .block {
          margin: 3rem 0;
          .content {
            margin-top: 1.3rem;
            width: 100%;
          }
        }
      `}</style>
    </>
  );
};

export default InformationBlock;
