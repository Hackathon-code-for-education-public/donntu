import useSWR from "swr";
import { fetcherWithData } from "./fetcher";
import { University } from "@/api/university";

export function useUniversity(id: string) {
  const { data, error, isLoading } = useSWR<University>(
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
