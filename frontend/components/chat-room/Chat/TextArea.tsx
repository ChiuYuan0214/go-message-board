import { chatCtx } from "@/context/chat";
import { FormEvent, useContext, useEffect, useState } from "react";

interface Props {
  targetId: number;
  scroll: (b: "instant" | "smooth") => void;
}

const TextArea: React.FC<Props> = ({ targetId, scroll }) => {
  const [text, setText] = useState("");
  const { sendMessage } = useContext(chatCtx);

  const submitHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (text === "") return;
    await sendMessage(targetId, text);
    setText("");
  };

  return (
    <>
      <form onSubmit={submitHandler}>
        <textarea
          onChange={(e) => setText(e.target.value)}
          value={text}
        ></textarea>
        <button type="submit">送出</button>
      </form>
      <style jsx>{`
        form {
          display: flex;
          align-items: center;
          padding: 0.3rem;
          textarea {
            flex-grow: 1;
            outline: none;
            padding: 0.3rem;
            border: 0px solid transparent;
            border-radius: 5px;
            background-color: #cbcbcb;
            color: black;
            font-size: 0.8rem;
          }
          button {
            cursor: pointer;
            margin-left: 0.5rem;
            padding: 0.2rem 0.7rem;
            border-radius: 5px;
            border: 0px solid transparent;
            opacity: 0.7;
            &:hover {
              opacity: 1;
            }
          }
        }
      `}</style>
    </>
  );
};

export default TextArea;
