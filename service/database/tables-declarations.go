package database

var userTable = `CREATE TABLE IF NOT EXISTS User
(
	id INTEGER NOT NULL UNIQUE, 
	username VARCHAR(16) NOT NULL UNIQUE,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id)
);`

var postTable = `CREATE TABLE IF NOT EXISTS Post 
(
	id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id, user_id),
	CONSTRAINT postFK_user
		FOREIGN KEY (user_id) REFERENCES User(id)
		ON DELETE CASCADE
);`

var commentTable = `CREATE TABLE IF NOT EXISTS Comment
(
	id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	owner_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	text TEXT NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id, owner_id, post_id),
	CONSTRAINT commentFK_post
		FOREIGN KEY (post_id, owner_id) REFERENCES Post(id, user_id)
		ON DELETE CASCADE,
	CONSTRAINT commentFK_user
		FOREIGN KEY (user_id) REFERENCES User(id)
		ON DELETE CASCADE
);`

var likeTable = `CREATE TABLE IF NOT EXISTS Like
(
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	owner_id INTEGER NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(user_id, post_id, owner_id),
	CONSTRAINT likeFK_post
		FOREIGN KEY (post_id, owner_id) REFERENCES Post(id, user_id)
		ON DELETE CASCADE,
	CONSTRAINT likeFK_user
		FOREIGN KEY (user_id) REFERENCES User(id)
		ON DELETE CASCADE
);`

var followTable = `CREATE TABLE IF NOT EXISTS Follow
(
	follower_id INTEGER NOT NULL,
	followed_id INTEGER NOT NULL,
	PRIMARY KEY(follower_id, followed_id),
	CONSTRAINT followFK_follower
		FOREIGN KEY (follower_id) REFERENCES User(id)
		ON DELETE CASCADE,
	CONSTRAINT followFK_followed
		FOREIGN KEY (followed_id) REFERENCES User(id)
		ON DELETE CASCADE
);`

var banTable = `CREATE TABLE IF NOT EXISTS Ban
(
	banner_id INTEGER NOT NULL,
	banned_id INTEGER NOT NULL,
	PRIMARY KEY(banner_id, banned_id),
	CONSTRAINT banFK_banner
		FOREIGN KEY (banner_id) REFERENCES User(id)
		ON DELETE CASCADE,
	CONSTRAINT banFK_banned
		FOREIGN KEY (banned_id) REFERENCES User(id)
		ON DELETE CASCADE
);`
