//@ts-ignore
//@ts-nocheck

import useSWR from "swr";
import { fetcher } from "./fetcher";

export type OpenDayData = {
  id: string;
  universityName: string;
  description: string;
  place: string;
  link: string;
};

type Response = {
  data: unknown
}

export function useOpenDaysByUniversity(id: string) {
  const { data, error, isLoading } = useSWR<Response>(
    `/api/v1/universities/open?universityId=${id}`,
    fetcher,
  );
  // TODO

  return {
    data: data ? data.data.map((x) => ({
      id: '',
      universityName: x.UniversityName,
      description: x.Description,
      place: x.Address,
      link: x.Link,
    })) : undefined,
    isLoading,
    error,
  };
}
