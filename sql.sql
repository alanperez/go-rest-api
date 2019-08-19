create database droom;
-- creates table named droom
\c droom;

-- table struct
create table if not exists users (
  id serial primary key,
  email varchar(40) not null unique,
  password varchar(100) not null,
  status char(1) default '0'
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);

-- might need a seperate table.
create table if not exists companies (

)