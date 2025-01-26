<script lang="ts">
	import type { IApplication } from '$lib/services/use-application-by-id-query';
	import ProviderCard from './(components)/provider-card.svelte';

	type Props = { application?: IApplication | null };

	let { application }: Props = $props();
</script>

<div class="grid grid-cols-3 gap-3">
	{#each application?.oauthProviders || [] as provider (provider.id)}
		<ProviderCard
			title={provider.name}
			description={provider.description}
			isEnabled={provider.isEnabled}
			clientId={provider.clientId}
			clientSecret={provider.clientSecret}
		/>
	{/each}
</div>
