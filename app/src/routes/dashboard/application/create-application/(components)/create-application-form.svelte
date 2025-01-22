<script lang="ts">
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';

	import * as Form from '$lib/components/ui/form';
	import { Textarea } from '$lib/components/ui/textarea';
	import Input from '$lib/components/ui/input/input.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import MultiSelectInput from '$lib/components/ui/multi-select-input/multi-select-input.svelte';

	import { formSchema, type FormSchema } from '../schema';
	import StrongPasswordModal from './strong-password-modal.svelte';

	let {
		data
	}: {
		data: { form: SuperValidated<Infer<FormSchema>> };
	} = $props();

	let isLoading = $state(false);

	const form = superForm(data.form, { validators: zodClient(formSchema) });
	const { form: formData } = form;

	async function onSubmit(
		e: SubmitEvent & {
			currentTarget: EventTarget & HTMLFormElement;
		}
	) {
		e.preventDefault();
		isLoading = true;

		setTimeout(() => (isLoading = false), 3000);
	}

	let badges: string[] = $state([]);
	let passwordHashSecret = $state('');

	$effect(() => {
		$formData.passwordHashSecret = passwordHashSecret
	})
</script>

<form onsubmit={onSubmit} class="mt-4 max-w-[600px]">
	<div class="grid gap-2">
		<Form.Field {form} name="name">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>
						Name
						<span class="text-red-500">*</span>
					</Form.Label>
					<Input
						{...props}
						bind:value={$formData.name}
						placeholder="Type the application name"
						autocomplete="name"
					/>
				{/snippet}
			</Form.Control>
			<Form.Description></Form.Description>
			<Form.FieldErrors></Form.FieldErrors>
		</Form.Field>

		<Form.Field {form} name="Description">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Description</Form.Label>
					<Textarea
						{...props}
						bind:value={$formData.Description}
						placeholder="Type the application description"
					/>
				{/snippet}
			</Form.Control>
			<Form.Description></Form.Description>
			<Form.FieldErrors></Form.FieldErrors>
		</Form.Field>

		<Separator />

		<Form.Field {form} name="passwordHashSecret">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>
						Password Hash Secret
						<span class="text-red-500">*</span>
					</Form.Label>

					<Input
						{...props}
						bind:value={passwordHashSecret}
						placeholder="Type the password hash secret"
						type="text"
					/>
				{/snippet}
			</Form.Control>

			<StrongPasswordModal setPassword={value => (passwordHashSecret = value)} />

			<Form.Description>
				This is the secret that will be used to hash all the passwords from users that are
				registered to this application.
			</Form.Description>
			<Form.FieldErrors></Form.FieldErrors>
		</Form.Field>

		<Form.Field {form} name="badges">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Badges</Form.Label>
					<MultiSelectInput items={badges} />
				{/snippet}
			</Form.Control>

			<Form.Description>
				This is used to categorize the application. You can add multiple badges to the application.
			</Form.Description>
			<Form.FieldErrors></Form.FieldErrors>
		</Form.Field>

		<Button type="submit" disabled={isLoading} class="ml-auto w-fit">
			Create Application
		</Button>
	</div>
</form>
