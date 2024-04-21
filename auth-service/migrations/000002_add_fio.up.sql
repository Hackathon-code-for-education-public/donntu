ALTER TABLE credentials
    ADD COLUMN last_name text not null default 'Иванов',
    ADD COLUMN first_name text not null default 'Иван',
    ADD COLUMN middle_name text not null default 'Иванович';