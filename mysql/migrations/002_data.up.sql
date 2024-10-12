UNLOCK TABLES;

LOCK TABLES `feeds` WRITE;

/*!40000 ALTER TABLE `feeds` DISABLE KEYS */;

INSERT INTO `feeds` (`id`, `created_at`, `updated_at`, `deleted_at`, `url`, `name`)
VALUES
	(1,'2024-10-12 13:44:40.000',NULL,NULL,'https://news.ycombinator.com/rss','hacker news');

/*!40000 ALTER TABLE `feeds` ENABLE KEYS */;

UNLOCK TABLES;

LOCK TABLES `users` WRITE;

/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_uuid`, `provider_user_id`, `username`, `name`, `email`, `photo_url`)
VALUES
	(1,'2024-10-12 13:43:25.000',NULL,NULL,'user','user','user','name','email@example.com',NULL);

/*!40000 ALTER TABLE `users` ENABLE KEYS */;

UNLOCK TABLES;

LOCK TABLES `articles` WRITE;

/*!40000 ALTER TABLE `articles` DISABLE KEYS */;

INSERT INTO `articles` (`id`, `created_at`, `updated_at`, `deleted_at`, `feed_id`, `url`, `title`, `image_url`, `preview`, `content`, `guid`, `author_name`, `author_email`)
VALUES
	(1,'2024-10-12 13:42:14.000',NULL,NULL,1,'https://mainichi.jp/english/articles/20220916/p2a/00m/0sc/017000c','2-in-1 calculator app adds up to surprise hit for retired engineer',NULL,NULL,'<a href=\\\"https://news.ycombinator.com/item?id=32902520\\\">Comments</a>',NULL,NULL,NULL);

/*!40000 ALTER TABLE `articles` ENABLE KEYS */;

UNLOCK TABLES;

LOCK TABLES `user_articles` WRITE;

/*!40000 ALTER TABLE `user_articles` DISABLE KEYS */;

INSERT INTO
    `user_articles` (`feed_id`, `user_id`, `article_id`, `viewed_at`)
VALUES
    (1, 1, 1, '2022-10-09 23:17:22'),
    (1, 1, 2, '2022-10-09 17:25:44'),
    (1, 1, 3, '2022-10-09 17:25:01');

/*!40000 ALTER TABLE `user_articles` ENABLE KEYS */;

UNLOCK TABLES;

LOCK TABLES `user_feeds` WRITE;

/*!40000 ALTER TABLE `user_feeds` DISABLE KEYS */;

INSERT INTO
    `user_feeds` (
        `user_id`,
        `feed_id`,
        `created_at`,
        `viewed_at`,
        `unread_start_at`
    )
VALUES
    (
        1,
        1,
        '2022-10-06 10:57:54',
        '2022-11-13 14:57:49',
        '2022-10-11 00:17:27'
    );

/*!40000 ALTER TABLE `user_feeds` ENABLE KEYS */;

UNLOCK TABLES;
