CREATE TABLE `feeds` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `url` longtext,
  `name` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_feeds_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `articles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `feed_id` bigint unsigned DEFAULT NULL,
  `url` longtext,
  `title` longtext,
  `image_url` longtext,
  `preview` longtext,
  `content` longtext,
  `guid` longtext,
  `author_name` longtext,
  `author_email` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_articles_deleted_at` (`deleted_at`),
  KEY `fk_articles_feed` (`feed_id`),
  CONSTRAINT `fk_articles_feed` FOREIGN KEY (`feed_id`) REFERENCES `feeds` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_uuid` longtext,
  `provider_user_id` longtext,
  `username` longtext,
  `name` longtext,
  `email` longtext,
  `photo_url` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `user_feeds` (
  `user_id` int NOT NULL,
  `feed_id` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `viewed_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `unread_start_at` datetime NOT NULL,
  PRIMARY KEY (`user_id`, `feed_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `user_articles` (
  `feed_id` int NOT NULL,
  `user_id` int NOT NULL,
  `article_id` int NOT NULL,
  `viewed_at` datetime NOT NULL,
  PRIMARY KEY (`feed_id`, `user_id`, `article_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3;