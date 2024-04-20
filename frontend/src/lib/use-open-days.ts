import useSWR from "swr";

export type OpenDayData = {
  id: string;
  universityName: string;
  description: string;
  place: string;
  link: string;
};

// Mock data array
const mockReviews: OpenDayData[] = [
  {
    id: "1",
    universityName: "ДонНТУ",
    description: "День открытых дверей факультета компьютерных технологий и информатики",
    place: "Онлайн",
    link: "https://donntu.ru/news/id202401270929",
  },
  {
    id: "2",
    universityName: "ДонНТУ",
    description: "День открытых дверей факультета компьютерных технологий и информатики",
    place: "Онлайн",
    link: "https://donntu.ru/news/id202401270929",
  },
  {
    id: "3",
    universityName: "ДонНТУ",
    description: "День открытых дверей факультета компьютерных технологий и информатики",
    place: "Онлайн",
    link: "https://donntu.ru/news/id202401270929",
  },
  {
    id: "4",
    universityName: "ДонНТУ",
    description: "День открытых дверей факультета компьютерных технологий и информатики",
    place: "Онлайн",
    link: "https://donntu.ru/news/id202401270929",
  },
];

// Mock fetcher function
const mockFetcher = async (url: string): Promise<OpenDayData[]> => {
  // Simulate network delay
  await new Promise((resolve) => setTimeout(resolve, 1000));

  // Optionally, you can use the URL to differentiate responses if needed
  return mockReviews;
};

const fetcher = (url: string) => fetch(url).then((r) => r.json());

export function useOpenDaysByUniversity(id: string) {
  const { data, error, isLoading } = useSWR<OpenDayData[]>(
    `/api/v1/open-days?university=${id}`,
    mockFetcher
  );
  // TODO

  return {
    data,
    isLoading,
    error,
  };
}
