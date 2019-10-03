DROP TABLE IF EXISTS `recipes`;

CREATE TABLE `recipes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `title_ru` varchar(100) DEFAULT NULL,
  `title_en` varchar(100) DEFAULT NULL,
  `intro_en` varchar(210) DEFAULT NULL,
  `intro_ru` varchar(210) DEFAULT NULL,
  `ingredients_ru` text,
  `ingredients_en` text,
  `text_ru` text,
  `text_en` text,
  `ready_ru` bit(1) DEFAULT b'0',
  `ready_en` bit(1) DEFAULT b'0',
  `approved_ru` bit(1) DEFAULT b'0',
  `approved_en` bit(1) DEFAULT b'0',
  `published_ru` bit(1) DEFAULT b'0',
  `published_en` bit(1) DEFAULT b'0',
  `slug` varchar(100) DEFAULT NULL,
  `time` smallint(5) unsigned DEFAULT '0',
  `image` varchar(220) DEFAULT 'default.jpg',
  `simple` bit(1) DEFAULT b'0',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;