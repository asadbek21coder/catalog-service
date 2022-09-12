create table "books" (
    "id" serial not null primary key,
    "name" varchar(256) not null,
    "author" varchar(128) not null,
    "category" int  not null
);

create table "book_categories" (
    "category_id" serial primary key not null,
    "name" varchar(256) not null
);