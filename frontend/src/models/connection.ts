export default class ConnectionModel {
	id: string = 'id';
	platform_id: string = 'platform_id';
	platform_name: string = 'platform_name';
	user_id: string = 'user_id';

	constructor(id: string, platform_id: string, platform_name: string, user_id: string) {
		this.id = id;
		this.platform_id = platform_id;
		this.platform_name = platform_name;
		this.user_id = user_id;
	}
}
