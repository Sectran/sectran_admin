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

 Date: 22/10/2023 10:32:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for st_user
-- ----------------------------
DROP TABLE IF EXISTS `st_user`;
CREATE TABLE `st_user` (
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `account` varchar(255) NOT NULL COMMENT '用户账号',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `username` varchar(255) DEFAULT NULL COMMENT '用户姓名',
  `dept_id` int(11) NOT NULL COMMENT '用户所属部门ID',
  `disable` tinyint(1) NOT NULL COMMENT '账号是否禁用',
  `description` varchar(255) DEFAULT NULL COMMENT '账号描述',
  `create_time` datetime(6) NOT NULL COMMENT '创建时间',
  `create_by_uid` int(11) NOT NULL COMMENT '创建人',
  `is_delete` tinyint(1) NOT NULL COMMENT '是否被删除',
  `role_id` int(11) NOT NULL COMMENT '用户角色ID',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
