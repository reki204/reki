import axios from "axios";
import { MessageType } from "@/types";

const apiUrl = process.env.NEXT_PUBLIC_API_URL;

if (!apiUrl) {
  throw new Error("API URL is not defined.");
}

export const getMessage = async (): Promise<MessageType> => {
  try {
    const response = await axios.get(apiUrl);
    const result = response.data;
    console.log(result);
    return result;
  } catch (error) {
    console.error("エラーが発生しました:", error);
    throw error;
  }
};
