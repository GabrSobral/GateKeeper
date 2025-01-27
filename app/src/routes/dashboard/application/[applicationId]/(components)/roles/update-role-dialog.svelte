<script lang="ts">
	import type { Writable } from 'svelte/store';

	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';

	import { buttonVariants } from '$lib/components/ui/button';
	import Button from '$lib/components/ui/button/button.svelte';
	import { cn } from '$lib/utils';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import type { IApplication } from '$lib/services/use-application-by-id-query';

	type Props = {
		role: IApplication["roles"]["data"][number] | null;
		isDeleteModalOpened: Writable<boolean>;
	}

	let { role, isDeleteModalOpened }: Props = $props();


	let isLoading = $state(false);

	let nameDraft = $state(role?.name || "");
	let descriptionDraft = $state(role?.description || "");

	$effect(() => {
		nameDraft = role?.name || "";
		descriptionDraft = role?.description || "";
	});
	

	function handle() {
		isLoading = true;

		isLoading = false;
	}
</script>

<Dialog.Root open={$isDeleteModalOpened} onOpenChange={isOpened => isDeleteModalOpened.set(isOpened)}>
	<Dialog.Content class="sm:max-w-[450px]">
		<Dialog.Header>
			<Dialog.Title>Update Application Role</Dialog.Title>
			<Dialog.Description>
				Update a role for your application. Handle permissions and access.
			</Dialog.Description>
		</Dialog.Header>

		<div class="grid gap-4 py-4">
			<div class="flex flex-col gap-3">
				<Label for="name">Name <span class="text-red-500">*</span></Label>
				<Input id="name" placeholder="Type the role name" bind:value={nameDraft}/>
			</div>

			<div class="flex flex-col gap-3">
				<Label for="description">Description <span class="text-red-500">*</span> ({120 - descriptionDraft.length})</Label>
				<Textarea id="description" placeholder="Type the role description" bind:value={descriptionDraft} maxlength={120} />
			</div>
		</div>

		<Dialog.Footer>
			<Button type="submit" onclick={handle}>
                {#if isLoading}
                    Updating...
                {:else}
                    Apply changes
                {/if}
            </Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
