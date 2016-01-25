CREATE DATABASE IF NOT EXISTS `slugapp`;
USE slugapp;
CREATE TABLE IF NOT EXISTS `event` (
    PRIMARY KEY (`id`),
    `name` VARCHAR(255),
    `description` VARCHAR(255)
    `date` TIMESTAMP
);
