import useSWR from "swr";

export type UniversityData = {
  id: string;
  name: string;
  longName: string;
  logoUrl: string;
};

// Mock data array
const mockUniversity: UniversityData[] = [{
    id: '1',
    name: 'ДонНТУ',
    longName: 'Донецкий национальный технический университет',
    logoUrl: 'https://donntu.ru/sites/all/themes/donntu/logo.png',
}]

// Mock fetcher function
const mockFetcher = async (url: string): Promise<UniversityData[]> => {
  // Simulate network delay
  await new Promise((resolve) => setTimeout(resolve, 1000));

  // Optionally, you can use the URL to differentiate responses if needed
  return mockUniversity;
};

const fetcher = (url: string) => fetch(url).then((r) => r.json());

export function useUniversities(searchQuery: string) {
  const { data, error, isLoading } = useSWR<UniversityData[]>(
    `/api/v1/university?q=${searchQuery}`,
    mockFetcher, {
        keepPreviousData: true
    }
  );
  // TODO

  return {
    data,
    isLoading,
    error,
  };
}
