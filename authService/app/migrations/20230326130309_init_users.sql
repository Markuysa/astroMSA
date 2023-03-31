-- +goose Up
-- +goose StatementBegin
create table users
(
    id         bigserial
        primary key,
    email      text      not null
        unique,
    sign       text      not null,
    name       text      not null,
    password   text      not null,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    birth_info timestamp not null
);
create table roles
(
    userid bigserial
        primary key
        references users,
    role   text not null
);

create index usersEmail on users (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users cascade;
drop table roles cascade;
-- +goose StatementEnd
