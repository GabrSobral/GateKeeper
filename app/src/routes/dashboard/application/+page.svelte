<script lang="ts">
	import { goto } from '$app/navigation';
	import { useQuery } from '@sveltestack/svelte-query';

	import { cn } from '$lib/utils';
	import * as Tabs from '$lib/components/ui/tabs';
	import * as Card from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { buttonVariants } from '$lib/components/ui/button/button.svelte';
	import { getApplicationsService } from '$lib/services/use-applications-query';

	import Breadcrumbs from '../(components)/breadcrumbs.svelte';
	import { organizationStore } from '$lib/stores/organization';

	let applicationsResult = useQuery(
		['list-applications', $organizationStore?.id],
		() => getApplicationsService({ organizationId: $organizationStore?.id }, { accessToken: '' }),
		{ refetchOnWindowFocus: false }
	);

	$inspect($organizationStore?.id);
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

		<a
			href="/dashboard/application/create-application"
			class={cn('float-right', buttonVariants({ variant: 'default' }))}
		>
			New Application
		</a>

		<Tabs.Content value="overview" class="flex flex-1 flex-wrap gap-3">
			{#if $applicationsResult.isLoading}
				<div class="flex flex-1 items-center justify-center">Loading...</div>
			{:else}
				{#if $applicationsResult.error}
					<div class="flex flex-1 items-center justify-center">Failed to load applications</div>
				{/if}

				{#if $applicationsResult.data?.length === 0}
					<div class="flex flex-1 items-center justify-center">No applications found</div>
				{/if}

				{#each $applicationsResult.data || [] as application (application.id)}
					<Card.Root
						onclick={() =>
							goto(
								`/dashboard/application/${application.id}?organizationId=${$organizationStore?.id}`
							)}
						class="w-[calc(33.333%-8px)] min-w-[400px] transition-all hover:scale-[1.01] hover:cursor-pointer hover:shadow-lg"
					>
						<Card.Header>
							<Card.Title>{application.name}</Card.Title>
							<Card.Description class="line-clamp-4">{application.description}</Card.Description>
						</Card.Header>

						<Card.Footer class="mt-3">
							{#each application.badges as tag}
								<Badge variant="outline">{tag}</Badge>
							{/each}
						</Card.Footer>
					</Card.Root>
				{/each}
			{/if}
		</Tabs.Content>
	</Tabs.Root>
</main>
