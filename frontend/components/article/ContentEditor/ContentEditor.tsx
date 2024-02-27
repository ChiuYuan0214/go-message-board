import Dialog from "@/components/UI/Dialog";
import EditContentBlock, { EditType } from "./EditContentBlock";
import { Article } from "@/types/article";
import { useRouter } from "next/router";
import { deleteArticle } from "@/api/article";
import { deleteComment } from "@/api/comment";

export type DeleteType = "article" | "comment" | "";

interface Props {
  articleId: number;
  data: Article;
  editType: EditType;
  selectedComment: number;
  deleteType: DeleteType;
  cancelEditHandler: () => void;
  cancelDeletionHandler: () => void;
  refreshArticle: () => void;
  refreshComments: () => void;
}

const ContentEditor: React.FC<Props> = ({
  articleId,
  data,
  editType,
  selectedComment,
  deleteType,
  cancelEditHandler,
  cancelDeletionHandler,
  refreshArticle,
  refreshComments,
}) => {
  const router = useRouter();

  const doDeleteArticle = async () => {
    const { status, message } = await deleteArticle(articleId);
    if (status !== "success") {
      alert(message);
    } else {
      router.push("/profile?mode=article");
    }
  };

  const doDeleteComment = async () => {
    if (!selectedComment) return;
    const { status, message } = await deleteComment(selectedComment);
    if (status !== "success") {
      alert(message);
    } else {
      cancelDeletionHandler();
      refreshComments();
    }
  };

  return (
    <>
      <EditContentBlock
        editType={editType}
        articleId={articleId}
        articleData={data}
        onClose={cancelEditHandler}
        refreshArticle={refreshArticle}
        refreshComments={refreshComments}
      />
      <Dialog
        title={
          deleteType === "article"
            ? "確定要刪除此篇文章嗎？"
            : deleteType === "comment"
            ? "確定要刪除這則評論嗎？"
            : ""
        }
        desc=""
        confirmText="確認"
        cancelText="取消"
        onConfirm={deleteType === "comment" ? doDeleteComment : doDeleteArticle}
        onCancel={cancelDeletionHandler}
      />
      <style jsx>{`
        section {
        }
      `}</style>
    </>
  );
};

export default ContentEditor;
function setComments(list: any) {
  throw new Error("Function not implemented.");
}
