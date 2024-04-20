import Panorama from "@/components/panorama";
import Image from "next/image";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center">
      <div className="flex w-full h-80 items-center justify-center bg-gradient-to-r from-purple-500 to-blue-500 mb-10">
        <p className="w-2/3 font-bold text-white text-center text-4xl">
          Добро пожаловать Сервис по онлайн экскурсиям в университеты
        </p>
      </div>
      <Panorama />
    </main>
  );
}
