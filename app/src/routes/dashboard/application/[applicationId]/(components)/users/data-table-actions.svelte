<script lang="ts">
	import Ellipsis from 'lucide-svelte/icons/ellipsis';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import { goto } from '$app/navigation';
	import { copy } from '$lib/utils';
	import { page } from '$app/state';

	type Props = { 
		id: string;
		selectUserToDelete: () => void; 
	}
	
	let { id, selectUserToDelete }: Props = $props();

	let applicationId = page.params.applicationId;
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
			<DropdownMenu.Item onclick={() => copy(id)}>
				Copy user ID
			</DropdownMenu.Item>
		</DropdownMenu.Group>
		
		<DropdownMenu.Separator />

		<DropdownMenu.Item onclick={() => goto(`/dashboard/application/${applicationId}/user/${id}?edit=true`)}>
			Update User
		</DropdownMenu.Item>

		<DropdownMenu.Item class="text-red-500 font-bold" onclick={selectUserToDelete}>
			Remove User
		</DropdownMenu.Item>
	</DropdownMenu.Content>
</DropdownMenu.Root>
