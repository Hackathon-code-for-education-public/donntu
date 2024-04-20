import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { UniversityOpenDays } from "@/components/university-open-days";
import { UniversityReviews } from "@/components/university-reviews";
import Image from "next/image";

interface Params {
  id: string;
}

export default function Page({ params }: { params: Params }) {
  const universityName = "ДонНТУ";
  const universityLongName = "Донецкий национальный технический университет";

  return (
    <main className="min-h-screen">
      <div className="bg-white p-6 shadow-lg rounded-lg max-w-4xl mx-auto">
        <div className="flex items-center justify-between mb-4">
          <div className="flex items-center">
            <Image
              alt="Logo"
              className="h-12 w-12 mr-3"
              height="50"
              src="https://donntu.ru/sites/all/themes/donntu/logo.png"
              style={{
                aspectRatio: "50/50",
                objectFit: "contain",
              }}
              width="50"
            />
            <div>
              <h1 className="text-xl font-bold">{universityName}</h1>
              <h2 className="text-lg">{universityLongName}</h2>
              <div className="flex items-center mt-1"></div>
            </div>
          </div>
          <div className="flex items-center"></div>
        </div>
        <div className="flex items-center justify-between mb-6">
          <div className="flex items-center">
            <span>
              <span className="text-xl font-bold">8.1</span>
              <span className="text-sm">/10</span>
            </span>
          </div>
          <div className="text-sm text-gray-600">5703 оценок</div>
        </div>
        <Tabs defaultValue="about">
          <TabsList className="flex justify-between">
            <div className="flex">
              <TabsTrigger
                value="about"
                className="px-4 py-2 text-sm font-medium"
              >
                О вузе
              </TabsTrigger>
              <TabsTrigger
                value="reviews"
                className="px-4 py-2 text-sm font-medium text-gray-700"
              >
                Отзывы
              </TabsTrigger>
              <TabsTrigger
                value="open-day"
                className="px-4 py-2 text-sm font-medium text-gray-700"
              >
                Дни открытых дверей
              </TabsTrigger>
              <TabsTrigger
                value="dorm"
                className="px-4 py-2 text-sm font-medium text-gray-700"
              >
                Общежитие
              </TabsTrigger>
            </div>
          </TabsList>
          <TabsContent value="about" className="m-5">
            <div className="grid grid-cols-3 gap-4 text-center py-4 bg-gray-200 rounded-b-lg">
              <div>
                <div className="text-3xl font-bold">22</div>
                <div className="text-sm">направлений подготовки</div>
              </div>
              <div>
                <div className="text-3xl font-bold">59</div>
                <div className="text-sm">образовательных программ</div>
              </div>
              <div>
                <div className="text-3xl font-bold">1879</div>
                <div className="text-sm">бюджетных мест</div>
              </div>
              <div className="col-span-3">
                <div className="text-3xl font-bold">72.3</div>
                <div className="text-sm">средневзвешенный проходной балл</div>
              </div>
            </div>
            <div className="mt-4 text-sm">
              Результаты представленны самим ВУЗом.
            </div>
            <div className="mt-4"></div>
          </TabsContent>
          <TabsContent value="reviews" className="m-5">
            <h2 className="text-lg">Отзывы</h2>
            <UniversityReviews universityId={params.id} />
          </TabsContent>
          <TabsContent value="open-day" className="m-5">
            <h2 className="text-lg">Дни открытых дверей в {universityName}</h2>
            <UniversityOpenDays universityId={params.id} />
          </TabsContent>
          <TabsContent value="dorm">
            <div>В этом вузе есть общежитие</div>
            <div>
              Оценка общежития: <span className="font-bold">6.63</span>/10
            </div>
          </TabsContent>
        </Tabs>
      </div>
    </main>
  );
}
