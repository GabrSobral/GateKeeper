<script lang="ts">
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';

	import { Toaster } from '$lib/components/ui/sonner';
	import * as Sidebar from '$lib/components/ui/sidebar';
	import DashboardSideBar from './(components)/dashboard-side-bar.svelte';

	let { children } = $props();

	const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
	<div class="flex h-screen w-screen">
		<Sidebar.Provider>
			<DashboardSideBar />

			<div class="flex-1">
				<Toaster />
				{@render children?.()}
			</div>
		</Sidebar.Provider>
	</div>
</QueryClientProvider>
