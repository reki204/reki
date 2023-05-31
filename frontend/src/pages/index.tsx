import { getMessage } from "@/services/getApi";

export default function Home() {
  getMessage();
  return (
    <>
      <h1>Hello World</h1>
    </>
  );
}
