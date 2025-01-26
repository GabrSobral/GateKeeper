<script lang="ts">
	import * as Tooltip from '$lib/components/ui/tooltip';
	import Label from '$lib/components/ui/label/label.svelte';
	import Checkbox from '$lib/components/ui/checkbox/checkbox.svelte';

	type Props = { isEditEnabled: boolean };

	let { isEditEnabled }: Props = $props();
</script>

<div class="flex flex-col gap-1">
	<span class="text-sm font-medium"> Application Roles </span>

	<span class="text-muted-foreground my-2 text-sm">
		Choose the roles that will be assigned to this user.
	</span>

	<div class="flex items-center space-x-2">
		<Checkbox id="asd1" disabled={!isEditEnabled} />

		<Tooltip.Provider>
			<Tooltip.Root>
				<Tooltip.Trigger>
					<Label
						for="asd1"
						data-isEditEnabled={isEditEnabled}
						class="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 data-[isEditEnabled=false]:font-normal"
					>
						User
					</Label>
				</Tooltip.Trigger>

				<Tooltip.Content>User Description</Tooltip.Content>
			</Tooltip.Root>
		</Tooltip.Provider>
	</div>

	<div class="flex items-center space-x-2">
		<Checkbox id="asd2" disabled={!isEditEnabled} />

		<Tooltip.Provider>
			<Tooltip.Root>
				<Tooltip.Trigger>
					<Label
						for="asd2"
						data-isEditEnabled={isEditEnabled}
						class="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 data-[isEditEnabled=false]:font-normal"
					>
						Admin
					</Label>
				</Tooltip.Trigger>

				<Tooltip.Content>Admin Description</Tooltip.Content>
			</Tooltip.Root>
		</Tooltip.Provider>
	</div>
</div>
