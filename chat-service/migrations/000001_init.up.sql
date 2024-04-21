create table if not exists messages
(
    id         bigserial NOT NULL primary key,
    chat_id    text      not null,
    user_id    text      not null,
    text       text      not null,
    read       boolean   not null default false,
    created_at timestamp not null default now()
);

create table if not exists chats
(
    id text not null primary key
);

create table if not exists users_to_chats
(
    chat_id text not null references chats (id),
    user_id text not null,
    primary key (user_id, chat_id)
)