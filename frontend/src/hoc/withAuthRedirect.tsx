"use client";
import { useEffect } from "react";
import { redirect } from "next/navigation";
import { useUser } from "@/lib/use-user";

export default function withAuthRedirect(Component: any) {
  return function WithAuthRedirect(props: any) {
    const { user, loading } = useUser();

    useEffect(() => {
      if (!loading && user) {
        return redirect("/account");
      }
    }, [loading, user]);

    if (loading) {
      return <main className="flex min-h-screen flex-col items-center"></main>;
    }

    return <Component {...props} />;
  };
}
