SET NAMES utf8;
# SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

CREATE DATABASE redditclone;
CREATE USER 'service_admin' IDENTIFIED BY 'qwerty';

GRANT ALL PRIVILEGES ON redditclone.* TO 'service_admin';

USE redditclone;


