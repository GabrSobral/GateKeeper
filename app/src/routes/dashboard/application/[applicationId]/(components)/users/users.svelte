<script lang="ts">
	import { goto } from '$app/navigation';

	import ChevronDown from 'lucide-svelte/icons/chevron-down';
	import {
		type ColumnDef,
		type ColumnFiltersState,
		type PaginationState,
		type RowSelectionState,
		type SortingState,
		type VisibilityState,
		getCoreRowModel,
		getFilteredRowModel,
		getPaginationRowModel,
		getSortedRowModel
	} from '@tanstack/table-core';
	import { createRawSnippet } from 'svelte';
	
	import DataTableCheckbox from './data-table-checkbox.svelte';
	import DataTableEmailButton from './data-table-email-button.svelte';
	import DataTableActions from './data-table-actions.svelte';
	import * as Table from '$lib/components/ui/table';
	import { Button, buttonVariants } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { Input } from '$lib/components/ui/input';
	import {
		FlexRender,
		createSvelteTable,
		renderComponent,
		renderSnippet
	} from '$lib/components/ui/data-table';
	
	import type { IApplication } from '$lib/services/use-application-by-id-query';
	import DataTableUserRoles from './data-table-user-roles.svelte';
	import DataTableUserStatus from './data-table-user-status.svelte';
	import { cn } from '$lib/utils';
	import DeleteUserDialog from './delete-user-dialog.svelte';
	import { writable } from 'svelte/store';

	type Props = { application?: IApplication | null }
	type ApplicationUser = IApplication["users"]["data"][number];

	let isDeleteModalOpened = writable(false);
	let selectedUser = $state<ApplicationUser | null>(null);

	let { application }: Props = $props();

	const columns: ColumnDef<ApplicationUser>[] = [
		{
			id: 'select',
			header: ({ table }) =>
				renderComponent(DataTableCheckbox, {
					checked: table.getIsAllPageRowsSelected(),
					indeterminate: table.getIsSomePageRowsSelected() && !table.getIsAllPageRowsSelected(),
					onCheckedChange: (value) => table.toggleAllPageRowsSelected(!!value),
					'aria-label': 'Select all'
				}),
			cell: ({ row }) =>
				renderComponent(DataTableCheckbox, {
					checked: row.getIsSelected(),
					onCheckedChange: (value) => row.toggleSelected(!!value),
					'aria-label': 'Select row'
				}),
			enableSorting: false,
			enableHiding: false
		},

		{
			accessorKey: 'displayName',
			header: 'Display Name',
			cell: ({ row }) => {
				const statusSnippet = createRawSnippet<[string]>((getStatus) => {
					const status = getStatus();
					return {
						render: () => `<div class="capitalize">${status}</div>`
					};
				});
				return renderSnippet(statusSnippet, row.getValue('displayName'));
			}
		},

		{
			accessorKey: 'email',
			header: ({ column }) =>
				renderComponent(DataTableEmailButton, {
					onclick: () => column.toggleSorting(column.getIsSorted() === 'asc')
				}),
			cell: ({ row }) => {
				const emailSnippet = createRawSnippet<[string]>((getEmail) => {
					const email = getEmail();
					return {
						render: () => `<div class="lowercase">${email}</div>`
					};
				});

				return renderSnippet(emailSnippet, row.getValue('email'));
			}
		},

		{
			accessorKey: 'roles',
			header: "Roles",
			cell: ({ row }) => renderComponent(DataTableUserRoles, {
				roles: row.getValue('roles') as ApplicationUser['roles']
			})
		},

		{
			accessorKey: 'deactivatedAt',
			header: 'Status',
			cell: ({ row }) => renderComponent(DataTableUserStatus, {
				deactivatedAt: row.getValue('deactivatedAt') as ApplicationUser['deactivatedAt']
			}),
		},
		{
			id: 'actions',
			enableHiding: false,
			cell: ({ row }) => renderComponent(DataTableActions, { 
				id: row.original.id,
				selectUserToDelete: () => {
					selectedUser = row.original;
					isDeleteModalOpened.set(true);
				} 
			})
		}
	];

	let pagination = $state<PaginationState>({ pageIndex: 0, pageSize: 10 });
	let sorting = $state<SortingState>([]);
	let columnFilters = $state<ColumnFiltersState>([]);
	let rowSelection = $state<RowSelectionState>({});
	let columnVisibility = $state<VisibilityState>({});

	const table = createSvelteTable({
		get data() {
			return application?.users.data ?? [];
		},
		columns,
		state: {
			get pagination() { return pagination },
			get sorting() { return sorting },
			get columnVisibility() { return columnVisibility },
			get rowSelection() { return rowSelection },
			get columnFilters() { return columnFilters }
		},
		getCoreRowModel: getCoreRowModel(),
		getPaginationRowModel: getPaginationRowModel(),
		getSortedRowModel: getSortedRowModel(),
		getFilteredRowModel: getFilteredRowModel(),
		onPaginationChange: (updater) => {
			if (typeof updater === 'function') {
				pagination = updater(pagination);
			} else {
				pagination = updater;
			}
		},
		onSortingChange: (updater) => {
			if (typeof updater === 'function') {
				sorting = updater(sorting);
			} else {
				sorting = updater;
			}
		},
		onColumnFiltersChange: (updater) => {
			if (typeof updater === 'function') {
				columnFilters = updater(columnFilters);
			} else {
				columnFilters = updater;
			}
		},
		onColumnVisibilityChange: (updater) => {
			if (typeof updater === 'function') {
				columnVisibility = updater(columnVisibility);
			} else {
				columnVisibility = updater;
			}
		},
		onRowSelectionChange: (updater) => {
			if (typeof updater === 'function') {
				rowSelection = updater(rowSelection);
			} else {
				rowSelection = updater;
			}
		}
	});
</script>

<div class="mx-auto w-full">
	<div class="flex items-center py-4">
		<Input
			placeholder="Filter emails..."
			value={(table.getColumn('email')?.getFilterValue() as string) ?? ''}
			oninput={(e) => table.getColumn('email')?.setFilterValue(e.currentTarget.value)}
			onchange={(e) => {
				table.getColumn('email')?.setFilterValue(e.currentTarget.value);
			}}
			class="max-w-sm"
		/>

		<DropdownMenu.Root>
			<a 
				href={`/dashboard/application/${application?.id}/user/create-user`} 
				class={cn(buttonVariants({ variant: 'default' }), "ml-4")}
			>
				Add User
			</a>

			<DropdownMenu.Trigger>
				{#snippet child({ props })}
					<Button {...props} variant="outline" class="ml-auto">
						Columns <ChevronDown class="ml-2 size-4" />
					</Button>
				{/snippet}
			</DropdownMenu.Trigger>

			<DropdownMenu.Content align="end">
				{#each table.getAllColumns().filter((col) => col.getCanHide()) as column}
					<DropdownMenu.CheckboxItem
						class="capitalize"
						bind:checked={() => column.getIsVisible(), (v) => column.toggleVisibility(!!v)}
					>
						{column.id}
					</DropdownMenu.CheckboxItem>
				{/each}
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	</div>

	<div class="rounded-md border">
		<Table.Root>
			<Table.Header>
				{#each table.getHeaderGroups() as headerGroup (headerGroup.id)}
					<Table.Row>
						{#each headerGroup.headers as header (header.id)}
							<Table.Head class="[&:has([role=checkbox])]:pl-3">
								{#if !header.isPlaceholder}
									<FlexRender
										content={header.column.columnDef.header}
										context={header.getContext()}
									/>
								{/if}
							</Table.Head>
						{/each}
					</Table.Row>
				{/each}
			</Table.Header>

			<Table.Body>
				{#each table.getRowModel().rows as row (row.id)}
					<Table.Row 
						data-state={row.getIsSelected() && 'selected'} 
						onclick={() => goto(`/dashboard/application/${application?.id}/user/${row.original.id}`)}
						class="hover:cursor-pointer"
					>
						{#each row.getVisibleCells() as cell (cell.id)}
							<Table.Cell class="[&:has([role=checkbox])]:pl-3">
								<FlexRender content={cell.column.columnDef.cell} context={cell.getContext()} />
							</Table.Cell>
						{/each}
					</Table.Row>
				{:else}
					<Table.Row>
						<Table.Cell colspan={columns.length} class="h-24 text-center">No results.</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	</div>

	<div class="flex items-center justify-end space-x-2 pt-4">
		<div class="text-muted-foreground flex-1 text-sm">
			{table.getFilteredSelectedRowModel().rows.length} of
			{table.getFilteredRowModel().rows.length} row(s) selected.
		</div>

		<div class="space-x-2">
			<Button
				variant="outline"
				size="sm"
				onclick={() => table.previousPage()}
				disabled={!table.getCanPreviousPage()}
			>
				Previous
			</Button>
			
			<Button
				variant="outline"
				size="sm"
				onclick={() => table.nextPage()}
				disabled={!table.getCanNextPage()}
			>
				Next
			</Button>
		</div>
	</div>
</div>

<DeleteUserDialog user={selectedUser} {isDeleteModalOpened}/>
