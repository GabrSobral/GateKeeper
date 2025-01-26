<script lang="ts">
	import { page } from '$app/state';
	import { goto } from '$app/navigation';

	import Pencil from 'lucide-svelte/icons/pencil';
	import ChevronLeft from "lucide-svelte/icons/chevron-left";

	import { cn } from '$lib/utils';
	import * as Tabs from '$lib/components/ui/tabs';
	import { Badge } from '$lib/components/ui/badge';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import { buttonVariants } from '$lib/components/ui/button';

	import Users from './(components)/users/users.svelte';
	import Roles from './(components)/roles/roles.svelte';
	import Overview from './(components)/overview/overview.svelte';
	import Providers from './(components)/providers/providers.svelte';
	import Breadcrumbs from '../../(components)/breadcrumbs.svelte';
	import DeleteApplicationDialog from './(components)/delete-application-dialog.svelte';
	import { useApplicationByIdQuery } from '$lib/services/use-application-by-id-query';

	let currentTab = $derived(page.url.searchParams.get('tab') || "overview") as "overview" | "users" | "roles" | "providers";
	let applicationId = $derived(page.params.applicationId);

	let application = $derived(useApplicationByIdQuery({ applicationId }, { accessToken: "" }));	
</script>

<Breadcrumbs
	items={[
		{ name: 'Dashboard', path: '/dashboard' },
		{ name: 'Applications', path: '/dashboard/application' },
		{ name: $application?.data?.name || "-", path: `/dashboard/application/${$application?.data?.id}` }
	]}
/>

<main class="flex flex-col p-4">
	<a href="/dashboard/application" class="hover:underline flex items-center text-md gap-2 text-gray-600 hover:text-gray-800 mb-4">
		<ChevronLeft size={24} />
		Go back to applications list
	</a>

	<div class="flex items-center justify-between gap-4">
		<h2 class="text-3xl font-bold tracking-tight">{$application?.data?.name}</h2>

		<div class="flex gap-1">
			<Tooltip.Provider>
				<Tooltip.Root delayDuration={0}>
					<Tooltip.Trigger>
						<DeleteApplicationDialog />
					</Tooltip.Trigger>

					<Tooltip.Content>
						<p>Delete Application</p>
					</Tooltip.Content>
				</Tooltip.Root>
			</Tooltip.Provider>

            <Tooltip.Provider>
				<Tooltip.Root delayDuration={0}>
					<Tooltip.Trigger 
						class={cn(buttonVariants({ variant: 'outline' }))} 
						onclick={() => goto(`/dashboard/application/${$application?.data?.id}/edit-application`)}
					>
						<Pencil />
					</Tooltip.Trigger>

					<Tooltip.Content>
						<p>Update Application</p>
					</Tooltip.Content>
				</Tooltip.Root>
			</Tooltip.Provider>
		</div>
	</div>

	<span class="mt-3 text-sm tracking-tight text-gray-600">
		{$application?.data?.description}
	</span>

	<div class="mt-4 flex flex-wrap">
		{#each $application?.data?.badges || [] as badge (badge)}
			<Badge variant="outline">{badge}</Badge>
			
		{/each}
	</div>

	<Tabs.Root value={currentTab} class="mt-4">
		<Tabs.List>
			<Tabs.Trigger value="overview" onclick={() => goto(`/dashboard/application/${$application?.data?.id}?tab=overview`)}>Overview</Tabs.Trigger>
			<Tabs.Trigger value="users" onclick={() => goto(`/dashboard/application/${$application?.data?.id}?tab=users`)}>Users</Tabs.Trigger>
			<Tabs.Trigger value="roles" onclick={() => goto(`/dashboard/application/${$application?.data?.id}?tab=roles`)}>Roles</Tabs.Trigger>
			<Tabs.Trigger value="providers" onclick={() => goto(`/dashboard/application/${$application?.data?.id}?tab=providers`)}>OAuth Providers</Tabs.Trigger>
		</Tabs.List>
        
		<Tabs.Content value="overview">
			<Overview application={$application.data}/>
		</Tabs.Content>

		<Tabs.Content value="users">
			<Users application={$application.data} />
		</Tabs.Content>

		<Tabs.Content value="roles">
			<Roles application={$application.data} />
		</Tabs.Content>

		<Tabs.Content value="providers">
			<Providers application={$application.data} />
		</Tabs.Content>
	</Tabs.Root>
</main>
