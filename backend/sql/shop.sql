/*
店家-列表
*/
CREATE TABLE IF NOT EXISTS `Shop` (
    `UUID` VARCHAR(36) UNIQUE KEY NOT NULL,
    `Name` VARCHAR(16) UNIQUE KEY NOT NULL,
    `Mobile` VARCHAR(16) UNIQUE KEY NOT NULL,
    `Actived` TINYINT UNSIGNED DEFAULT 1
);