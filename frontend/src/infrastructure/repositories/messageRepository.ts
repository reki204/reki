import { apiClient } from "../apiClient";
import { Message } from "@/domain/Message";

export const messageRepository = {
  async getMessage(): Promise<Message> {
    try {
      const response = await apiClient.get("/message");
      const result = response.data;
      console.log(result);
      return result;
    } catch (err) {
      console.error("エラーが発生しました:", err);
      throw err;
    }
  },
};
