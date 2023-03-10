-- +goose Up
-- +goose StatementBegin
create table users
(
    id         bigserial,
    email      text,
    sign       text,
    name       text,
    password text,
    created_at timestamp,
    updated_at timestamp,
    birth_info timestamp,
    primary key (id)
);

-- +goose StatementEnd
create index id_users on users(id);
-- +goose Down
-- +goose StatementBegin
drop index id_users;
drop table users;
-- +goose StatementEnd
