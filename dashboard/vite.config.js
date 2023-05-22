import { sveltekit } from '@sveltejs/kit/vite';
import { vitePreprocess } from '@sveltejs/kit/vite';

/** @type {import('vite').UserConfig} */
const config = {
	plugins: [sveltekit()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	preprocess: [vitePreprocess()]
};

export default config;
