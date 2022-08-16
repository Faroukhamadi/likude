import { writable } from 'svelte/store';

export const usernameForProfile = writable<string>('');
