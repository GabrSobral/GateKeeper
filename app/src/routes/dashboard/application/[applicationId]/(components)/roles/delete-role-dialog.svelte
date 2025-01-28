<script lang="ts">
	
	import type { Writable } from "svelte/store";

	import * as Dialog from '$lib/components/ui/dialog';
	import { buttonVariants } from '$lib/components/ui/button';
	import Button from '$lib/components/ui/button/button.svelte';
	import type { IApplication } from "$lib/services/use-application-by-id-query";

	let isLoading = $state(false);

	type Props = {
		role: IApplication["roles"]["data"][number] | null;
		isDeleteModalOpened: Writable<boolean>;
	}

	let { role, isDeleteModalOpened }: Props = $props();

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
			<Dialog.Title>Delete Application Role</Dialog.Title>
			<Dialog.Description>v 
				On deleting this role ({role?.name}), it will be permanently removed from the server. Are you sure?
			</Dialog.Description>
		</Dialog.Header>


		<Dialog.Footer>
			<Dialog.Close class={buttonVariants({ variant: 'outline' })}>
				Cancel
			</Dialog.Close>

			<Button type="submit" onclick={handler}>
                {#if isLoading}
                    Deleting...
                {:else}
                    Delete
                {/if}
            </Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root> 
