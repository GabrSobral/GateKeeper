<script lang="ts">
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';

	import * as Form from '$lib/components/ui/form';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';

	import { cn } from '$lib/utils.js';
	import { formSchema, type FormSchema } from '../schema';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import type { PageData } from '../$types';

	let {
		data,
		class: className
	}: {
		data: PageData;
		class?: string | null;
	} = $props();

	let isLoading = $state(false);
	let applicationId = page.params.applicationId;

	const form = superForm(data.form, { validators: zodClient(formSchema) });
	const { form: formData, enhance } = form;

	async function onSubmit() {
		isLoading = true;

		goto(`/auth/${applicationId}/one-time-password`);

		setTimeout(() => {
			isLoading = false;
		}, 3000);
	}
</script>

<div class={cn('grid gap-4', className)}>
	<form on:submit|preventDefault={onSubmit}>
		<div class="grid gap-2">
			<Form.Field {form} name="email">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>E-mail</Form.Label>
						<Input
							{...props}
							bind:value={$formData.email}
							placeholder="example@email.com"
							autocomplete="email"
						/>
					{/snippet}
				</Form.Control>
				<Form.Description />
				<Form.FieldErrors />
			</Form.Field>

			<Form.Field {form} name="password">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Password</Form.Label>
						<Input
							{...props}
							bind:value={$formData.password}
							placeholder="********"
							autocomplete="new-password"
						/>
					{/snippet}
				</Form.Control>
				<Form.Description />
				<Form.FieldErrors />
			</Form.Field>

			<div class="flex items-center justify-between gap-2">
				<a
					href={`/auth/${applicationId}/forgot-password`}
					class="text-md mb-2 text-center hover:underline">Forgot password</a
				>

				{#if data.applicationData?.canSelfRegister}
					<a href={`/auth/${applicationId}/sign-up`} class="font-semibold text-md text-center hover:underline">
						Create an account
					</a>
				{/if}
			</div>

			<Button type="submit" disabled={isLoading}>Sign In with Email</Button>
		</div>
	</form>

	<div class="relative">
		<div class="absolute inset-0 flex items-center">
			<span class="w-full border-t"></span>
		</div>
		<div class="relative flex justify-center text-xs uppercase">
			<span class="bg-background text-muted-foreground px-2"> Or continue with </span>
		</div>
	</div>

	<div class="flex flex-col gap-1">
		{#each data.applicationData?.oauthProviders || [] as provider}
			<Button variant="outline" type="button" disabled={isLoading}>
				{provider.name}
			</Button>
		{/each}
	</div>
</div>
