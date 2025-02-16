import { Fragment } from "react";

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { SidebarTrigger } from "@/components/ui/sidebar";

type Props = {
  items: {
    name: string;
    path: string;
  }[];
};

export function Breadcrumbs({ items }: Props) {
  return (
    <div className="flex gap-4 items-center">
      <SidebarTrigger />

      <Breadcrumb>
        <BreadcrumbList>
          {items.map((item, index) => (
            <Fragment key={index}>
              <BreadcrumbItem>
                <BreadcrumbLink href={item.path}>{item.name}</BreadcrumbLink>
              </BreadcrumbItem>

              {index !== items.length - 1 && <BreadcrumbSeparator />}
            </Fragment>
          ))}
        </BreadcrumbList>
      </Breadcrumb>
    </div>
  );
}
