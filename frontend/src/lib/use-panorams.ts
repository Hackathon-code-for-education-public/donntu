import useSWR from "swr";
import { fetcherWithData } from "./fetcher";
import { University } from "@/api/university";
import { Panorama } from "@/api/panorama";

export function usePanorams(id: string) {
  const { data, error, isLoading, mutate } = useSWR<Panorama[]>(
    id && `/api/v1/panoramas?universityId=${id}&category=Прочее`,
    fetcherWithData
  );

  return {
    data,
    isLoading,
    error,
    mutate
  };
}
