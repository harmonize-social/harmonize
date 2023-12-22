// playlist.ts
import type SongModel from "./song";

export default class PlaylistModel {
  id: string = 'id';
  name: string = 'name';
  songs: SongModel[] = [];
  url: string = 'url';

  constructor(id: string, name: string, songs: SongModel[], url: string) {
    this.id = id;
    this.name = name;
    this.songs = songs;
    this.url = url;
  }
}
