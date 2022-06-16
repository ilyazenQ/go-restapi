DROP TABLE IF EXISTS status;

DROP TABLE IF EXISTS currency;

DROP TABLE IF EXISTS transactions;

create table status (
     id serial not null unique,
     title varchar(64) not null unique,
     primary key(id)
);
create table currency (
     id serial not null unique,
     title varchar(64) not null unique,
     primary key(id)
);
create table transactions (
     id serial not null unique,
     user_id int not null,
     user_email varchar(255) not null,
     total float not null,
     currency_id int references currency (id) not null,
     status_id int references status (id) not null,
     created_at timestamp not null,
     updated_at timestamp not null,
     primary key(id)
);
insert into status(title)
values
    ('NEW'),
    ('SUCCESS'),
    ('FAIL'),
    ('ERROR'),
    ('CLOSED');

insert into currency(title)
values
    ('RUB'),
    ('EUR'),
    ('USD');
