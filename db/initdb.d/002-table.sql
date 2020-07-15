\c sns
CREATE TABLE app_users
(
    id bigserial,
    handle text NOT NULL UNIQUE,
    name text NOT NULL,
    birthday date,
    profile text,
    image text,
    is_private boolean NOT NULL DEFAULT false,
    disabled boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    PRIMARY KEY (id)
);

CREATE TRIGGER set_update_app_users
BEFORE UPDATE ON app_users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE follow
(
    to_user bigserial NOT NULL,
    by_user bigserial NOT NULL,
    permit  smallint NOT NULL default 0,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (to_user)
    REFERENCES app_users(id),
    FOREIGN KEY (by_user)
    REFERENCES app_users(id),
);

CREATE TRIGGER set_update_follow
BEFORE UPDATE ON follow
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE post
(
    id bigserial,
    user_id bigserial NOT NULL,
    to_post bigserial DEFAULT NULL,
    body text NOT NULL,
    disabled boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (user_id)
    REFERENCES app_users(id),
    FOREIGN KEY (to_post)
    REFERENCES post(id),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_update_post
BEFORE UPDATE ON post
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE praise
(
    user_id bigserial NOT NULL,
    post_id bigserial NOT NULL,
    disabled boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (user_id)
    REFERENCES app_users(id),
    FOREIGN KEY (post_id)
    REFERENCES post(id),
);

CREATE TRIGGER set_update_praise
BEFORE UPDATE ON praise
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE diffusion
(
    user_id bigserial NOT NULL,
    post_id bigserial NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (user_id)
    REFERENCES app_users(id),
    FOREIGN KEY (post_id)
    REFERENCES post(id),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_update_diffusion
BEFORE UPDATE ON diffusion
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();