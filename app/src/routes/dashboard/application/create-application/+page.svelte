<script lang="ts">
	import ChevronLeft from "lucide-svelte/icons/chevron-left";
    
    import Breadcrumbs from "../../(components)/breadcrumbs.svelte";
	import type { PageData } from "./$types";
	import CreateApplicationForm from "./(components)/create-application-form.svelte";

    let { data }: { data: PageData } = $props();
</script>

<Breadcrumbs
	items={[
		{ name: 'Dashboard', path: '/dashboard' },
		{ name: 'Applications', path: '/dashboard/application' },
		{ name: 'Create Application', path: '/dashboard/application/create-application' }
	]}
/>

<main class="flex flex-col p-4">
	<a href="/dashboard/application" class="hover:underline flex items-center text-md gap-2 text-gray-600 hover:text-gray-800 mb-4">
		<ChevronLeft size={24} />
		Go back to applications list
	</a>

    <h2 class="text-3xl font-bold tracking-tight">
        Create Application
    </h2>

	<span class="mt-3 text-sm tracking-tight text-gray-600">
        Here you can create a new application. Fill in the form below and click the "Create Application" button to create a new application.
	</span>

    <CreateApplicationForm data={data} />
</main>