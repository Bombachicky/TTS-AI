"use client";
import React, { useState } from "react";
import { UserMessage, AIMessage } from "./message";

interface messageLog {
  log: string[];
}

function ChatBox({ log }: messageLog) {
  const [message, setMessage] = useState("");
  const [messagelog, setMessageLog] = useState(log);
  const [disable, setDisable] = useState(false);

  let disabled = disable ? " bg-gray-300" : "";

  const handleMessage = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    log.push(message);
    log.push(message);
    setMessageLog(log);

    let input = document.getElementById("message") as HTMLInputElement;
    input.value = "";

    setDisable(true);

    setTimeout(() => {
      setDisable(false);
    }, 2000);
  };

  let chat = messagelog.map((message: string, index: number) => {
    if (index % 2 === 0) {
      return <UserMessage message={message} />;
    }
    return <AIMessage message={message} />;
  });

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
            disabled={disable}
            className={"w-full h-10 px-4 rounded-xl" + disabled}
            onChange={(e) => setMessage(e.target.value)}
          />
          <button
            type="submit"
            className="flex items-center justify-center w-10 h-10 ml-2 rounded-xl bg-blue-500 hover:bg-blue-600">
            <svg
              className="text-white hover:rotate-90 duration-300"
              clip-rule="evenodd"
              fill-rule="evenodd"
              stroke-linejoin="round"
              stroke-miterlimit="2"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg">
              <path
                d="m10.211 7.155c-.141-.108-.3-.157-.456-.157-.389 0-.755.306-.755.749v8.501c0 .445.367.75.755.75.157 0 .316-.05.457-.159 1.554-1.203 4.199-3.252 5.498-4.258.184-.142.29-.36.29-.592 0-.23-.107-.449-.291-.591zm.289 7.563v-5.446l3.522 2.719z"
                fill-rule="nonzero"
              />
            </svg>
          </button>
        </form>
      </div>
      <div className="flex flex-col pt-8 items-center mb-32">
        <div className="flex flex-col w-full px-8">
          {/* This is where the chat logs will go */}
          {chat}
        </div>
      </div>
    </>
  );
}

export default ChatBox;
