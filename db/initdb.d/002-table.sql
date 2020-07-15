\c sns
CREATE TABLE app_users
(
    id bigserial,
    handle text NOT NULL UNIQUE,
    password text NOT NULL,
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

CREATE TABLE to_follows
(
    id      bigserial,
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

CREATE TRIGGER set_update_to_follows
BEFORE UPDATE ON to_follows
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE posts
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
    REFERENCES posts(id),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_update_posts
BEFORE UPDATE ON posts
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE praises
(
    id bigserial,
    user_id bigserial NOT NULL,
    post_id bigserial NOT NULL,
    disabled boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (user_id)
    REFERENCES app_users(id),
    FOREIGN KEY (post_id)
    REFERENCES post(id),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_update_praises
BEFORE UPDATE ON praises
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE diffusions
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

CREATE TRIGGER set_update_diffusions
BEFORE UPDATE ON diffusions
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE notifications
(
    id         bigserial,
    user_id    bigserial NOT NULL,
    type       smallint  NOT NULL,
    status     smallint  NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (user_id)
    REFERENCES app_users(id),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_update_notifications
BEFORE UPDATE ON notifications
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE notification_to_follows
(
    notification_id bigserial NOT NULL,
    to_follow_id bigserial NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (notification_id)
    REFERENCES notifications(id),
);

CREATE TRIGGER set_update_notification_to_follows
BEFORE UPDATE ON notification_to_follows
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE notification_praises
(
    notification_id bigserial NOT NULL,
    praise_id bigserial NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (notification_id)
    REFERENCES notifications(id),
    FOREIGN KEY (praise_id)
    REFERENCES praises(id),
);

CREATE TRIGGER set_update_notification_to_follows
BEFORE UPDATE ON notification_follows
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE notification_praises
(
    notification_id bigserial NOT NULL,
    praise_id bigserial NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (notification_id)
    REFERENCES notifications(id),
    FOREIGN KEY (praise_id)
    REFERENCES praises(id),
);

CREATE TRIGGER set_update_notification_praises
BEFORE UPDATE ON notification_praises
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE notification_diffusions
(
    notification_id bigserial NOT NULL,
    diffusion_id bigserial NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (notification_id)
    REFERENCES notifications(id),
    FOREIGN KEY (diffusion_id)
    REFERENCES diffusions(id),
);

CREATE TRIGGER set_update_notification_diffusions
BEFORE UPDATE ON notification_diffusions
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE notification_mentions
(
    notification_id bigserial NOT NULL,
    post_id bigserial NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    FOREIGN KEY (notification_id)
    REFERENCES notifications(id),
    FOREIGN KEY (post_id)
    REFERENCES posts(id),
);

CREATE TRIGGER set_update_notification_mentions
BEFORE UPDATE ON notification_mentions
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();