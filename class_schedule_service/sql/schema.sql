CREATE TABLE `classes` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `date` timestamp NOT NULL,
    `order` int NOT NULL,
    `subject` int unsigned NOT NULL,
    `teacher` varchar(32) NOT NULL,
    `group` varchar(64) NOT NULL,
    `class_subject` text,
    PRIMARY KEY (`id`)
);