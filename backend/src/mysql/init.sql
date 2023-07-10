CREATE DATABASE `fog` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;


USE fog;


CREATE TABLE `articles` (`id` bigint(20) NOT NULL AUTO_INCREMENT,
                          `created_at` int(11) NOT NULL,
                          `updated_at` int(11) default 0,
                          `title` varchar(255)  NOT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
