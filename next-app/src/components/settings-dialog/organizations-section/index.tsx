import { useState } from "react";

import { DefaultPage } from "./default-page";
import { AddOrganization } from "./add-organization";
import { EditOrganization } from "./edit-organization";
import { ViewOrganization } from "./view-organization";

export type OrganizationsPages =
  | "default"
  | "add-organization"
  | "edit-organization"
  | "view-organization";

export function OrganizationsSection() {
  const [page, setPage] = useState<OrganizationsPages>("default");

  return page === "default" ? (
    <DefaultPage setPage={setPage} />
  ) : page === "add-organization" ? (
    <AddOrganization setPage={setPage} />
  ) : page === "edit-organization" ? (
    <EditOrganization setPage={setPage} />
  ) : page === "view-organization" ? (
    <ViewOrganization setPage={setPage} />
  ) : null;
}
