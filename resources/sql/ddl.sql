create database quality;

create schema accounts;

create table roles
(
  code        varchar(10) not null
    constraint roles_pkey
    primary key,
  description varchar(50)
);
