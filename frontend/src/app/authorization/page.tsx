"use client";

import { useState } from "react";
import { CardHeader, CardContent, Card } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { ComboBox } from "@/components/combobox";
import { AuthAPI } from "@/lib/auth";

const items = [
  {
    value: "applicant",
    label: "Абитуриент",
  },
  {
    value: "student",
    label: "Студент",
  },
  {
    value: "university",
    label: "Представитель ВУЗа",
  },
];

function LoginForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleLogin = async (event) => {
    event.preventDefault();
    setLoading(true);
    setError("");

    try {
      const data = await AuthAPI.login(email, password);
      console.log("Logged in successfully:", data);
      setLoading(false);
    } catch (error) {
      if (error instanceof Error) {
        console.error("Login failed:", error.message);
        setError(error.message);
      } else {
        console.error("Login failed with unknown error:", error);
        setError("An unknown error occurred");
      }
      setLoading(false);
    }
  };

  return (
    <form className="space-y-4" onSubmit={handleLogin}>
      <div className="space-y-2">
        <Label htmlFor="email">Почта</Label>
        <Input
          id="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="m@example.com"
          required
          type="email"
        />
      </div>
      <div className="space-y-2">
        <Label htmlFor="password">Пароль</Label>
        <Input
          id="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          type="password"
        />
      </div>
      <Button className="w-full" type="submit" disabled={loading}>
        Войти
      </Button>
      {error && <div className="text-red-500">{error}</div>}
    </form>
  );
}

function RegisterForm() {
  const [lastName, setLastName] = useState("");
  const [firstName, setFirstName] = useState("");
  const [middleName, setMiddleName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleRegister = async (event) => {
    event.preventDefault();
    setLoading(true);
    setError("");

    try {
      const data = await AuthAPI.register(
        email,
        password,
        lastName,
        firstName,
        middleName
      );
      console.log("Registration successful:", data);
      setLoading(false);
    } catch (error) {
      if (error instanceof Error) {
        console.error("Registration failed:", error.message);
        setError(error.message);
      } else {
        console.error("Registration failed with unknown error:", error);
        setError("An unknown error occurred");
      }
      setLoading(false);
    }
  };

  return (
    <form className="space-y-4" onSubmit={handleRegister}>
      <div className="space-y-2">
        <Label htmlFor="lastName">Фамилия</Label>
        <Input
          id="lastName"
          value={lastName}
          onChange={(e) => setLastName(e.target.value)}
          placeholder="Иванов"
          required
          type="text"
        />
      </div>
      <div className="space-y-2">
        <Label htmlFor="firstName">Имя</Label>
        <Input
          id="firstName"
          value={firstName}
          onChange={(e) => setFirstName(e.target.value)}
          placeholder="Иван"
          required
          type="text"
        />
      </div>
      <div className="space-y-2">
        <Label htmlFor="middleName">Отчество</Label>
        <Input
          id="middleName"
          value={middleName}
          onChange={(e) => setMiddleName(e.target.value)}
          placeholder="Иванович"
          type="text"
        />
      </div>
      <div className="space-y-2">
        <Label htmlFor="email">Почта</Label>
        <Input
          id="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="m@example.com"
          required
          type="email"
        />
      </div>
      <div className="space-y-2">
        <Label htmlFor="password">Пароль</Label>
        <Input
          id="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          type="password"
        />
      </div>
      <Button className="w-full" type="submit" disabled={loading}>
        Регистрация
      </Button>
      {error && <div className="text-red-500">{error}</div>}
    </form>
  );
}
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
              <RegisterForm />
            </CardContent>
          </TabsContent>
          <TabsContent value="login">
            <CardContent>
              <LoginForm />
            </CardContent>
          </TabsContent>
        </Tabs>
      </Card>
    </main>
  );
}
