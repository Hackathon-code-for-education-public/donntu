create table if not exists universities
(
    id        text primary key not null,
    name      text             not null,
    long_name text             not null,
    logo      text             not null
);

create table if not exists university_owners
(
    university_id text references universities (id) not null,
    user_id       text                              not null
);

create table if not exists university_open_days
(
    university_id text references universities (id) not null,
    description   text                              not null,
    date          timestamp                         not null,
    address       text                              not null,
    link          text                              not null
);