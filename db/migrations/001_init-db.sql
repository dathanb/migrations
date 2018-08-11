-- +migrate Up
-- +migrate StatementBegin
create table orders (
  id uuid primary key,
  customer text,
  created_at timestamp,
  updated_at timestamp
);

create table line_items (
  id uuid primary key,
  order_id uuid not null,
  product_id uuid not null,
  created_at timestamp,
  updated_at timestamp
);

create table products (
  id uuid primary key,
  name text not null,
  description text not null,
  amount integer not null,
  created_at timestamp,
  updated_at timestamp
);

create index product_name_index on products using gin(to_tsvector('english', "name"));
create index product_description_index on products using gin(to_tsvector('english', description));
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
drop table line_items;
drop table orders;
drop table skus;
-- +migrate StatementEnd
