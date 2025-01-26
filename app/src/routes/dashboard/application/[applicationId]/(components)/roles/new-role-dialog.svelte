<script lang="ts">
	import CircleAlert from 'lucide-svelte/icons/circle-alert';

	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import * as Alert from '$lib/components/ui/alert';

	import { buttonVariants } from '$lib/components/ui/button';
	import Button from '$lib/components/ui/button/button.svelte';
	import DatePicker from '$lib/components/ui/date-picker/date-picker.svelte';
	import { cn } from '$lib/utils';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';

	let isLoading = $state(false);

	let name = $state("");
	let description = $state('');

	function generate() {
		isLoading = true;

		clear()

		isLoading = false;
	}

	function clear() {
		name = '';
		description = '';
	}
</script>

<Dialog.Root onOpenChange={isOpened => isOpened && clear()}>
	<Dialog.Trigger class={cn(buttonVariants({ variant: 'default' }), "ml-4")}>
		Add Role
	</Dialog.Trigger>

	<Dialog.Content class="sm:max-w-[450px]">
		<Dialog.Header>
			<Dialog.Title>New Application Role</Dialog.Title>
			<Dialog.Description>
				Create a new role for your application. Handle permissions and access.
			</Dialog.Description>
		</Dialog.Header>

		<div class="grid gap-4 py-4">
			<div class="flex flex-col gap-3">
				<Label for="name">Name <span class="text-red-500">*</span></Label>
				<Input id="name" placeholder="Type the role name" bind:value={name}/>
			</div>

			<div class="flex flex-col gap-3">
				<Label for="description">Description <span class="text-red-500">*</span> ({120 - description.length})</Label>
				<Textarea id="description" placeholder="Type the role description" bind:value={description} maxlength={120} />
			</div>
		</div>

		<Dialog.Footer>
			<Button type="submit" onclick={generate}>
                {#if isLoading}
                    Creating...
                {:else}
                    Create
                {/if}
            </Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
