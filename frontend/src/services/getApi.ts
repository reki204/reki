import axios from "axios";
import { MessageType } from "@/types";

const url = "http://127.0.0.1:8080/api/v1/";

export const getMessage = async (): Promise<MessageType> => {
  // const result = await axios.get(url).then((response) => response.data);
  // // return result;
  // console.log("res", result.data);

  // return result.data;
  try {
    const response = await axios.get(url);
    const result = response.data;
    console.log("res", result);
    return result;
  } catch (error) {
    // エラーハンドリング
    console.error("エラーが発生しました:", error);
    throw error;
  }
};
