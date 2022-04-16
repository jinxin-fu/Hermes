CREATE TABLE `hermesd`
(
    `aletname` varchar(255) NOT NULL COMMENT 'aletname',
    `aggeraterules` varchar(255) NOT NULL COMMENT 'aggerate rules',
    `receiveraddress` varchar(255) NOT NULL COMMENT 'receiver address',
    `returnvalueflag` varchar(255) NOT NULL COMMENT 'return value flag',
    PRIMARY KEY(`aletname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;