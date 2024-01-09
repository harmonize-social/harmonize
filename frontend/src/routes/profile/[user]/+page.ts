import type { PageLoad } from '../user/$types';

export const load: PageLoad = async ({ params }) => {
    const user = params.user;
    return{
        user,
    }
}
