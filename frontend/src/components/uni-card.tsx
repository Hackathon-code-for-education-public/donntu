import Image from "next/image";
import { Separator } from "@/components/ui/separator";
import Link from "next/link";

export default function UniCard() {
  const reviews_count = 52;
  const university = "Донецкий национальный технический университет";
  const id = 1;
  return (
    <div className="">
      <div className="grid grid-cols-3 grid-rows-1 gap-4">
        <div className="row-span-2 content-center justify-self-center">
          <Image
            src="/gerb_donntu.jpg"
            alt="avatar"
            width={100}
            height={100}
            className="rounded-full"
          />
        </div>
        <div className="col-span-2">
          <Link href={`/university/${id}`} className="text-xl font-bold mb-2">
            {university}
          </Link>
          <p className="text-sm text-slate-600">
            ДонНТУ – признанное в мире высшее учебное заведение, активно
            осуществляющее международное научно-техническое сотрудничество с
            более чем 80 известными университетами из 25-ти стран мира.
          </p>
          <Separator className="mt-2 mb-2" />
          <Link href="/">{reviews_count} Отзывов</Link>
        </div>
        <div className="col-span-2 col-start-2 row-start-2"></div>
      </div>
    </div>
  );
}
