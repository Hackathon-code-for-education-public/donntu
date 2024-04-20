import useSWR from "swr";

export type UserData = {
  id: string;
  role: string;
  firstName: string;
  middleName: string;
  lastName: string;
};

// Mock fetcher function
const mockFetcher = async (url: string): Promise<UserData | null> => {
  // Simulate network delay
  await new Promise((resolve) => setTimeout(resolve, 1000));

  // Optionally, you can use the URL to differentiate responses if needed
  /*
  return {
    id: '1',
    role: 'applicant',
    firstName: 'Test',
    middleName: 'Test',
    lastName: 'Test',
  };*/

  /*
  const error = new Error('An error occurred while fetching the data.')
  error.status = 401;
  throw error
  */
};

export function useUser() {
  const { data, mutate, error, isLoading } = useSWR("/api/v1/user", mockFetcher);

  // const loading = !data && !error;
  const loggedOut = error && error.status === 401;

  return {
    loading: isLoading,
    loggedOut,
    user: data,
    mutate,
  };
}
