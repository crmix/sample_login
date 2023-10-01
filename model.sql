create table users(
    id serial not null primary key,
    name character varying(32),
    phone int not null,
    username character varying(32)
);

insert INTO users (name, phone, username) values ('feruz', 996543414, 'feruz1');
insert into users(name, username, phone) values('Bexruz', 'bexruz', 910745351);
