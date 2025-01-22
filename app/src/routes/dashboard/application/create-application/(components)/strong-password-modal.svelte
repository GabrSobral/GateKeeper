<script lang="ts">
	import { Eye, Key, Copy, Check, EyeOff, RefreshCcw } from 'lucide-svelte';

	import * as Dialog from '$lib/components/ui/dialog';
	import Button from '$lib/components/ui/button/button.svelte';

	let { setPassword }: { setPassword: (valueText: string) => void } = $props();

	let isCopied = $state(false);
	let isModalOpened = $state(false);
	let isPasswordVisible = $state(true);
	let strongPassword = $state(generatePassword());

	function generatePassword(): string {
		const passwordLength = 32;
		const charSet =
			'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;:,.<>?';

		let password = '';
		for (let i = 0; i < passwordLength; i++) {
			const randomIndex = Math.floor(Math.random() * charSet.length);
			password += charSet[randomIndex];
		}

		return password;
	}

	function copyToClipboard() {
		navigator.clipboard.writeText(strongPassword);
		isCopied = true;

		setTimeout(() => (isCopied = false), 1000);
	}

	function handler() {
		setPassword(strongPassword);
		isModalOpened = false;
	}
</script>

<Dialog.Root open={isModalOpened} onOpenChange={(value) => (isModalOpened = value)}>
	<Dialog.Trigger 
		type="button" 
		class="text-sm text-blue-500 underline-offset-2 hover:underline"
	>
		Generate a strong password
	</Dialog.Trigger>

	<Dialog.Content class="max-w-[36rem]">
		<Dialog.Title>Generate a strong password</Dialog.Title>

		<Dialog.Description>
			A strong and random password helps protect your online accounts and personal information from
			cyber threats
		</Dialog.Description>

		<Dialog.Description>
			By using a mix of characters and symbols, you can reduce the risk of unauthorized access and
			identity theft.
		</Dialog.Description>

		<Dialog.Description>Protect the application by choosing a strong password.</Dialog.Description>

		<div
			class="bg-background mt-4 flex items-center justify-between rounded-[10px] p-2 px-4 shadow-inner brightness-90"
		>
			<strong class="text-primary text-[1.15rem] font-bold tracking-widest">
				{isPasswordVisible
					? strongPassword
					: strongPassword
							.split('')
							.map(() => '*')
							.join('')}
			</strong>

			<div class="flex items-center gap-3">
				<button
					type="button"
					onclick={() => (isPasswordVisible = !isPasswordVisible)}
					title="Show password"
					class="flex w-fit items-center justify-center"
				>
					{#if isPasswordVisible}
						<EyeOff size={28} />
					{:else}
						<Eye size={28} />
					{/if}
				</button>

				<button
					type="button"
					onclick={() => (strongPassword = generatePassword())}
					title="Regenerate the password"
					class="flex w-fit items-center justify-center transition-all active:scale-90"
				>
					<RefreshCcw size={28} />
				</button>
			</div>
		</div>

		<div class="ml-auto flex gap-2">
			<Button variant="outline" type="button" onclick={copyToClipboard}>
				{#if isCopied}
					<Check size={28} />
					Copied
				{:else}
					<Copy size={28} />
					Copy
				{/if}
			</Button>

			<Button type="button" onclick={handler}>
				<Key size={28} />
				Use this password
			</Button>
		</div>
	</Dialog.Content>
</Dialog.Root>
