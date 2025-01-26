<script lang="ts">
	import Badge from "$lib/components/ui/badge/badge.svelte";
	import type { IApplication } from "$lib/services/use-application-by-id-query";

    type Props = {
        roles: IApplication["users"]["data"][number]["roles"];	
    }

    let { roles }: Props = $props();
</script>

{#each roles as role (role.id)}
    <Badge variant="outline">{role.name}</Badge>
{/each}