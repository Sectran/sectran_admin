/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80400
 Source Host           : 127.0.0.1:3306
 Source Schema         : sectran

 Target Server Type    : MySQL
 Target Server Version : 80400
 File Encoding         : 65001

 Date: 21/06/2024 15:43:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP DATABASE IF EXISTS `sectran`;
CREATE DATABASE `sectran` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `sectran`;

-- ----------------------------
-- Table structure for accounts
-- ----------------------------
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `username` varchar(16) COLLATE utf8mb4_bin NOT NULL COMMENT 'account username|账号名称',
  `port` int unsigned NOT NULL COMMENT 'account port|端口',
  `protocol` tinyint unsigned NOT NULL COMMENT 'protocol of the this account.|账号协议',
  `password` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'account password|账号密码',
  `private_key` varchar(4096) COLLATE utf8mb4_bin NOT NULL COMMENT 'private_key of the this account.|账号私钥',
  `device_id` bigint unsigned DEFAULT NULL COMMENT 'account belong to|账号所属设备',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of accounts
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for casbin_rules
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rules`;
CREATE TABLE `casbin_rules` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `ptype` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v0` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v1` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v2` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v3` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v4` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `v5` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=67 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of casbin_rules
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rules` VALUES (38, 'p', '1', '/department', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (39, 'p', '1', '/department/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (40, 'p', '1', '/department/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (41, 'p', '1', '/department/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (42, 'p', '1', '/department/children', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (43, 'p', '1', ':department', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (44, 'p', '1', ':user', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (45, 'p', '1', '/user', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (46, 'p', '1', '/user/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (47, 'p', '1', '/user/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (48, 'p', '1', '/user/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (49, 'p', '1', '/user/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (50, 'p', '1', ':device', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (51, 'p', '1', '/device', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (52, 'p', '1', '/device/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (53, 'p', '1', '/device/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (54, 'p', '1', '/device/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (55, 'p', '1', '/device/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (56, 'p', '1', '/account/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (57, 'p', '1', '/account/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (58, 'p', '1', '/account/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (59, 'p', '1', '/account/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (60, 'p', '1', '/role', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (61, 'p', '1', '/role/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (62, 'p', '1', '/role/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (63, 'p', '1', '/role/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (64, 'p', '1', '/role/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (65, 'p', '1', ':role', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (66, 'p', '1', '/department/delete', 'POST', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for departments
-- ----------------------------
DROP TABLE IF EXISTS `departments`;
CREATE TABLE `departments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'The name of the department.|部门名称',
  `area` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'The area where the department is located.|部门所在地区',
  `description` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'Description of the department.|部门描述',
  `parent_department_id` bigint unsigned NOT NULL COMMENT 'parent department ID.|父亲部门id',
  `parent_departments` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Comma-separated list of parent department IDs in ascending order.|上级部门集合逗号分隔升序排列',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of departments
-- ----------------------------
BEGIN;
INSERT INTO `departments` VALUES (1, '2024-06-21 07:34:28', '2024-06-21 07:34:28', '山川科技', '北京', '北京山川科技股份有限公司根部门', 9223372036854775806, '0');
COMMIT;

-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'The name of the device.|设备名称',
  `host` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'login host|设备地址',
  `type` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'type of the device.|设备类型',
  `description` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'Description of the device.|设备描述',
  `department_id` bigint unsigned DEFAULT NULL COMMENT 'ID of the device''s department.|设备所属部门',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of devices
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for lable_trees
-- ----------------------------
DROP TABLE IF EXISTS `lable_trees`;
CREATE TABLE `lable_trees` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'lable name|标签名称',
  `type` bigint unsigned NOT NULL COMMENT 'lable type|标签类型',
  `icon` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'lable icon|标签图标',
  `parent_lable` bigint unsigned NOT NULL COMMENT 'parent lable id|父标签id',
  `parent_lables` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'parent lables id,split by '',''|父标签id集合升序排列,逗号分隔',
  `lable_owner` bigint unsigned NOT NULL COMMENT 'lable owner,user ID|标签所属者,用户ID',
  `inherit` tinyint(1) NOT NULL COMMENT 'child lable can inherit parents|标签是否可以继承',
  `related_labels` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'related labels id,split by '',''|关联标签id集合升序排列,逗号分隔',
  `description` varchar(1024) COLLATE utf8mb4_bin NOT NULL COMMENT 'label description|标签描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of lable_trees
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `name` varchar(32) COLLATE utf8mb4_bin NOT NULL COMMENT 'The name of the role.|角色名称',
  `weight` bigint NOT NULL COMMENT 'The weight of the role. Smaller values indicate higher priority.|角色优先级，值越小优先级越高',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO `roles` VALUES (1, '2024-06-21 07:34:28', '2024-06-21 07:34:28', '开发者管理员', 0);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL COMMENT 'Create Time | 创建日期',
  `updated_at` timestamp NOT NULL COMMENT 'Update Time | 修改日期',
  `account` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'User account.|用户账号',
  `name` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'User name.|用户姓名',
  `password` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'User password.|用户密码',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'User status (enabled(true) or disabled(false)).|用户账号状态',
  `description` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'User description.|用户账号描述',
  `email` varchar(64) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'User email.|用户邮箱',
  `phone_number` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'User phone number.|用户手机号码',
  `department_id` bigint unsigned DEFAULT NULL COMMENT 'ID of the user''s department.|用户所属部门',
  `role_id` bigint unsigned DEFAULT NULL COMMENT 'ID of the user''s role.|用户所属角色',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, '2024-06-21 07:34:28', '2024-06-21 07:34:28', 'administrator', 'admin', '0okm)OKM', 1, NULL, NULL, NULL, 1, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
