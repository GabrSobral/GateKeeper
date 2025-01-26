<script lang="ts">
	import Pencil from 'lucide-svelte/icons/pencil';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import ChevronLeft from 'lucide-svelte/icons/chevron-left';
	import Copy from 'lucide-svelte/icons/copy';
	import { superForm, type Infer, type SuperValidated } from 'sveltekit-superforms';

	import { cn, copy } from '$lib/utils';
	import * as Form from '$lib/components/ui/form';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { buttonVariants } from '$lib/components/ui/button';
	import Switch from '$lib/components/ui/switch/switch.svelte';
	import Separator from '$lib/components/ui/separator/separator.svelte';

	import Breadcrumbs from '../../../../(components)/breadcrumbs.svelte';
	import DeleteUserDialog from './(components)/delete-user-dialog.svelte';
	import MultiFactorAuthSection from './(components)/multi-factor-auth-section.svelte';
	import ApplicationRolesSection from './(components)/application-roles-section.svelte';

	import { formSchema, type FormSchema } from './schema';
	import Button from '$lib/components/ui/button/button.svelte';
	import Badge from '$lib/components/ui/badge/badge.svelte';
	import ResetPasswordDialog from './(components)/reset-password-dialog.svelte';
	import { page } from '$app/state';

	type Props = { data: { form: SuperValidated<Infer<FormSchema>> } };
	let { data }: Props = $props();

	let isLoading = $state(false);
	let isEditEnabled = $state(false);

	let applicationId = $derived(page.params.applicationId);

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

	{#if isEditEnabled}
		<Badge class="mb-4 w-fit text-sm" title="Edit is enabled">Editing</Badge>
	{/if}

	<div class="flex items-center justify-between gap-4">
		{#if isEditEnabled}
			<Form.Field {form} name="displayName" class="w-full max-w-[700px]">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>
							Display Name
							<span class="text-red-500">*</span>
						</Form.Label>

						<div class="flex gap-2">
							<Input
								{...props}
								type="text"
								style="font-size:1.75rem; font-weight:700; height:3.5rem; line-height:3.5rem;"
								bind:value={$formData.displayName}
								placeholder="Type the user display name"
							/>
							{#if !isEditEnabled}
								<Tooltip.Provider>
									<Tooltip.Root delayDuration={0}>
										<Tooltip.Trigger
											class={cn(
												buttonVariants({ variant: 'outline' }),
												'min-h-[3.5rem] min-w-[3.5rem]'
											)}
											onclick={() => copy($formData?.displayName || '')}
										>
											<Copy />
										</Tooltip.Trigger>

										<Tooltip.Content>Copy display name</Tooltip.Content>
									</Tooltip.Root>
								</Tooltip.Provider>
							{/if}
						</div>
					{/snippet}
				</Form.Control>
				<Form.Description>The name that will be displayed to the user.</Form.Description>
				<Form.FieldErrors></Form.FieldErrors>
			</Form.Field>
		{:else}
			<div class="flex gap-4">
				<h2 class="text-3xl font-bold tracking-tight">Ken O'Conner</h2>

				<Tooltip.Provider>
					<Tooltip.Root delayDuration={0}>
						<Tooltip.Trigger
							class={buttonVariants({ variant: 'outline' })}
							onclick={() => copy($formData?.email || '')}
						>
							<Copy />
						</Tooltip.Trigger>

						<Tooltip.Content>Copy display name</Tooltip.Content>
					</Tooltip.Root>
				</Tooltip.Provider>
			</div>
		{/if}

		<div class="flex gap-1">
			<Tooltip.Provider>
				<Tooltip.Root delayDuration={0}>
					<Tooltip.Trigger>
						<DeleteUserDialog />
					</Tooltip.Trigger>

					<Tooltip.Content>Delete User</Tooltip.Content>
				</Tooltip.Root>
			</Tooltip.Provider>

			<Tooltip.Provider>
				<Tooltip.Root delayDuration={0}>
					<Tooltip.Trigger
						class={cn(buttonVariants({ variant: 'outline' }))}
						onclick={() => (isEditEnabled = !isEditEnabled)}
					>
						<Pencil />
					</Tooltip.Trigger>

					<Tooltip.Content>Enable Changes</Tooltip.Content>
				</Tooltip.Root>
			</Tooltip.Provider>
		</div>
	</div>

	<div class="mt-4 flex flex-col gap-1">
		<Label class="text-foreground text-sm font-semibold" for="user-status-switch">Status</Label>

		<div class="flex items-center gap-2">
			<Switch
				checked={true}
				aria-labelledby="status-label"
				id="user-status-switch"
				disabled={!isEditEnabled}
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

							<div class="flex gap-2">
								<Input
									{...props}
									type="text"
									data-isEditEnabled={isEditEnabled}
									readonly={!isEditEnabled}
									class="data-[isEditEnabled=true]:outline-none"
									bind:value={$formData.firstName}
									placeholder="Type the user first name"
									autocomplete="given-name"
								/>

								{#if !isEditEnabled}
									<Tooltip.Provider>
										<Tooltip.Root delayDuration={0}>
											<Tooltip.Trigger
												class={buttonVariants({ variant: 'outline' })}
												onclick={() => copy($formData?.firstName || '')}
											>
												<Copy />
											</Tooltip.Trigger>

											<Tooltip.Content>Copy first name</Tooltip.Content>
										</Tooltip.Root>
									</Tooltip.Provider>
								{/if}
							</div>
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

							<div class="flex gap-2">
								<Input
									{...props}
									type="text"
									data-isEditEnabled={isEditEnabled}
									readonly={!isEditEnabled}
									class="data-[isEditEnabled=true]:outline-none"
									bind:value={$formData.lastName}
									placeholder="Type the last name"
									autocomplete="family-name"
								/>

								{#if !isEditEnabled}
									<Tooltip.Provider>
										<Tooltip.Root delayDuration={0}>
											<Tooltip.Trigger
												class={buttonVariants({ variant: 'outline' })}
												onclick={() => copy($formData?.lastName || '')}
											>
												<Copy />
											</Tooltip.Trigger>

											<Tooltip.Content>Copy last name</Tooltip.Content>
										</Tooltip.Root>
									</Tooltip.Provider>
								{/if}
							</div>
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

						<div class="flex gap-2">
							<Input
								{...props}
								type="email"
								data-isEditEnabled={isEditEnabled}
								readonly={!isEditEnabled}
								class="data-[isEditEnabled=true]:outline-none"
								bind:value={$formData.email}
								placeholder="Type the user e-mail"
								autocomplete="email"
							/>

							{#if !isEditEnabled}
								<Tooltip.Provider>
									<Tooltip.Root delayDuration={0}>
										<Tooltip.Trigger
											class={buttonVariants({ variant: 'outline' })}
											onclick={() => copy($formData?.email || '')}
										>
											<Copy />
										</Tooltip.Trigger>

										<Tooltip.Content>Copy e-mail</Tooltip.Content>
									</Tooltip.Root>
								</Tooltip.Provider>
							{/if}
						</div>
					{/snippet}
				</Form.Control>
				<Form.Description></Form.Description>
				<Form.FieldErrors></Form.FieldErrors>
			</Form.Field>

			<div class="flex flex-col gap-1">
				<span class="text-sm font-medium"> Reset User Password</span>
			
				<span class="text-muted-foreground my-2 text-sm">
					Reset the user password. On click, the user will receive an e-mail with the new password, and the user will be required to change it on the next login.
				</span>
				
				<ResetPasswordDialog />
			</div>

			<Separator class="my-2" />

			<MultiFactorAuthSection {isEditEnabled} />

			<Separator class="my-2" />

			<ApplicationRolesSection {isEditEnabled} />
		</div>

		<Button type="submit" class="float-right mt-4" disabled={isLoading}>Apply Changes</Button>
	</form>
</main>
