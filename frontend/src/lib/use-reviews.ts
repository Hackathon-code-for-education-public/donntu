import useSWRInfinite from "swr/infinite";
import { fetcher } from "./fetcher";
import { API_HOST } from "./auth";

export type ReviewData = {
  id: string;
  authorStatus: string;
  date: string;
  sentiment: string;
  text: string;
  repliesCount: number;
};

// const fetcher = (url: string) => fetch(url).then((r) => r.json());

const fetcherWithData = (url: string) => fetcher(url).then(r => r.data);

export function useReviewsByUniversity(universityId: string) {
  const getKey = (pageIndex: number, previousPageData: []) => {
    if (previousPageData && !previousPageData.length) return null;

    if (pageIndex === 0) return `/api/v1/reviews?universityId=${universityId}&offset=0&limit=10`;
    return `/api/v1/reviews?universityId=${universityId}&offset=${pageIndex * 10}&limit=10`;
  };

  const { data, size, setSize, isLoading, error } = useSWRInfinite(getKey, fetcherWithData);

  // Return the necessary values
  return {
    data: data ? data.flat() : data,
    isLoading,
    error,
    loadMore: () => setSize(size + 1)
  };
}
