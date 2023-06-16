import { messageRepository } from "@/infrastructure/repositories/messageRepository";

export const messageService = {
  async getMessage() {
    const message = await messageRepository.getMessage();
    return message;
  },
};
