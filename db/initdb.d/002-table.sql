CREATE DATABASE IF NOT EXISTS sns;
use sns;
CREATE TABLE app_users
(
    id bigint unsigned AUTO_INCREMENT,
    handle VARCHAR(255) NOT NULL UNIQUE,
    password text NOT NULL,
    name text NOT NULL,
    birthday date,
    profile text,
    image text,
    is_private boolean NOT NULL DEFAULT false,
    disabled boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE tokens
(
    id bigint unsigned AUTO_INCREMENT,
    user_id bigint unsigned NOT NULL,
    token text NOT NULL,
    created_at       timestamp NOT NULL DEFAULT current_timestamp,
    updated_at       timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id)
     REFERENCES app_users(id),
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE to_follows
(
    id      bigint unsigned,
    to_user bigint unsigned NOT NULL,
    by_user bigint unsigned NOT NULL,
    permission  int unsigned NOT NULL default 0,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (to_user)
    REFERENCES app_users(id),
    FOREIGN KEY (by_user)
    REFERENCES app_users(id),
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE posts
(
    id bigint unsigned AUTO_INCREMENT,
    user_id bigint unsigned NOT NULL,
    to_post bigint unsigned,
    body text NOT NULL,
    disabled boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id)
    REFERENCES app_users(id),
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE to_posts
(
    by_post bigint unsigned NOT NULL,
    to_post bigint unsigned NOT NULL,
    FOREIGN KEY (by_post)
    REFERENCES posts(id),
    FOREIGN KEY (to_post)
    REFERENCES posts(id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE praises
(
    id bigint unsigned AUTO_INCREMENT,
    user_id bigint unsigned NOT NULL,
    post_id bigint unsigned NOT NULL,
    disabled boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id)
    REFERENCES app_users(id),
    FOREIGN KEY (post_id)
    REFERENCES posts(id),
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE diffusions
(
    id bigint unsigned AUTO_INCREMENT,
    user_id bigint unsigned NOT NULL,
    post_id bigint unsigned NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id)
    REFERENCES app_users(id),
    FOREIGN KEY (post_id)
    REFERENCES posts(id),
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE notifications
(
    id         bigint unsigned,
    user_id    bigint unsigned NOT NULL,
    type       int unsigned  NOT NULL,
    status     int unsigned  NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (user_id)
    REFERENCES app_users(id),
    PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE notification_to_follows
(
    notification_id bigint unsigned NOT NULL,
    to_follow_id bigint unsigned NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (notification_id)
    REFERENCES notifications(id),
    FOREIGN KEY (to_follow_id)
    REFERENCES to_follows(id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE notification_praises
(
    notification_id bigint unsigned NOT NULL,
    praise_id bigint unsigned NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (notification_id)
    REFERENCES notifications(id),
    FOREIGN KEY (praise_id)
    REFERENCES praises(id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE notification_diffusions
(
    notification_id bigint unsigned NOT NULL,
    diffusion_id bigint unsigned NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (notification_id)
    REFERENCES notifications(id),
    FOREIGN KEY (diffusion_id)
    REFERENCES diffusions(id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE notification_mentions
(
    notification_id bigint unsigned NOT NULL,
    post_id bigint unsigned NOT NULL,
    created_at timestamp NOT NULL DEFAULT current_timestamp,
    updated_at timestamp NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (notification_id)
    REFERENCES notifications(id),
    FOREIGN KEY (post_id)
    REFERENCES posts(id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
