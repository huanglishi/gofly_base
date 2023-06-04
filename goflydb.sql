/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : goflydb

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 24/05/2023 10:50:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gofly_article
-- ----------------------------
DROP TABLE IF EXISTS `gofly_article`;
CREATE TABLE `gofly_article`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cid` int(11) NOT NULL DEFAULT 0 COMMENT '分类',
  `title` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `des` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '简述',
  `subtitle` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '副标题',
  `keywords` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '关键词',
  `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '跳转地址',
  `image` varchar(145) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '缩略图',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `istop` int(11) NOT NULL DEFAULT 0 COMMENT '置顶',
  `visits` int(11) NOT NULL DEFAULT 0 COMMENT '访问数量',
  `author` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'Admin' COMMENT '作者',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '权重',
  `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '状态0=正常，1=隐藏',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gofly_article
-- ----------------------------
INSERT INTO `gofly_article` VALUES (1, 1, 'ChatGpt', '人工智能ChteGPT', ' ', ' ', ' ', ' ', '美国就科技', 0, 1, 'Admin', 1, 0, 1667827895);

-- ----------------------------
-- Table structure for gofly_article_cate
-- ----------------------------
DROP TABLE IF EXISTS `gofly_article_cate`;
CREATE TABLE `gofly_article_cate`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL COMMENT '父级',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `weigh` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gofly_article_cate
-- ----------------------------
INSERT INTO `gofly_article_cate` VALUES (1, 0, '科技', 1);
INSERT INTO `gofly_article_cate` VALUES (2, 0, '体育', 2);

SET FOREIGN_KEY_CHECKS = 1;
