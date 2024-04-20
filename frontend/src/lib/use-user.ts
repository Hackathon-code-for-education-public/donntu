import useSWR from "swr";

export type UserData = {
  id: string;
  role: string;
  firstName: string;
  middleName: string;
  lastName: string;
};

// Mock fetcher function
const mockFetcher = async (url: string): Promise<UserData> => {
  // Simulate network delay
  await new Promise((resolve) => setTimeout(resolve, 1000));

  // Optionally, you can use the URL to differentiate responses if needed
  return {
    id: '1',
    role: 'applicant',
    firstName: 'Test',
    middleName: 'Test',
    lastName: 'Test',
  };
};

export function useUser() {
  const { data, mutate, error } = useSWR("/api/v1/user", mockFetcher);

  const loading = !data && !error;
  const loggedOut = error && error.status === 403;

  return {
    loading,
    loggedOut,
    user: data,
    mutate,
  };
}
