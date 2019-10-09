CREATE TABLE recipes
(
    id             int(10) UNSIGNED                 NOT NULL AUTO_INCREMENT,
    user_id        int(10) UNSIGNED                 NOT NULL,
    title_ru       varchar(100)         DEFAULT NULL,
    title_en       varchar(100)         DEFAULT NULL,
    intro_en       varchar(210)         DEFAULT NULL,
    intro_ru       varchar(210)         DEFAULT NULL,
    ingredients_ru text,
    ingredients_en text,
    text_ru        text,
    text_en        text,
    ready_ru       tinyint(1) UNSIGNED  DEFAULT '0' NOT NULL,
    ready_en       tinyint(1) UNSIGNED  DEFAULT '0' NOT NULL,
    approved_ru    tinyint(1) UNSIGNED  DEFAULT '0' NOT NULL,
    approved_en    tinyint(1) UNSIGNED  DEFAULT '0' NOT NULL,
    published_ru   tinyint(1) UNSIGNED  DEFAULT '0' NOT NULL,
    published_en   tinyint(1) UNSIGNED  DEFAULT '0' NOT NULL,
    simple         tinyint(1) UNSIGNED  DEFAULT '0' NOT NULL,
    slug           varchar(100)                     NOT NULL,
    time           smallint(5) UNSIGNED DEFAULT '0',
    image          varchar(220)         DEFAULT 'default.jpg',
    created_at     datetime             DEFAULT CURRENT_TIMESTAMP,
    updated_at     datetime             DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY slug (slug)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;