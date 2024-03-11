import FriendCmp from "@/components/chat-room/Friend";
import { Friend as FType } from "@/models/friend";
import { useRouter } from "next/router";

interface Props {
  list: FType[];
}

const SearchList: React.FC<Props> = ({ list }) => {
  const router = useRouter();
  if (list.length === 0) {
    return null;
  }
  return (
    <>
      <ul>
        {list.map((f) => (
          <FriendCmp
            key={f.userId}
            data={f}
            onClick={() => router.push(`/profile?userId=${f.userId}`)}
            isSearch
          />
        ))}
      </ul>
      <style jsx>{`
        ul {
          width: 100%;
          margin: 0;
          padding: 0;
          position: absolute;
          background-color: #bababa;
          border-radius: 0 0 5px 5px;
        }
      `}</style>
    </>
  );
};

export default SearchList;
