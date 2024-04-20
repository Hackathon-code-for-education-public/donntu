import * as React from "react";
import { ChevronsUpDown } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Command, CommandGroup, CommandItem } from "@/components/ui/command";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";

type Item = {
    value: string,
    label: string,
}

interface IProps {
    items: Item[],
    value?: string,  // Controlled value
    onChange?: (newValue: string) => void, // Callback when value changes
    defaultValue?: string,
    placeholder?: string,
}

export function ComboBox({
  items,
  value: externalValue,
  onChange,
  defaultValue = "",
  placeholder = "",
}: IProps) {
  const [open, setOpen] = React.useState(false);
  const isControlled = externalValue != null; // Check if the component is controlled
  const [value, setValue] = React.useState(defaultValue);

  const handleSelect = (selectedValue: string) => {
    if (!isControlled) {
      // Only update the internal state if it's not controlled
      setValue(selectedValue);
    }
    onChange?.(selectedValue); // Call the onChange prop if it exists
    setOpen(false);
  };

  const displayLabel = items.find((item) => item.value === (isControlled ? externalValue : value))?.label;

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between"
        >
          {displayLabel || placeholder}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="p-0" align="end">
        <Command>
          <CommandGroup>
            {items.map((item) => (
              <CommandItem
                key={item.value}
                value={item.value}
                onSelect={handleSelect}
              >
                {item.label}
              </CommandItem>
            ))}
          </CommandGroup>
        </Command>
      </PopoverContent>
    </Popover>
  );
}