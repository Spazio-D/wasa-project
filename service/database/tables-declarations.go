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
	FOREIGN KEY(user_id) REFERENCES User(id)
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
	FOREIGN KEY(post_id, owner_id) REFERENCES Post(id, user_id)
		ON DELETE CASCADE,
	FOREIGN KEY(user_id) REFERENCES User(user_id)
		ON DELETE CASCADE
);`

var likeTable = `CREATE TABLE IF NOT EXISTS Like
(
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	owner_id INTEGER NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(user_id, post_id, owner_id),
	FOREIGN KEY(post_id, owner_id) REFERENCES Post(id, user_id)
		ON DELETE CASCADE,
	FOREIGN KEY(user_id) REFERENCES User(id)
		ON DELETE CASCADE
);`

var followTable = `CREATE TABLE IF NOT EXISTS Follow
(
	follower_id INTEGER NOT NULL,
	followed_id INTEGER NOT NULL,
	PRIMARY KEY(follower_id, followed_id),
	FOREIGN KEY(follower_id) REFERENCES User(id)
		ON DELETE CASCADE,
	FOREIGN KEY(followed_id) REFERENCES User(id)
		ON DELETE CASCADE
);`

var banTable = `CREATE TABLE IF NOT EXISTS Ban
(
	banner_id INTEGER NOT NULL,
	banned_id INTEGER NOT NULL,
	PRIMARY KEY(banner_id, banned_id),
	FOREIGN KEY (banner_id) REFERENCES User(userID)
		ON DELETE CASCADE,
	FOREIGN KEY (banned_id) REFERENCES User(userID)
		ON DELETE CASCADE
);`
