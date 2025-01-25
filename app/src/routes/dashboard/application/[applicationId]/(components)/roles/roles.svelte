<script lang="ts">
	import Ellipsis from 'lucide-svelte/icons/ellipsis';

	import * as Table from '$lib/components/ui/table';
	import { Button } from '$lib/components/ui/button';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import Input from '$lib/components/ui/input/input.svelte';

	const invoices = [
		{
			id: 'INV001',
			name: 'User',
			description: 'A user of the application. Can view and manage data.'
		},
		{
			id: 'INV002',
			name: 'Admin',
			description: 'An admin of the application. Can manage users and roles.'
		}
	]
</script>

<Table.Root>
	<Table.Caption>A list of application roles.</Table.Caption>
	<Table.Header>
		<Table.Row>
			<Table.Head class="w-[200px]">ID</Table.Head>
			<Table.Head>Name</Table.Head>
			<Table.Head>Description</Table.Head>
			<Table.Head>Actions</Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		{#each invoices as role, i (i)}
			<Table.Row>
				<Table.Cell class="font-medium">{role.id}</Table.Cell>
				<Table.Cell class="font-bold">{role.name}</Table.Cell>
				<Table.Cell>{role.description || "-"}</Table.Cell>
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
