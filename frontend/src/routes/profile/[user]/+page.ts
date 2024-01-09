import { get } from '../../../fetch';
import type PostModel from '../../../models/post';
import type { PageLoad } from '../user/$types';



export async function _fetchUserPosts(user: string): Promise<PostModel[]> {
    try{
        const response: PostModel[] = await get<PostModel[]>(`/posts?username=${user}`);
        return response;
    }catch(error){
        console.log(error);
        return [];
    }
}

export const load: PageLoad = async ({ params }) => {
    console.log(params);
    const username = params.user;
    return{ 
        user: {
            username,
            posts: await _fetchUserPosts(params.user)
        }
    }
}