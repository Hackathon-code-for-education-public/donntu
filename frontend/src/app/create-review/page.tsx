"use client";

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { z } from "zod";
import { useUniversity } from "@/lib/use-university";
import { Button } from "@/components/ui/button";
import { API } from "@/lib/api";
import { useState } from "react";

const schema = z.object({
  sentiment: z.string({
    required_error: "Оценка обязательна",
  }),
  text: z
    .string({
      required_error: "Текст обязателен",
    })
    .min(1, "Текст обязателен"),
});

function CreateReview({
  searchParams,
}: {
  searchParams: { [key: string]: string };
}) {
  const universityId = searchParams.universityId;

  const { data } = useUniversity(universityId);

  const form = useForm({
    resolver: zodResolver(schema),
  });

  const [submitSuccess, setSubmitSuccess] = useState(false);

  // here
  const onSubmit = async (data) => {
    console.log(data);

    await API.createReview(universityId, data.sentiment, data.text);
    setSubmitSuccess(true);
  };

  return (
    <main className="min-h-screen">
      <div className="bg-white p-6 shadow-lg rounded-lg max-w-4xl mx-auto">
        <h1 className="text-xl font-bold">Оставьте отзыв о {data?.name}</h1>
        {submitSuccess ? (
          <p className="text-green-500">Отзыв успешно отправлен!</p>
        ) : (
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-4">
              <FormField
                control={form.control}
                name="sentiment"
                render={({ field, fieldState }) => (
                  <FormItem>
                    <FormLabel>Оценка</FormLabel>
                    <Select
                      onValueChange={field.onChange}
                      defaultValue={field.value}
                    >
                      <FormControl>
                        <SelectTrigger>
                          <SelectValue placeholder="Выберите оценку" />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        <SelectItem value="positive">
                          👍 Положительная
                        </SelectItem>
                        <SelectItem value="neutral">😐 Нейтральная</SelectItem>
                        <SelectItem value="negative">
                          👎 Отрицательная
                        </SelectItem>
                      </SelectContent>
                    </Select>
                    {fieldState.error && (
                      <FormMessage>{fieldState.error.message}</FormMessage>
                    )}
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="text"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Текст отзыва</FormLabel>
                    <FormControl>
                      <Textarea placeholder="Введите текст" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <Button type="submit">Отправить</Button>
            </form>
          </Form>
        )}
      </div>
    </main>
  );
}

export default CreateReview;
