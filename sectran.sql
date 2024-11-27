-- MySQL dump 10.13  Distrib 8.4.0, for Linux (aarch64)
--
-- Host: localhost    Database: sectran
-- ------------------------------------------------------
-- Server version	8.4.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `sectran`
--

/*!40000 DROP DATABASE IF EXISTS `sectran`*/;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `sectran` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `sectran`;

--
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `username` varchar(16) COLLATE utf8mb4_bin NOT NULL COMMENT 'account username|账号名称',
  `port` int unsigned NOT NULL COMMENT 'account port|端口',
  `protocol` tinyint unsigned NOT NULL COMMENT 'protocol of the this account.|账号协议',
  `password` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'account password|账号密码',
  `private_key` varchar(4096) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'private_key of the this account.|账号私钥',
  `private_key_password` varchar(4096) COLLATE utf8mb4_bin DEFAULT NULL COMMENT 'private_key password of the this account.|私钥口令',
  `device_id` bigint unsigned NOT NULL COMMENT 'account belong to|账号所属设备',
  `department_id` bigint unsigned NOT NULL COMMENT 'account belong to|账号所属部门',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `casbin_rules`
--

DROP TABLE IF EXISTS `casbin_rules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rules`
--

LOCK TABLES `casbin_rules` WRITE;
/*!40000 ALTER TABLE `casbin_rules` DISABLE KEYS */;
INSERT INTO `casbin_rules` VALUES (1, 'p', '1', '/department', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (2, 'p', '1', '/department/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (3, 'p', '1', '/department/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (4, 'p', '1', '/department/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (5, 'p', '1', '/department/children', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (6, 'p', '1', ':department', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (7, 'p', '1', ':role', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (8, 'p', '1', '/get_menu_authority_list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (9, 'p', '1', '/file/upload', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (10, 'p', '1', '/file/download', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (11, 'p', '1', ':user', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (12, 'p', '1', '/user', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (13, 'p', '1', '/user/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (14, 'p', '1', '/user/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (15, 'p', '1', '/user/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (16, 'p', '1', '/user/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (17, 'p', '1', '/role', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (18, 'p', '1', '/role/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (19, 'p', '1', '/role/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (20, 'p', '1', '/role/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (21, 'p', '1', '/role/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (22, 'p', '1', '/update_authority_api', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (23, 'p', '1', '/department/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (24, 'p', '1', ':device', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (25, 'p', '1', '/device', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (26, 'p', '1', '/device/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (27, 'p', '1', '/device/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (28, 'p', '1', '/device/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (29, 'p', '1', '/device/delete', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (30, 'p', '1', '/account/list', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (31, 'p', '1', '/account/create', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (32, 'p', '1', '/account/update', 'POST', '', '', '');
INSERT INTO `casbin_rules` VALUES (33, 'p', '1', '/account/delete', 'POST', '', '', '');
/*!40000 ALTER TABLE `casbin_rules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `departments`
--

DROP TABLE IF EXISTS `departments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `departments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `name` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'The name of the department.|部门名称',
  `area` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'The area where the department is located.|部门所在地区',
  `description` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'Description of the department.|部门描述',
  `parent_department_id` bigint unsigned NOT NULL COMMENT 'parent department ID.|父亲部门id',
  `parent_departments` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'Comma-separated list of parent department IDs in ascending order.|上级部门集合逗号分隔升序排列',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `departments`
--

LOCK TABLES `departments` WRITE;
/*!40000 ALTER TABLE `departments` DISABLE KEYS */;
INSERT INTO `departments` VALUES (1,'2024-11-27 13:20:23','2024-11-27 13:20:23','sectran','Beijing','sectran admin root department',0,'0');
/*!40000 ALTER TABLE `departments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `devices`
--

DROP TABLE IF EXISTS `devices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `devices` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `name` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'The name of the device.|设备名称',
  `host` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'login host|设备地址',
  `type` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'type of the device.|设备类型',
  `description` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT 'Description of the device.|设备描述',
  `department_id` bigint unsigned NOT NULL COMMENT 'ID of the device''s department.|设备所属部门',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `devices`
--

LOCK TABLES `devices` WRITE;
/*!40000 ALTER TABLE `devices` DISABLE KEYS */;
/*!40000 ALTER TABLE `devices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lable_trees`
--

DROP TABLE IF EXISTS `lable_trees`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lable_trees` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `name` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT 'lable name|标签名称',
  `type` bigint unsigned NOT NULL COMMENT 'lable type|标签类型（分组标签、控制标签、授权标签）',
  `icon` varchar(32) COLLATE utf8mb4_bin NOT NULL COMMENT 'lable icon|标签图标',
  `content` varchar(1024) COLLATE utf8mb4_bin NOT NULL COMMENT 'lable content|标签内容',
  `ownership` tinyint unsigned NOT NULL COMMENT 'lable ownership Level (Department Level/User Level)|标签所有权级别（部门级别/用户级别）',
  `owner_id` bigint unsigned NOT NULL COMMENT 'lable owner,user ID,dept ID|标签所属者,用户ID,部门ID',
  `parent_id` bigint unsigned NOT NULL COMMENT 'parent lable id|父标签id',
  `description` varchar(256) COLLATE utf8mb4_bin NOT NULL COMMENT 'label description|标签描述',
  `target_type` smallint unsigned NOT NULL COMMENT 'lable target type|标签目标类型',
  `parents` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'parent lables id,split by '','',lable tree deep cannot gather than 32|父标签id集合升序排列,逗号分隔,限制最多不可超过64级',
  `inherit` tinyint(1) NOT NULL COMMENT 'child lable can inherit parents|标签是否可以继承',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lable_trees`
--

LOCK TABLES `lable_trees` WRITE;
/*!40000 ALTER TABLE `lable_trees` DISABLE KEYS */;
/*!40000 ALTER TABLE `lable_trees` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `name` varchar(32) COLLATE utf8mb4_bin NOT NULL COMMENT 'The name of the role.|角色名称',
  `weight` bigint NOT NULL COMMENT 'The weight of the role. Smaller values indicate higher priority.|角色优先级，值越小优先级越高',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'2024-11-27 13:20:23','2024-11-27 13:20:23','开发者',0);
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2024-11-27 13:20:23','2024-11-27 13:20:23','administrator','开发者管理员','Passwordryan@0',1,NULL,NULL,NULL,1,1);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-11-27 13:24:02
