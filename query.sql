CREATE TABLE libraries (
	id_libraries serial PRIMARY KEY,
	name VARCHAR ( 50 )  NOT NULL,
	route VARCHAR ( 50 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	scanned_at TIMESTAMP NOT NULL
);
CREATE TABLE songs (
	id_songs serial PRIMARY KEY,
	id_libraries int,
	name VARCHAR ( 50 )  NOT NULL,
	created_at TIMESTAMP NOT NULL,
	FOREIGN KEY (id_libraries) REFERENCES libraries (id_libraries)
);
CREATE TABLE metadata (
	id_metadata serial PRIMARY KEY,
	id_songs int,
	artist VARCHAR ( 50 )  NOT NULL,
	album VARCHAR ( 50 )  NOT NULL,
	track VARCHAR ( 50 )  NOT NULL,
	genre VARCHAR ( 50 )  NOT NULL,
	year VARCHAR ( 50 )  NOT NULL,
	created_at TIMESTAMP NOT NULL,
	FOREIGN KEY (id_songs) REFERENCES songs (id_songs)
);


