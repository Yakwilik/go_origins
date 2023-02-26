CREATE TABLE IF NOT EXISTS `users` (
    `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL UNIQUE ,
    `password_hash` VARCHAR(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;