<script lang="ts">
	import { goto } from '$app/navigation';

	import * as Tabs from '$lib/components/ui/tabs';
	import * as Card from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';

	import Breadcrumbs from '../(components)/breadcrumbs.svelte';
	import Button, { buttonVariants } from '$lib/components/ui/button/button.svelte';
	import { cn } from '$lib/utils';
</script>

<Breadcrumbs
	items={[
		{ name: 'Dashboard', path: '/dashboard' },
		{ name: 'Applications', path: '/dashboard/application' }
	]}
/>

<main class="flex flex-col p-4">
	<h2 class="text-3xl font-bold tracking-tight">Applications</h2>

	<span class="mt-3 text-sm tracking-tight text-gray-600">
		Applications are the projects you have created. You can add, edit, and delete them here.
	</span>

	<Tabs.Root value="overview" class="mt-4">
		<Tabs.List>
			<Tabs.Trigger value="overview">Overview</Tabs.Trigger>
		</Tabs.List>

		<a href="/dashboard/application/create-application" class={cn("float-right", buttonVariants({ variant: "default" }))}>
			New Application
		</a>

		<Tabs.Content value="overview" class="flex flex-1 flex-wrap gap-3">
			<Card.Root
				onclick={() => goto('/dashboard/application/<random-uuid-here>')}
				class="w-[calc(33.333%-8px)] min-w-[400px] transition-all hover:scale-105 hover:cursor-pointer"
			>
				<Card.Header>
					<Card.Title>ProxyMity</Card.Title>
					<Card.Description class="line-clamp-4">
						ProxyMity is a proxy server that allows you to access the internet securely.
					</Card.Description>
				</Card.Header>

				<Card.Footer class="mt-3">
					<Badge variant="outline">Chat</Badge>
					<Badge variant="outline">C#</Badge>
					<Badge variant="outline">Javascript</Badge>
				</Card.Footer>
			</Card.Root>
		</Tabs.Content>
	</Tabs.Root>
</main>
