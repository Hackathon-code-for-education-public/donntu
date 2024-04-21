import useSWR from "swr";
import { fetcherWithData } from "./fetcher";

export type UniversityData = {
  id: string;
  name: string;
  longName: string;
  logo: string;
};


export function useUniversities(searchQuery: string) {
  const { data, error, isLoading } = useSWR<UniversityData[]>(
    `/api/v1/universities/search?name=${searchQuery}`,
    fetcherWithData, {
        keepPreviousData: true
    }
  );

  return {
    data,
    isLoading,
    error,
  };
}
