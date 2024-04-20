"use client";
import { Button } from "@/components/ui/button";
import { AuthAPI } from "@/lib/auth";
import { useUser } from "@/lib/use-user";

export function LogoutButton() {
    const { mutate } = useUser()

    const onClick = async () => {
        await AuthAPI.signOut();

        mutate(undefined, false);
    }

    return <Button onClick={onClick}>Выйти</Button>
}