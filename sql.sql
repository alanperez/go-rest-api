create database droom;

\c droom;

create table if not exists users (
  id bigserial primary key,
  email varchar(40) not null unique,
  password varchar(100) not null,
  status char(1) default '0'
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);

create table if not exists companies (

)