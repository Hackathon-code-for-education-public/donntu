"use client";

import { Separator } from "@/components/ui/separator";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { Table, TableBody, TableCell, TableRow } from "@/components/ui/table";
import isAuth from "@/hoc/isAuth";
import { useUser } from "@/lib/use-user";

function Page() {
  const { user } = useUser();

  return (
    <main className="flex min-h-screen flex-col items-center">
      <div className="w-2/3 flex flex-col items-center">
        {/*
        <Avatar>
          <AvatarImage src="https://github.com/shadcn.png" />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
        */}
        <div className="flex gap-2 font-bold text-4xl">
          <p>{user?.lastName}</p>
          <p>{user?.firstName}</p>
          <p>{user?.middleName}</p>
        </div>
        <p className="font-bold text-xl self-start">Личная информация:</p>
        <Table className="w-full">
          <TableBody>
            <TableRow>
              <TableCell className="font-medium">Фамилия</TableCell>
              <TableCell className="text-right">{user?.lastName}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">Имя</TableCell>
              <TableCell className="text-right">{user?.firstName}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">Отчество</TableCell>
              <TableCell className="text-right">{user?.middleName}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">E-mail</TableCell>
              <TableCell className="text-right">{user?.email}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell className="font-medium">Роль</TableCell>
              <TableCell className="text-right">{user?.role}</TableCell>
            </TableRow>
          </TableBody>
        </Table>
        {/*
          <p className="font-bold text-xl self-start">Ваши отзывы: </p>
        */}
      </div>
    </main>
  );
}

export default isAuth(Page);