import axios from "axios";
import { MessageType } from "@/types";

const API_URL = "http://localhost:8080/api/v1/";

export const getMessage = async (): Promise<MessageType> => {
  try {
    const response = await axios.get(API_URL);
    const result = response.data;
    console.log(result);
    return result;
  } catch (error) {
    console.error("エラーが発生しました:", error);
    throw error;
  }
};
