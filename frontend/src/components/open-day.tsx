import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

import { MapPinIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { OpenDayData } from "@/lib/use-open-days";


interface IProps {
    day: OpenDayData;
}

export function OpenDay({ day }: IProps) {
  return (
    <Card key={day.link}>
      <CardHeader>
        <CardTitle>{day.universityName}</CardTitle>
        <CardDescription>{day.description}</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="flex gap-2">
          <MapPinIcon />
          <p>{day.place}</p>
        </div>
      </CardContent>
      <CardFooter>
        <Button variant="outline" className="w-full" asChild>
          <Link href={day.link}>Подробнее</Link>
        </Button>
      </CardFooter>
    </Card>
  );
}
