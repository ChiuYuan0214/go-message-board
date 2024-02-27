import { FormEvent, useCallback, useContext, useEffect, useState } from "react";
import Form from "../Form/Form";
import { authCtx } from "@/context/auth";
import { updateProfile } from "@/api/profile";
import Input from "@/components/UI/Input";

interface Props {
  open: boolean;
}

const InfoSection: React.FC<Props> = ({ open }) => {
  const { userInfo, isAuthInit, getMyProfile } = useContext(authCtx);
  const [username, setUsername] = useState("");
  const [phone, setPhone] = useState("");
  const [job, setJob] = useState("");
  const [address, setAddress] = useState("");
  const [error, setError] = useState("");
  const initInfo = useCallback(() => {
    setUsername(userInfo.username);
    setPhone(userInfo.phone);
    setJob(userInfo.job);
    setAddress(userInfo.address);
  }, [userInfo]);

  const SubmitHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError("");

    if (username.trim() === "") {
      return setError("用戶名稱不可為空");
    }
    const { status, message } = await updateProfile({
      username,
      phone,
      job,
      address,
    });
    if (status !== "success") {
      return setError(message || "伺服器異常");
    }
    getMyProfile(false);
  };

  const resetHandler = () => {
    initInfo();
  };

  useEffect(() => {
    if (!isAuthInit) return;
    initInfo();
  }, [isAuthInit, initInfo]);

  useEffect(() => {
    if (!open) setError("");
  }, [open]);

  if (!open) return null;

  return (
    <>
      <Form onSubmit={SubmitHandler} onCancel={resetHandler} error={error}>
        <Input
          type="text"
          title={<>用戶名稱</>}
          id="username"
          placeholder="請輸入你的名稱"
          value={username}
          onChange={setUsername}
        />
        <Input
          type="phone"
          id="phone"
          title="電話"
          placeholder="請輸入你的電話號碼"
          value={phone}
          onChange={setPhone}
        />
        <Input
          type="text"
          id="job"
          title="工作"
          placeholder="請輸入你的工作"
          value={job}
          onChange={setJob}
        />
        <Input
          type="text"
          id="address"
          title="地址"
          placeholder="請輸入你的地址"
          value={address}
          onChange={setAddress}
        />
      </Form>
    </>
  );
};

export default InfoSection;
