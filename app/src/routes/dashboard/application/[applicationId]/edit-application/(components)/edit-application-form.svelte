<script lang="ts">
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	
	import * as Form from '$lib/components/ui/form';
	import { Checkbox } from "$lib/components/ui/checkbox";
	import { Textarea } from '$lib/components/ui/textarea';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import MultiSelectInput from '$lib/components/ui/multi-select-input/multi-select-input.svelte';

	import { formSchema, type FormSchema } from '../schema';
	import StrongPasswordModal from './strong-password-modal.svelte';
	import { toast } from 'svelte-sonner';

	type Props = { data: { form: SuperValidated<Infer<FormSchema>> } } 
	let { data }: Props = $props();

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

		setTimeout(() => {
			isLoading = false
			toast.success('Application created successfully!')
		}, 3000);
	}

	let badges: string[] = $state([]);
	let passwordHashSecret = $state('');

	$effect(() => {
		$formData.passwordHashSecret = passwordHashSecret;
	});
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

			<StrongPasswordModal setPassword={(value) => (passwordHashSecret = value)} />

			<Form.Description>
				This is the secret that will be used to hash all the passwords from users that are
				registered to this application.
			</Form.Description>
			<Form.FieldErrors></Form.FieldErrors>
		</Form.Field>

		<Separator />

		<div class="flex flex-col gap-3">
			<span class="font-medium text-sm">
				Multi Factor Authentication
			</span>

			<span class="text-sm text-muted-foreground">
				Choose the methods that will be used for multi factor authentication.
			</span>

			<div class="flex items-center space-x-2">
				<Checkbox id="e-mail-mfa" checked={false} aria-labelledby="terms-label" />
	
				<Label
				  id="e-mail-mfa-label"
				  for="e-mail-mfa"
				  class="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
				>
				  E-mail
				</Label>
			</div>
	
			<div class="flex items-center space-x-2">
				<Checkbox id="auth-app-mfa" checked={false} aria-labelledby="terms-label" />
	
				<Label
				  id="auth-app-mfa-label"
				  for="auth-app-mfa"
				  class="text-sm leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
				>
				  Authenticator App (Microsoft, Google, etc)
				</Label>
			</div>
		</div>

		<Button type="submit" disabled={isLoading} class="ml-auto w-fit">
			{#if isLoading}
				Creating Application...
			{:else}
				Create Application
			{/if}
		</Button>
	</div>
</form>
