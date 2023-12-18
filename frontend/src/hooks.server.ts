//TODO: set session cookie for logged in users and clear session cookie for logged out users. Also redirect to login page if not logged in.
import { redirect, type Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
    const user = event.cookies.get('credentials');
    if(!user && event.url.pathname === '/'){
        throw redirect(302, '/login');
    }
    else if(!user && event.url.pathname === '/register' || event.url.pathname === '/login') {
        const cookies = await resolve(event);
        cookies.headers.set('credentials', 'checked');
    }   
        const response = await resolve(event);
        response.headers.set('session', 'started');
        return response;
}