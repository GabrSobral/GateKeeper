<script lang="ts">
	import { page } from '$app/state';
	import ArrowLeftIcon from 'lucide-svelte/icons/arrow-left';
	import type { PageData } from '../sign-in/$types';
	import AuthForm from './(components)/auth-form.svelte';
	
	let { data }: { data: PageData } = $props();
	let applicationId = page.params.applicationId;
</script>

<a href="/auth/{applicationId}/sign-in" class="text-primary flex items-center gap-3 hover:underline absolute top-4 left-4">
	<div class="flex items-center justify-center rounded-md bg-secondary text-primary p-2 hover:brightness-90 transition-all">
		<ArrowLeftIcon />
	</div>
	Back to sign in
</a>

<div class="flex flex-col space-y-2 text-center">
	<h1 class="text-2xl font-semibold tracking-tight">Forgot your password?</h1>
	<p class="text-muted-foreground text-sm">Enter your e-mail below to continue</p>
</div>

<AuthForm {data} />
