

``` mysql
CREATE TABLE video_server.users (
	id INT UNSIGNED PRIMARY KEY auto_increment NULL,
	login_name varchar(64) UNIQUE KEY NULL,
	pwd TEXT NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;
```

``` mysql
CREATE TABLE video_server.video_info (
	id varchar(64) PRIMARY KEY NOT NULL,
	author_id INT UNSIGNED NULL,
	name TEXT NULL,
	display_ctime TEXT NULL,
	create_time DATETIME NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;
```

``` mysql
CREATE TABLE video_server.comments (
	id varchar(64) PRIMARY KEY NOT NULL,
	video_id varchar(64) NULL,
	author_id INT UNSIGNED NULL,
	content TEXT NULL,
	`time` DATETIME NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;
```

``` mysql
CREATE TABLE video_server.sessions (
	session_id TINYTEXT NOT NULL,
	TTL TINYTEXT NULL,
	login_name varchar(64) NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COLLATE=utf8_general_ci;
```

