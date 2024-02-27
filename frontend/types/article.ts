export interface Article {
  articleId: number;
  userId: number;
  author: string;
  authorImage: string;
  title: string;
  content: string;
  topCommentId: number;
  edited: boolean;
  viewCount: number;
  voteUp: number;
  voteDown: number;
  myScore: number;
  hasCollec: boolean;
  publishTime: string;
  creationTime: string;
  updateTime: string;
  tags: string[];
}
