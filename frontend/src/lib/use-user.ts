import useSWR from "swr";
import { fetcherWithData } from "./fetcher";

export type UserData = {
  id: string;
  email: string;
  role: string;
  firstName: string;
  middleName: string;
  lastName: string;
};

export function useUser() {
  const { data, mutate, error, isLoading } = useSWR<UserData>("/api/v1/profile", fetcherWithData);

  // const loading = !data && !error;
  // const loggedOut = error && error.status === 401;

  // const loggedOut = !data && !error && !isLoading;

  const loggedOut = !data //&& !isLoading;

  const loading = !data && !error;

  return {
    loading,
    loggedOut,
    user: data,
    mutate,
  };
}
