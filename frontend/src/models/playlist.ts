// playlist.ts
import type SongModel from "./song";

export default class PlaylistModel {
  id: string = 'id';
  title: string = 'name';
  songs: SongModel[] = [];
  mediaUrl: string = 'url';

  constructor(id: string, name: string, songs: SongModel[], url: string) {
    this.id = id;
    this.title = name;
    this.songs = songs;
    this.mediaUrl = url;
  }
}
