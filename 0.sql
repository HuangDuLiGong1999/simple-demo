create database douyin;
use douyin;
SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `user_info`;
DROP TABLE IF EXISTS `user_follower`;
DROP TABLE IF EXISTS `video_info`;
DROP TABLE IF EXISTS `comment_info`;
DROP TABLE IF EXISTS `video_favorite`;
SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `user_info` (
    `id` bigint(20) NOT NULL,
    `username` VARCHAR(20) NOT NULL,
    `password` VARCHAR(64) NOT NULL,
    `follow_count` INTEGER(11) NOT NULL,
    `follower_count` INTEGER(11) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `user_follower` (
    `id` bigint(20) NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `follower_id` VARCHAR(64) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `video_info` (
    `id` bigint(20) NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `play_url` VARCHAR(64) NOT NULL,
    `cover_url` VARCHAR(64) NOT NULL,
    `favorite_count` INTEGER(11) NOT NULL,
    `comment_count` INTEGER(11) NOT NULL,
    `publish_time` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `comment_info` (
    `id` bigint(20) NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `vedio_id` bigint(20) NOT NULL,
    `content` TEXT NOT NULL,
    `create_date` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `video_favorite` (
    `id` bigint(20) NOT NULL,
    `user_id` bigint(20) NOT NULL,
    `video_id` bigint(20) NOT NULL,
    PRIMARY KEY (`id`)
);
