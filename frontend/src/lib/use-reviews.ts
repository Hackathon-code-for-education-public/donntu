import useSWR from "swr";

export type ReviewData = {
  id: string;
  authorStatus: string;
  date: string;
  sentiment: string;
  text: string;
  repliesCount: number;
};

// Mock data array
const mockReviews: ReviewData[] = [
  {
    id: "1",
    authorStatus: "Студент этого вуза",
    date: "2023-04-01",
    sentiment: "positive",
    text: "Great product, fast shipping!",
    repliesCount: 2,
  },
  {
    id: "2",
    authorStatus: "Выпускник этого вуза",
    date: "2023-04-02",
    sentiment: "negative",
    text: "Took too long to arrive and the product was damaged.",
    repliesCount: 0,
  },
  {
    id: "3",
    authorStatus: "Отчисленный",
    date: "2023-04-03",
    sentiment: "neutral",
    text: "Product is okay, not what I expected but it works.",
    repliesCount: 1,
  },
  {
    id: "4",
    authorStatus: "Некто",
    date: "2023-04-03",
    sentiment: "neutral",
    text: `Какое-то время назад я училась в этом университете и созрела
    все-таки написать свой отзыв. Начну с неприятных моментов, а
    закончу уже плюсами. Меня отчислили из этого университета, нарушив
    при этом несколько пунктов регламента. (При этом я училась на
    платной основе). Как так получилось? Сейчас расскажу, но моя
    история может показаться небылицей и, если честно, то я сама в
    шоке с данной ситуации и никогда не предполагала, что это может
    произойти именно со мной :) У меня был преподаватель по статистике, у которой рейтинг на
    сайте, котором оценивают преподавателей, составлял (по крайней
    мере в то время) 1.3-2.0 из 5. Но это не суть, так вышло, что она
    отправила меня на комиссию, поскольку почему-то она меня
    возненавидела, хотя я была всегда вежливой и очень уважительно
    относилась ко всем. Может быть, ей именно это и не понравилось :)
    Честно, я не грув в этом предмете, но я старалась выполнять все
    работы и какие-то мои работы она оценивала достаточно хорошо :)`,
    repliesCount: 1,
  },
  
];

// Mock fetcher function
const mockFetcher = async (url: string): Promise<ReviewData[]> => {
  // Simulate network delay
  await new Promise((resolve) => setTimeout(resolve, 1000));

  // Optionally, you can use the URL to differentiate responses if needed
  return mockReviews;
};

const fetcher = (url: string) => fetch(url).then((r) => r.json());

export function useReviewsByUniversity(id: string) {
  const { data, error, isLoading } = useSWR<ReviewData[]>(
    `/api/v1/reviews?university=${id}`,
    mockFetcher
  );
  // TODO

  return {
    data,
    isLoading,
    error,
  };
}
