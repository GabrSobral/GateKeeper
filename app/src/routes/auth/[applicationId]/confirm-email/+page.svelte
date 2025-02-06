<script lang="ts">
	import type { PageData } from '../one-time-password/$types';
	import AuthForm from './(components)/auth-form.svelte';

	let { data }: { data: PageData } = $props();
</script>

<div class="flex flex-col space-y-2 text-center">
	<h1 class="text-2xl font-semibold tracking-tight">Confirm E-mail</h1>
	<p class="text-muted-foreground text-sm">We sent a confirmation code to your e-mail. Enter the code below to confirm your e-mail</p>
</div>

<AuthForm {data} />
