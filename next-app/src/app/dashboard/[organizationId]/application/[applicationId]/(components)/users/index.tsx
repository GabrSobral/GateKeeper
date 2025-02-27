"use client";

import Link from "next/link";
import { useState } from "react";
import { ArrowUpDown, ChevronDown, MoreHorizontal } from "lucide-react";
import { useParams } from "next/navigation";
import {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table";

import { Button, buttonVariants } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { Checkbox } from "@/components/ui/checkbox";

import { cn, copy } from "@/lib/utils";

import { DeleteUserDialog } from "./delete-user-dialog";
import { IApplication } from "@/services/dashboard/get-application-by-id";
import { deleteApplicationUserApi } from "@/services/dashboard/delete-application-user";

export type ApplicationUser = IApplication["users"]["data"][number];

type Props = {
  application: IApplication | null;
};

export function Users({ application }: Props) {
  const [selectedUser, setSelectedUser] = useState<ApplicationUser | null>(
    null
  );
  const [isDeleteModalOpened, setIsDeleteModalOpened] = useState(false);

  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});
  const [users, setUsers] = useState<ApplicationUser[]>(
    application?.users.data || []
  );

  const { applicationId, organizationId } = useParams() as {
    applicationId: string;
    organizationId: string;
  };

  const columns: ColumnDef<ApplicationUser>[] = [
    {
      id: "select",
      header: ({ table }) => (
        <Checkbox
          checked={
            table.getIsAllPageRowsSelected() ||
            (table.getIsSomePageRowsSelected() && "indeterminate")
          }
          onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
          aria-label="Select all"
        />
      ),
      cell: ({ row }) => (
        <Checkbox
          checked={row.getIsSelected()}
          onCheckedChange={(value) => row.toggleSelected(!!value)}
          aria-label="Select row"
        />
      ),
      enableSorting: false,
      enableHiding: false,
    },

    {
      accessorKey: "displayName",
      header: "Display Name",
      cell: ({ row }) => (
        <div className="capitalize">{row.getValue("displayName")}</div>
      ),
    },

    {
      accessorKey: "email",
      header: ({ column }) => {
        return (
          <Button
            variant="ghost"
            onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
          >
            Email
            <ArrowUpDown />
          </Button>
        );
      },
      cell: ({ row }) => (
        <div className="lowercase">{row.getValue("email")}</div>
      ),
    },

    {
      accessorKey: "roles",
      header: "Roles",
      cell: ({ row }) => {
        const roles = row.getValue("roles") as ApplicationUser["roles"];

        return (
          <div className="flex gap-1">
            {roles?.map((role) => (
              <Badge key={role.id} variant="default" color="green">
                {role.name}
              </Badge>
            ))}
          </div>
        );
      },
    },
    {
      id: "actions",
      enableHiding: false,
      cell: ({ row }) => {
        const user = row.original;

        return (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" className="h-8 w-8 p-0">
                <span className="sr-only">Open menu</span>
                <MoreHorizontal />
              </Button>
            </DropdownMenuTrigger>

            <DropdownMenuContent align="end">
              <DropdownMenuLabel>Actions</DropdownMenuLabel>
              <DropdownMenuItem onClick={() => copy(user.id)}>
                Copy user ID
              </DropdownMenuItem>

              <DropdownMenuSeparator />

              <DropdownMenuItem>Update User</DropdownMenuItem>
              <DropdownMenuItem
                className="text-red-500 font-bold"
                onClick={() => {
                  setSelectedUser(user);
                  setIsDeleteModalOpened(true);
                }}
              >
                Remove User
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        );
      },
    },
  ];

  const table = useReactTable({
    data: users,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
  });

  async function deleteSelection() {
    const currentUsers = table
      .getFilteredSelectedRowModel()
      .rows.map((item) => item.original);

    await Promise.all(
      currentUsers.map(async (row) => {
        await deleteApplicationUserApi(
          { applicationId, organizationId, userId: row.id },
          { accessToken: "" }
        );
      })
    );

    table.setRowSelection({}); // Clear selection

    setUsers((state) => state.filter((role) => !currentUsers.includes(role)));
  }

  return (
    <>
      <div className="mx-auto w-full">
        <div className="flex items-center py-4">
          <Input
            placeholder="Filter emails..."
            value={(table.getColumn("email")?.getFilterValue() as string) ?? ""}
            onInput={(e) =>
              table.getColumn("email")?.setFilterValue(e.currentTarget.value)
            }
            onChange={(e) => {
              table.getColumn("email")?.setFilterValue(e.currentTarget.value);
            }}
            className="max-w-sm"
          />

          <DropdownMenu>
            <Link
              href={`/dashboard/${organizationId}/application/${applicationId}/user/create-user`}
              className={cn(buttonVariants({ variant: "default" }), "ml-4")}
            >
              Add User
            </Link>

            {table.getFilteredSelectedRowModel().rows.length !== 0 && (
              <Button
                type="button"
                variant="destructive"
                onClick={deleteSelection}
                className="ml-4"
              >
                Delete Selection
              </Button>
            )}

            <DropdownMenuTrigger asChild>
              <Button variant="outline" className="ml-auto">
                Columns <ChevronDown />
              </Button>
            </DropdownMenuTrigger>

            <DropdownMenuContent align="end">
              {table
                .getAllColumns()
                .filter((column) => column.getCanHide())
                .map((column) => {
                  return (
                    <DropdownMenuCheckboxItem
                      key={column.id}
                      className="capitalize"
                      checked={column.getIsVisible()}
                      onCheckedChange={(value) =>
                        column.toggleVisibility(!!value)
                      }
                    >
                      {column.id}
                    </DropdownMenuCheckboxItem>
                  );
                })}
            </DropdownMenuContent>
          </DropdownMenu>
        </div>

        <div className="rounded-md border">
          <Table>
            <TableHeader>
              {table.getHeaderGroups().map((headerGroup) => (
                <TableRow key={headerGroup.id}>
                  {headerGroup.headers.map((header) => {
                    return (
                      <TableHead key={header.id}>
                        {header.isPlaceholder
                          ? null
                          : flexRender(
                              header.column.columnDef.header,
                              header.getContext()
                            )}
                      </TableHead>
                    );
                  })}
                </TableRow>
              ))}
            </TableHeader>

            <TableBody>
              {table.getRowModel().rows?.length ? (
                table.getRowModel().rows.map((row) => (
                  <TableRow
                    key={row.id}
                    data-state={row.getIsSelected() && "selected"}
                  >
                    {row.getVisibleCells().map((cell) => (
                      <TableCell key={cell.id}>
                        {flexRender(
                          cell.column.columnDef.cell,
                          cell.getContext()
                        )}
                      </TableCell>
                    ))}
                  </TableRow>
                ))
              ) : (
                <TableRow>
                  <TableCell
                    colSpan={columns.length}
                    className="h-24 text-center"
                  >
                    No results.
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>

        <div className="flex items-center justify-end space-x-2 py-4">
          <div className="flex-1 text-sm text-muted-foreground">
            {table.getFilteredSelectedRowModel().rows.length} of{" "}
            {table.getFilteredRowModel().rows.length} row(s) selected.
          </div>

          <div className="space-x-2">
            <Button
              variant="outline"
              size="sm"
              onClick={() => table.previousPage()}
              disabled={!table.getCanPreviousPage()}
            >
              Previous
            </Button>

            <Button
              variant="outline"
              size="sm"
              onClick={() => table.nextPage()}
              disabled={!table.getCanNextPage()}
            >
              Next
            </Button>
          </div>
        </div>
      </div>

      <DeleteUserDialog
        isOpened={isDeleteModalOpened}
        onOpenChange={setIsDeleteModalOpened}
        user={selectedUser}
        removeUser={(user) =>
          setUsers((state) => state.filter((item) => item.id !== user.id))
        }
      />
    </>
  );
}
