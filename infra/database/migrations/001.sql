create table "youtube" (
    id int primary key,
    title varchar(100),
    description varchar(300),
    channel_title varchar(100),
    published_at timestamp,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);


alter table "youtube" add column "url" varchar(100);