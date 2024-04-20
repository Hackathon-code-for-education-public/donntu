import Image from "next/image";
import { Separator } from "@/components/ui/separator";

export default function UniCard() {
  return (
    <div className="">
      <div className="grid grid-cols-3 grid-rows-1 gap-4">
        <div className="row-span-2">
          <Image
            src="/gerb_donntu.jpg"
            alt="avatar"
            width={100}
            height={100}
            className="rounded-full"
          />
        </div>
        <div className="col-span-2">
          <p className="text-lg font-bold">Название</p>
        </div>
        <div className="col-span-2 col-start-2 row-start-2">
          <p className="text-sm">
            ДонНТУ – признанное в мире высшее учебное заведение, активно
            осуществляющее международное научно-техническое сотрудничество с
            более чем 80 известными университетами из 25-ти стран мира и 32
            иностранными фирмами, являющееся членом 24-х ведущих международных
            образовательных ассоциаций.
          </p>
        </div>
      </div>
    </div>
  );
}
