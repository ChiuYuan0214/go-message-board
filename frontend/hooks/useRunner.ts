import { useState, useEffect, useRef } from "react";

const getSum = (num: number) => {
  let sum = 0;
  const middle = num / 2;
  for (let i = 1; i <= middle; i++) {
    sum += i + (num - i + 1);
  }
  if (num % 2 !== 0) {
    sum += Math.floor(middle) + 1;
  }

  return sum;
};

const useRunner = (num: number, duration: number, delayTime?: number) => {
  const [curNum, setCurNum] = useState(0);
  const sumRef = useRef(getSum(num));

  useEffect(() => {
    sumRef.current = getSum(num);
    setCurNum(0);
  }, [num]);

  useEffect(() => {
    let timer: ReturnType<typeof setTimeout>;

    if (curNum < num) {
      const suffix = num - curNum < 5 ? curNum / 30 : 1;
      const curTimeout = ((duration * curNum) / sumRef.current) * suffix;
      const delay = curNum === 0 && delayTime ? delayTime : 0;
      timer = setTimeout(() => {
        setCurNum((prevNum) => prevNum + 1);
      }, curTimeout + delay);
    }

    return () => clearTimeout(timer);
  }, [curNum, delayTime, duration, num]);

  return { curNum };
};

export default useRunner;
