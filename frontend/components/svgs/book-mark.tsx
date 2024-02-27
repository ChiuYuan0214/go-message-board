interface Props {
  color: string;
  isMark: boolean;
  onClick?: () => void;
}

const BookMarkIcon: React.FC<Props> = ({ color, isMark, onClick }) => {
  const markColor = color || "black";
  return (
    <svg
      onClick={onClick}
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      width="30"
      height="30"
      fill={isMark ? markColor : "none"}
      stroke={markColor}
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M6 2v20l6-6 6 6V2H6z" />
    </svg>
  );
};

export default BookMarkIcon;
