import { SetState } from "@/types/general";
import { ModeType } from "../types";

interface Props {
  type: ModeType;
  title: string;
  activeType: string;
  onClick: SetState<ModeType>;
}

const Tab: React.FC<Props> = ({ type, title, activeType, onClick }) => {
  return (
    <>
      <li
        className={activeType === type ? "isActive" : ""}
        onClick={() => onClick(type)}
      >
        {title}
      </li>
      <style jsx>{`
        li {
          margin-right: 1rem;
          cursor: pointer;
          &.isActive {
            color: black;
          }
        }
      `}</style>
    </>
  );
};

export default Tab;
