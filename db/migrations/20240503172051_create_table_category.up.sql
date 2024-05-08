CREATE TABLE category
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL
);

insert into category(name) values ('Gadget');
insert into category(name) values ('Fashion');