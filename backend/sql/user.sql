/*
使用者-列表
*/
CREATE TABLE IF NOT EXISTS `User` (
  ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `UName` VARCHAR(8) UNIQUE KEY NOT NULL,
  `UPassword` VARCHAR(8) NOT NULL,
  `UNickname` VARCHAR(16) NOT NULL,
  `UKey` VARCHAR(32) NOT NULL,
  `ULv` VARCHAR(1) DEFAULT '1',
  LastIP VARCHAR(15) DEFAULT '',
  CreatedAt DATETIME NOT NULL
);