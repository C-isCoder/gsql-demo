CREATE DATABASE db_gsql;
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee`
(
    `id`   int(6) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(30)     NOT NULL,
    `city` varchar(30)     not null,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;