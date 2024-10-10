package models

import "gorm.io/gorm"

type Feed struct {
	/*
		CREATE TABLE `feeds` (
		  `id` int unsigned NOT NULL AUTO_INCREMENT,
		  `url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
		  `name` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
		  `is_active` bit(1) NOT NULL DEFAULT b'1',
		  `date_added` datetime NOT NULL,
		  `date_fetched` datetime NOT NULL,
		  `article_count` int NOT NULL DEFAULT '0',
		  PRIMARY KEY (`id`)
		) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	*/
	gorm.Model
	Url  string
	Name string
}

type Article struct {
	/*
	   CREATE TABLE `articles` (
	   	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	   	`feed_id` int NOT NULL,
	   	`url` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'RFC sates no lenght, but let''s be real',
	   	`title` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
	   	`image_url` varchar(2048) NOT NULL,
	   	`preview` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
	   	`content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
	   	`guid` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'Observe patterns over time',
	   	`author_name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
	   	`author_email` varchar(64) NOT NULL,
	   	`date_added` datetime NOT NULL,
	   	`date_published` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	   	`date_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	   	PRIMARY KEY (`id`)

	   ) ENGINE=InnoDB AUTO_INCREMENT=75800 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	*/
	gorm.Model
	FeedID      uint
	Feed        Feed `gorm:"foreignKey:FeedID"`
	Url         string
	Title       string
	ImageUrl    string
	Preview     string
	Content     string
	Guid        string
	AuthorName  string
	AuthorEmail string
}

type User struct {
	/*
		CREATE TABLE `users` (
		  `id` int unsigned NOT NULL AUTO_INCREMENT,
		  `user_uuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
		  `provider_user_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
		  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
		  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
		  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
		  `photo_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
		  `is_active` bit(1) NOT NULL DEFAULT b'1',
		  `date_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  `date_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		  PRIMARY KEY (`id`)
		) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
	*/
	gorm.Model
	UserUUID       string
	ProviderUserID string
	Username       string
	Name           string
	Email          string
	PhotoURL       string
}

/*
CREATE TABLE `feed_fetch` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `feed_id` int NOT NULL,
  `date_started` int NOT NULL,
  `date_complete` int DEFAULT NULL,
  `article_count` int NOT NULL,
  `status` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


CREATE TABLE `user_articles` (
  `feed_id` int NOT NULL,
  `user_id` int NOT NULL,
  `article_id` int NOT NULL,
  `date_viewed` datetime NOT NULL,
  PRIMARY KEY (`feed_id`,`user_id`,`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `users_feeds` (
  `user_id` int NOT NULL,
  `feed_id` int NOT NULL,
  `date_added` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `date_viewed` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `date_unread_start` datetime NOT NULL,
  PRIMARY KEY (`user_id`,`feed_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/
