"use client";

import * as React from "react";
import Link from "next/link";

import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu";
import { Button } from "@/components/ui/button";
import { useUser } from "@/lib/use-user";
import { LogoutButton } from "./logout-button";
import RoleProtected from "./RoleProtected";

export function MainNavigation() {
  const { user, loading, loggedOut } = useUser();

  return (
    <header className="flex justify-around p-2 border border-b-1">
      <NavigationMenu>
        <NavigationMenuList>
          <NavigationMenuItem>
            <Link href="/" legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                Главная
              </NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <Link href="/university" legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                Университеты
              </NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <Link href="/compare" legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                Сравнить университеты
              </NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
          <RoleProtected>
            <NavigationMenuItem>
              <Link href="/account" legacyBehavior passHref>
                <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                  Профиль
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
          </RoleProtected>
          <RoleProtected requiredRoles={["STUDENT", "APPLICANT"]}>
            <NavigationMenuItem>
              <Link href="/chat" legacyBehavior passHref>
                <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                  Чат
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
          </RoleProtected>
          {/*
          <NavigationMenuItem>
            <Link href="/for-applicants" legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                Абитуриентам
              </NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
          */}
          {/*
          <NavigationMenuItem>
            <Link href="/for-universities" legacyBehavior passHref>
              <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                Университетам
              </NavigationMenuLink>
            </Link>
          </NavigationMenuItem>
          */}
        </NavigationMenuList>
      </NavigationMenu>
      {loggedOut && (
        <Button>
          <Link href="/authorization" legacyBehavior passHref>
            Вход
          </Link>
        </Button>
      )}
      {!loggedOut && <LogoutButton />}
    </header>
  );
}
