<script lang="ts">
	import type CommentModel from '../models/comment';
	import type PostModel from '../models/post';
	import Comment from './Comment.svelte';
	import Song from './Song.svelte';
	import Album from './Album.svelte';
	import Playlist from './Playlist.svelte';
	import Artist from './Artist.svelte';
	import ActionButton from './ActionButton.svelte';
	import { delete_, post, throwError } from '../fetch';
	import ErrorPopup from './ErrorPopup.svelte';
	import { errorMessage } from '../store';

	export let content: any;
	let comments: CommentModel[] = [];
	export let caption: PostModel['caption'];
	export let likes: PostModel['likeCount'];
	export let id: PostModel['id'];
	export let typez: string;
	export let isLiked: PostModel['hasLiked'];
	export let isSaved: PostModel['hasSaved'];
	let error = '';
	errorMessage.subscribe((value) => {
		error = value;
	});

	async function postLike(): Promise<string> {
		try {
			const response: string = await post<string, string>(`/me/liked?id=${id}`, id);
			return response;
		} catch (e) {
			throwError('Error posting like');
			return e as string;
		}
	}
	async function postSave(): Promise<string> {
		try {
			const response: string = await post<string, string>(`/me/saved?id=${id}`, id);
			return response;
		} catch (e) {
			throwError('Error posting save');
			return e as string;
		}
	}

	async function deleteLike(): Promise<string> {
		try {
			const response: string = await delete_<string>(`/me/liked?id=${id}`);
			return response;
		} catch (e) {
			throwError('Error deleting like');
			return e as string;
		}
	}

	async function deleteSave(): Promise<string> {
		try {
			const response: string = await delete_<string>(`/me/saved?id=${id}`);
			return response;
		} catch (e) {
			throwError('Error deleting save');
			return e as string;
		}
	}

	function toggleLikeButton() {
		if (isLiked) {
			deleteLike();
			likes--;
		} else {
			postLike();
			likes++;
		}
		isLiked = !isLiked;
	}

	function toggleSaveButton() {
		if (isSaved) {
			deleteSave();
		} else {
			postSave();
		}
		isSaved = !isSaved;
	}
</script>

<div class="post">
	<h3>{caption}</h3>
	{#if typez == 'song'}
		<Song {content} />
	{:else if typez == 'album'}
		<Album {content} />
	{:else if typez == 'playlist'}
		<Playlist {content} />
	{:else if typez == 'artist'}
		<Artist {content} />
	{:else}
		<p>Invalid content type</p>
	{/if}
	<h4>Comments:</h4>
	{#if comments.length == 0}
		<p>No comments yet</p>
	{/if}
	{#each comments as _comment}
		<Comment {content} />
	{/each}
	<h4>Likes: {likes}</h4>
	<div class="action-buttons">
		<ActionButton state={isLiked} type="like" action={toggleLikeButton} />

		<ActionButton state={isSaved} type="save" action={toggleSaveButton} />
	</div>

	{#if error}
		<ErrorPopup message={error}></ErrorPopup>
	{/if}
</div>

<style>
	.post {
		border: 1px solid black;
		padding: 0px;
		margin: 50px;
		border-radius: 10px;
		color: black;
		text-align: center;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	h3 {
		max-width: 75%;
		justify-items: flex-start;
	}

    .action-buttons {
        display: flex;
        flex-direction: row;
        justify-content: space-evenly;
        width: 100%;
    }
</style>
