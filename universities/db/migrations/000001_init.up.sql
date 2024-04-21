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

create type statuses as enum ('Студент этого вуза',
    'Выпускник этого вуза', 'Отчисленный', 'Некто');

create type sentiments as enum ('positive', 'negative', 'neutral');

create table if not exists university_reviews
(
    university_id text references universities (id) not null,
    author_status statuses                          not null,
    sentiment     sentiments                        not null,
    date          timestamp                         not null,
    text          text                              not null,
    repliesCount  integer                           not null
);

create type panorama_types as enum ('Корпуса',
    'Общежития',
    'Столовые',
    'Прочее');

create table if not exists university_panoramas
(
    university_id  text references universities (id) not null,
    address        text                              not null,
    name           text                              not null,
    firstLocation  text                              not null,
    secondLocation text                              not null,
    type           panorama_types                    not null
);