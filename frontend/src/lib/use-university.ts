import useSWR from "swr";
import { fetcherWithData } from "./fetcher";

export type UniversityData = {
  id: string;
  name: string;
  longName: string;
  logo: string;
};

export function useUniversity(id: string) {
  const { data, error, isLoading } = useSWR<UniversityData>(
    `/api/v1/universities/${id}`,
    fetcherWithData
  );
  // TODO

  return {
    data,
    isLoading,
    error,
  };
}
