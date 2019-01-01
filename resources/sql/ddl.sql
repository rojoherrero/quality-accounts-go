create database "quality-go"
  with owner postgres;

create table users
(
  id bigserial not null
    constraint users_pk
      primary key,
  full_name varchar(100),
  user_name varchar(50),
  password varchar(100),
  created timestamp default now(),
  updated timestamp default now()
);

alter table users owner to postgres;

create unique index users_id_uindex
  on users (id);

create unique index users_username_uindex
  on users (user_name);

create table roles
(
  code varchar(50) not null
    constraint roles_pk
      primary key,
  name varchar(50) not null
);

alter table roles owner to postgres;

create table departments
(
  code varchar(50) not null
    constraint departments_pk
      primary key,
  name varchar(50) not null
);

alter table departments owner to postgres;

create table user_roles_departments
(
  user_id bigint not null
    constraint user_roles_departments_users_fk
      references users,
  role_code varchar(50) not null
    constraint user_roles_departments_roles_fk
      references roles,
  departmente_code varchar(50) not null
    constraint user_roles_departments_departments_fk
      references departments,
  constraint user_roles_departments_pk
    primary key (user_id, role_code, departmente_code)
);

alter table user_roles_departments owner to postgres;

