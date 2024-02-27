import { FormEvent } from "react";

interface Props {
  children: React.ReactNode;
  onSubmit: (e: FormEvent<HTMLFormElement>) => void;
  onCancel: () => void;
  error: string;
}

const Form: React.FC<Props> = ({ children, onSubmit, onCancel, error }) => {
  return (
    <>
      <form onSubmit={onSubmit}>
        {children}
        <p className="error">{error}</p>
        <div className="actions">
          <button type="submit">確認修改</button>
          <button type="button" onClick={onCancel}>
            取消修改
          </button>
        </div>
      </form>
      <style jsx>{`
        form {
          margin-left: 2rem;
          .error {
            height: 1rem;
            color: #e30000;
            font-size: 0.8rem;
          }
          .actions {
            display: flex;
            align-items: center;
            button {
              cursor: pointer;
              padding: 0.2rem 1rem;
              height: 30px;
              margin: 0 0.5rem 0 0;
            }
          }
        }
      `}</style>
    </>
  );
};

export default Form;
