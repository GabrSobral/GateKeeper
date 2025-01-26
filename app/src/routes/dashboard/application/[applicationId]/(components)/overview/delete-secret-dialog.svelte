<script lang="ts">
	import Trash from "lucide-svelte/icons/trash-2"
	
	import * as Dialog from '$lib/components/ui/dialog';
	import { buttonVariants } from '$lib/components/ui/button';
	import Button from '$lib/components/ui/button/button.svelte';
	import type { IApplication } from "$lib/services/use-application-by-id-query";

	let isLoading = $state(false);

	type Props = {
		secret: IApplication["secrets"][number];
	}

	let { secret }: Props = $props();

	function handler() {
		isLoading = true;

		// Logic here
		isLoading = false;
	}
</script>

<Dialog.Root>
	<Dialog.Trigger class={buttonVariants({ variant: 'outline' })} >
		<Trash />
	</Dialog.Trigger>

	<Dialog.Content class="sm:max-w-[450px]">
		<Dialog.Header>
			<Dialog.Title>Delete Application Secret</Dialog.Title>
			<Dialog.Description>
				On deleting this secret ({secret.name}), it will be permanently removed from the server. Are you sure?
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
