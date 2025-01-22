<script lang="ts">
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import DashboardSideBar from './(components)/dashboard-side-bar.svelte';

	let { children } = $props();
</script>

<div class="flex h-screen w-screen">
	<Sidebar.Provider>
		<DashboardSideBar/>

		<div class="flex-1">
			{@render children?.()}
		</div>
	</Sidebar.Provider>
</div>
