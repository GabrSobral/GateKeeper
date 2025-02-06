<script lang="ts">
	import * as InputOTP from "$lib/components/ui/input-otp";

	import { zodClient } from 'sveltekit-superforms/adapters';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';

	import Button from '$lib/components/ui/button/button.svelte';

	import { formSchema, type FormSchema } from '../schema';

	let { data }: { data: { form: SuperValidated<Infer<FormSchema>> }} = $props();

	let isLoading = $state(false);

	const form = superForm(data.form, { validators: zodClient(formSchema) });

	async function onSubmit() {
		isLoading = true;

		setTimeout(() => {
			isLoading = false;
		}, 3000);
	}
</script>

<div class='grid gap-4'>
	<form on:submit|preventDefault={onSubmit}>
		<div class="grid gap-6">
			<InputOTP.Root maxlength={6} class="mx-auto">
				{#snippet children({ cells })}
				  <InputOTP.Group>
					{#each cells.slice(0, 3) as cell}
					  <InputOTP.Slot {cell} />
					{/each}
				  </InputOTP.Group>
				  
				  <InputOTP.Separator />

				  <InputOTP.Group>
					{#each cells.slice(3, 6) as cell}
					  <InputOTP.Slot {cell} />
					{/each}
				  </InputOTP.Group>
				{/snippet}
			  </InputOTP.Root>

			<Button type="submit" disabled={isLoading}>Continue</Button>
		</div>
	</form>
</div>
