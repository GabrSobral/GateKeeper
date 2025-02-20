"use client";

import Link from "next/link";
import { useParams } from "next/navigation";
import { LayoutPanelLeft, LogOut, Plus } from "lucide-react";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupAction,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";

import { OrganizationList } from "./OrganizationList";
import { useApplicationsSWR } from "@/services/dashboard/use-applications-swr";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";

import { ThemeToggle } from "../ui/theme-togle";

export function DashboardSidebar() {
  const organizationId = useParams().organizationId as string;
  const { data } = useApplicationsSWR({ organizationId }, { accessToken: "" });

  const items = [
    {
      title: "Applications",
      url: `/dashboard/${organizationId}/application`,
      icon: LayoutPanelLeft,
    },
  ];

  return (
    <Sidebar collapsible="icon">
      <SidebarHeader>
        <SidebarMenu>
          <OrganizationList />
        </SidebarMenu>
      </SidebarHeader>

      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>General</SidebarGroupLabel>

          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item, index) => (
                <SidebarMenuItem key={index}>
                  <SidebarMenuButton asChild tooltip={item.title}>
                    <Link href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        <SidebarGroup className="group-data-[collapsible=icon]:hidden">
          <SidebarGroupLabel>Applications</SidebarGroupLabel>

          <SidebarGroupAction title="Add Application" onClick={() => {}}>
            <Plus /> <span className="sr-only">Add Application</span>
          </SidebarGroupAction>

          <SidebarGroupContent>
            <SidebarMenu>
              {data?.map((application) => (
                <SidebarMenuItem key={application.id}>
                  <SidebarMenuButton asChild>
                    <Link
                      href={`/dashboard/${organizationId}/application/${application.id}`}
                    >
                      {application.name}
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton>
              <LogOut />
              Logout
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>

        <ThemeToggle />

        <div className="flex gap-2">
          <Avatar>
            <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>

          <div className="flex flex-col">
            <span className="text-sm font-semibold">John Doe</span>
            <span className="text-sm">Johndoe@email.com</span>
          </div>
        </div>
      </SidebarFooter>
    </Sidebar>
  );
}
