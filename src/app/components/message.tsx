import { useEffect, useState } from "react";
import axios from "axios";

interface MessageProps {
  message: string;
}

export function UserMessage({ message }: MessageProps) {
  return (
    <>
      <div className="flex justify-end py-2 px-8 animate-fadeDown">
        <div className="text-black bg-gray-100 rounded-xl p-4 max-w-2xl break-words shadow-md shadow-gray-500">
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
      });
  }, []);

  return (
    <>
      {loading || error ? (
        loading ? (
          <div className="text-white">Loading...</div>
        ) : (
          <div className="flex justify-start py-2 px-8 animate-fadeDown">
            <div className="text-white bg-gray-600 rounded-xl p-4 max-w-lg break-words shadow-md shadow-gray-100">
              <div>Error in generating response</div>
            </div>
          </div>
        )
      ) : (
        <div className="flex justify-start py-2 px-8 animate-fadeDown">
          <div className="text-white bg-gray-600 rounded-xl p-4 max-w-lg break-words shadow-md shadow-gray-100">
            <div>{AImessage}</div>
          </div>
        </div>
      )}
    </>
  );
}
