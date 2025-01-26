<script lang="ts">
	import { cn } from '$lib/utils';

	import * as Dialog from '$lib/components/ui/dialog';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { buttonVariants } from '$lib/components/ui/button';
	import Button from '$lib/components/ui/button/button.svelte';
	import { toast } from 'svelte-sonner';

	let isLoading = $state(false);

	function handler() {
		isLoading = true;

		// Logic here
		isLoading = false;

		toast.success('Password reset successfully');
	}
</script>

<Dialog.Root>
	<Dialog.Trigger class={cn(buttonVariants({ variant: 'secondary' }), 'w-fit')}>
		Reset Password
	</Dialog.Trigger>

	<Dialog.Content class="sm:max-w-[550px]">
		<Dialog.Header>
			<Dialog.Title>Reset Password</Dialog.Title>
			<Dialog.Description>
				On confirm, the user will receive an e-mail with the new password, and the user will be
				required to change it on the next login.
			</Dialog.Description>
		</Dialog.Header>

		<div class="flex flex-col gap-3">
			<Label for="temp-password-input">Temporary Password</Label>
			<Input id="temp-password-input" type="password" placeholder="Type the temporary password" />
		</div>

		<Dialog.Footer>
			<Dialog.Close class={buttonVariants({ variant: 'outline' })}>Cancel</Dialog.Close>

			<Button type="submit" onclick={handler}>
				{#if isLoading}
					Resetting...
				{:else}
					Confirm
				{/if}
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
