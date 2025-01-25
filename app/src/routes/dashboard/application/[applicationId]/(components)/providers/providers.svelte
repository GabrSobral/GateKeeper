<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import Badge from '$lib/components/ui/badge/badge.svelte';
	import ProviderCard from './(components)/provider-card.svelte';
</script>


<div class="grid grid-cols-3 gap-3">
	<ProviderCard 
		title="Google" 
		description="Google authentication allows users to sign in to your application using their Google account, providing a secure and convenient way to authenticate." 
		isEnabled={true} 
		clientId={window.crypto.randomUUID()} 
		clientSecret={window.crypto.getRandomValues(new Uint8Array(16)).join('')}
	/>
</div>
