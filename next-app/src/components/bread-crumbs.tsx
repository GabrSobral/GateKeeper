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
    path?: string;
  }[];
  disableSideBar?: boolean;
};

export function Breadcrumbs({ items, disableSideBar = false }: Props) {
  return (
    <div className="flex gap-4 items-center">
      {!disableSideBar && <SidebarTrigger />}

      <Breadcrumb>
        <BreadcrumbList>
          {items.map((item, index) => (
            <Fragment key={index}>
              <BreadcrumbItem>
                {item?.path ? (
                  <BreadcrumbLink href={item.path}>{item.name}</BreadcrumbLink>
                ) : (
                  <span>{item.name}</span>
                )}
              </BreadcrumbItem>

              {index !== items.length - 1 && <BreadcrumbSeparator />}
            </Fragment>
          ))}
        </BreadcrumbList>
      </Breadcrumb>
    </div>
  );
}
