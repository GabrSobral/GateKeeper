<script lang="ts">
	import Ellipsis from 'lucide-svelte/icons/ellipsis';
	
	import { copy } from '$lib/utils';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import type { IApplication } from '$lib/services/use-application-by-id-query';

	type Props = { 
		role: IApplication["roles"]["data"][number] 
		selectRoleToDelete: () => void
		selectRoleToUpdate: () => void
	};

	let { role, selectRoleToDelete, selectRoleToUpdate }: Props = $props();
</script>

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
			<DropdownMenu.Item onclick={() => copy(role.id)}>
				Copy role ID
			</DropdownMenu.Item>
		</DropdownMenu.Group>
		
		<DropdownMenu.Separator />

		<DropdownMenu.Item onclick={selectRoleToUpdate}>
			Update Role
		</DropdownMenu.Item>

		<DropdownMenu.Item class="text-red-500 font-bold" onclick={selectRoleToDelete}>
			Remove Role
		</DropdownMenu.Item>
	</DropdownMenu.Content>
</DropdownMenu.Root>