alter table universities
    add rating float not null default 0;
alter table universities
    add region text not null;
alter table universities
    add budget_places int not null default 0;
create type university_type as enum ('Государственный', 'Частный');
alter table universities
    add type university_type not null default 'Государственный';
alter table universities
    add study_fields int not null default 0;