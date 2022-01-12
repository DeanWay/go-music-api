create table album (
  uuid UUID primary key,
  title varchar(255) not null,
  artist varchar(255) not null,
  price NUMERIC(20, 3) null,
  created_at timestamp with time zone not null
);

create table song (
  uuid UUID primary key,
  title varchar(255) not null,
  artist varchar(255) not null,
  duration_seconds integer not null,
  audio_file varchar(255) null
);

create table album_song (
  album_uuid UUID references album (uuid),
  song_uuid UUID references song (uuid)
);
