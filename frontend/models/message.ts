export class Message {
  constructor(
    public senderId: number,
    public receiverId: number,
    public content: string,
    public time: number
  ) {}
}
