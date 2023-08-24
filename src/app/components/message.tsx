interface MessageProps {
  message: string;
}

export function UserMessage({ message }: MessageProps) {
  return (
    <>
      <div className="flex justify-end py-2 px-2 animate-fadeDown">
        <div className="text-black bg-gray-100 rounded-xl p-4 max-w-2xl break-words">
          <div>{message}</div>
        </div>
      </div>
    </>
  );
}

export function AIMessage({ message }: MessageProps) {
  return (
    <>
      <div className="flex justify-start py-2 px-2 animate-fadeDown">
        <div className="text-white bg-blue-600 rounded-xl p-4 max-w-2xl break-words">
          <div>{message}</div>
        </div>
      </div>
    </>
  );
}
