export const convertDate = (ISOstr: string) => {
  if (!ISOstr || ISOstr.trim() === "") {
    return "";
  }
  const [dateStr, timeStr] = ISOstr.split("T");
  const [year, month, date] = dateStr.split("-");
  const [hour, minute] = timeStr.split(":");
  return `${year}年${month}月${date}日 ${hour}:${minute}`;
};

const getPrefixTime = (time: number) => {
  return time < 10 ? "0" + time : time;
};
("2024-02-22T20:51:00Z");
