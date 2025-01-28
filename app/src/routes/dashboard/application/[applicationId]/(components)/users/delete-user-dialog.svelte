<script lang="ts">
	import Trash from 'lucide-svelte/icons/trash-2';

	import * as Dialog from '$lib/components/ui/dialog';
	import { buttonVariants } from '$lib/components/ui/button';
	import Button from '$lib/components/ui/button/button.svelte';
	import type { IApplication } from '$lib/services/use-application-by-id-query';
	import type { Writable } from 'svelte/store';

	type Props = {
		user: IApplication["users"]["data"][number] | null;
		isDeleteModalOpened: Writable<boolean>;
	}

	let { user, isDeleteModalOpened }: Props = $props();

	let isLoading = $state(false);

	function handler() {
		isLoading = true;

		// Logic here
		isLoading = false;
		
		isDeleteModalOpened.set(false);
	}
</script>

<Dialog.Root open={$isDeleteModalOpened} onOpenChange={(value) => isDeleteModalOpened.set(value)}>
	<Dialog.Content class="sm:max-w-[450px]">
		<Dialog.Header>
			<Dialog.Title>Delete User</Dialog.Title>
			<Dialog.Description>
				On deleting this user, it will be permanently removed from the application. Are you sure?
			</Dialog.Description>
		</Dialog.Header>

		<Dialog.Footer>
			<Dialog.Close class={buttonVariants({ variant: 'outline' })}>Cancel</Dialog.Close>

			<Button type="submit" onclick={handler} variant="destructive">
				{#if isLoading}
					Deleting...
				{:else}
					Delete
				{/if}
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
