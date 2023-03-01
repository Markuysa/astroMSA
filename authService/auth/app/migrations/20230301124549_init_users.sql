-- +goose Up
-- +goose StatementBegin
create table users
(
    id         bigserial,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    email      text,
    birth_info timestamp,
    sign       text,
    name       text,
    primary key (id)
);

-- +goose StatementEnd
create index id_users on users(id);
-- +goose Down
-- +goose StatementBegin
drop index id_users;
drop table users;
-- +goose StatementEnd
