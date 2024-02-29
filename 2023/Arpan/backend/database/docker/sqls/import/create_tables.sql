DROP TABLE IF EXISTS `todos`;
CREATE TABLE `todos` (
  `id` varchar(36) NOT NULL,
  `text` text,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
