import { updatePassword } from "@/api/profile";
import Input from "@/components/UI/Input";
import { FormEvent, useEffect, useState } from "react";
import Form from "../Form/Form";

interface Props {
  open: boolean;
}

const PasswordSection: React.FC<Props> = ({ open }) => {
  const [oldPassword, setOldPassword] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");

  const initPassword = () => {
    setOldPassword("");
    setPassword("");
    setConfirmPassword("");
  };

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError("");
    if (oldPassword.trim() === "" || password.trim() === "") {
      return setError("新舊密碼不可為空");
    }
    if (password !== confirmPassword) {
      return setError("請重新確認新密碼");
    }
    const { status, message } = await updatePassword({
      oldPassword,
      newPassword: password,
    });
    if (status !== "success") {
      return setError(message || "伺服器異常");
    }
    return initPassword();
  };

  const resetHandler = () => {
    initPassword();
  };

  useEffect(() => {
    if (!open) initPassword();
  }, [open]);

  if (!open) return null;

  return (
    <>
      <Form onSubmit={onSubmit} onCancel={resetHandler} error={error}>
        <Input
          type="password"
          title={<>現有密碼</>}
          id="old-password"
          placeholder="請輸入你現在的密碼"
          value={oldPassword}
          onChange={setOldPassword}
        />
        <Input
          type="password"
          title={<>新密碼</>}
          id="password"
          placeholder="請輸入你的新密碼"
          value={password}
          onChange={setPassword}
        />
        <Input
          type="password"
          id="confirm-password"
          title="確認密碼"
          placeholder="請再次輸入你的新密碼"
          value={confirmPassword}
          onChange={setConfirmPassword}
        />
      </Form>
    </>
  );
};

export default PasswordSection;
