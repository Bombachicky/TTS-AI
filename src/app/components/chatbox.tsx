"use client";
import { useState } from "react";
import axios from "axios";
import { UserMessage, AIMessage } from "./message";

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
        message
      );
      console.log(response.data);
    } catch (error) {
      console.log("Error: " + error);
    }
  };

  return (
    <>
      <div className="flex flex-col pt-8 items-center">
        <div className="flex flex-col w-full px-8">
          {/* This is where the chat logs will go */}
          <UserMessage message="Hello, I am Bomba" />
          <AIMessage message="[GPT Message]" />
          <UserMessage message="How are you?" />
          <AIMessage message="[GPT Message]" />
        </div>
      </div>
      <form className="flex justify-center fixed w-full mt-20">
        <input
          id="message"
          name="message"
          type="text"
          placeholder="Type a message"
          required
          className="w-2/3 h-10 px-4 rounded-xl"
        />
      </form>
    </>
  );
}

export default ChatBox;
