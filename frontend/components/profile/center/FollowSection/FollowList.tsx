import FriendCmp from "@/components/chat-room/Friend";
import { Friend } from "@/models/friend";

interface Props {
  list: Friend[];
  title: string;
  isOther: boolean;
  onClick: (f: Friend) => void;
}

const FollowList: React.FC<Props> = ({ title, list, onClick, isOther }) => {
  return (
    <>
      <div>
        <h3>{title}</h3>
        <ul>
          {list.map((f) => (
            <FriendCmp
              key={f.userId}
              data={f}
              onClick={() => onClick(f)}
              isOther={isOther}
            />
          ))}
        </ul>
      </div>
      <style jsx>{`
        div {
          width: 100%;
          height: 100%;
          padding: 0 1rem;
          border-radius: 0;
          border: 0px solid transparent;
          background-color: #e5e3e3;
          overflow: hidden;
          white-space: nowrap;
          &:first-of-type {
            border-right: 1px dotted black;
          }
        }
        h3 {
          margin: 0;
          text-align: center;
          color: #2d2d2e;
        }
        ul {
          text-decoration: none;
          list-style-type: none;
          padding: 0;
          background-color: white;
          border-radius: 5px;
          overflow: hidden;
        }
      `}</style>
    </>
  );
};

export default FollowList;
