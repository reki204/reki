import axios from "axios";

const apiUrl = process.env.NEXT_PUBLIC_API_URL;

if (!apiUrl) {
  throw new Error("API URL is not defined.");
};

export const apiClient = axios.create({
  baseURL: apiUrl,
  timeout: 1000,
});
