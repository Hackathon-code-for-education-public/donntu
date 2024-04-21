"use client";
import useSWR from "swr";
import { ReviewData } from "./use-reviews";

export interface ReplyData {
  id: string;
  name: string;
  date: string;
  text: string;
}

export type ReviewDataFull = Omit<ReviewData, 'repliesCount'> & Partial<Pick<ReviewData, 'repliesCount'>>;

export interface ReviewWithReplies {
  review: ReviewDataFull;
  replies: ReplyData[];
}

const mockData: ReviewWithReplies = {
  review: {
    reviewId: "",
    authorStatus: "Некто",
    date: "",
    sentiment: "positive",
    text: "Текст Текст Текст Текст Текст Текст Текст Текст",
  },
  replies: [
    {
      id: "",
      name: "Аноним",
      date: "2020-02-20",
      text: "test",
    },
  ],
};

const mockFetcher = async (url: string): Promise<ReviewWithReplies> => {
  await new Promise((resolve) => setTimeout(resolve, 2000));
  return mockData;
};

const fetcher = (url: string) => fetch(url).then((r) => r.json());

export function useReview(id: string) {
  const { data, error, isLoading } = useSWR<ReviewWithReplies>(
    `/api/v1/reviews/${id}`,
    mockFetcher
  );
  // TODO

  return {
    data,
    isLoading,
    error,
  };
}
