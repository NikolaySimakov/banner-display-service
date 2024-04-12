create table tags (
	id SERIAL PRIMARY KEY,
	name TEXT
);

create table features (
	id SERIAL PRIMARY KEY,
	name TEXT
);

drop table tags;

drop table features;