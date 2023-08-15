function ChatBox() {
  return (
    <>
      <div className="flex flex-col gap-28">
        <div className="text-white border-2 border-gray-800 rounded-xl m-8 h-96">
          <div>HIIIi</div>
          <div>ok</div>
        </div>
        <div className="flex flex-row justify-center">
          <input
            type="text"
            placeholder="Type a message"
            className="w-1/2 h-10 px-4 rounded-xl"
          />
        </div>
      </div>
    </>
  );
}

export default ChatBox;
