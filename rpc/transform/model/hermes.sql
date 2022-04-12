CREATE TABLE `hermesd`
(
  `hermes` varchar(255) NOT NULL COMMENT 'hermes key',
  `url` varchar(255) NOT NULL COMMENT 'original url',
  PRIMARY KEY(hermes)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
