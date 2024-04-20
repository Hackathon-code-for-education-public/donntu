import { Separator } from "@/components/ui/separator";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { Table, TableBody, TableCell, TableRow } from "@/components/ui/table";

export default function Page() {
  const user = {
    name: "John",
    surname: "Doe",
    middlename: "Doedovich",
    role: "univercity",
    email: "haxrelt@gmail.com",
  };

  return (
    <main className="flex min-h-screen flex-col items-center">
      <div className="w-2/3 flex flex-col items-center">
        <Avatar>
          <AvatarImage src="https://github.com/shadcn.png" />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
        <div className="flex gap-2 font-bold text-4xl">
          <p>{user.surname}</p>
          <p>{user.name}</p>
          <p>{user.middlename}</p>
        </div>
        <p className="font-bold text-xl self-start">Личная информация:</p>
        <div className="justify-stretch">
          <Table>
            <TableBody>
              <TableRow>
                <TableCell className="font-medium">Фамилия</TableCell>
                <TableCell className="text-right">{user.surname}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell className="font-medium">Имя</TableCell>
                <TableCell className="text-right">{user.name}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell className="font-medium">Отчество</TableCell>
                <TableCell className="text-right">{user.middlename}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell className="font-medium">E-mail</TableCell>
                <TableCell className="text-right">{user.email}</TableCell>
              </TableRow>
              <TableRow>
                <TableCell className="font-medium">Роль</TableCell>
                <TableCell className="text-right">{user.role}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </div>

        <p className="font-bold text-xl self-start">Ваши отзывы: </p>
      </div>
    </main>
  );
}
