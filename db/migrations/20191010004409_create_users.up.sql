CREATE TABLE users
(
    id           int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    name         varchar(50)   DEFAULT NULL,
    status       varchar(250)  DEFAULT NULL,
    email        varchar(250)     NOT NULL,
    xp           smallint      DEFAULT '0',
    streak_days  int(8)        DEFAULT '0',
    popularity   decimal(8, 1) DEFAULT '0',
    active       tinyint(1)    DEFAULT '1',
    photo        varchar(200)  DEFAULT 'default.jpg',
    password     varchar(250)     NOT NULL,
    streak_check datetime      DEFAULT CURRENT_TIMESTAMP,
    notif_check  datetime      DEFAULT CURRENT_TIMESTAMP,
    online_check datetime      DEFAULT CURRENT_TIMESTAMP,
    created_at   datetime      DEFAULT CURRENT_TIMESTAMP,
    updated_at   datetime      DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY (email)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;