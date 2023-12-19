<script lang="ts">
    import Song from './Song.svelte';
    import Album from './Album.svelte';
    import Playlist from './Playlist.svelte';
    import SongModel from '../models/song';
    import AlbumModel from '../models/album';
    import PlaylistModel from '../models/playlist';
    export let content: SongModel | AlbumModel | PlaylistModel;
	let contentOptions = [
		{format: SongModel, component: Song },
		{format: AlbumModel, component: Album },
		{format: PlaylistModel, component: Playlist }
	];
    let selectedOption = contentOptions.find(type  => type.format == content?.constructor)?.component;
    let content_type = contentOptions.find(type  => type.format == content?.constructor)?.format;
    $: component = selectedOption;
    $: content = content_type;

</script>
<svelte:component this={component} content={content}/>
