"use client";
import useSWR from "swr";

export interface Chat {
  id: string;
  user: string;
  lastMessage: string;
  time: string;
  avatar: string;
}

const mockData: Chat[] = [
  {
    id: "1",
    user: "Алиса",
    lastMessage: "Мы все еще встречаемся в пятницу?",
    time: "Вчера",
    avatar: "/alice-avatar.jpg",
  },
  {
    id: "2",
    user: "Боб",
    lastMessage: "Получил файлы, спасибо!",
    time: "10:15 утра",
    avatar: "/bob-avatar.jpg",
  },
  {
    id: "3",
    user: "Кэрол",
    lastMessage: "Можешь проверить отчет?",
    time: "08:30 утра",
    avatar: "/carol-avatar.jpg",
  },
];

const mockFetcher = async (url: string): Promise<Chat[]> => {
  await new Promise((resolve) => setTimeout(resolve, 1000));
  return mockData;
};

export function useChats() {
  const { data, error, isLoading } = useSWR<Chat[]>(
    `/api/v1/chats`,
    mockFetcher
  );
  // TODO

  return {
    data,
    isLoading,
    error,
  };
}
