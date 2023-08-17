"use client";
import { useState } from "react";
import axios from "axios";

function ChatBox() {
  const [message, setMessage] = useState("");

  const handleMessage = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const message = formData.get("message") as string;
    console.log(message);

    try {
      const response = await axios.post(
        "https://5f0ek1er9i.execute-api.us-east-1.amazonaws.com/prod/users/message",
        {
          message: message,
        }
      );
      console.log(response.data);
    } catch (error) {
      console.log("Error: " + error);
    }
  };

  return (
    <>
      <div className="flex flex-col gap-28">
        <div className="text-white border-2 border-gray-800 rounded-xl m-8 h-96">
          <div>HIIIi</div>
          <div>ok</div>
        </div>
        <div>
          <form className="flex justify-center" onSubmit={handleMessage}>
            <input
              id="message"
              name="message"
              type="text"
              placeholder="Type a message"
              required
              className="w-2/3 h-10 px-4 rounded-xl"
            />
          </form>
        </div>
      </div>
    </>
  );
}

export default ChatBox;
