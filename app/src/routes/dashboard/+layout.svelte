<script lang="ts">
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import DashboardSideBar from './(components)/dashboard-side-bar.svelte';

	let { children } = $props();
</script>

<div class="flex h-screen w-screen">
	<Sidebar.Provider>
		<DashboardSideBar/>

		<main>
			<Sidebar.Trigger />
			{@render children?.()}
		</main>
	</Sidebar.Provider>
</div>
