CREATE DATABASE IF NOT EXISTS `mental_health`;

USE `mental_health`;

CREATE TABLE `user` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `sid`        VARCHAR(10)  NOT NULL UNIQUE COMMENT "学生学号",
  `username`   VARCHAR(25)  NOT NULL ,
  `avatar`     VARCHAR(255),
  `introduction`  VARCHAR(100) , COMMENT "个性签名" ,
  `phone`  VARCHAR(50) ,
  `back_avatar`  VARCHAR(255) COMMENT "个人主页照片"  ,
  `is_teacher`  TINYINT(1)   NOT NULL  COMMENT "0为学生，1为老师",

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `mood` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `time`   DATETIME    NOT NULL    COMMENT "记录时间",
  `star`   TINYINT(1)   NOT NULL    COMMENT "星数评级",
  `notes` VARCHAR(50) COMMENT "心情记录",

  `user_id` INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`)
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `problem` (
  `id`                    INT UNSIGNED NOT NULL AUTO_INCREMENT,

  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `parent_comment` (
  `id`              VARCHAR(40) NOT NULL           COMMENT "uuid",
  `time`            DATETIME    NOT NULL           COMMENT "评论时间",
  `content`         TEXT                           COMMENT "评论内容",
  `sub_comment_num` INT         NOT NULL DEFAULT 0 COMMENT "子评论数",

  `user_id`         INT UNSIGNED NOT NULL,
  `problem_id`   INT UNSIGNED NOT NULL COMMENT "问题id",

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `problem_id` (`problem_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `sub_comment` (
  `id`           VARCHAR(40) NOT NULL           COMMENT "uuid",
  `time`         DATETIME    NOT NULL           COMMENT "评论时间",
  `content`      TEXT                           COMMENT "评论内容",

  `parent_id`      VARCHAR(40) NOT NULL,
  `user_id`        INT UNSIGNED  NOT NULL,
  `target_user_id` INT UNSIGNED  NOT NULL COMMENT "评论的目标用户id",

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `parent_id` (`parent_id`),
  KEY `target_user_id` (`target_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

INSERT INTO `user` VALUES(0,2018212691,'hjm','asd');