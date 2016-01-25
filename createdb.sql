CREATE DATABASE IF NOT EXISTS `slugapp`;
USE slugapp;
CREATE TABLE IF NOT EXISTS `event` (
    PRIMARY KEY (`id`),
    `name` varchar(255),
    `description` varchar(255)
    `date` TIMESTAMP
);
