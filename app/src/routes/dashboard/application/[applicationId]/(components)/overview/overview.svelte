<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import Badge from '$lib/components/ui/badge/badge.svelte';

	import NewSecretDialog from './new-secret-dialog.svelte';
	import DeleteSecretDialog from './delete-secret-dialog.svelte';
</script>

<section class="flex w-full flex-col gap-y-4">
	<Card.Root class="w-full transition-all">
		<Card.Header>
			<Card.Title>Info</Card.Title>
			<Card.Description>Information about the application.</Card.Description>
		</Card.Header>

		<Card.Content class="flex flex-wrap gap-x-8 gap-y-4">
			<div class="flex flex-col">
				<span class="text-md font-semibold">Application ID</span>
				<span class="text-sm">{window.crypto.randomUUID()}</span>
			</div>

			<div class="flex flex-col">
				<span class="text-md font-semibold">Status</span>
				<Badge variant="default">Active</Badge>
			</div>

			<div class="flex flex-col">
				<span class="text-md font-semibold">Multi Factor Auth</span>
				<Badge variant="outline" class="w-fit">Yes</Badge>
			</div>
		</Card.Content>
	</Card.Root>

	<Card.Root class="w-full transition-all">
		<Card.Header>
			<Card.Title class="flex flex-wrap justify-between gap-4">
				Secrets

				<NewSecretDialog />
			</Card.Title>
			<Card.Description>
				Secrets are used to authenticate your application with the server. Keep them safe.
			</Card.Description>
		</Card.Header>

		<Card.Content class="flex flex-col gap-y-4">
			<div class="flex items-center gap-4">
				<div class="space-y-1">
					<p class="text-sm font-medium leading-none">Secret Name</p>
					<p class="text-muted-foreground">ysd!da?4*********************</p>
				</div>

				<div class="ml-auto text-sm">
					Expiração: <span class="text-md font-medium">30/03/2032</span>
				</div>

				<Tooltip.Provider>
					<Tooltip.Root>
						<Tooltip.Trigger>
							<DeleteSecretDialog />
						</Tooltip.Trigger>

						<Tooltip.Content>
							<p>Delete secret</p>
						</Tooltip.Content>
					</Tooltip.Root>
				</Tooltip.Provider>
			</div>

			<div class="flex items-center gap-4">
				<div class="space-y-1">
					<p class="text-sm font-medium leading-none">Secret Name</p>
					<p class="text-muted-foreground">ysd!da?4*********************</p>
				</div>

				<div class="ml-auto text-sm">
					Expiração: <span class="text-md font-medium">30/03/2032</span>
				</div>

				<Tooltip.Provider>
					<Tooltip.Root>
						<Tooltip.Trigger>
							<DeleteSecretDialog />
						</Tooltip.Trigger>

						<Tooltip.Content>
							<p>Delete secret</p>
						</Tooltip.Content>
					</Tooltip.Root>
				</Tooltip.Provider>
			</div>
		</Card.Content>
	</Card.Root>
</section>
