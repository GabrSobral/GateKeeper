<script lang="ts">
	import CircleAlert from 'lucide-svelte/icons/circle-alert';

	import * as Dialog from '$lib/components/ui/dialog';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import * as Alert from '$lib/components/ui/alert';

	import { buttonVariants } from '$lib/components/ui/button';
	import Button from '$lib/components/ui/button/button.svelte';
	import DatePicker from '$lib/components/ui/date-picker/date-picker.svelte';
	import { createApplicationSecretApi } from '$lib/services/create-application-secret';
	import { toast } from 'svelte-sonner';
	import { organizationStore } from '$lib/stores/organization';
	import { page } from '$app/state';
	import type { DateValue } from '@internationalized/date';

	let isLoading = $state(false);

	let copied = $state(false);
	let secret = $state('');

	let secretName = $state('');
	let expiresAt = $state<DateValue | undefined>(undefined);

	let applicationId = $derived(page.params.applicationId);

	$inspect(expiresAt);

	async function generate() {
		isLoading = true;

		const [response, err] = await createApplicationSecretApi(
			{
				name: secretName,
				expiresAt: expiresAt?.toDate('pt-br') || null,
				applicationId: applicationId,
				organizationId: $organizationStore?.id
			},
			{ accessToken: '' }
		);

		if (err) {
			toast.error('Failed to generate secret');
			console.error(err);
			return;
		}

		secret = response?.value || '';

		isLoading = false;
	}

	function copySecret(secret: string) {
		navigator.clipboard.writeText(secret).then(() => {
			copied = true;

			setTimeout(() => (copied = false), 1000);
		});
	}
</script>

<Dialog.Root onOpenChange={(isOpened) => isOpened && (secret = '')}>
	<Dialog.Trigger class={buttonVariants({ variant: 'default' })}>New Secret</Dialog.Trigger>

	<Dialog.Content class="top-[40%] sm:max-w-[450px]">
		<Dialog.Header>
			<Dialog.Title>New Application Secret</Dialog.Title>
			<Dialog.Description>
				Generate a new secret for your application. Keep it safe.
			</Dialog.Description>
		</Dialog.Header>

		<div class="grid gap-4 py-4">
			<div class="flex flex-col gap-3">
				<Label for="name">Secret Name</Label>
				<Input
					id="name"
					placeholder="E.g: My Ultra Application Secret"
					class="col-span-3"
					bind:value={secretName}
				/>
			</div>

			<div class="flex flex-col gap-3">
				<Label for="username">Expires At</Label>
				<DatePicker value={expiresAt} />
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
