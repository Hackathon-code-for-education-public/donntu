import UniCard from "@/components/uni-card";

export default function Page() {
  return (
    <main className="flex min-h-screen flex-col items-center">
      <div className="w-2/3 mt-20">
        <p className="text-4xl font-bold mb-10">Наши рекомендации</p>
        <div className="grid grid-cols-2 grid-rows-2 gap-4">
          <div className="row-span-2">
            <UniCard />
          </div>
          <div>
            <UniCard />
          </div>
          <div className="col-start-2">
            <UniCard />
          </div>
        </div>
      </div>
    </main>
  );
}
