\c sns

INSERT INTO app_uses(id, handle, name, birthday, profile, image, is_private) VALUES
    (1, "test_01", "テスト1", "2000-01-01", "こんにちは", "", false),
    (2, "test_02", "テスト2", "2000-01-02", "やあ", "", false),
    (3, "test_03", "テスト3", "2000-01-03", "３", "", false),
    (4, "test_04", "テスト4", "2000-01-04", "４", "", false),
    (5, "test_05", "テスト5", "2000-01-05", "５", "", false),
    (6, "test_06", "テスト6", "2000-01-06", "６", "", true);

/* 0;未確認 1:許可 2:拒否 */
INSERT INTO follow(to_user, by_user, permit) VALUES
    (1, 2, 1 ),
    (1, 3, 1 ),
    (1, 4, 1 ),
    (1, 5, 1 ),
    (1, 6, 1 ),
    (2, 1, 1 ),
    (2, 5, 1 ),
    (3, 2, 1 ),
    (4, 6, 0 ),
    (5, 6, 2 );

INSERT INTO post(id, user_id, to_post, body) VALUES
    (1, 1, 0, "投稿1"),
    (2, 2, 0, "投稿2"),
    (3, 3, 0, "投稿3"),
    (4, 4, 0, "投稿4"),
    (5, 5, 0, "投稿5"),
    (6, 6, 0, "投稿6"),
    (7, 1, 2, "投稿2へ送信");

INSERT INTO praise(user_id, post_id) VALUES
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

INSERT INTO diffusion(user_id, post_id) VALUES
    (1, 2),
    (1, 3),
    (1, 4),
    (1, 5),
    (3, 1),
    (6, 1);