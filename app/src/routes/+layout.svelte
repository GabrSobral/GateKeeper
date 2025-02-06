<script lang="ts">
	import { i18n } from '$lib/i18n';
	import { ParaglideJS } from '@inlang/paraglide-sveltekit';
	import '../app.css';
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';

	let { children } = $props();

	const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
	<ParaglideJS {i18n}>
		{@render children()}
	</ParaglideJS>
</QueryClientProvider>
