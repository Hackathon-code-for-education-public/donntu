"use client";
import useSWR from "swr";
import { fetcher, fetcherWithData } from "./fetcher";

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

export function useMessages(chatId?: string) {
  const { data, error, isLoading, mutate } = useSWR<Message[]>(
    chatId && `/api/v1/chats/history/${chatId}`,
    fetcher
  );
  // TODO

  return {
    data,
    isLoading,
    error,
    mutate
  };
}
