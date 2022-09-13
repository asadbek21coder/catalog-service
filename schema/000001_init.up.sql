create table "book_categories" (
    "category_id" serial primary key not null,
    "name" varchar(256) not null
);

create table "books" (
    "id" serial not null,
    "name" varchar(256),
    "author" varchar(128),
    "category" int references "book_categories" ("category_id") on delete cascade not null,
    "price" int not null,
    primary key("name", "author")
);

INSERT INTO "book_categories" (name) VALUES ('detective'), ('romance'), ('poetry'), ('adventure');