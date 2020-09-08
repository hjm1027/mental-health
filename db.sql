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
  `watch_num`  INT                    NOT NULL DEFAULT 0 COMMENT "观看数",
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


CREATE TABLE `message` (
  `id`          INT UNSIGNED NOT NULL auto_increment,
  `pub_user_id` INT UNSIGNED NOT NULL DEFAULT 0,
  `sub_user_id` INT UNSIGNED NOT NULL DEFAULT 0,
  `kind`        TINYINT(1) UNSIGNED   NOT NULL DEFAULT 0  COMMENT "消息提醒的种类，0是点赞，1是收藏，2是评论",
  `is_read`     TINYINT(1)   NOT NULL DEFAULT 0,
  `reply`       VARCHAR(255),
  `time`        DATETIME  NOT NULL,
  `hole_id` INT UNSIGNED,
  `content`     VARCHAR(255),
  `sid`         VARCHAR(255),
  `parent_comment_id`     VARCHAR(255),

  PRIMARY KEY (`id`),
  KEY sub_user_id (`sub_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `poster` (
  `id`          INT UNSIGNED NOT NULL auto_increment,
  `home` VARCHAR(255) ,
  `platform`         VARCHAR(255),
  `hole`     VARCHAR(255),

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `reserve` (
  `id`          INT UNSIGNED NOT NULL auto_increment COMMENT "id最大值固定，无限覆盖记录",
  `weekday` TINYINT(1) UNSIGNED   NOT NULL COMMENT "1-7代表周一到周日，这七天永远都是未来的七天", 
  `schedule` TINYINT(1) UNSIGNED NOT NULL COMMENT "1-6，一天的六个时间段",
  `teacher`     VARCHAR(255) NOT NULL COMMENT "在这个时间段值班的老师，此字段通常不会变化",
  `teacher_id`     VARCHAR(20) NOT NULL COMMENT "老师的用户id",
  `reserve`     TINYINT(1) NOT NULL COMMENT "预约状态，0/1/2为 可预约/审核中/预约成功",
  `time`     DATETIME  COMMENT "这个时间段上一次提交预约的时间",
  `advance_time`  TINYINT(1) NOT NULL DEFAULT 0 COMMENT "提前x天预约，2<=x<=8",
  `type`     TINYINT(1) NOT NULL  DEFAULT 0 COMMENT "预约类别，1-6为环境适应，人际关系，学业学习，生活经济，求职择业，其他",
  `method`     TINYINT(1) NOT NULL DEFAULT 0 COMMENT "0/1=线上预约/线下预约",

  `user_id`    INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "提交预约的用户id",

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `teacher_id` (`teacher_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `record` (
  `id`          INT UNSIGNED NOT NULL auto_increment ,
  `teacher`     VARCHAR(255) NOT NULL COMMENT "预约的老师",
  `time`     DATETIME  NOT NULL,
  `type`     TINYINT(1) NOT NULL COMMENT "0/1/2为发起预约/接受预约/拒绝预约",

  `user_id`    INT UNSIGNED NOT NULL ,

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE `test` (
  `id`          INT UNSIGNED NOT NULL auto_increment ,
  `url`     VARCHAR(255) NOT NULL,
  `header`     VARCHAR(255)  NOT NULL,
  `content`     VARCHAR(255) NOT NULL,

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

/*
INSERT INTO `user` VALUES(0,2018212691,'hjm','0','asd','','','');
*/

-- mock data
INSERT INTO `user` (sid, username, is_teacher) VALUES ('2018212691', 'Hjm1027',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('1234568890', '随便',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('3787546378', '不知道取什么好',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('1047395326', 'Wow, IGNB',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('9247128475', '信息管理学院',  1);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('7204901939', 'GITHUB',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('8705469760', '中华人民共和国湖北省武汉市',  0);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('0000000000', '这是个老师',  1);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('0111111110', '这也是老师',  1);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('0122222210', '这还是老师',  1);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('0123333210', '又一个老师',  1);
INSERT INTO `user` (sid, username, is_teacher) VALUES ('0123443210', '最后的老师',  1);

INSERT INTO `hole` (name, content, comment_num,type,time,user_id) VALUES ('第一个问题', '作业好多写不完',3,  1,'2020-08-04 09:16:50',1);
INSERT INTO `hole` (name, content, type,time,user_id) VALUES ('水', '经验+3，告辞',  3,'2019-01-01 14:18:2',1);
INSERT INTO `hole` (name, content, type,time,user_id) VALUES ('哈哈哈哈哈', '我又来水了',  5,'2020-07-03 22:27:01',1);
INSERT INTO `hole` (name, content, type,time,user_id) VALUES ('过年啦', '到2020了',  2,'2020-01-01 00:00:01',2);
INSERT INTO `hole` (name, content, type,time,user_id) VALUES ('不许水评论！', '小心封号',  2,'2020-08-02 23:56:59',5);

INSERT INTO `parent_comment` (time, content, sub_comment_num,user_id,hole_id) VALUES ('2019-06-04 09:16:50', '我也是',  0,2,1);
INSERT INTO `parent_comment` (time, content, sub_comment_num,user_id,hole_id) VALUES ('2020-01-09 18:25:12', '可以试试**方法',  1,3,1);
INSERT INTO `parent_comment` (time, content, sub_comment_num,user_id,hole_id) VALUES ('2020-07-01 09:16:50', '+1',  0,4,1);

INSERT INTO `sub_comment` (time, content, parent_id,user_id,target_user_id) VALUES ('2020-08-02 22:51:02', '果然效果显著',  2,6,3);

INSERT INTO `course` (name,url, source, summary,time) VALUES ('自信培养', 'www.baidu.com','CCNU心理站','培养自信','2018-06-07 12:56:01');
INSERT INTO `course` (name,url, source, summary,time) VALUES ('心理与生活', 'www.google.com','2级心理站','大致介绍','2019-11-30 07:23:18');
INSERT INTO `course` (name,url, source, summary,time) VALUES ('谈话的艺术', 'www.bing.com','心理健康中心','如何谈话','2020-02-12 19:09:22');
INSERT INTO `course` (name,url, source, summary,time) VALUES ('发展心理学', 'www.asjhjesh.com','校医院','心理学史','2020-07-12 23:18:00');
INSERT INTO `course` (name,url, source, summary,time) VALUES ('yugbyusfb', 'www.asd.com','zxcsdf','ytjtrherg','2020-08-07 19:09:22');
INSERT INTO `course` (name,url, source, summary,time) VALUES ('21443242', 'www.6575.com','123123','98066566','2020-08-12 23:18:00');
INSERT INTO `course` (name,url, source, summary,time) VALUES ('$^**((^$', 'www.((*)).com','#%^','!@@@!!','2020-08-12 23:18:00');

INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2019.11.13', 2019,11,13,5,'erg', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2019.11.16', 2019,11,16,2,'erg2', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2019.11.21', 2019,11,21,3,'erg3', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2019.11.27', 2019,11,27,1,'erg4', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2019.12.09', 2019,12,09,4,'asd', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2019.12.25', 2019,12,25,3,'asd2', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.07.02', 2020,7,2,5,'zxc', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.07.04', 2020,7,4,1,'zxc2', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.07.07', 2020,7,7,4,'zxc3', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.07.14', 2020,7,14,2,'zxc4', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.07.15', 2020,7,15,1,'zxc5', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.07.18', 2020,7,18,3,'zxc6', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.07.22', 2020,7,22,5,'zxc7', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.07.28', 2020,7,28,4,'zxc8', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.08.13', 2020,8,13,5,'还可以', 1);
INSERT INTO `mood` (date, year, month,day,score,note,user_id) VALUES ('2020.08.14', 2020,8,14,4,'今日心情测试', 1);

-- true data

INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (1, 1,'第1个老师', 0,10001);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (1, 2,'第2个老师', 1,10002);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (1, 3,'第3个老师', 2,10003);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (1, 4,'第4个老师', 0,10004);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (1, 5,'第5个老师', 1,10005);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (1, 6,'第6个老师', 2,10006);

INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (2, 1,'第7个老师', 2,10007);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (2, 2,'第8个老师', 0,10008);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (2, 3,'第9个老师', 1,10009);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (2, 4,'第10个老师', 2,10010);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (2, 5,'第11个老师', 0,10011);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (2, 6,'第12个老师', 1,10012);

INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (3, 1,'第13个老师', 1,10013);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (3, 2,'第14个老师', 2,10014);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (3, 3,'第15个老师', 0,10015);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (3, 4,'第16个老师', 1,10016);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (3, 5,'第17个老师', 2,10017);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (3, 6,'第18个老师', 0,10018);

INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (4, 1,'第19个老师', 0,10019);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (4, 2,'第20个老师', 0,10020);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (4, 3,'第21个老师', 0,10021);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (4, 4,'第22个老师', 0,10022);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (4, 5,'第23个老师', 0,10023);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (4, 6,'第24个老师', 0,10024);

INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (5, 1,'第25个老师', 0,10025);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (5, 2,'第26个老师', 0,10026);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (5, 3,'第27个老师', 0,10027);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (5, 4,'第28个老师', 0,10028);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (5, 5,'第29个老师', 0,10029);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (5, 6,'第30个老师', 0,10030);

INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (6, 1,'第31个老师', 0,10031);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (6, 2,'第32个老师', 0,10032);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (6, 3,'第33个老师', 0,10033);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (6, 4,'第34个老师', 0,10034);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (6, 5,'第35个老师', 0,10035);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (6, 6,'第36个老师', 0,10036);

INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (7, 1,'第37个老师', 0,10037);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (7, 2,'第38个老师', 0,10038);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (7, 3,'第39个老师', 0,10039);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (7, 4,'第40个老师', 0,10040);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (7, 5,'第41个老师', 0,10041);
INSERT INTO `reserve` (weekday, schedule, teacher,reserve,teacher_id) VALUES (7, 6,'第42个老师', 0,10042);
