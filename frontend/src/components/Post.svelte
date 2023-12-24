<script lang="ts">
	import type CommentModel from '../models/comment';
	import type PostModel from '../models/post';
	import Comment from './Comment.svelte';
	import Song from './Song.svelte';
	import Album from './Album.svelte';
	import Playlist from './Playlist.svelte';
	import Artist from './Artist.svelte';
	import ActionButton from './ActionButton.svelte';
	import {  post, throwError } from '../fetch.js';
	import type ArtistModel from '../models/artist';
	export let content: any;
	let comments: CommentModel[] = [];
	export let caption: PostModel['caption'];
	export let likes: PostModel['likeCount'];
	export let id: string;
	export let typez: string;

	// async function getArtists(): Promise<ArtistModel[]> {
	// 	try {
	// 		const response: ArtistModel[] = await get<ArtistModel[]>(`/artists?id=${id}`);
	// 		return response;
	// 	} catch (error) {
	// 		throwError('Error fetching artists');
	// 		return [];
	// 	}
	// }
	// async function getComments(): Promise<CommentModel[]> {
	// 	try {
	// 		const response: CommentModel[] = await get<CommentModel[]>(`/comments?id=${id}`);
	// 		return response;
	// 	} catch (error) {
	// 		throwError('Error fetching comments');
	// 		return [];
	// 	}
	// }
	async function postLike(): Promise<number> {
		try {
			const response: number = await post<number, number>(`/likes?id=${id}`, 0);
			return response;
		} catch (error) {
			throwError('Error posting like');
			return 0;
		}
	}
	async function postSave(): Promise<number> {
		try {
			const response: number = await post<number, number>(`/me/saved?id=${id}`, 0);
			return response;
		} catch (error) {
			throwError('Error posting save');
			return 0;
		}
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
	{#each comments as _comment}
		<Comment {content} />
	{/each}
	<h4>Likes: {likes}</h4>
	<ActionButton type="like" on:click={postLike} />
	<ActionButton type="save" on:click={postSave} />
</div>

<style>
	.post {
		border: 1px solid black;
		padding: 50px;
		margin: 50px;
		border-radius: 10px;
		background-color: grey;
		color: black;
		text-align: center;
	}
</style>
