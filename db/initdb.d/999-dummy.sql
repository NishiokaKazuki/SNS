\c sns

INSERT INTO app_users(id, handle, password, name, birthday, profile, image, is_private) VALUES
    (1, "test_01", "secret", "テスト1", "2000-01-01", "こんにちは", "", false),
    (2, "test_02", "secret", "テスト2", "2000-01-02", "やあ", "", false),
    (3, "test_03", "secret", "テスト3", "2000-01-03", "３", "", false),
    (4, "test_04", "secret", "テスト4", "2000-01-04", "４", "", false),
    (5, "test_05", "secret", "テスト5", "2000-01-05", "５", "", false),
    (6, "test_06", "secret", "テスト6", "2000-01-06", "６", "", true);

/* 0;未確認 1:許可 2:拒否 */
INSERT INTO to_follows(id, to_user, by_user, permit) VALUES
    (1, 1, 2, 1 ),
    (2, 1, 3, 1 ),
    (3, 1, 4, 1 ),
    (4, 1, 5, 1 ),
    (5, 1, 6, 1 ),
    (6, 2, 1, 1 ),
    (7, 2, 5, 1 ),
    (8, 3, 2, 1 ),
    (9, 6, 4, 0 ),
    (10, 6, 3, 1 ),
    (11, 6, 5, 2 );

INSERT INTO posts(id, user_id, to_post, body) VALUES
    (1, 1, 0, "投稿1"),
    (2, 2, 0, "投稿2"),
    (3, 3, 0, "投稿3"),
    (4, 4, 0, "投稿4"),
    (5, 5, 0, "投稿5"),
    (6, 6, 0, "投稿6"),
    (7, 1, 2, "@test_02 投稿2へ送信");

INSERT INTO praises(user_id, post_id) VALUES
    (1, 2),
    (1, 3),
    (1, 4),
    (1, 5),
    (1, 6),
    (2, 1),
    (3, 1),
    (4, 1),
    (6, 1),
    (3, 7);

INSERT INTO diffusions(user_id, post_id) VALUES
    (1, 2),
    (1, 3),
    (1, 4),
    (1, 5),
    (3, 1),
    (6, 1);

/* type   0:to_follows 1:praises 2:diffusions 3:mentions*/
/* status 0:未確認 1:確認済み 2:その他 */
INSERT INTO notifications(id, user_id, type, status) VALUES
    (1, 1, 0, 0),
    (2, 1, 1, 0),
    (3, 2, 2, 0),
    (4, 1, 3, 0),
    (5, 1, 0, 1),
    (6, 6, 0, 0),
    (7, 6, 0, 1);

INSERT INTO notification_to_follows(notification_id, to_follow_id) VALUES
    (1, 1),
    (5, 2),
    (6, 9),
    (7, 10);

INSERT INTO notification_to_praises(notification_id, praise_id) VALUES
    (2, 1);

INSERT INTO notification_diffusions(notification_id, diffusion_id) VALUES
    (3, 1);

INSERT INTO notification_mentions(notification_id, post_id) VALUES
    (4, 7);