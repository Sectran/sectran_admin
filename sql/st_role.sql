/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : 127.0.0.1:3306
 Source Schema         : sectran

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 22/10/2023 10:32:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for st_role
-- ----------------------------
DROP TABLE IF EXISTS `st_role`;
CREATE TABLE `st_role` (
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `name` varchar(11) NOT NULL COMMENT '角色名称',
  `description` varchar(1024) DEFAULT NULL COMMENT '角色描述',
  `create_by_uid` int(11) NOT NULL COMMENT '创建者',
  `is_delete` tinyint(1) NOT NULL COMMENT '是否被删除',
  `create_time` datetime(6) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
