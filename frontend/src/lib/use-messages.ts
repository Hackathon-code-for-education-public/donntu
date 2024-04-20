"use client";
import useSWR from "swr";

export type Message = {
  sender: string;
  text: string;
  isYou: boolean;
};

const mockData: Message[] = [
  {
    sender: "Алиса",
    text: "Привет, когда мы встречаемся?",
    isYou: false,
  },
  {
    sender: "Вы",
    text: "Давай в пятницу в 6 вечера.",
    isYou: true,
  },
];

const mockFetcher = async (url: string): Promise<Message[]> => {
  await new Promise((resolve) => setTimeout(resolve, 1000));  
  return mockData;
};

const fetcher = (url: string) => fetch(url).then((r) => r.json());

export function useMessages(chatId?: string) {
  const { data, error, isLoading } = useSWR<Message[]>(
    chatId && `/api/v1/messages/${chatId}`,
    mockFetcher
  );
  // TODO

  return {
    data,
    isLoading,
    error,
  };
}