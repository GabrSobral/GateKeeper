<script lang="ts">
	import Badge from "$lib/components/ui/badge/badge.svelte";
	import type { IApplication } from "$lib/services/use-application-by-id-query";

    type Props = {
        deactivatedAt: IApplication["users"]["data"][number]["deactivatedAt"];	
    }

    let { deactivatedAt }: Props = $props();

    $inspect(deactivatedAt);
</script>

{#if deactivatedAt}
    <Badge variant="secondary" color="red">Deactivated</Badge>
{:else}
    <Badge variant="default" color="green">Active</Badge>
{/if}
