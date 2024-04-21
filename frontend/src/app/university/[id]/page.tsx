"use client";

import { Panorama } from "@/api/panorama";
import { PanoramsTab } from "@/components/panorams-tab";
import { Button } from "@/components/ui/button";

import { Skeleton } from "@/components/ui/skeleton";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { UniversityOpenDays } from "@/components/university-open-days";
import { UniversityReviews } from "@/components/university-reviews";
import RoleProtected from "@/components/RoleProtected";
import { useUniversity } from "@/lib/use-university";
import Link from "next/link";

interface Params {
  id: string;
}

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
            <div className="grid grid-cols-3 gap-4 text-center py-4 bg-gray-200 rounded-lg">
              {isLoading ? (
                <div className="flex justify-center items-center flex-col">
                  <Skeleton className="h-6 w-1/4" />
                  <Skeleton className="h-6 w-1/2 mt-2" />
                </div>
              ) : (
                <div>
                  <div className="text-3xl font-bold">
                    {data?.rating.toFixed(2)}
                  </div>
                  <div className="text-sm">рейтинг</div>
                </div>
              )}
              {isLoading ? (
                <div className="flex justify-center items-center flex-col">
                  <Skeleton className="h-6 w-1/4" />
                  <Skeleton className="h-6 w-1/2 mt-2" />
                </div>
              ) : (
                <div>
                  <div className="text-3xl font-bold">{data?.studyFields}</div>
                  <div className="text-sm">направлений подготовки</div>
                </div>
              )}
              {isLoading ? (
                <div className="flex justify-center items-center flex-col">
                  <Skeleton className="h-6 w-1/4" />
                  <Skeleton className="h-6 w-1/2 mt-2" />
                </div>
              ) : (
                <div>
                  <div className="text-3xl font-bold">{data?.budgetPlaces}</div>
                  <div className="text-sm">бюджетных мест</div>
                </div>
              )}
            </div>
            <div className="mt-4 text-sm">
              Данные представленны самим ВУЗом.
            </div>
          </TabsContent>
          <TabsContent value="reviews" className="m-5">
            <div className="flex justify-between">
              <h2 className="text-lg">Отзывы</h2>
              <RoleProtected requiredRoles={"STUDENT"}>
                <Link
                  href={`/create-review?universityId=${params.id}`}
                  passHref
                  legacyBehavior
                >
                  <Button>Добавить отзыв</Button>
                </Link>
              </RoleProtected>
            </div>
            <div className="pt-5">
              <UniversityReviews universityId={params.id} />
            </div>
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
            <PanoramsTab universityId={params.id} />
          </TabsContent>
        </Tabs>
      </div>
    </main>
  );
}
