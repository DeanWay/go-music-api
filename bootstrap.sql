create database music;

create table album (
  uuid UUID primary key,
  title varchar(255) not null,
  artist varchar(255) not null,
  price NUMERIC(20, 3) null,
  createdAt timestamp with time zone not null
);

create table song (
  uuid UUID primary key,
  albumUuid UUID references album (uuid),
  title varchar(255) not null,
  durationSeconds integer not null
);
