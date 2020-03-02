/*
 Navicat Premium Data Transfer

 Source Server         : third-party-service
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3308
 Source Schema         : kendo

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 02/03/2020 09:50:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `account` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `expire_at` date DEFAULT NULL,
  `phone_number` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `deleted_at` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'elsa', 'admin', '12345', '2020-02-29', '0972889009', '2020-02-26', '2020-02-26', NULL);
INSERT INTO `users` VALUES (2, '1234', '1114', '', '2020-02-26', '', '2020-02-26', '2020-02-27', NULL);
INSERT INTO `users` VALUES (3, 'winston', 'winston', 'winston', '2021-02-26', '0912345678', '2020-02-26', '2020-02-27', NULL);
INSERT INTO `users` VALUES (4, 'amy', 'amy', 'amy', '2020-02-26', '', '2020-02-26', '2020-02-26', '2020-02-27');
INSERT INTO `users` VALUES (5, 'GGSU', 'GGSU', 'GGSU', '2020-02-26', '', '2020-02-26', '2020-02-26', '2020-02-27');
INSERT INTO `users` VALUES (6, 'peggy', 'peggy', '12345', '2020-02-27', '', '2020-02-27', '2020-02-27', NULL);
INSERT INTO `users` VALUES (7, 'angel', 'angel', '111', '2020-02-27', '', '2020-02-27', '2020-02-27', NULL);
INSERT INTO `users` VALUES (8, 'test', 'test', 'test', '2020-02-27', '', '2020-02-27', '2020-02-27', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
