import { useContext, useEffect, useState } from "react";
import FollowList from "./FollowList";
import { Friend as FType } from "@/models/friend";
import { chatCtx } from "@/context/chat";
import Dialog from "@/components/UI/Dialog";
import { getFollowers, getFollows } from "@/api/follower";
import SearchBar from "./SearchBar";

interface Props {
  otherUserId: number;
}

const FollowSection: React.FC<Props> = ({ otherUserId }) => {
  const {
    followList,
    followerList,
    removeFollowFromList,
    removeFollowerFromList,
  } = useContext(chatCtx);
  const [selectedUser, setSelectedUser] = useState<FType | null>(null);
  const [title, setTitle] = useState("");
  const [isFollow, setIsFollow] = useState(true);
  const [otherFollowList, setOtherFollowList] = useState<FType[]>([]);
  const [otherFollowerList, setOtherFollowerList] = useState<FType[]>([]);

  const removeFollowHandler = async () => {
    if (!selectedUser || otherUserId) return;
    removeFollowFromList(selectedUser.userId);
    onCancel();
  };

  const removeFollowerHandler = async () => {
    if (!selectedUser || otherUserId) return;
    removeFollowerFromList(selectedUser.userId);
    onCancel();
  };

  const onCancel = () => {
    setSelectedUser(null);
    setTitle("");
  };

  const onFollowClick = (f: FType) => {
    if (otherUserId) return;
    setSelectedUser(f);
    setIsFollow(true);
    setTitle(`確定要取消追蹤${f.username}嗎?`);
  };

  const onFollowerClick = (f: FType) => {
    if (otherUserId) return;
    setSelectedUser(f);
    setIsFollow(false);
    setTitle(`確定要移除${f.username}的追蹤嗎?`);
  };

  useEffect(() => {
    if (!otherUserId) return;
    (async () => {
      const data = await getFollows(otherUserId);
      if (data.status !== "success") return;
      setOtherFollowList(data.list);
    })();
    (async () => {
      const data = await getFollowers(otherUserId);
      if (data.status !== "success") return;
      setOtherFollowerList(data.list);
    })();
  }, [otherUserId]);

  return (
    <>
      <SearchBar />
      <div className="container">
        <FollowList
          title="追蹤清單"
          isOther={!!otherUserId}
          list={otherUserId ? otherFollowList : followList}
          onClick={onFollowClick}
        />
        <FollowList
          title="粉絲清單"
          isOther={!!otherUserId}
          list={otherUserId ? otherFollowerList : followerList}
          onClick={onFollowerClick}
        />
      </div>
      <Dialog
        title={title}
        desc=""
        onConfirm={isFollow ? removeFollowHandler : removeFollowerHandler}
        onCancel={onCancel}
      />
      <style jsx>{`
        .container {
          display: flex;
          margin-top: 30px;
          width: 750px;
          padding: 0 2rem;
        }
      `}</style>
    </>
  );
};

export default FollowSection;
