"use client";
import { useEffect } from "react";
import { redirect } from "next/navigation";
import { useUser } from "@/lib/use-user";


export default function isAuth(Component: any) {
  return function IsAuth(props: any) {
    const { user, loading } = useUser();


    useEffect(() => {
      if (!loading && !user) {
        return redirect("/");
      }
    }, [loading, user]);


    if (loading) {
        return null;
    }

    return <Component {...props} />;
  };
}