import { useState } from "react";
import axios from "axios";

function ChatBox() {
  return (
    <>
      <div className="flex flex-col gap-28">
        <div className="text-white border-2 border-gray-800 rounded-xl m-8 h-96">
          <div>HIIIi</div>
          <div>ok</div>
        </div>
        <div>
          <form className="flex justify-center">
            <input
              id="message"
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
