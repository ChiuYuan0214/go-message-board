import React from "react";
import ActionButtonList from "./ActionButtonList";

interface Props {
  toggle: (e: React.MouseEvent<HTMLDivElement>) => void;
}

const AddNewContentBlock: React.FC<Props> = ({ toggle }) => {
  const BUTTON_LIST = [{ text: "新增文章", onClick: toggle }];
  return (
    <>
      <div className="new-content-block">
        <div className="btn-list">
          <ActionButtonList list={BUTTON_LIST} />
        </div>
      </div>
      <style jsx>{`
        .new-content-block {
          .btn-list {
            margin-top: 1.3rem;
          }
        }
      `}</style>
    </>
  );
};

export default AddNewContentBlock;
