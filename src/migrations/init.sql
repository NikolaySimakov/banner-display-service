create table tags (
	id SERIAL PRIMARY KEY,
	name TEXT
);

create table features (
	id SERIAL PRIMARY KEY,
	name TEXT
);

CREATE TABLE banners (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    text TEXT NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    last_version BOOLEAN NOT NULL DEFAULT TRUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
		tag_id INTEGER[] NOT NULL,
		feature_id INTEGER NOT NULL
);

CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    hash_key VARCHAR,
		user_status VARCHAR(15)
);


DROP TABLE api_keys;

DROP TABLE banners;

DROP TABLE tags;

DROP TABLE features;