import {
  login,
  register,
  resendVerificationCode,
  verifyCode,
} from "@/api/auth";
import {
  setLocalToken,
  setLocalTokenExpire,
  setLocalUserId,
} from "@/api/utils";
import CountDown from "@/components/UI/CountDown";
import Input from "@/components/UI/Input";
import { authCtx } from "@/context/auth";
import { setCookies } from "@/utils/cookie";
import { GetServerSideProps, NextPage } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
import { useContext, useEffect, useState } from "react";

interface Props {}

const Login: NextPage<Props> = ({}) => {
  const {
    userInfo: { username: authUsername },
    isAuthInit,
    getMyProfile,
  } = useContext(authCtx);
  const [isRegister, setIsRegister] = useState(false);
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [phone, setPhone] = useState("");
  const [job, setJob] = useState("");
  const [address, setAddress] = useState("");
  const [error, setError] = useState("");
  const [tempUserId, setTempUserId] = useState(0);
  const [veriCode, setVeriCode] = useState("");
  const [hasSentCode, setHasSentCode] = useState(false);
  const [codeExpireTime, setCodeExpireTime] = useState(0);
  const router = useRouter();

  const registerHandler = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (
      username.trim() === "" ||
      email.trim() === "" ||
      password.trim() === "" ||
      confirmPassword.trim() === ""
    ) {
      return setError("用戶名稱、信箱和密碼不可為空");
    }

    if (password !== confirmPassword) {
      return setError("請重新確認密碼是否輸入錯誤");
    }
    const data = await register({
      username,
      email,
      password,
      phone,
      job,
      address,
    });
    if (data.status !== "success") {
      return setError(data.message);
    }

    setHasSentCode(true);
    setTempUserId(data.userId);
    const start =
      (new Date(data.expireTime).getTime() - new Date().getTime()) / 1000;
    setCodeExpireTime(start);
  };

  const resendCodeHandler = async () => {
    const data = await resendVerificationCode({ email, password });
    const start =
      (new Date(data.expireTime).getTime() - new Date().getTime()) / 1000;
    setCodeExpireTime(start);
  };

  const setInfo = (data: {
    token: string;
    expireTime: number;
    userId: number;
  }) => {
    setLocalToken(data.token);
    setLocalTokenExpire(data.expireTime * 1000);
    setLocalUserId(data.userId);
    setCookies(tempUserId, data.token);
    getMyProfile(true);
    router.push("/profile");
  };

  const veriCodeHandler = async () => {
    const data = await verifyCode({ userId: tempUserId, code: +veriCode });
    if (data.status !== "success") {
      return setError(data.message);
    }
    setInfo({ userId: tempUserId, ...data });
  };

  const loginHandler = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (email === "" || password === "") {
      return setError("信箱和密碼不可為空");
    }
    const data = await login({ email, password });
    if (data.status !== "success") {
      return setError(data.message);
    }
    setInfo(data);
  };

  useEffect(() => {
    if (!isAuthInit) return;
    if (authUsername) {
      router.push("/profile");
    }
  }, [isAuthInit, authUsername, router]);

  if (!isAuthInit || authUsername) return null;

  return (
    <>
      <Head>
        <title>登入會員</title>
      </Head>
      <main className="login-page">
        <h1>{isRegister ? "註冊帳號" : "登入帳號"}</h1>
        <form onSubmit={isRegister ? registerHandler : loginHandler}>
          {isRegister && (
            <Input
              type="text"
              title={
                <>
                  用戶名稱<span>(必填)</span>
                </>
              }
              id="username"
              placeholder="請輸入你的名稱"
              value={username}
              onChange={setUsername}
            />
          )}
          <Input
            type="email"
            title={<>信箱{isRegister && <span>(必填)</span>}</>}
            id="email"
            placeholder="請輸入你的信箱"
            value={email}
            onChange={setEmail}
          />
          <Input
            type="password"
            title={<>密碼{isRegister && <span>(必填)</span>}</>}
            id="password"
            placeholder="請輸入你的密碼"
            value={password}
            onChange={setPassword}
          />
          {isRegister && (
            <>
              <Input
                type="password"
                id="confirm-password"
                title="確認密碼"
                placeholder="請輸入你的密碼"
                value={confirmPassword}
                onChange={setConfirmPassword}
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
              {hasSentCode && (
                <>
                  <div className="veri-code-block">
                    <input
                      type="text"
                      placeholder="請輸入驗證碼"
                      value={veriCode}
                      onChange={(e) => setVeriCode(e.target.value)}
                    />
                    <button type="button" onClick={veriCodeHandler}>
                      確認
                    </button>
                    <button type="button" onClick={resendCodeHandler}>
                      重新發送
                    </button>
                  </div>
                  {codeExpireTime && <CountDown start={codeExpireTime} />}
                </>
              )}
            </>
          )}
          <p className="error">{error}</p>
          {!hasSentCode && <button type="submit">確認送出</button>}
          <p
            className="toggle-mode"
            onClick={() => {
              setIsRegister((prev) => !prev);
              setError("");
            }}
          >
            {isRegister ? "已有帳號？直接登入" : "還沒有帳號？現在註冊"}
          </p>
        </form>
      </main>
      <style jsx>{`
        .login-page {
          min-height: 100vh;
          background-color: #d9d9d9;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
          h1 {
            color: #5f5e5e;
            letter-spacing: 1.5px;
          }
          .error {
            height: 1rem;
            color: #e30000;
            font-size: 0.8rem;
          }
          .veri-code-block {
            display: flex;
            align-items: center;
            input {
              color: black;
              background-color: white;
              border: 0px solid transparent;
              border-radius: 5px;
              padding-left: 1rem;
              outline: none;
              height: 30px;
              &::placeholder {
                color: #8e8e8e;
              }
            }
            button {
              height: 30px;
              margin: 0 0.5rem;
            }
          }
          button {
            cursor: pointer;
            padding: 0.5rem 1.5rem;
            margin: 1rem auto 0;
            display: block;
            background-color: #7c7c7c;
            outline: none;
            border-radius: 5px;
            border: 0px solid transparent;
          }
          .toggle-mode {
            cursor: pointer;
            color: #7c7c7c;
            text-align: center;
          }
        }
      `}</style>
    </>
  );
};

export const getServerSideProps: GetServerSideProps = async () => {
  return { props: {} };
};

export default Login;
