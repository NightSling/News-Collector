CREATE TABLE IF NOT EXISTS `news` (
    `uniqId` text NOT NULL,
    `sourceid` text DEFAULT NULL,
    `sourcename` text DEFAULT NULL,
    `author` text DEFAULT NULL,
    `title` text DEFAULT NULL,
    `description` text DEFAULT NULL,
    `url` text DEFAULT NULL,
    `urlToImage` text DEFAULT NULL,
    `publishedAt` text DEFAULT NULL,
    `content` text DEFAULT NULL,
    UNIQUE KEY `uniqId` (`uniqId`) USING HASH
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;