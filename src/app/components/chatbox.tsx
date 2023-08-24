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
    const input = formData.get("message") as string;
    setMessage(input);

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
            name="message"
            type="text"
            placeholder="Type a message"
            required
            className="w-full h-10 px-4 rounded-xl"
          />
          <button
            type="submit"
            className="flex items-center justify-center w-10 h-10 ml-2 rounded-xl bg-blue-500 hover:bg-blue-600">
            <svg
              className="w-6 h-6 text-white fill-current"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24">
              <path d="M2 12l2-2h14v4H4l-2 2z" />
              <path d="M21 12c0-4.97-4.03-9-9-9s-9 4.03-9 9 4.03 9 9 9 9-4.03 9-9zm-2 0c0 3.86-3.14 7-7 7s-7-3.14-7-7 3.14-7 7-7 7 3.14 7 7z" />
            </svg>
          </button>
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
