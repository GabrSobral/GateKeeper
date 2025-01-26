<script lang="ts">
	import ChevronLeft from "lucide-svelte/icons/chevron-left";
    
    import Breadcrumbs from "../../../(components)/breadcrumbs.svelte";
	import type { PageData } from "./$types";
	import EditApplicationForm from "./(components)/edit-application-form.svelte";
	import { page } from "$app/state";

    let { data }: { data: PageData } = $props();

	$inspect(data);
</script>

<Breadcrumbs
	items={[
		{ name: 'Dashboard', path: '/dashboard' },
		{ name: 'Applications', path: '/dashboard/application' },
		{ name: 'Edit Application', path: '/dashboard/application/1/edit-application' }
	]}
/>

<main class="flex flex-col p-4">
	<a href="/dashboard/application/" class="hover:underline flex items-center text-md gap-2 text-gray-600 hover:text-gray-800 mb-4">
		<ChevronLeft size={24} />
		Go back to applications details
	</a>

    <h2 class="text-3xl font-bold tracking-tight">
        Edit Application
    </h2>

	<span class="mt-3 text-sm tracking-tight text-gray-600">
        Here you can edit the application. Fill in the form below and click the "Apply Changes" button to edit the application.
	</span>

    <EditApplicationForm data={data} />
</main>