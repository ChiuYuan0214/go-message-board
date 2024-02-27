import { useEffect, useState } from "react";

interface Props {
  start: number;
}

const CountDown: React.FC<Props> = ({ start }) => {
  const [count, setCount] = useState(Math.floor(start));

  useEffect(() => {
    const timer = setTimeout(() => setCount((prev) => prev - 1), 1000);
    if (count === 0) return clearTimeout(timer);
  }, [count]);

  useEffect(() => {
    setCount(Math.floor(start));
  }, [start]);

  return (
    <>
      <p>{count === 0 ? "已逾時，請重新發送" : `剩餘時間: ${count}秒`}</p>
      <style jsx>{`
        p {
          color: black;
        }
      `}</style>
    </>
  );
};

export default CountDown;
