import { useEffect, useState } from "react";
import axios from "axios";

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
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);
  const [AImessage, setAImessage] = useState("");

  useEffect(() => {
    axios
      .post("http://localhost:3001/message", JSON.stringify(message))
      .then((res) => {
        setAImessage(res.data);
        setLoading(false);
      })
      .catch((err) => {
        setError(true);
        setLoading(false);
        setAImessage("Error in generating response");
      });
  }, [message]);

  return (
    <>
      {loading || error ? (
        loading ? (
          <div className="text-white">Loading...</div>
        ) : (
          <div className="text-white">{AImessage}</div>
        )
      ) : (
        <div className="flex justify-start py-2 px-2 animate-fadeDown">
          <div className="text-white bg-blue-600 rounded-xl p-4 max-w-2xl break-words">
            <div>{AImessage}</div>
          </div>
        </div>
      )}
    </>
  );
}
