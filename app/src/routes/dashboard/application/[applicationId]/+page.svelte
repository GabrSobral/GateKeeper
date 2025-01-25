<script lang="ts">
	import Trash from 'lucide-svelte/icons/trash-2';
	import Pencil from 'lucide-svelte/icons/pencil';
	import ChevronLeft from "lucide-svelte/icons/chevron-left";

	import * as Tabs from '$lib/components/ui/tabs';
	import { Badge } from '$lib/components/ui/badge';
	import * as Tooltip from '$lib/components/ui/tooltip';

	import Users from './(components)/users/users.svelte';
	import Roles from './(components)/roles/roles.svelte';
	import Overview from './(components)/overview/overview.svelte';
	import Breadcrumbs from '../../(components)/breadcrumbs.svelte';
	import Providers from './(components)/providers/providers.svelte';
	
	import { cn } from '$lib/utils';
	import { buttonVariants } from '$lib/components/ui/button';
</script>

<Breadcrumbs
	items={[
		{ name: 'Dashboard', path: '/dashboard' },
		{ name: 'Applications', path: '/dashboard/application' },
		{ name: 'ProxyMity', path: '/dashboard/application/<random-uuid-here>' }
	]}
/>

<main class="flex flex-col p-4">
	<a href="/dashboard/application" class="hover:underline flex items-center text-md gap-2 text-gray-600 hover:text-gray-800 mb-4">
		<ChevronLeft size={24} />
		Go back to applications list
	</a>

	<div class="flex items-center justify-between gap-4">

		<h2 class="text-3xl font-bold tracking-tight">ProxyMity</h2>

		<div class="flex gap-1">
			<Tooltip.Provider>
				<Tooltip.Root delayDuration={0}>
					<Tooltip.Trigger class={cn(buttonVariants({ variant: 'destructive' }))}>
						<Trash />
					</Tooltip.Trigger>

					<Tooltip.Content>
						<p>Delete Application</p>
					</Tooltip.Content>
				</Tooltip.Root>
			</Tooltip.Provider>

            <Tooltip.Provider>
				<Tooltip.Root delayDuration={0}>
					<Tooltip.Trigger class={cn(buttonVariants({ variant: 'outline' }))}>
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
		ProxyMity is a proxy server that allows you to access the internet securely.
	</span>

	<div class="mt-4 flex flex-wrap">
		<Badge variant="outline">Chat</Badge>
		<Badge variant="outline">C#</Badge>
		<Badge variant="outline">Javascript</Badge>
	</div>

	<Tabs.Root value="overview" class="mt-4 ">
		<Tabs.List>
			<Tabs.Trigger value="overview">Overview</Tabs.Trigger>
			<Tabs.Trigger value="users">Users</Tabs.Trigger>
			<Tabs.Trigger value="roles">Roles</Tabs.Trigger>
			<Tabs.Trigger value="providers">OAuth Providers</Tabs.Trigger>
		</Tabs.List>
        
		<Tabs.Content value="overview">
			<Overview />
		</Tabs.Content>

		<Tabs.Content value="users">
			<Users />
		</Tabs.Content>

		<Tabs.Content value="roles">
			<Roles />
		</Tabs.Content>

		<Tabs.Content value="providers">
			<Providers />
		</Tabs.Content>
	</Tabs.Root>
</main>
