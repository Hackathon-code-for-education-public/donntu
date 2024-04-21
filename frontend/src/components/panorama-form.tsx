"use client";
import { z } from "zod";
import { Panorama } from "@/api/panorama";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { PanoramaAPI } from "@/lib/panoramas";
import { useState } from "react";

interface PanoramaFormProps {
  universityId: string;
}


export function PanoramaForm({ universityId }: PanoramaFormProps) {
  const [panoramaToPost, setPanoramaToPost] = useState<Panorama>({
    name: "",
    address: "",
    type: "",
    firstLocation: "",
    secondLocation: "",
  });

  return (
    <Card className="mt-20">
      <CardHeader>
        <CardTitle>Добавьте панораму</CardTitle>
        <CardDescription>
          Заполните название здания, его адрес, выберите тип и загрузите фото
          локаций
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form className="grid gap-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="name">Имя</Label>
              <Input
                id="name"
                placeholder="Введите название"
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="address">Address</Label>
              <Input id="address" placeholder="Введите адрес здания" />
            </div>
          </div>
          <div className="space-y-2">
            <Label htmlFor="category">Категория</Label>
            <Select>
              <SelectTrigger>
                <SelectValue placeholder="Выберите тип здания" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="Корпус">Корпус</SelectItem>
                <SelectItem value="Общежитие">Общежитие</SelectItem>
                <SelectItem value="Столовая">Прочее</SelectItem>
              </SelectContent>
            </Select>
          </div>
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <Label htmlFor="loc1">Выберите фото локации</Label>
              <Input id="loc1" type="file" />
            </div>
            <div className="space-y-2">
              <Label htmlFor="loc2">Выберите фото локации</Label>
              <Input id="loc2" type="file" />
            </div>
          </div>
        </form>
      </CardContent>
      <CardFooter>
        <Button
          onClick={() => PanoramaAPI.postPanorama(universityId, panoramaToPost)}
        >
          Загрузить
        </Button>
      </CardFooter>
    </Card>
  );
}
