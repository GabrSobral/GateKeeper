<script lang="ts">
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';

	import * as Form from '$lib/components/ui/form';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';

	import { cn } from '$lib/utils.js';
	import { formSchema, type FormSchema } from '../schema';
	import { goto } from '$app/navigation';

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

	async function onSubmit() {
		isLoading = true;

		goto("/auth/one-time-password");

		setTimeout(() => {
			isLoading = false;
		}, 3000);
	}
</script>

<div class={cn('grid gap-6', className)}>
	<form on:submit|preventDefault={onSubmit}>
		<div class="grid gap-2">
			<Form.Field {form} name="password">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>New Password</Form.Label>
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

            <Form.Field {form} name="confirmPassword">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Confirm Your Password</Form.Label>
						<Input
							{...props}
							bind:value={$formData.confirmPassword}
							placeholder="********"
							autocomplete="new-password"
						/>
					{/snippet}
				</Form.Control>
				<Form.Description />
				<Form.FieldErrors />
			</Form.Field>

			<Button type="submit" disabled={isLoading}>Save New Password</Button>
		</div>
	</form>
</div>
