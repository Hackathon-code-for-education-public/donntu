import { useUser } from "@/lib/use-user";

export default function RoleProtected({
  children,
  requiredRoles,
}: {
  children: React.ReactNode;
  requiredRoles?: string | string[];
}) {
  const { user, loading } = useUser();

  if (loading) {
    return <></>;
  }

  const rolesArray =
    requiredRoles !== null
      ? Array.isArray(requiredRoles)
        ? requiredRoles
        : [requiredRoles]
      : [];

  if (
    !user?.role ||
    (rolesArray.length !== 0 && !rolesArray.includes(user.role))
  ) {
    return <></>;
  }

  // Если пользователь авторизован и имеет подходящую роль, отображаем дочерние элементы
  return <>{children}</>;
}
