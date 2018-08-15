-- +migrate Up
-- +migrate StatementBegin
create table users (
  id integer primary key,
  display_name text
);

create table posts (
  id integer primary key,
  post_type integer not null,
  post_type_id integer not null,
  body text not null,
  user_id integer not null references users(id)
);

create table comments (
  id integer primary key,
  post_id integer not null references posts(id),
  "text" text not null,
  user_id integer not null references users(id)
);
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
drop table users;
drop table posts;
-- +migrate StatementEnd
