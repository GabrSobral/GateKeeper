<script lang="ts">
	import type { PageData } from './$types';
	import AuthForm from './(components)/auth-form.svelte';
	

	let { data }: { data: PageData } = $props();
</script>

<div class="flex flex-col space-y-2 text-center">
	<h1 class="text-2xl font-semibold tracking-tight">Multi Factor</h1>
	<p class="text-muted-foreground text-sm">Enter your one time password below to sign in</p>
</div>

<AuthForm {data} />
