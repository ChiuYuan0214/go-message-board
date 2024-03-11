import { SOCKET_URI } from "@/constants/env";

interface Handlers {
  onOpen: (e: Event) => void;
  onMessage: (e: MessageEvent) => void;
  onClose: (e: CloseEvent) => void;
  onError: (e: Event) => void;
}

export class Socket {
  private socket: WebSocket | null = null;
  private handlers: Handlers | null = null;
  private timer: NodeJS.Timeout | null = null;
  constructor(token: string) {
    this.createSocket(token);
  }

  private createSocket(token: string) {
    this.socket = new WebSocket(SOCKET_URI + `?token=${token}`);
    this.timer = setInterval(
      () => this.socket?.send(JSON.stringify({ type: "ping" })),
      3000
    );
  }

  private async tryState() {
    let state = 0;
    const doTry = () =>
      new Promise<number>((res) => {
        setTimeout(() => res(this.socket?.readyState || 0), 1000);
      });
    while (!state) {
      state = await doTry();
    }
  }

  public async isOpen() {
    await this.tryState();
    return this.socket?.readyState === 1;
  }

  public async isClosed() {
    await this.tryState();

    return this.socket?.readyState === 3;
  }

  public async init(callback: () => void) {
    callback();
    if (await this.isOpen()) {
      return true;
    } else {
      this.logout();
      return false;
    }
  }

  public logout() {
    if (this.timer) {
      clearTimeout(this.timer);
    }
    this.socket?.close();
  }

  public initHandlers(handlers: Handlers) {
    this.handlers = handlers;
    if (!this.socket) return;
    this.socket.addEventListener("open", handlers.onOpen);
    this.socket.addEventListener("message", handlers.onMessage);
    this.socket.addEventListener("close", handlers.onClose);
    this.socket.addEventListener("error", handlers.onError);
  }

  private send(data: any) {
    this.socket?.send(JSON.stringify(data));
  }

  public refresh(token: string) {
    this.logout();
    this.createSocket(token);
    this.initHandlers(this.handlers!);
  }

  public getHistroy(senderId: number, receiverId: number, time: number) {
    this.send({
      userId: senderId,
      targetUserId: receiverId,
      type: "history",
      content: time + "",
    });
  }

  public sendMessage(senderId: number, receiverId: number, content: string) {
    this.send({
      userId: senderId,
      targetUserId: receiverId,
      type: "message",
      content,
    });
  }

  public addFollow(userId: number, targetUserId: number) {
    this.send({
      userId,
      targetUserId,
      type: "add-follow",
      content: "",
    });
  }

  public removeFollow(userId: number, targetUserId: number) {
    this.send({
      userId,
      targetUserId,
      type: "remove-follow",
      content: "",
    });
  }

  public removeFollower(userId: number, targetUserId: number) {
    this.send({
      userId,
      targetUserId,
      type: "remove-follower",
      content: "",
    });
  }
}
