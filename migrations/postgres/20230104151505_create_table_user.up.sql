CREATE TABLE users (
	id varchar(50) NOT NULL,
	username varchar(50) NOT NULL,
	"password" text NOT NULL,
	fullname text NOT NULL,
    
	PRIMARY KEY (id),
	UNIQUE (username)
);
