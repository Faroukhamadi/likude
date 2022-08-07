/// <reference types="@sveltejs/kit" />

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
declare namespace App {
	interface Locals {
		user: {
			id: number;
			username: string;
		};
	}
	interface Platform {}
	interface Session {
		user: {
			id: number;
			username: string;
		};
	}
	interface Stuff {}
}
