import ChatBox from "../components/chatbox";

function ChatPage() {
  let messageLog: string[] = [];
  return (
    <div>
      <ChatBox log={messageLog}/>
    </div>
  );
}

export default ChatPage;
