<script lang="ts">
	import type { PageData } from '../sign-in/$types';
	import AuthForm from './(components)/auth-form.svelte';
	
	let { data }: { data: PageData } = $props();
</script>

<div class="flex flex-col space-y-2 text-center">
	<h1 class="text-2xl font-semibold tracking-tight">Forgot your password?</h1>
	<p class="text-muted-foreground text-sm">Enter your e-mail below to continue</p>
</div>

<AuthForm {data} />
