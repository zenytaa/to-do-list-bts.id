\c postgres

DROP DATABASE IF EXISTS todolist;
CREATE DATABASE todolist;

\c todolist;


CREATE TABlE users(
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE checklists(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE checklist_items(
    id BIGSERIAL PRIMARY KEY,
    item_name VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    checklist_id BIGINT NOT NULL REFERENCES checklists(id),
    is_done BOOLEAN NOT NULL DEFAULT FALSE
);