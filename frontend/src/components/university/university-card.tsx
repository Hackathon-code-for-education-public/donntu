import { UniversityIcon } from "lucide-react";
import Link from "next/link";

interface UniversityCardProps {
  id: string;
  name: string;
  longName: string;
  logoUrl?: string;
  onClick?: () => void; // Optional onClick handler
}

export function UniversityCard({
  id,
  name,
  longName,
  logoUrl,
  onClick,
}: UniversityCardProps) {
  const cardContent = (
    <>
      {logoUrl ? (
        <img
          alt="University Logo"
          className="mr-4"
          height={60}
          src={logoUrl}
          style={{ aspectRatio: "60/60", objectFit: "contain" }}
          width={60}
        />
      ) : (
        <div
          className="mr-4 flex items-center justify-center"
          style={{ width: 60, height: 60 }}
        >
          <UniversityIcon size="60px" />
        </div>
      )}
      <div>
        <h3 className="text-lg font-bold">{name}</h3>
        <p className="text-gray-500 text-sm">{longName}</p>
      </div>
    </>
  );

  return onClick ? (
    <div
      onClick={onClick}
      className="bg-white rounded-md shadow-md p-4 flex items-start cursor-pointer"
    >
      {cardContent}
    </div>
  ) : (
    <Link href={`/university/${id}`} passHref legacyBehavior>
      <a className="bg-white rounded-md shadow-md p-4 flex items-start">
        {cardContent}
      </a>
    </Link>
  );
}
