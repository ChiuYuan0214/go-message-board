import { updateImage } from "@/api/profile";
import Input from "@/components/UI/Input";
import { IMAGE_PATH } from "@/constants/env";
import { authCtx } from "@/context/auth";
import { useContext, useEffect, useState } from "react";

interface Props {
  open: boolean;
}

const ImageSection: React.FC<Props> = ({ open }) => {
  const { userInfo, getMyProfile } = useContext(authCtx);
  const [error, setError] = useState("");

  const onFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    setError("");
    const newFile = (e.target.files as FileList)[0];
    const { status, message } = await updateImage(newFile);
    if (status !== "success") {
      return setError(message);
    }
    getMyProfile(false);
  };

  useEffect(() => {
    if (!open) setError("");
  }, [open]);

  if (!open) return null;

  return (
    <>
      <section>
        <Input
          type="file"
          title="選擇圖片"
          id="image"
          onFileChange={onFileChange}
          value=""
        />
        <p style={{ color: "black", fontSize: "0.8rem" }}>圖檔不可超過5MB</p>
        <img src={IMAGE_PATH + userInfo.imagePath} alt={userInfo.imagePath} />
      </section>
      <p className="error">{error}</p>
      <style jsx>{`
        section {
        }
        img {
          max-width: 300px;
          max-height: 500px;
          display: ${userInfo.imagePath.split("?")[0] ? "block" : "none"};
        }
        .error {
          height: 1rem;
          color: #e30000;
          font-size: 0.8rem;
        }
      `}</style>
    </>
  );
};

export default ImageSection;
