"use client";

import { NavProps } from "..";
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "../../ui/sidebar";

type Props = {
  nav: NavProps[];
  selectSection: (section: NavProps) => void;
  selectedSection?: NavProps | null;
};

export function SettingsSidebar({
  nav,
  selectSection,
  selectedSection,
}: Props) {
  return (
    <Sidebar collapsible="none" className="hidden md:flex">
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarGroupLabel>General</SidebarGroupLabel>
            <SidebarMenu>
              {nav.map((item) => (
                <SidebarMenuItem key={item.name}>
                  <SidebarMenuButton
                    isActive={item.name === selectedSection?.name}
                    onClick={() => selectSection(item)}
                  >
                    <>
                      <item.icon />
                      <span>{item.name}</span>
                    </>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
}
