<script lang="ts">
import ErrorPopup from '../../../components/ErrorPopup.svelte';
import Button from '../../../components/Button.svelte';
import Panel from '../../../components/Panel.svelte';
import NavBar from '../../../components/NavBar.svelte';
import {
    delete_,
    get,
    post,
    throwError
} from '../../../fetch';
import type PostModel from '../../../models/post';
import Post from '../../../components/Post.svelte';
import {
    errorMessage
} from '../../../store';
import {
    _fetchUserPosts
} from './+page';
import type {
    PageData
} from '../user/$types';
import {
    onMount
} from 'svelte';

let follows: boolean;
let loading = false;
let error = '';
let text = '';
export let data: PageData;
export let posts: PostModel[] = data.user.posts;
export let name: string = data.user.username;

errorMessage.subscribe((value) => {
    error = value;
});

async function isFollowing(username: string){
    try {
        const response = await get < any > (`/me/following?username=${username}`);
        if(response.includes(username)){
            follows = true;
            text = 'Unfollow';
        } else {
            follows = false;
            text = 'Follow';
        }

    } catch (e) {
        throwError('Error fetching user posts');
        follows = false;
    }
    return follows;
}
async function deleteFollow(username: string) {
    try {
        const response: string = await delete_ < string > (`/me/follow?username=${username}`);
        return response;
    } catch (e) {
        throwError('Error deleting follow');
        follows = true;
    }
}
async function postFollow(username: string) {
    try {
        const response: string = await post < string, string > (`/me/follow?username=${username}`, username);
        return response;
    } catch (e) {
        throwError('Error posting follow');
        follows = false;
    }
}

// function onScroll(event: Event) {
//     const target = event.target as HTMLElement;
//     if (target.scrollHeight - target.scrollTop === target.clientHeight) {
//         loadMorePosts(name);
//     }
// }
// async function loadMorePosts(name: string) {
//     if (loading) return;
//     const morePosts = await _fetchUserData(name);
//     posts = [...posts, ...morePosts];
// }

let isClicked = false;
const handleButtonClick = async () => {
    if (follows) {
        await deleteFollow(name);
        text = 'Follow';
        follows = false;
    } else {
        await postFollow(name);
        text = 'Unfollow';
        follows = true;
    }
    isClicked = !isClicked;

};
onMount(async () => {
    await isFollowing(name);
});
</script>

<!-- navbar -->
<div class="nav">
    <NavBar current_page={`/profile/${name}`}></NavBar>
</div>
<!-- profile -->
<div class="profile-container">
    <h2 class="username">{name}</h2>

    <!-- personal feed -->
    <div class="feed-container">
        <Panel title={`${name}'s posts`}>
            <div class="feed">
                {#if posts.length === 0}
                <p>{name} did not post yet!</p>
                {/if}
                {#each posts as post}
                <div class="post">
                    <Post
                        content={post.content}
                        caption={post.caption}
                        likes={post.likeCount}
                        id={post.id}
                        typez={post.type}
                        isLiked={post.hasLiked}
                        isSaved={post.hasSaved}
                        />
                        </div>
                        {/each}
                        {#if loading}
                        <p>Loading more posts...</p>
                        {/if}
                        {#if error}
                        <ErrorPopup message={error}></ErrorPopup>
                        {/if}
                        </div>
                        <div class="follow_button">
                            <Button buttonText={text} action={handleButtonClick}></Button>
                        </div>
                        </Panel>
                        </div>
                        </div>

<style>
.profile-container {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
}

.username {
    grid-area: username;
}

.feed {
    display: flex;
}
</style>

