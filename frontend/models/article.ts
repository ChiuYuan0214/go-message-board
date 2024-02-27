export class ArticleListData {
  constructor(
    public articleId: number,
    public userId: number,
    public title: string,
    public content: string,
    public voteUp: number,
    public voteDown: number,
    public author: string,
    public authorImage: string,
    public myScore: number,
    public hasCollec: boolean,
    public commentTitle: string,
    public commentContent: string,
    public commentUser: string,
    public commentUserImage: string,
    public publishTime: string,
    public tags: string[]
  ) {}
}
