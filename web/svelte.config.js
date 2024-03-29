import adapter from '@sveltejs/adapter-auto';
import preprocess from 'svelte-preprocess';
import houdini from 'houdini/preprocess';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: [
		preprocess({
			postcss: true,
		}), houdini()
	],

	kit: {
		adapter: adapter(),
		alias: {
			$houdini: './$houdini',
		}
	}
};

export default config;
