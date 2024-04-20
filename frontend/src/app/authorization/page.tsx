import {
    CardTitle,
    CardDescription,
    CardHeader,
    CardContent,
    Card,
  } from "@/components/ui/card";
  import { Label } from "@/components/ui/label";
  import { Input } from "@/components/ui/input";
  import Link from "next/link";
  import { Button } from "@/components/ui/button";
  import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
  import { ComboBox } from "@/components/combobox";
  
  const items = [
    {
      value: "applicant",
      label: "Абитуриент",
    },
    {
      value: "university",
      label: "Представитель ВУЗа",
    },
  ];
  
  export default function Page() {
    return (
      <main className="flex min-h-screen items-center justify-center bg-gray-100 dark:bg-gray-950 p-20">
        <Card className="mx-auto max-w-sm min-h-96 min-w-96">
          <Tabs className="w-full" defaultValue="login">
            <CardHeader className="space-y-1">
              <TabsList className="grid w-full grid-cols-2">
                <TabsTrigger value="login">Авторизация</TabsTrigger>
                <TabsTrigger value="register">Регистрация</TabsTrigger>
              </TabsList>
            </CardHeader>
            <TabsContent value="register">
              <CardContent>
                <form className="space-y-4">
                  <div className="space-y-2">
                    <ComboBox items={items} defaultValue="applicant" />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="surname">Фамилия</Label>
                    <Input id="surname" placeholder="Иванов" required type="text" />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="firstName">Имя</Label>
                    <Input id="firstName" placeholder="Иван" required type="text" />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="patronymic">Отчество</Label>
                    <Input id="patronymic" placeholder="Иванович" type="text" />
                  </div>
                  <div className="space-y-2">
                    <Label htmlFor="email">Почта</Label>
                    <Input
                      id="email"
                      placeholder="m@example.com"
                      required
                      type="email"
                    />
                  </div>
                  <div className="space-y-2">
                    <div className="flex items-center justify-between">
                      <Label htmlFor="password">Пароль</Label>
                    </div>
                    <Input id="password" required type="password" />
                  </div>
                  <Button className="w-full" type="submit">
                    Регистрация
                  </Button>
                </form>
              </CardContent>
            </TabsContent>
            <TabsContent value="login">
              <CardContent>
                <form className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="email">Почта</Label>
                    <Input
                      id="email"
                      placeholder="m@example.com"
                      required
                      type="email"
                    />
                  </div>
                  <div className="space-y-2">
                    <div className="flex items-center justify-between">
                      <Label htmlFor="password">Пароль</Label>
                    </div>
                    <Input id="password" required type="password" />
                  </div>
                  <Button className="w-full" type="submit">
                    Войти
                  </Button>
                </form>
              </CardContent>
            </TabsContent>
          </Tabs>
        </Card>
      </main>
    );
  }
  