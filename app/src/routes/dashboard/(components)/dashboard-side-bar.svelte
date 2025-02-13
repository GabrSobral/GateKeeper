<script lang="ts">
	import { page } from '$app/state';
	import { goto } from '$app/navigation';

	import Plus from 'lucide-svelte/icons/plus';
	import Logout from 'lucide-svelte/icons/log-out';
	import LayoutPanelLeft from 'lucide-svelte/icons/layout-panel-left';

	import * as Avatar from '$lib/components/ui/avatar';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';

	import { useApplicationsQuery } from '$lib/services/use-applications-query';
	import OrganizationsInput from './organizations-input.svelte';

	const items = [
		{
			title: 'Applications',
			url: '/dashboard/application',
			icon: LayoutPanelLeft
		}
	];

	const path = $derived(page.url.pathname);

	let applications = useApplicationsQuery({ accessToken: '' });
</script>

<Sidebar.Root variant="sidebar">
	<Sidebar.Header>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<div>
					<OrganizationsInput />
				</div>
			</Sidebar.MenuItem>
		</Sidebar.Menu>
	</Sidebar.Header>

	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupLabel>General</Sidebar.GroupLabel>

			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each items as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton isActive={item.url === path}>
								{#snippet child({ props })}
									<a href={item.url} {...props}>
										<item.icon />
										<span>{item.title}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>

		<Sidebar.Group>
			<Sidebar.GroupLabel>Applications</Sidebar.GroupLabel>
			<Sidebar.GroupAction
				title="Add Project"
				onclick={() => goto('/dashboard/application/create-application')}
			>
				<Plus /> <span class="sr-only">Add Project</span>
			</Sidebar.GroupAction>

			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each $applications?.data || [] as application (application?.id)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton>
								{#snippet child({ props })}
									<a href={`/dashboard/application/${application.id}`} {...props}>
										<span>{application.name}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
	</Sidebar.Content>

	<Sidebar.Footer>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<Sidebar.MenuButton>
					<Logout />
					Logout
				</Sidebar.MenuButton>
			</Sidebar.MenuItem>
		</Sidebar.Menu>

		<Sidebar.Separator />

		<div class="flex gap-2">
			<Avatar.Root>
				<Avatar.Image src="https://github.com/shadcn.png" alt="@shadcn" />
				<Avatar.Fallback>CN</Avatar.Fallback>
			</Avatar.Root>

			<div class="flex flex-col">
				<span class="text-sm font-semibold">John Doe</span>
				<span class="text-sm">Johndoe@email.com</span>
			</div>
		</div>
	</Sidebar.Footer>
</Sidebar.Root>
