DROP DATABASE IF EXISTS `mental_health`;

CREATE DATABASE `mental_health`;

/*CREATE DATABASE IF NOT EXISTS `mental_health`;*/

USE `mental_health`;

CREATE TABLE `user` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `sid`        VARCHAR(10)  NOT NULL UNIQUE COMMENT "学生学号",
  `username`   VARCHAR(25)  NOT NULL ,
  `is_teacher`  BOOLEAN   NOT NULL  DEFAULT false ,
  `avatar`     VARCHAR(255),
  `introduction`  VARCHAR(255)  COMMENT "个性签名",
  `phone`  VARCHAR(50),
  `back_avatar`  VARCHAR(255) COMMENT "个人主页照片",

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `mood` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `date`   VARCHAR(50) NOT NULL COMMENT "记录时间(2020.08.02)",
  `year`   INT    NOT NULL    COMMENT "记录时间(年)",
  `month`   INT    NOT NULL    COMMENT "记录时间(月)",
  `day`   INT    NOT NULL    COMMENT "记录时间(日)",
  `score`   TINYINT(1)   NOT NULL    COMMENT "星数评级",
  `note` VARCHAR(255) COMMENT "心情记录",

  `user_id` INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `hole` (
  `id`                    INT UNSIGNED   NOT NULL AUTO_INCREMENT,
  `name`           VARCHAR(255)     NOT NULL,
  `content`      VARCHAR(255)     NOT NULL COMMENT "问题内容",
  `like_num`              INT                 NOT NULL DEFAULT 0 COMMENT "点赞数",
  `favorite_num`  INT                    NOT NULL DEFAULT 0 COMMENT "收藏数",
  `comment_num`           INT       NOT NULL DEFAULT 0 COMMENT "一级评论数",
  `read_num`           INT                  NOT NULL DEFAULT 0 COMMENT "浏览数",
  `type`                  TINYINT(1)        NOT NULL COMMENT "1/2/3/4/5分别为 环境适应、人际关系、学业学习、生活经济、求职择业",
  `time`                  DATETIME         NOT NULL COMMENT "发布时间",

  `user_id`               INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  FULLTEXT KEY (`name`,`content`) WITH PARSER ngram
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `parent_comment` (
  `id`                    INT UNSIGNED   NOT NULL AUTO_INCREMENT,
  /*`id`              VARCHAR(40) NOT NULL           COMMENT "uuid",*/
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
  `id`                    INT UNSIGNED   NOT NULL AUTO_INCREMENT,
  `time`         DATETIME    NOT NULL           COMMENT "评论时间",
  `content`      TEXT                           COMMENT "评论内容",

  `parent_id`      INT UNSIGNED NOT NULL,
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

CREATE TABLE `hole_favorite` (
  `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `hole_id` INT UNSIGNED NOT NULL COMMENT "问题id",
  `user_id`       INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `hole_id` (`hole_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `hole_read` (
  `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `hole_id` INT UNSIGNED NOT NULL COMMENT "问题id",
  `user_id`       INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `hole_id` (`hole_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `comment_like` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `comment_id` INT UNSIGNED  NOT NULL COMMENT "评论id",
  `user_id`    INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `comment_id` (`comment_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `course` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `url`    VARCHAR(255) NOT NULL  COMMENT "视频地址",
  `name`  VARCHAR(255) NOT NULL,
  `source` VARCHAR(255) NOT NULL,
  `summary` VARCHAR(255) NOT NULL,
  `like_num`              INT                 NOT NULL DEFAULT 0 COMMENT "点赞数",
  `favorite_num`  INT                    NOT NULL DEFAULT 0 COMMENT "收藏数",
  `watch_num`  INT                    NOT NULL DEFAULT 0 COMMENT "收藏数",
  `time`                  DATETIME         NOT NULL COMMENT "发布时间",

  PRIMARY KEY (`id`),
  FULLTEXT KEY (`name`, `source`) WITH PARSER ngram
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `course_like` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `course_id` INT UNSIGNED  NOT NULL COMMENT "课程id",
  `user_id`    INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `course_id` (`course_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `course_favorite` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `course_id` INT UNSIGNED  NOT NULL COMMENT "课程id",
  `user_id`    INT UNSIGNED NOT NULL,

  PRIMARY KEY (`id`),
  KEY `course_id` (`course_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

/*
INSERT INTO `user` VALUES(0,2018212691,'hjm','0','asd','','','');
*/

-- mock data
INSERT INTO `user` (sid, username, is_teacher) VALUES ('2018212691', 'Hjm1027',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('1234568890', '随便',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('3787546378', '不知道取什么好',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('1047395326', 'Wow, IGNB',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('9247128475', '信息管理学院',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('7204901939', 'GITHUB',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('8705469760', '中华人民共和国湖北省武汉市',  0);

INSERT INTO `hole` (name, content, type,time,user_id) VALUES ('第一个问题', '作业好多写不完',  1,'2020-08-04 09:16:50',1);
INSERT INTO `hole` (name, content, type,time,user_id) VALUES ('水', '经验+3，告辞',  3,'2019-01-01 14:18:2',1);
INSERT INTO `hole` (name, content, type,time,user_id) VALUES ('哈哈哈哈哈', '我又来水了',  5,'2020-07-03 22:27:01',1);
INSERT INTO `hole` (name, content, type,time,user_id) VALUES ('过年啦', '到2020了',  2,'2020-01-01 00:00:01',2);
INSERT INTO `hole` (name, content, type,time,user_id) VALUES ('不许水评论！', '小心封号',  2,'2020-08-02 23:56:59',5);

INSERT INTO `course` (name,url, source, summary,time) VALUES ('自信培养', 'www.baidu.com','CCNU心理站','培养自信','2018-06-07 12:56:01');
INSERT INTO `course` (name,url, source, summary,time) VALUES ('心理与生活', 'www.google.com','2级心理站','大致介绍','2019-11-30 07:23:18');
INSERT INTO `course` (name,url, source, summary,time) VALUES ('谈话的艺术', 'www.bing.com','心理健康中心','如何谈话','2020-02-12 19:09:22');
INSERT INTO `course` (name,url, source, summary,time) VALUES ('发展心理学', 'www.asjhjesh.com','校医院','心理学史','2020-07-12 23:18:00');