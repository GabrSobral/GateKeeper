<script lang="ts">
	import CircleAlert from 'lucide-svelte/icons/circle-alert';

	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import * as Alert from '$lib/components/ui/alert';

	import { buttonVariants } from '$lib/components/ui/button';
	import Button from '$lib/components/ui/button/button.svelte';
	import DatePicker from '$lib/components/ui/date-picker/date-picker.svelte';

	let isLoading = $state(false);

	let copied = $state(false);
	let secret = $state('');

	function generate() {
		isLoading = true;

		secret = 'aksdjask';

		isLoading = false;
	}

	function copySecret(secret: string) {
		navigator.clipboard.writeText(secret).then(() => {
			copied = true;

			setTimeout(() => {
				copied = false;
			}, 1000);
		});
	}
</script>

<Dialog.Root onOpenChange={isOpened => isOpened && (secret = '')}>
	<Dialog.Trigger class={buttonVariants({ variant: 'default' })}>New Secret</Dialog.Trigger>

	<Dialog.Content class="top-[40%] sm:max-w-[450px]">
		<Dialog.Header>
			<Dialog.Title>New Application Secret</Dialog.Title>
			<Dialog.Description>
				Generate a new secret for your application. Keep it safe.
			</Dialog.Description>
		</Dialog.Header>

		<div class="grid gap-4 py-4">
			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="name" class="text-right">Secret Name</Label>
				<Input id="name" placeholder="E.g: My Ultra Application Secret" class="col-span-3" />
			</div>

			<div class="grid grid-cols-4 items-center gap-4">
				<Label for="username" class="text-right">Expires At</Label>
				<DatePicker value={undefined} />
			</div>
		</div>

		{#if secret}
			<Alert.Root class="bg-orange-500/10">
				<CircleAlert />

				<Alert.Title>Wait!</Alert.Title>
				<Alert.Description>
					The secret will be only visible at this moment. Save it on another safe place or copy to
					use now.
				</Alert.Description>
			</Alert.Root>

			<div class="flex items-center justify-between gap-4 rounded-md bg-gray-100 p-4">
				<span class="w-full text-center text-lg font-bold"> daslçdasd!@#123123sad0912ld </span>

				<Button onclick={() => copySecret('aksdjask')} class="min-w-[75px]">
					{#if copied}
						Copied!
					{:else}
						Copy
					{/if}
				</Button>
			</div>
		{/if}

		<Dialog.Footer>
			<Button type="submit" onclick={generate}>
                {#if isLoading}
                    Generating...
                {:else}
                    Generate
                {/if}
            </Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
