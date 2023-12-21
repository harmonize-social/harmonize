import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import replace from '@rollup/plugin-replace';

export default defineConfig({
	plugins: [
		sveltekit(),
		replace({
			'process.env.API_URL': JSON.stringify(process.env.API_URL ||'http://localhost:8080'),
			preventAssignments: true
		})
	],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
