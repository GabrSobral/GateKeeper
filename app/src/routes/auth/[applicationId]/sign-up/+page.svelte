<script lang="ts">
	import type { PageData } from './$types';
	import AuthForm from './(components)/auth-form.svelte';

	let { data }: { data: PageData } = $props();
</script>

<svelte:head>
   <title>Sign Up - {data.applicationData?.name}</title>
</svelte:head>

<div class="flex flex-col space-y-2 text-center">
	<h1 class="text-2xl font-semibold tracking-tight">Create an account</h1>
	<p class="text-muted-foreground text-sm">Enter your data below to create your account</p>
</div>

<AuthForm {data} />
