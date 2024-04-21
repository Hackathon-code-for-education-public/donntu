import { Input } from "@/components/ui/input";

export function SearchInput({ onChange }: { onChange: (value: string) => void }) {
  return (
    <Input
      className="w-full rounded-md py-2 px-4 border border-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-500"
      placeholder="Название университета"
      type="text"
      onChange={(event) => onChange(event.target.value)}
    />
  );
}
