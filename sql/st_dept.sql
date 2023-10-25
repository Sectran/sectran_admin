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

 Date: 22/10/2023 10:31:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for st_dept
-- ----------------------------
DROP TABLE IF EXISTS `st_dept`;
CREATE TABLE `st_dept` (
  `dept_id` int(11) NOT NULL COMMENT '部门ID',
  `name` varchar(48) NOT NULL COMMENT '部门名称',
  `description` varchar(1024) DEFAULT NULL COMMENT '部门描述',
  `parent_id` int(11) NOT NULL COMMENT '上级部门ID',
  `child_ids` varchar(4096) NOT NULL COMMENT '下级部门ID集合，用逗号分隔',
  `create_by_uid` int(11) DEFAULT NULL COMMENT '创建者',
  `region` varchar(1024) NOT NULL COMMENT '部门所在地区',
  `is_delete` tinyint(1) NOT NULL COMMENT '是否被删除',
  PRIMARY KEY (`dept_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
