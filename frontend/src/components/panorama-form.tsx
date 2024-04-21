"use client";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
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
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form";
import { API } from "@/lib/api";
import { usePanorams } from "@/lib/use-panorams";

interface PanoramaFormProps {
  universityId: string;
}

const ACCEPTED_IMAGE_TYPES = [
  "image/jpeg",
  "image/jpg",
  "image/png",
  "image/webp",
];

const panoramaSchema = z.object({
  name: z
    .string({
      required_error: "Название обязательно",
    })
    .min(1, "Название обязательно"),
  address: z
    .string({
      required_error: "Адрес обязателен",
    })
    .min(1, "Адрес обязателен"),
  category: z.string({
    required_error: "Категория обязательна",
  }),
  firstLocation: z
    .custom<File>()
    .refine((file) => file !== undefined, "Файл обязателен")
    .refine((file) => {
      return file !== undefined && ACCEPTED_IMAGE_TYPES.includes(file.type);
    }, "Только .jpg, .jpeg, .png разрешены"),
  // secondLocation: z.any(), // Uncomment if using the second location.
});

export function PanoramaForm({ universityId }: PanoramaFormProps) {
  const { mutate } = usePanorams(universityId);

  const form = useForm({
    resolver: zodResolver(panoramaSchema),
  });

  // @ts-ignore
  const onSubmit = async (data) => {
    console.log(data);

    // uploadPanorama(data)
    await API.uploadPanorama(
      data.firstLocation,
      data.firstLocation,
      universityId,
      data.name,
      data.address,
      data.category
    );

    mutate(undefined);
  };

  return (
    <Card className="mt-20">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-4">
          <CardHeader>
            <CardTitle>Добавьте панораму</CardTitle>
            <CardDescription>
              Заполните название здания, его адрес, выберите тип и загрузите
              фото локаций
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-2">
                <FormField
                  control={form.control}
                  name="name"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Название</FormLabel>
                      <FormControl>
                        <Input placeholder="Введите название" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <div className="space-y-2">
                <FormField
                  control={form.control}
                  name="address"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Адрес</FormLabel>
                      <FormControl>
                        <Input placeholder="Введите адрес" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
            </div>
            <div className="space-y-2">
              <FormField
                control={form.control}
                name="category"
                render={({ field, fieldState }) => (
                  <FormItem>
                    <FormLabel>Категория</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger>
                          <SelectValue placeholder="Выберите категорию" />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        <SelectItem value="Корпуса">Корпус</SelectItem>
                        <SelectItem value="Общежития">Общежитие</SelectItem>
                        <SelectItem value="Столовые">Столовая</SelectItem>
                        <SelectItem value="Прочее">Прочее</SelectItem>
                      </SelectContent>
                    </Select>
                    {fieldState.error && (
                      <FormMessage>{fieldState.error.message}</FormMessage>
                    )}
                  </FormItem>
                )}
              />
            </div>
            <div className="grid grid-cols-2 gap-4">
              <FormField
                control={form.control}
                name="firstLocation"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Выберите фото локации</FormLabel>
                    <FormControl>
                      <Input
                        type="file"
                        accept="image/png, image/jpg, image/jpeg, image/webp"
                        onChange={(e) => {
                          field.onChange(e.target.files && e.target.files[0]);
                        }}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
          </CardContent>
          <CardFooter>
            <Button type="submit">Загрузить</Button>
          </CardFooter>
        </form>
      </Form>
    </Card>
  );
}
