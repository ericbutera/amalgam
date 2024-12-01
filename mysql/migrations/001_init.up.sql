CREATE TABLE
  `feeds` (
    `id` VARCHAR(36) NOT NULL DEFAULT (UUID ()),
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` datetime (3) DEFAULT NULL,
    `url` longtext,
    `name` longtext,
    `is_active` tinyint (1) DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `idx_feeds_deleted_at` (`deleted_at`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE
  `articles` (
    `id` VARCHAR(36) NOT NULL DEFAULT (UUID ()),
    `feed_id` VARCHAR(36) NOT NULL,
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` datetime (3) DEFAULT NULL,
    `url` longtext,
    `title` longtext,
    `description` longtext,
    `image_url` longtext,
    `preview` longtext,
    `content` longtext,
    `guid` longtext,
    `author_name` longtext,
    `author_email` longtext,
    PRIMARY KEY (`id`),
    KEY `fk_articles_feed` (`feed_id`),
    KEY `idx_articles_updated_at` (`updated_at`),
    KEY `idx_articles_deleted_at` (`deleted_at`),
    CONSTRAINT `fk_articles_feed` FOREIGN KEY (`feed_id`) REFERENCES `feeds` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE
  `users` (
    `id` VARCHAR(36) NOT NULL DEFAULT (UUID ()),
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` datetime (3) DEFAULT NULL,
    `user_uuid` longtext,
    `provider_user_id` longtext,
    `username` longtext,
    `name` longtext,
    `email` longtext,
    `photo_url` longtext,
    PRIMARY KEY (`id`),
    KEY `idx_users_deleted_at` (`deleted_at`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE
  `user_feeds` (
    `user_id` VARCHAR(36) NOT NULL,
    `feed_id` VARCHAR(36) NOT NULL,
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `viewed_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `unread_start_at` datetime NOT NULL,
    `unread_count` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`user_id`, `feed_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE
  `user_articles` (
    `feed_id` VARCHAR(36) NOT NULL,
    `user_id` VARCHAR(36) NOT NULL,
    `article_id` VARCHAR(36) NOT NULL,
    `viewed_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`feed_id`, `user_id`, `article_id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3;
