import { ChangeEvent, useState } from "react";
import SearchList from "./SearchList";
import { Friend } from "@/models/friend";
import { getUsers } from "@/api/users";

interface Props {}

const SearchBar: React.FC<Props> = ({}) => {
  const [searchList, setSearchList] = useState<Friend[]>([]);

  const onChangeHandler = async (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.value.trim() === "") {
      return setSearchList([]);
    }
    const data = await getUsers(e.target.value);
    if (data.status === "success") {
      setSearchList(data.users);
    }
  };

  return (
    <>
      <section>
        <input type="text" placeholder="搜尋好友" onChange={onChangeHandler} />
        <SearchList list={searchList} />
      </section>
      <style jsx>{`
        section {
          margin: 100px 0 0 48px;
          position: relative;
          z-index: 1000;
          width: 250px;
          input {
            border-radius: 5px;
            outline: none;
            width: 100%;
            height: 30px;
          }
        }
      `}</style>
    </>
  );
};

export default SearchBar;
