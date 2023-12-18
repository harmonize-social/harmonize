//TODO: set session cookie for logged in users and clear session cookie for logged out users. Also redirect to login page if not logged in.
import { redirect, type Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
    const user = await event.cookies.get('credentials')
    if(!user) {
        redirect(302, '/login');
    }
    const response = await resolve(event);
    response.headers.set('session', 'true');
    return response;
}