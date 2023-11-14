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

 Date: 27/10/2023 15:17:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for st_dept
-- ----------------------------
DROP TABLE IF EXISTS `st_dept`;
CREATE TABLE `st_dept` (
  `dept_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `name` varchar(48) NOT NULL COMMENT '部门名称',
  `description` varchar(1024) DEFAULT NULL COMMENT '部门描述',
  `parent_id` int(11) NOT NULL COMMENT '上级部门ID',
  `child_ids` varchar(4096) NOT NULL COMMENT '下级部门ID集合，用逗号分隔',
  `create_by_uid` int(11) DEFAULT NULL COMMENT '创建者',
  `region` varchar(1024) NOT NULL COMMENT '部门所在地区',
  `create_time` datetime(6) DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for st_role
-- ----------------------------
DROP TABLE IF EXISTS `st_role`;
CREATE TABLE `st_role` (
  `role_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(11) NOT NULL COMMENT '角色名称',
  `description` varchar(1024) DEFAULT NULL COMMENT '角色描述',
  `create_by_uid` int(11) NOT NULL COMMENT '创建者',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for st_user
-- ----------------------------
DROP TABLE IF EXISTS `st_user`;
CREATE TABLE `st_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `account` varchar(255) NOT NULL COMMENT '用户账号',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `username` varchar(255) DEFAULT NULL COMMENT '用户姓名',
  `dept_id` int(11) NOT NULL COMMENT '用户所属部门ID',
  `disable` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '账号是否禁用',
  `description` varchar(255) DEFAULT NULL COMMENT '账号描述',
  `create_time` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `create_by_uid` int(11) NOT NULL COMMENT '创建人',
  `role_id` int(11) NOT NULL COMMENT '用户角色ID',
  `telephone` varchar(15) DEFAULT NULL COMMENT '用户手机号码',
  `email` varchar(100) DEFAULT NULL COMMENT '用户邮箱',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
