DROP DATABASE IF EXISTS `mental_health`;

CREATE DATABASE `mental_health`;

/*CREATE DATABASE IF NOT EXISTS `mental_health`;*/

USE `mental_health`;

CREATE TABLE `user` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `sid`        VARCHAR(10)  NOT NULL UNIQUE COMMENT "学生学号",
  `username`   VARCHAR(25)  NOT NULL ,
  `is_teacher`  TINYINT(1)   NOT NULL  DEFAULT 0 COMMENT "0为学生,1为老师",
  `avatar`     VARCHAR(255),
  `introduction`  VARCHAR(255)  COMMENT "个性签名",
  `phone`  VARCHAR(50),
  `back_avatar`  VARCHAR(255) COMMENT "个人主页照片",

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `mood` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `time`   DATETIME    NOT NULL    COMMENT "记录时间",
  `star`   TINYINT(1)   NOT NULL    COMMENT "星数评级",
  `notes` VARCHAR(50) COMMENT "心情记录",

  `user_id` INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `hole` (
  `id`                    INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `hole_name`           VARCHAR(50)  NOT NULL,
  `content`               TEXT           COMMENT "问题内容",
  `like_num`              INT          NOT NULL DEFAULT 0 COMMENT "点赞数",
  `comment_num`           INT          NOT NULL DEFAULT 0 COMMENT "一级评论数",
  `look_num`           INT          NOT NULL DEFAULT 0 COMMENT "浏览数",
  `type`                  TINYINT(1)               COMMENT "0/1/2/ 环境适应，人际关系，学业学习",
  `time`                  DATETIME     NOT NULL           COMMENT "发布时间",

  `user_id`               INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `parent_comment` (
  `id`              VARCHAR(40) NOT NULL           COMMENT "uuid",
  `time`            DATETIME    NOT NULL           COMMENT "评论时间",
  `content`         TEXT                           COMMENT "评论内容",
  `sub_comment_num` INT         NOT NULL DEFAULT 0 COMMENT "子评论数",

  `user_id`         INT UNSIGNED NOT NULL,
  `hole_id`   INT UNSIGNED NOT NULL COMMENT "问题id",

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `hole_id` (`hole_id`)
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

CREATE TABLE `hole_like` (
  `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `hole_id` INT UNSIGNED NOT NULL COMMENT "问题id",
  `user_id`       INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `hole_id` (`hole_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `comment_like` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `comment_id` VARCHAR(40)  NOT NULL COMMENT "评论id",
  `user_id`    INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `comment_id` (`comment_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

INSERT INTO `user` VALUES(0,2018212691,'hjm','0','asd','','','');