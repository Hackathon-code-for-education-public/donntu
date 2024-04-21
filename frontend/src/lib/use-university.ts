import useSWR from "swr";
import { fetcherWithData } from "./fetcher";
import { Univercity } from "@/api/university";

export function useUniversity(id: string) {
  const { data, error, isLoading } = useSWR<Univercity>(
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
