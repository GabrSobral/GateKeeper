<script lang="ts">
	import Settings from 'lucide-svelte/icons/settings';
	import House from 'lucide-svelte/icons/house';
	import Plus from 'lucide-svelte/icons/plus';
	import ChevronDown from 'lucide-svelte/icons/chevron-down';
	import Logout from 'lucide-svelte/icons/log-out';

	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import * as Avatar from '$lib/components/ui/avatar';

	const items = [
		{
			title: 'Home',
			url: '#',
			icon: House
		},
		{
			title: 'Settings',
			url: '#',
			icon: Settings
		}
	];
</script>

<Sidebar.Root variant="sidebar">
	<Sidebar.Header>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<DropdownMenu.Root>
					<DropdownMenu.Trigger>
						{#snippet child({ props })}
							<Sidebar.MenuButton {...props}>
								Select Organization
								<ChevronDown class="ml-auto" />
							</Sidebar.MenuButton>
						{/snippet}
					</DropdownMenu.Trigger>

					<DropdownMenu.Content class="w-[--bits-dropdown-menu-anchor-width]">
						<DropdownMenu.Item>
							<span>Acme Inc</span>
						</DropdownMenu.Item>
						<DropdownMenu.Item>
							<span>Acme Corp.</span>
						</DropdownMenu.Item>
					</DropdownMenu.Content>
				</DropdownMenu.Root>
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
							<Sidebar.MenuButton>
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
			<Sidebar.GroupAction title="Add Project">
				<Plus /> <span class="sr-only">Add Project</span>
			</Sidebar.GroupAction>

			<Sidebar.GroupContent>
				<Sidebar.Menu>
					<Sidebar.MenuItem>
						<Sidebar.MenuButton>
							{#snippet child({ props })}
								<a href="#" {...props}>
									<span>ProxyMity</span>
								</a>
							{/snippet}
						</Sidebar.MenuButton>
					</Sidebar.MenuItem>

					<Sidebar.MenuItem>
						<Sidebar.MenuButton>
							{#snippet child({ props })}
								<a href="#" {...props}>
									<span>GateKeeper</span>
								</a>
							{/snippet}
						</Sidebar.MenuButton>
					</Sidebar.MenuItem>
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
