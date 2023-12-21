import { redirect, type Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	const value = event.cookies.get('session');
	if (!value) {
		throw redirect(302, '/login');
	}
	const value = event.cookies.get('session');
	if (!value && (event.url.pathname !== '/auth/login' && event.url.pathname !== '/auth/register')) {
		throw redirect(302, '/auth/login');
	}
	const response = await resolve(event);
	return response;
};

//https://svelte.dev/repl/171505499759483bb9f069c1e07cf54d?version=4.1.2




