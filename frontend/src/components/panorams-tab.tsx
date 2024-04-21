import { usePanorams } from "@/lib/use-panorams";
import { PanoramaForm } from "./panorama-form";
import {
  Accordion,
  AccordionItem,
  AccordionTrigger,
  AccordionContent,
} from "./ui/accordion";

import dynamic from "next/dynamic";
import { useUser } from "@/lib/use-user";
import RoleProtected from "./RoleProtected";

const PanoramaView = dynamic(
  () => import("@/components/panorama").then((module) => module) as any,
  { ssr: false }
) as any;

const categories: string[] = ["Общежития", "Корпуса", "Столовые", "Прочее"];

// @ts-ignore
export function PanoramsTab({ universityId }) {
  const { data, isLoading } = usePanorams(universityId);
  const { loggedOut } = useUser();

  if (isLoading) {
    return <div>Загрузка...</div>
  }

  return (
    <>
      <Accordion type="single" collapsible className="w-full">
        {categories.map((category, categoryIndex) => (
          <AccordionItem key={categoryIndex} value={category}>
            <AccordionTrigger>{category}</AccordionTrigger>
            <AccordionContent>
              {data
                ?.filter((item) => item.type === category)
                ?.map((panorama, panoramaIndex) => (
                  <PanoramaView key={panoramaIndex} panorama={panorama} />
                ))}
            </AccordionContent>
          </AccordionItem>
        ))}
      </Accordion>
      {
        !loggedOut && <RoleProtected requiredRoles={"UNIVERSITY"}><PanoramaForm universityId={universityId} /></RoleProtected>
      }
    </>
  );
}
