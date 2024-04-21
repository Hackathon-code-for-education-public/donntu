"use client";
import { useEffect } from "react";
import { redirect } from "next/navigation";
import { useUser } from "@/lib/use-user";

export default function isAuth(
  Component: any,
  requiredRoles?: string | string[]
) {
  return function IsAuth(props: any) {
    const { user, loading } = useUser();

    useEffect(() => {
      if (!loading) {
        if (!user) {
          // If there is no user, redirect to the login page
          return redirect("/authorization");
        } else if (requiredRoles) {
          // Convert requiredRoles to an array if it's not already one
          const rolesArray = Array.isArray(requiredRoles)
            ? requiredRoles
            : [requiredRoles];

          // Check if user has any of the required roles
          if (!user.role || !rolesArray.includes(user.role)) {
            // If the user does not have any of the required roles, redirect to an unauthorized page
            return redirect("/unauthorized");
          }
        }
        // If requiredRoles is undefined, do not restrict access by roles
      }
    }, [loading, user, requiredRoles]);

    if (loading) {
      // Show a loading state while the user data is loading
      return (
        <main className="flex min-h-screen flex-col items-center justify-center">
          <p>Loading...</p>
        </main>
      );
    }

    // If the user is logged in and has the correct role(s), or if no specific roles are required, render the component
    return <Component {...props} />;
  };
}
