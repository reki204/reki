export class Message {
  constructor(
    public message: Message,
  ) {}
}

export type MessageType = {
  message: string;
};
