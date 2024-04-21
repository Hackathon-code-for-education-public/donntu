"use client";
import useSWR from "swr";
import { ReviewData } from "./use-reviews";
import { fetcherWithData } from "./fetcher";

export interface ReplyData {
  id: string;
  name: string;
  date: string;
  text: string;
}

export type ReviewDataFull = Omit<ReviewData, "repliesCount"> &
  Partial<Pick<ReviewData, "repliesCount">>;

export interface ReviewWithReplies {
  review: ReviewDataFull;
  replies: ReplyData[];
}

export function useReview(id: string) {
  const { data, error, isLoading } = useSWR<ReviewWithReplies>(
    `/api/v1/reviews/${id}`,
    fetcherWithData
  );
  // TODO

  return {
    data,
    isLoading,
    error,
  };
}
