<script lang="ts">
	import { tick } from 'svelte';
	import Check from 'lucide-svelte/icons/check';
	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down';

	import { Button } from '$lib/components/ui/button/index';
	import * as Command from '$lib/components/ui/command/index';
	import * as Popover from '$lib/components/ui/popover/index';

	import { cn } from '$lib/utils';
	import { organizationStore } from '$lib/stores/organization';
	import { useOrganizationsQuery } from '$lib/services/use-organizations-query';
	import { page } from '$app/state';

	let open = $state(false);
	let triggerRef = $state<HTMLButtonElement>(null!);

	let organizations = useOrganizationsQuery({ accessToken: '' });
	let organizationId = $derived(page.url.searchParams.get('organizationId') || '');

	$effect(() => {
		if (organizationId && $organizationStore === null && $organizations.data) {
			const organization = $organizations.data.find(
				(organization) => organization.id === organizationId
			);

			if (organization) {
				organizationStore.set(organization);
			}
		}
	});

	function closeAndFocusTrigger() {
		open = false;
		tick().then(() => triggerRef.focus());
	}

	$effect(() => {
		if ($organizationStore == null && $organizations?.data?.[0]) {
			organizationStore.set($organizations.data[0]);
		}
	});
</script>

<Popover.Root bind:open>
	<Popover.Trigger bind:ref={triggerRef}>
		{#snippet child({ props })}
			<Button
				variant="outline"
				class="w-full justify-between"
				{...props}
				role="combobox"
				aria-expanded={open}
			>
				{$organizationStore?.name || 'Select a organization...'}
				<ChevronsUpDown class="opacity-50" />
			</Button>
		{/snippet}
	</Popover.Trigger>

	<Popover.Content class="w-[240px] p-0">
		<Command.Root>
			<Command.Input placeholder="Search organization..." />
			<Command.List>
				<Command.Empty>No organization found.</Command.Empty>
				<Command.Group>
					{#each $organizations.data || [] as organization}
						<Command.Item
							value={organization.id}
							onSelect={() => {
								organizationStore.set(organization);
								closeAndFocusTrigger();
							}}
						>
							<Check class={cn($organizationStore?.id !== organization.id && 'text-transparent')} />
							{organization.name}
						</Command.Item>
					{/each}
				</Command.Group>
			</Command.List>
		</Command.Root>
	</Popover.Content>
</Popover.Root>
