create table if not exists messages (
    id bigserial NOT NULL primary key,
    chatId text not null,
    userId text not null,
    text text not null,
    read boolean not null default false,
    createdAt timestamp not null
);

create table if not exists chats (
    id text not null primary key
);

create table if not exists users_to_chats (
    userId text not null,
    chatId text not null references chats(id),
    primary key (userId, chatId)
)