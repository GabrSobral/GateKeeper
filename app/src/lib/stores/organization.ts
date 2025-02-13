import { writable } from "svelte/store";

interface IOrganization {
    id: string;
    name: string;
    createdAt: Date;
}

export const organizationStore = writable<IOrganization | null>(null);