"use client";

import { Panorama } from "@/api/panorama";
import dynamic from "next/dynamic";

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { Skeleton } from "@/components/ui/skeleton";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { UniversityOpenDays } from "@/components/university-open-days";
import { UniversityReviews } from "@/components/university-reviews";
import { useUniversity } from "@/lib/use-university";

interface Params {
  id: string;
}

const PanoramaView = dynamic(
  () => import('@/components/panorama').then(module => module) as any,
  { ssr: false },
) as any;

const mokPanoramas: Panorama[] = [
  {
    name: "Главный корпус",
    address: "г. Донецк, ул. Пушкина, д.1",
    type: "Корпус",
    loc1: "/room.jpg",
    loc2: "/mus.jpg",
  },
  {
    name: "Корпус №8",
    address: "г. Донецк, ул. Артема, д.2",
    type: "Корпус",
    loc1: "/alma.jpg",
    loc2: "/lib.jpg",
  },
  {
    name: "Общежитие №3",
    address: "г. Донецк, ул. И. Ткаченко, д.3",
    type: "Общежитие",
    loc1: "/mus.jpg",
    loc2: "/lib.jpg",
  },
  {
    name: "Стадион №3",
    address: "г. Донецк, ул. Университетская, д.3",
    type: "Прочее",
    loc1: "/mus.jpg",
    loc2: "/lib.jpg",
  },
];

export default function Page({ params }: { params: Params }) {
  const { data, isLoading, error } = useUniversity(params.id);

  return (
    <main className="min-h-screen">
      <div className="bg-white p-6 shadow-lg rounded-lg max-w-4xl mx-auto">
        <div className="flex items-center justify-between mb-4">
          <div className="flex items-center w-full">
            {isLoading ? (
              <Skeleton className="h-12 w-12 mr-3" />
            ) : (
              <img
                alt="Logo"
                className="h-12 w-12 mr-3"
                height="50"
                src={data?.logo}
                style={{
                  aspectRatio: "50/50",
                  objectFit: "contain",
                }}
                width="50"
              />
            )}
            {isLoading ? (
              <div className="w-full">
                <Skeleton className="h-6 w-1/4" />
                <Skeleton className="h-6 w-1/2 mt-2" />
              </div>
            ) : (
              <div>
                <h1 className="text-xl font-bold">{data?.name}</h1>
                <h2 className="text-lg">{data?.longName}</h2>
              </div>
            )}
          </div>
        </div>
        <Tabs defaultValue="about">
          <TabsList className="flex w-full">
            <TabsTrigger
              value="about"
              className="px-4 py-2 text-sm font-medium flex-1"
            >
              О вузе
            </TabsTrigger>
            <TabsTrigger
              value="reviews"
              className="px-4 py-2 text-sm font-medium flex-1"
            >
              Отзывы
            </TabsTrigger>
            <TabsTrigger
              value="open-day"
              className="px-4 py-2 text-sm font-medium flex-1"
            >
              Дни открытых дверей
            </TabsTrigger>
            <TabsTrigger
              value="panorams"
              className="px-4 py-2 text-sm font-medium flex-1"
            >
              Панорамы
            </TabsTrigger>
          </TabsList>
          <TabsContent value="about" className="m-5">
            {/*
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
            </div>
            </div>
            */}
          </TabsContent>
          <TabsContent value="reviews" className="m-5">
            <h2 className="text-lg">Отзывы</h2>
            <UniversityReviews universityId={params.id} />
          </TabsContent>
          <TabsContent value="open-day" className="m-5">
            {isLoading ? (
              <Skeleton className="h-6 w-1/4" />
            ) : (
              <h2 className="text-lg">Дни открытых дверей в {data?.name}</h2>
            )}
            <UniversityOpenDays universityId={params.id} />
          </TabsContent>
          <TabsContent value="panorams">
            <Accordion type="single" collapsible className="w-full">
              <AccordionItem value="item-1">
                <AccordionTrigger>Корпуса</AccordionTrigger>
                <AccordionContent>
                  {mokPanoramas
                    .filter((item) => item.type === "Корпус")
                    .map((panorama) => (
                      <PanoramaView panorama={panorama} />
                    ))}
                </AccordionContent>
              </AccordionItem>
              <AccordionItem value="item-2">
                <AccordionTrigger>Общежития</AccordionTrigger>
                <AccordionContent>
                  {mokPanoramas
                    .filter((item) => item.type === "Общежитие")
                    .map((panorama) => (
                      <PanoramaView panorama={panorama} />
                    ))}
                </AccordionContent>
              </AccordionItem>
              <AccordionItem value="item-3">
                <AccordionTrigger>Столовые</AccordionTrigger>
                <AccordionContent>
                  {mokPanoramas
                    .filter((item) => item.type === "Столовая")
                    .map((panorama) => (
                      <PanoramaView panorama={panorama} />
                    ))}
                </AccordionContent>
              </AccordionItem>
              <AccordionItem value="item-4">
                <AccordionTrigger>Прочее</AccordionTrigger>
                <AccordionContent>
                  {mokPanoramas
                    .filter((item) => item.type === "Прочее")
                    .map((panorama) => (
                      <PanoramaView panorama={panorama} />
                    ))}
                </AccordionContent>
              </AccordionItem>
            </Accordion>
            {/* {mokPanoramas.map((panorama) => (
              <PanoramaView panorama={panorama} />
            ))} */}
          </TabsContent>
        </Tabs>
      </div>
    </main>
  );
}
