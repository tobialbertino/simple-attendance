CREATE TABLE attendances (
	id varchar(50) NOT NULL,
	user_id varchar(50) NOT NULL,
    activity TEXT,
	location Text,
	check_in INT8,
    check_out INT8,
	
	PRIMARY KEY (id)
);

ALTER TABLE attendances ADD CONSTRAINT attendances_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;



