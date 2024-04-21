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

export function MainNavigation() {
  const { user, loading, loggedOut } = useUser();

  return (
    <header className="flex justify-between p-2">
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
          {user && (
            <NavigationMenuItem>
              <Link href="/account" legacyBehavior passHref>
                <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                  Профиль
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
          )}
          {user && (
            <NavigationMenuItem>
              <Link href="/chat" legacyBehavior passHref>
                <NavigationMenuLink className={navigationMenuTriggerStyle()}>
                  Чат
                </NavigationMenuLink>
              </Link>
            </NavigationMenuItem>
          )}
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
      {!loading && !loggedOut && (
        <Button>
          <Link href="/authorization" legacyBehavior passHref>
            Вход
          </Link>
        </Button>
      )}
      {!loading && loggedOut && <Button>Выйти</Button>}
    </header>
  );
}
