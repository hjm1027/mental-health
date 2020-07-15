CREATE DATABASE IF NOT EXISTS `mental_health`;

USE `mental_health`;

CREATE TABLE `user` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `sid`        VARCHAR(10)  NOT NULL UNIQUE COMMENT "学生学号",
  `username`   VARCHAR(25)  ,
  `avatar`     VARCHAR(255) ,

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

INSERT INTO `user` VALUES(0,2018212691,'hjm','asd');