interface Props {
  title: string;
  desc: string;
  onConfirm: () => void;
  onCancel: () => void;
  confirmText?: string;
  cancelText?: string;
}

const Dialog: React.FC<Props> = ({
  title,
  desc,
  onConfirm,
  onCancel,
  confirmText,
  cancelText,
}) => {
  if (!title) return null;
  return (
    <>
      <section>
        <h3>{title}</h3>
        <p>{desc}</p>
        <div className="actions">
          <button onClick={onConfirm}>{confirmText || "Confirm"}</button>
          <button onClick={onCancel}>{cancelText || "Cancel"}</button>
        </div>
      </section>
      <div className="backdrop" onClick={onCancel}></div>
      <style jsx>{`
        section {
          display: flex;
          flex-direction: column;
          position: fixed;
          z-index: 1000;
          top: 50%;
          left: 50%;
          transform: translate(-50%, -50%);
          width: 500px;
          padding: 1rem;
          border-radius: 5px;
          background-color: white;
          color: black;
          .actions {
            margin-top: auto;
            button {
              margin-right: 1rem;
              cursor: pointer;
              padding: 0.2rem 1rem;
              &:last-of-type {
                background-color: #840505;
                border: 2px solid #840505;
                border-radius: 3px;
                &:hover {
                  opacity: 0.7;
                }
              }
            }
          }
        }
        .backdrop {
          position: fixed;
          top: 0;
          left: 0;
          z-index: 999;
          width: 100vw;
          height: 100vh;
          background-color: black;
          opacity: 0.5;
        }
      `}</style>
    </>
  );
};

export default Dialog;
