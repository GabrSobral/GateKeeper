"use client";

import * as React from "react";
import { useTheme } from "next-themes";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import { Label } from "@/components/ui/label";

export function ThemeSwitcher() {
  const { theme, setTheme } = useTheme();

  return (
    <RadioGroup
      value={theme}
      onValueChange={(val) => setTheme(val)}
      className="flex gap-4 mt-4"
    >
      <div className="relative flex flex-col gap-4">
        <RadioGroupItem
          id="theme-light"
          value="light"
          className="peer sr-only"
        />

        <Label
          htmlFor="theme-light"
          className="cursor-pointer rounded-lg border shadow w-[330px] h-[190px] p-2 flex gap-3 transition-all peer-data-[state=checked]:border-primary peer-data-[state=checked]:ring-6 peer-data-[state=checked]:ring-primary/30 bg-white"
        >
          <div className="h-full bg-gray-200 w-[60px] rounded-md" />
          <div className="w-full flex flex-col gap-2 ">
            <div className="h-3 w-[90px] bg-gray-200 rounded" />
            <div className="h-3 w-[190px] bg-gray-200 rounded" />
            <div className="h-8 w-full bg-gray-200 rounded" />
            <div className="flex gap-2">
              <div className="h-[30px] w-[70px] bg-gray-200 rounded" />
              <div className="h-[30px] w-[70px] bg-gray-200 rounded" />
              <div className="h-[30px] w-[70px] bg-gray-200 rounded" />
            </div>
            <div className="h-3 w-[150px] bg-gray-200 rounded" />
            <div className="h-3 w-[226px] bg-gray-200 rounded" />
          </div>
        </Label>

        <span className="text-lg font-semibold">Light</span>
      </div>

      <div className="relative flex flex-col gap-4">
        <RadioGroupItem id="theme-dark" value="dark" className="peer sr-only" />
        <Label
          htmlFor="theme-dark"
          className="cursor-pointer rounded-lg border border-gray-500 shadow w-[330px] h-[190px] p-3 flex gap-2 bg-gray-900 transition-all peer-data-[state=checked]:border-primary peer-data-[state=checked]:ring-6 peer-data-[state=checked]:ring-primary/30"
        >
          <div className="h-full bg-gray-700 w-[60px] rounded-md" />
          <div className="w-full flex flex-col gap-2">
            <div className="h-3 w-[90px] bg-gray-700 rounded" />
            <div className="h-3 w-[190px] bg-gray-700 rounded" />
            <div className="h-8 w-full bg-gray-700 rounded" />
            <div className="flex gap-2">
              <div className="h-[30px] w-[70px] bg-gray-700 rounded" />
              <div className="h-[30px] w-[70px] bg-gray-700 rounded" />
              <div className="h-[30px] w-[70px] bg-gray-700 rounded" />
            </div>
            <div className="h-3 w-[150px] bg-gray-700 rounded" />
            <div className="h-3 w-[226px] bg-gray-700 rounded" />
          </div>
        </Label>

        <span className="text-lg font-semibold">Dark</span>
      </div>

      <div className="relative flex flex-col gap-4">
        <RadioGroupItem
          id="theme-system"
          value="system"
          className="peer sr-only"
        />
        <Label
          htmlFor="theme-system"
          className="relative cursor-pointer rounded-lg  shadow-lg w-[330px] h-[190px] p-3 flex gap-2 bg-black transition-all peer-data-[state=checked]:border-primary peer-data-[state=checked]:ring-6 peer-data-[state=checked]:ring-primary/30"
        >
          <div className="h-full bg-gray-800 w-[60px] rounded-md" />
          <div className="w-full flex flex-col gap-2">
            <div className="h-3 w-[90px] bg-gray-800 rounded" />
            <div className="h-3 w-[190px] bg-gray-800 rounded" />
            <div className="h-8 w-full bg-gray-800 rounded" />
            <div className="flex gap-2">
              <div className="h-[30px] w-[70px] bg-gray-800 rounded" />
              <div className="h-[30px] w-[70px] bg-gray-800 rounded" />
              <div className="h-[30px] w-[70px] bg-gray-800 rounded" />
            </div>
            <div className="h-3 w-[150px] bg-gray-800 rounded" />
            <div className="h-3 w-[226px] bg-gray-800 rounded" />
          </div>

          <div className="absolute w-[176px] h-[190px] top-0 left-0 rounded-l-lg backdrop-invert w-50%"></div>
        </Label>

        <span className="text-lg font-semibold">System</span>
      </div>
    </RadioGroup>
  );
}
