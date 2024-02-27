import { SetState } from "@/types/general";

export type InputType = "text" | "email" | "password" | "phone" | "file";

interface Props {
  type: InputType;
  title: string | React.ReactElement;
  id: string;
  value: string;
  onChange?: (val: string) => void | SetState<string>;
  onFileChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
}

const Input: React.FC<Props> = ({
  type,
  title,
  id,
  value,
  onChange,
  onFileChange,
  placeholder,
}) => {
  return (
    <>
      <div className="control">
        <label className={type === "file" ? "isFile" : ""} htmlFor={id}>
          {title}
        </label>
        <input
          id={id}
          type={type}
          placeholder={placeholder || ""}
          value={value}
          onChange={
            type === "file"
              ? onFileChange
              : (e) => onChange && onChange(e.target.value)
          }
          accept={type === "file" ? "image/png, image/jpeg, image/webp" : ""}
        />
      </div>
      <style jsx>{`
        .control {
          display: flex;
          align-items: center;
          margin-bottom: 1rem;
          label {
            cursor: pointer;
            display: block;
            color: black;
            margin-right: 1rem;
            width: 100px;
            span {
              margin-left: 0.5rem;
              font-size: 0.6rem;
            }
            &.isFile {
              background-color: #505050;
              color: white;
              padding: 0.2rem 1rem;
              border-radius: 5px;
              &:hover {
                opacity: 0.7;
              }
            }
          }
          input {
            display: ${type === "file" ? "none" : "block"};
            outline: none;
            width: 250px;
            height: 30px;
            padding-left: 0.5rem;
          }
        }
      `}</style>
    </>
  );
};

export default Input;
