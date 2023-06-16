import { messageService } from "@/services/messageService";

export default function Home() {
  messageService.getMessage();
  return (
    <>
      <h1>Hello World</h1>
    </>
  );
}
