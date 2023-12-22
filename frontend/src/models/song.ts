import { onMount } from 'svelte';
import { get } from '../fetch'; 

export let title: string;
export let image: string;
export let alt: string;
export let album: string;
export let artist: string;
export let url: string;

export default class SongModel{
    id: string = 'id'
    title: string = 'title'
    url: string = 'url'

    constructor(id: string, title: string, url: string){
        this.id = id
        this.title = title
        this.url = url
    }
}

async function fetchData() {
    try {
      const content = await get<{ title: string, image: string, alt: string, album: string, artist: string, url: string }>('URL_VAN_DE_API');
  
      title = content.title;
      image = content.image;
      alt = content.alt;
      album = content.album;
      artist = content.artist;
      url = content.url;
    } catch (error) {
      console.error('Error fetching data:', (error as any).message);
    }
  }
  
  onMount(() => {
    fetchData();
  });
  