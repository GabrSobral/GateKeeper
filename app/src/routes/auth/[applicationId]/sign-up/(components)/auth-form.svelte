<script lang="ts">
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';

	import * as Form from '$lib/components/ui/form';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';

	import { cn } from '$lib/utils.js';

	import { formSchema, type FormSchema } from '../schema';
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

		setTimeout(() => {
			isLoading = false;
		}, 3000);
	}
</script>

<div class={cn('grid gap-4', className)}>
	<form on:submit|preventDefault={onSubmit}>
		<div class="grid gap-2">
			<Form.Field {form} name="firstName">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>First Name</Form.Label>
						<Input
							{...props}
							bind:value={$formData.firstName}
							placeholder="Type your first name"
							autocomplete="given-name"
						/>
					{/snippet}
				</Form.Control>
				<Form.Description />
				<Form.FieldErrors />
			</Form.Field>

			<Form.Field {form} name="lastName">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Last Name</Form.Label>
						<Input
							{...props}
							bind:value={$formData.lastName}
							placeholder="Type your last name"
							autocomplete="family-name"
						/>
					{/snippet}
				</Form.Control>
				<Form.Description />
				<Form.FieldErrors />
			</Form.Field>

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

			<a
				href={`/auth/${applicationId}/sign-in`}
				class="text-md text-center font-semibold hover:underline mb-2">Already has an account? Sign in</a
			>

			<Button type="submit" disabled={isLoading}>Sign up with Email</Button>
		</div>
	</form>
</div>
