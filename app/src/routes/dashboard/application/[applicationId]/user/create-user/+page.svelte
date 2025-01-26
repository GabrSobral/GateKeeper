<script lang="ts">
	import { zodClient } from 'sveltekit-superforms/adapters';
	import ChevronLeft from 'lucide-svelte/icons/chevron-left';
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';

	import * as Form from '$lib/components/ui/form';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import Switch from '$lib/components/ui/switch/switch.svelte';
	import Separator from '$lib/components/ui/separator/separator.svelte';

	import Breadcrumbs from '../../../../(components)/breadcrumbs.svelte';
	import MultiFactorAuthSection from './(components)/multi-factor-auth-section.svelte';
	import ApplicationRolesSection from './(components)/application-roles-section.svelte';

	import { formSchema, type FormSchema } from './schema';
	import { page } from '$app/state';

	type Props = { data: { form: SuperValidated<Infer<FormSchema>> } };
	let { data }: Props = $props();

	let applicationId = $derived(page.params.applicationId);

	let isLoading = $state(false);

	const form = superForm(data.form, { validators: zodClient(formSchema) });
	const { form: formData } = form;

	function onSubmit() {
		isLoading = true;

		// Logic here
		isLoading = false;
	}
</script>

<Breadcrumbs
	items={[
		{ name: 'Dashboard', path: '/dashboard' },
		{ name: 'Applications', path: '/dashboard/application' },
		{ name: applicationId, path: `/dashboard/application/${applicationId}?tab=users` },
		{ name: 'Create User', path: `/dashboard/application/${applicationId}/user/create-user` }
	]}
/>

<main class="flex flex-col p-4">
	<a
		href={`/dashboard/application/${applicationId}`}
		class="text-md mb-4 flex items-center gap-2 text-gray-600 hover:text-gray-800 hover:underline"
	>
		<ChevronLeft size={24} />
		Go back to application detail
	</a>

	<div class="flex items-center justify-between gap-4">
		<Form.Field {form} name="displayName" class="w-full max-w-[700px]">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>
						Display Name
						<span class="text-red-500">*</span>
					</Form.Label>

					<Input
						{...props}
						type="text"
						style="font-size:1.75rem; font-weight:700; height:3.5rem; line-height:3.5rem;"
						bind:value={$formData.displayName}
						placeholder="Type the user display name"
					/>
				{/snippet}
			</Form.Control>
			<Form.Description>The name that will be displayed to the user.</Form.Description>
			<Form.FieldErrors></Form.FieldErrors>
		</Form.Field>
	</div>

	<div class="mt-4 flex flex-col gap-1">
		<Label class="text-foreground text-sm font-semibold" for="user-status-switch">Status</Label>

		<div class="flex items-center gap-2">
			<Switch
				checked={true}
				aria-labelledby="status-label"
				id="user-status-switch"
			/>
			<span class="text-muted-foreground text-xs">Enabled</span>
		</div>
	</div>

	<form onsubmit={onSubmit} class="mt-4 max-w-[700px]">
		<div class="grid gap-2">
			<fieldset class="flex gap-2">
				<Form.Field {form} name="firstName" class="flex-1">
					<Form.Control>
						{#snippet children({ props })}
							<Form.Label>
								First Name
								<span class="text-red-500">*</span>
							</Form.Label>

							<Input
								{...props}
								type="text"
								bind:value={$formData.firstName}
								placeholder="Type the user first name"
								autocomplete="given-name"
							/>
						{/snippet}
					</Form.Control>
					<Form.FieldErrors></Form.FieldErrors>
				</Form.Field>

				<Form.Field {form} name="lastName" class="flex-1">
					<Form.Control>
						{#snippet children({ props })}
							<Form.Label>
								Last Name
								<span class="text-red-500">*</span>
							</Form.Label>

							<Input
								{...props}
								type="text"
								bind:value={$formData.lastName}
								placeholder="Type the last name"
								autocomplete="family-name"
							/>
						{/snippet}
					</Form.Control>
					<Form.FieldErrors></Form.FieldErrors>
				</Form.Field>
			</fieldset>

			<Separator class="my-2" />

			<Form.Field {form} name="email">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>
							E-mail
							<span class="text-red-500">*</span>
						</Form.Label>

						<Input
							{...props}
							type="email"
							bind:value={$formData.email}
							placeholder="Type the user e-mail"
							autocomplete="email"
						/>
					{/snippet}
				</Form.Control>
				<Form.Description></Form.Description>
				<Form.FieldErrors></Form.FieldErrors>
			</Form.Field>

			<Separator class="my-2" />

			<MultiFactorAuthSection />

			<Separator class="my-2" />

			<ApplicationRolesSection />
		</div>

		<Button type="submit" class="float-right mt-4" disabled={isLoading}>Create User</Button>
	</form>
</main>
