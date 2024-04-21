"use client";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { API } from "@/lib/api";
import { useChats } from "@/lib/use-chats";
import { useMessages } from "@/lib/use-messages";
import { useUser } from "@/lib/use-user";
import { PlusIcon } from "lucide-react";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { Suspense, useState } from "react";

interface Message {
  sender: string;
  text: string;
  isYou: boolean;
}

// Define props types for components
interface ChatHistoryProps {
  setActiveChatId: (str: string) => void;
}

interface ChatSectionProps {
  activeChatId?: string;
}

const ChatHistory: React.FC<ChatHistoryProps> = ({ setActiveChatId }) => {
  const { data: chats, error } = useChats();

  if (error) return <div>Failed to load chats</div>;
  if (!chats) return <div>Loading...</div>;

  return (
    <div className="border-r bg-gray-100/40 dark:bg-gray-800/40">
      <div className="flex h-[60px] items-center justify-between px-6">
        <h2 className="text-lg font-semibold">История чата</h2>
        <Button size="icon" variant="ghost">
          <PlusIcon className="h-5 w-5" />
          <span className="sr-only">Новый чат</span>
        </Button>
      </div>
      <div className="flex-1 overflow-y-auto p-4">
        <div className="space-y-4">
          {chats.map((chat, index) => (
            <div
              key={index}
              className="flex items-start gap-4 rounded-lg bg-white p-4 shadow-sm dark:bg-gray-950 cursor-pointer"
              /** @ts-ignore */
              onClick={() => setActiveChatId(chat.chatId)}
            >
              <div className="flex-1 space-y-1">
                <p className="text-sm font-medium">{chat.user}</p>
                <p className="text-sm text-gray-500 dark:text-gray-400">
                  {chat.lastMessage}
                </p>
                <div>Кавырлик</div>
              </div>
              <div className="text-xs text-gray-500 dark:text-gray-400">
                {chat.time}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

const ChatSection: React.FC<ChatSectionProps> = ({ activeChatId }) => {
  const { data: messages, error, mutate } = useMessages(activeChatId);
  const { user } = useUser();

  const onSendMessage = async (message: string) => {
    await API.sendChatMessage(activeChatId!, message);
    mutate(undefined);
  };

  if (!activeChatId) {
    return (
      <div className="flex flex-col justify-center items-center h-full">
        <p className="text-gray-500 text-lg">Выберите чат</p>
      </div>
    );
  }
  if (error) return <div>Failed to load messages</div>;
  if (!messages) return <div>Loading...</div>;

  return (
    <div className="flex flex-col">
      <div className="flex h-[60px] items-center justify-between border-b bg-gray-100/40 px-6 dark:bg-gray-800/40">
        <h2 className="text-lg font-semibold">Чат</h2>
      </div>
      <div className="flex-1 overflow-y-auto p-4 ">
        <div className="space-y-4">
          {messages.map((msg, index) => {
            /** @ts-ignore */
            const isYou = msg.userId === user?.id;
            /** @ts-ignore */
            return (
              <MessageBubble
                key={index}
                message={{
                  ...msg,
                  isYou,
                  /** @ts-ignore */
                  sender: isYou ? user?.firstName : "Кавырлик",
                }}
              />
            );
          })}
        </div>
      </div>
      <InputArea onSendChatCallback={onSendMessage} />
    </div>
  );
};

interface Message {
  sender: string;
  text: string;
  isYou: boolean;
}

interface MessageBubbleProps {
  message: Message;
}

const MessageBubble: React.FC<MessageBubbleProps> = ({ message }) => (
  <div
    className={`flex items-start gap-4 ${message.isYou ? "justify-end" : ""}`}
  >
    <div className={`space-y-1 ${message.isYou ? "text-right" : ""}`}>
      <p className="text-sm font-medium">{message.sender}</p>
      <p
        className={`text-sm rounded-lg ${
          message.isYou ? "bg-slate-100" : "bg-gray-100"
        } p-4`}
      >
        {/** @ts-ignore */}
        {message.message}
      </p>
    </div>
  </div>
);

const InputArea = ({
  onSendChatCallback,
}: {
  onSendChatCallback: (message: string) => void;
}) => {
  const [message, setMessage] = useState("");

  const onSendChat = async (message: string) => {
    if (message.length < 1) return;

    onSendChatCallback(message);
  };

  return (
    <div className="border-t bg-gray-100/40 px-6 py-4 dark:bg-gray-800/40">
      <div className="flex items-center gap-4">
        <Textarea
          className="flex-1 resize-none"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="Введите ваше сообщение..."
        />
        <Button
          onClick={(e) => onSendChat(message)}
          disabled={message.length < 1}
        >
          Отправить
        </Button>
      </div>
    </div>
  );
};

function Page() {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();

  const setActiveChatId = (chatId?: string) => {
    const params = new URLSearchParams(searchParams);
    if (chatId) {
      params.set('chatId', chatId);
    } else {
      params.delete('chatId');
    }
    replace(`${pathname}?${params.toString()}`);
  }

  const activeChatId = searchParams.get('chatId') || undefined;

  return (
    <div className="grid min-h-screen w-full grid-cols-[300px_1fr] overflow-hidden">
      <ChatHistory setActiveChatId={setActiveChatId} />
      <ChatSection activeChatId={activeChatId} />
    </div>
  );
}

export default function PPage() {
  return (
    <Suspense>
      <Page />
    </Suspense>
  );
}
