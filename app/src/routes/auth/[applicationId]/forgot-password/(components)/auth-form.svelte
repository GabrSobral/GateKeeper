<script lang="ts">
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';

	import * as Form from '$lib/components/ui/form';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';

	import { cn } from '$lib/utils.js';
	import { formSchema, type FormSchema } from '../schema';

	let {
		data,
		class: className
	}: {
		data: { form: SuperValidated<Infer<FormSchema>> };
		class?: string | null;
	} = $props();

	let isLoading = $state(false);

	const form = superForm(data.form, { validators: zodClient(formSchema) });
	const { form: formData, enhance } = form;

	let isSent = $state(false);

	async function onSubmit() {
		isLoading = true;

		setTimeout(() => {
			isLoading = false;
			isSent = true;
		}, 1000);
	}
</script>

{#if isSent}
	<div class={cn('grid gap-6', className)}>
		<p class="text-center text-md bg-green-100 p-4 rounded-lg">
			Check your email for a link to reset your password. If it doesn't appear within a few minutes,
			check your spam folder.
		</p>
	</div>

	<Button onclick={() => (isSent = false)} class="mt-4">Go Back</Button>
{:else}
	<div class={cn('grid gap-6', className)}>
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

				<Button type="submit" disabled={isLoading}>Send Mail</Button>
			</div>
		</form>
	</div>
{/if}
