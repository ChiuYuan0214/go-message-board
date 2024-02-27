export interface Comment {
  commentId: number;
  userId: number;
  commenter: string;
  commenterImage: string;
  title: string;
  content: string;
  creationTime: string;
  voteUp: number;
  voteDown: number;
}
