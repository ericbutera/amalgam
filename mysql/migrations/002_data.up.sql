UNLOCK TABLES;

LOCK TABLES `feeds` WRITE;

/*!40000 ALTER TABLE `feeds` DISABLE KEYS */;

INSERT INTO
    `feeds` (
        `id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `url`,
        `name`,
        `is_active`
    )
VALUES
    (
        "0e597e90-ece5-463e-8608-ff687bf286da",
        '2024-10-12 13:44:40.000',
        '2024-10-12 13:44:40.000',
        NULL,
        'https://news.ycombinator.com/rss',
        'hacker news',
        0
    ),
    (
        "9677b210-2211-49f1-902c-dbd5ee357ac9",
        '2024-10-12 13:44:40.000',
        '2024-10-12 13:44:40.000',
        NULL,
        'http://faker:8080/feed/9677b210-2211-49f1-902c-dbd5ee357ac9',
        'fake feed',
        1
    );

/*!40000 ALTER TABLE `feeds` ENABLE KEYS */;

UNLOCK TABLES;

LOCK TABLES `users` WRITE;

/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO
    `users` (
        `id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `user_uuid`,
        `provider_user_id`,
        `username`,
        `name`,
        `email`,
        `photo_url`
    )
VALUES
    (
        "2e597e90-ece5-463e-8608-ff687bf286da",
        '2024-10-12 13:43:25.000',
        '2024-10-12 13:43:25.000',
        NULL,
        'user',
        'user',
        'user',
        'name',
        'email@example.com',
        NULL
    );

/*!40000 ALTER TABLE `users` ENABLE KEYS */;

UNLOCK TABLES;

LOCK TABLES `articles` WRITE;

/*!40000 ALTER TABLE `articles` DISABLE KEYS */;

INSERT INTO
    `articles` (
        `id`,
        `created_at`,
        `updated_at`,
        `deleted_at`,
        `feed_id`,
        `url`,
        `title`,
        `image_url`,
        `preview`,
        `content`,
        `guid`,
        `author_name`,
        `author_email`
    )
VALUES
    (
        "2a1a227d-0375-4d9a-943f-85017ac4f930",
        '2024-10-12 13:42:14.000',
        '2024-10-12 13:42:14.000',
        NULL,
        "9677b210-2211-49f1-902c-dbd5ee357ac9",
        'https://mainichi.jp/english/articles/20220916/p2a/00m/0sc/017000c',
        '2-in-1 calculator app adds up to surprise hit for retired engineer',
        NULL,
        NULL,
        '<a href=\\\"https://news.ycombinator.com/item?id=32902520\\\">Comments</a>',
        NULL,
        NULL,
        NULL
    );

/*!40000 ALTER TABLE `articles` ENABLE KEYS */;

UNLOCK TABLES;

LOCK TABLES `user_feeds` WRITE;

/*!40000 ALTER TABLE `user_feeds` DISABLE KEYS */;

INSERT INTO
    `user_feeds` (
        `user_id`,
        `feed_id`,
        `created_at`,
        `viewed_at`,
        `unread_start_at`,
        `unread_count`
    )
VALUES
    (
        "2e597e90-ece5-463e-8608-ff687bf286da",
        "0e597e90-ece5-463e-8608-ff687bf286da",
        '2022-10-06 10:57:54',
        '2022-11-13 14:57:49',
        '2022-10-11 00:17:27',
        0
    ),
    (
        "2e597e90-ece5-463e-8608-ff687bf286da",
        "9677b210-2211-49f1-902c-dbd5ee357ac9",
        '2024-11-11 11:11:11',
        '2024-11-11 11:11:11',
        '2024-11-11 11:11:11',
        1
    );

/*!40000 ALTER TABLE `user_feeds` ENABLE KEYS */;

UNLOCK TABLES;
