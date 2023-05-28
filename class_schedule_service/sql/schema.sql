CREATE TABLE `classes` (
       `id` bigint NOT NULL AUTO_INCREMENT,
       `date` timestamp NOT NULL,
       `order` int NOT NULL,
       `subject` int unsigned NOT NULL,
       `teacher` varchar(32) NOT NULL,
       `group` varchar(64) NOT NULL,
       `class_subject` text,
       PRIMARY KEY (`id`),
       UNIQUE KEY `classes_UN` (`date`,`group`)
);

CREATE TABLE `attendance` (
      `class_id` bigint NOT NULL,
      `user_uid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
      `value` int NOT NULL,
      UNIQUE KEY `attendance_UN` (`user_uid`),
      KEY `attendance_FK` (`class_id`),
      CONSTRAINT `attendance_FK` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`)
);