import { SetState } from "@/types/general";
import { ModeType } from "../types";
import Tab from "./Tab";

interface Props {
  mode: ModeType;
  setMode: SetState<ModeType>;
}

const Tabs: React.FC<Props> = ({ mode, setMode }) => {
  return (
    <>
      <ul>
        <Tab
          type="info"
          title="修改個人資料"
          activeType={mode}
          onClick={setMode}
        />
        <Tab
          type="image"
          title="更新大頭貼"
          activeType={mode}
          onClick={setMode}
        />
        <Tab
          type="password"
          title="修改密碼"
          activeType={mode}
          onClick={setMode}
        />
      </ul>
      <style jsx>{`
        ul {
          display: flex;
          color: #5f5e5e;
          margin-bottom: 3rem;
        }
      `}</style>
    </>
  );
};

export default Tabs;
