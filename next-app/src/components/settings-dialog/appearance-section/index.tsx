import { Label } from "@/components/ui/label";
import { Separator } from "@/components/ui/separator";
import { Breadcrumbs } from "@/components/bread-crumbs";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";

import { ThemeSwitcher } from "./theme-switcher";

export function AppearanceSection() {
  return (
    <main className="flex flex-col p-4 md:max-h-[700px] overflow-auto overflow-x-hidden">
      <Breadcrumbs
        items={[{ name: "Settings" }, { name: "Appearance" }]}
        disableSideBar
      />
      <h2 className="text-3xl font-bold tracking-tight mt-4">Appearance</h2>

      <span className="mt-3 text-sm tracking-tight text-gray-600 dark:text-gray-300">
        Here you can manage the appearance settings of the application,
        including theme selection, font size adjustment, and layout options.
      </span>

      <Separator className="my-4" />

      <section className="flex flex-col gap-2">
        <h3 className="text-2xl font-semibold tracking-tight">Theming</h3>
        <span className="text-sm tracking-tight text-gray-600 dark:text-gray-300">
          Here you can manage the theme settings of the application.
        </span>

        <ThemeSwitcher />
      </section>

      <Separator className="my-4" />

      <section className="flex flex-col gap-2">
        <h3 className="text-2xl font-semibold tracking-tight">Primary Color</h3>
        <span className="text-sm tracking-tight text-gray-600 dark:text-gray-300">
          Here you can manage the theme settings of the application.
        </span>

        <RadioGroup className="flex mt-3">
          <div className="relative flex flex-col gap-4">
            <RadioGroupItem
              id="neutral"
              value="neutral"
              className="peer sr-only"
            />
            <Label
              htmlFor="neutral"
              className="relative cursor-pointer rounded-lg shadow-lg p-1 flex gap-2 transition-all peer-data-[state=checked]:border-primary peer-data-[state=checked]:ring-6 peer-data-[state=checked]:ring-primary/30"
            >
              <div className="h-[50px] bg-gray-600 w-[50px] rounded-md" />
            </Label>
          </div>

          <div className="relative flex flex-col gap-4">
            <RadioGroupItem id="blue" value="blue" className="peer sr-only" />
            <Label
              htmlFor="blue"
              className="relative cursor-pointer rounded-lg shadow-lg p-1 flex gap-2 transition-all peer-data-[state=checked]:border-primary peer-data-[state=checked]:ring-6 peer-data-[state=checked]:ring-primary/30"
            >
              <div className="h-[50px] bg-blue-600 w-[50px] rounded-md" />
            </Label>
          </div>
        </RadioGroup>
      </section>
    </main>
  );
}
