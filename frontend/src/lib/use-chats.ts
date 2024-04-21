"use client";
import useSWR from "swr";
import { fetcherWithData } from "./fetcher";

export interface Chat {
  id: string;
  user: string;
  lastMessage: string;
  time: string;
  avatar: string;
}

export function useChats() {
  const { data, error, isLoading } = useSWR<Chat[]>(
    `/api/v1/chats`,
    fetcherWithData
  );
  // TODO

  return {
    data,
    isLoading,
    error,
  };
}
