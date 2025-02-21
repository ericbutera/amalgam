CREATE TABLE
  `feeds` (
    `id` CHAR(36) NOT NULL DEFAULT (UUID ()),
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` datetime (3) DEFAULT NULL,
    `url` longtext,
    `url_hash` CHAR(64) GENERATED ALWAYS AS (SHA2 (url, 256)) STORED,
    `name` longtext,
    `is_active` tinyint (1) DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `idx_feeds_deleted_at` (`deleted_at`),
    UNIQUE (`url_hash`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE
  `articles` (
    `id` CHAR(36) NOT NULL DEFAULT (UUID ()),
    `feed_id` CHAR(36) NOT NULL,
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
    `id` CHAR(36) NOT NULL DEFAULT (UUID ()),
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
    `user_id` CHAR(36) NOT NULL,
    `feed_id` CHAR(36) NOT NULL,
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `viewed_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `unread_start_at` datetime NOT NULL,
    `unread_count` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`user_id`, `feed_id`),
    CONSTRAINT `fk_user_feeds_feeds_id` FOREIGN KEY (`feed_id`) REFERENCES `feeds` (`id`),
    CONSTRAINT `fk_user_feeds_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
    -- TODO: index for ordering
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE
  `user_articles` (
    `user_id` CHAR(36) NOT NULL,
    `article_id` CHAR(36) NOT NULL,
    `viewed_at` datetime (3) NULL,
    PRIMARY KEY (`user_id`, `article_id`),
    CONSTRAINT `fk_user_articles_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    CONSTRAINT `fk_user_articles_article_id` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE
  `feed_verification` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `url` longtext,
    `url_hash` CHAR(64) GENERATED ALWAYS AS (SHA2 (url, 256)) STORED,
    `user_id` CHAR(36) NULL, -- user who requested the verification
    `status` enum ('pending', 'success', 'failed') NOT NULL DEFAULT 'pending',
    `workflow_id` longtext,
    `message` longtext,
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    UNIQUE (`url_hash`),
    CONSTRAINT `fk_fv_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- fetch_history is append only
CREATE TABLE
  `fetch_history` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `feed_id` CHAR(36) COMMENT 'existing feed',
    `feed_verification_id` bigint COMMENT 'only for new feed verification fetching',
    `response_code` INT,
    `etag` longtext,
    `workflow_id` longtext,
    `bucket` longtext,
    `message` longtext,
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_fh_feed_id` FOREIGN KEY (`feed_id`) REFERENCES `feeds` (`id`),
    CONSTRAINT `fk_fh_feed_verification_id` FOREIGN KEY (`feed_verification_id`) REFERENCES `feed_verification` (`id`),
    CONSTRAINT `chk_feed_or_url` CHECK (
      `feed_id` IS NOT NULL
      OR `feed_verification_id` IS NOT NULL
    )
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
