import Image from "next/image";
import Link from "next/link";
import dynamic from "next/dynamic";

const Panorama = dynamic(
  () => import('@/components/panorama').then(module => module) as any,
  { ssr: false },
) as any;

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center">
      <div className="flex w-full h-80 items-center justify-center bg-gradient-to-r from-purple-500 to-blue-500 mb-10">
        <p className="w-2/3 font-bold text-white text-center text-4xl">
          Добро пожаловать Сервис по онлайн экскурсиям в университеты
        </p>
      </div>
      <div className="w-2/3 mt-20 mb-20">
        <p className="text-4xl font-bold mb-20">Что доступно на сайте?</p>
        <div className="flex flex-col items-center mb-20">
          <Image
            src="/lib.jpg"
            width="1300"
            height="500"
            alt="event-picture"
            className="bg-slate-200 rounded-xl"
          />
          <div>
            <p className="text-xl font-bold mb-4 mt-4">
              Онлайн-экскурсии в виде интерактивных панорам
            </p>
            <p className="leading-relaxed text-slate-600 text-justify">
              На странице ВУЗа можно попасть на <b>раздел с панорамами</b> для
              детального знакомства с заведением. Есть возможность перемещаться
              по нескольким панорамам и переходить по меткам Пользователь может
              загружать несколько панорам для нескольких помещений ВУЗа.
            </p>
          </div>
        </div>
        <div className="grid grid-cols-2 gap-4 justify-content-center mb-20">
          <div>
            <p className="text-xl font-bold mb-4">Система отзывов</p>
            <p className="leading-relaxed text-slate-600 text-justify">
              Студент, или бывший студент ВУЗа может оставить{" "}
              <b>позитивный, негативный или нейтральный отзыв</b> на странице
              ВУЗа в соответствующем разделе. Для абитуриентов будет видна дата,
              статус пользователя, оставившего отзыв (студент, отчисленный,
              выпускник)
            </p>
          </div>
          <div className=" justify-self-center">
            <Image
              src="/reviews.png"
              width="500"
              height="500"
              alt="event-picture"
              className="bg-slate-200 rounded-xl "
            />
          </div>
        </div>
        <div className="grid grid-cols-2 gap-4 justify-content-center mb-20">
          <Image
            src="/chat.png"
            width="500"
            height="500"
            alt="event-picture"
            className="bg-slate-200 rounded-xl"
          />
          <div>
            <p className="text-xl font-bold mb-4">
              Чат студентов с абитуриентами
            </p>
            <p className="leading-relaxed text-slate-600 text-justify">
              Абитуриенты могут спросить студентов о ВУЗе на прямую через{" "}
              <Link href={"/chat"}> 
                <b>систему чатов</b>
              </Link>
              .Число чатов не ограничено. 
            </p>
          </div>
        </div>
      </div>
    </main>
  );
}

