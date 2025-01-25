<script lang="ts">
	import * as Sheet from '$lib/components/ui/sheet/';
	import * as Card from '$lib/components/ui/card';
    import { Switch } from "$lib/components/ui/switch";

	import Badge from '$lib/components/ui/badge/badge.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import { buttonVariants } from '$lib/components/ui/button';

    type Props = {
        title: string;
        description: string;
        isEnabled: boolean;
        clientId: string;
        clientSecret: string;
    }

    let { title, description, clientId, clientSecret, isEnabled }: Props = $props()

	let draftTitle = $derived(title);
	let draftDescription = $derived(description);
	let draftClientId = $derived(clientId);
	let draftClientSecret = $derived(clientSecret);
	let draftIsEnabled = $derived(isEnabled);

</script>

<Sheet.Root>
	<Sheet.Trigger>
		<Card.Root class="transition-all hover:scale-[1.01] hover:cursor-pointer hover:shadow-lg">
			<Card.Header>
				<Card.Title class="flex flex-wrap justify-between gap-4">{draftTitle}</Card.Title>

				<Card.Description>
					{draftDescription}
				</Card.Description>
			</Card.Header>

			<Card.Content class="flex gap-1">
				<Badge variant="outline" class="w-fit">{(draftClientId && draftClientSecret) ? "Configured" : "Not configured"}</Badge>
				<Badge variant={draftIsEnabled ? "default" : "secondary"} class="w-fit">{draftIsEnabled ? "Enabled" : "Disabled"}</Badge>
			</Card.Content>
		</Card.Root>
	</Sheet.Trigger>

	<Sheet.Content side="right">
		<Sheet.Header>
			<Sheet.Title>Configure Provider</Sheet.Title>
			<Sheet.Description>
				Make changes to your Google authentication provider configuration. Then click "Save changes" to apply them.
			</Sheet.Description>
		</Sheet.Header>
        
		<div class="flex flex-col gap-4 my-4">
			<div class="flex flex-col gap-2">
				<Label for="client-id">Client ID</Label>
				<Input id="client-id" value={draftClientId} type="text" placeholder="Type the client ID" class="col-span-3" />
			</div>

			<div class="flex flex-col gap-2">
				<Label for="client-secret">Client Secret</Label>
				<Input id="client-secret" type="password" placeholder="Type the client secret" value={draftClientSecret} class="col-span-3" />
			</div>

            <div class="flex items-center space-x-2">
                <Switch id="provider-enabled" checked={draftIsEnabled} />
                <Label for="provider-enabled">{draftIsEnabled ? "Enabled" : "Disabled"}</Label>
              </div>
		</div>

		<Sheet.Footer>
			<Sheet.Close class={buttonVariants({ variant: 'default' })}>
                Save changes
            </Sheet.Close>
		</Sheet.Footer>
	</Sheet.Content>
</Sheet.Root>
