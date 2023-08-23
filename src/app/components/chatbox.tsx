"use client";
import { useState } from "react";
import axios from "axios";
import { UserMessage, AIMessage } from "./message";

function ChatBox() {
  const [message, setMessage] = useState("");
  const [AImessage, setAIMessage] = useState("");

  const handleMessage = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const message = formData.get("message") as string;
    setMessage(message);

    try {
      const response = await axios.post(
        "http://localhost:3000/message",
        JSON.stringify(message)
      );
      setAIMessage(response.data);
    } catch (error) {
      console.log("Error: " + error);
    }
  };

  return (
    <>
      <div className="flex flex-col fixed z-20 justify-end items-center w-full h-fit py-10 bottom-0 backdrop-blur-sm">
        <form
          className="flex flex-row justify-center w-2/3 backdrop-blur-lg"
          onSubmit={handleMessage}>
          <input
            id="message"
            name="message"
            type="text"
            placeholder="Type a message"
            required
            className="w-full h-10 px-4 rounded-xl"
          />
        </form>
      </div>
      <div className="flex flex-col pt-8 items-center">
        <div className="flex flex-col w-full px-8">
          {/* This is where the chat logs will go */}
          <UserMessage message={message} />
          <AIMessage message={AImessage} />
        </div>
      </div>
    </>
  );
}

export default ChatBox;
