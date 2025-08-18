import { useEffect, useState, useTransition } from "react";

import { Label } from "@/components/ui/label";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";

import { getCookieValue } from "@/lib/utils";

const COLORS = [
  { id: "default", color: "hsl(0 0% 9%)" },
  { id: "blueberry", color: "hsl(220 70% 65%)" },
  { id: "mint", color: "hsl(160 55% 60%)" },
  { id: "peach", color: "hsl(20 75% 70%)" },
  { id: "lavender", color: "hsl(265 50% 70%)" },
  { id: "sunny", color: "hsl(45 90% 70%)" },
  { id: "pink", color: "hsl(330 70% 75%)" },
];

export function ColorSwitcher() {
  const [value, setValue] = useState("default");
  const [, startTransition] = useTransition();

  useEffect(() => {
    const color = getCookieValue("color-scheme") || "default";

    console.log(color);
    setValue(color);

    document.documentElement.setAttribute("data-color", color);
  }, []);

  const setColor = async (color: string) => {
    startTransition(async () => {
      await fetch("/api/set-color", {
        method: "POST",
        body: JSON.stringify({ color }),
      });

      setValue(color);

      document.documentElement.setAttribute("data-color", color);
    });
  };

  return (
    <RadioGroup value={value} onValueChange={setColor} className="flex gap-4">
      {COLORS.map(({ id, color }) => (
        <div key={id} className="relative flex flex-col">
          <RadioGroupItem id={id} value={id} className="peer sr-only" />
          <Label
            htmlFor={id}
            className="relative cursor-pointer rounded-lg shadow-lg p-1 flex items-center justify-center transition-all border peer-data-[state=checked]:border-primary peer-data-[state=checked]:ring-4 peer-data-[state=checked]:ring-primary/30"
          >
            <div
              className="h-[50px] w-[50px] rounded-md"
              style={{ backgroundColor: color }}
            />
            <span className="sr-only">{id}</span>
          </Label>
        </div>
      ))}
    </RadioGroup>
  );
}
