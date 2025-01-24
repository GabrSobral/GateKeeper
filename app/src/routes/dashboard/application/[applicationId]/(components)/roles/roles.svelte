<script lang="ts">
	import Ellipsis from 'lucide-svelte/icons/ellipsis';

	import * as Table from '$lib/components/ui/table';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import Badge from '$lib/components/ui/badge/badge.svelte';

	const invoices = [
		{
			id: 'INV001',
			name: 'User',
			permissions: [
				{ id: "12345", name: "Consult Status" }
			]
		},
		{
			id: 'INV002',
			name: 'Admin',
			permissions: [
				{ id: "123456", name: "Read Status" },
				{ id: "123457", name: "Consult Status" }
			]
		}
	]
</script>

<Table.Root>
	<Table.Caption>A list of application roles.</Table.Caption>
	<Table.Header>
		<Table.Row>
			<Table.Head class="w-[200px]">ID</Table.Head>
			<Table.Head>Name</Table.Head>
			<Table.Head>Permissions</Table.Head>
			<Table.Head>Actions</Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		{#each invoices as role, i (i)}
			<Table.Row>
				<Table.Cell class="font-medium">{role.id}</Table.Cell>
				<Table.Cell>{role.name}</Table.Cell>
				<Table.Cell class="flex gap-[1px]">
					{#each role.permissions as permission (permission.id)}
						<Badge>{permission.name}</Badge>
					{/each}
				</Table.Cell>
				<Table.Cell>
					<DropdownMenu.Root>
						<DropdownMenu.Trigger>
							{#snippet child({ props })}
								<Button {...props} variant="ghost" size="icon" class="relative size-8 p-0">
									<span class="sr-only">Open menu</span>
									<Ellipsis />
								</Button>
							{/snippet}
						</DropdownMenu.Trigger>
						
						<DropdownMenu.Content>
							<DropdownMenu.Group>
								<DropdownMenu.GroupHeading>Actions</DropdownMenu.GroupHeading>
								<DropdownMenu.Item onclick={() => navigator.clipboard.writeText(role.id)}>
									Copy role ID
								</DropdownMenu.Item>
							</DropdownMenu.Group>
							<DropdownMenu.Separator />
							<DropdownMenu.Item>View role details</DropdownMenu.Item>
						</DropdownMenu.Content>
					</DropdownMenu.Root>
				</Table.Cell>
			</Table.Row>
		{/each}
	</Table.Body>
</Table.Root>
