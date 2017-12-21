/*
Navicat MySQL Data Transfer

Source Server         : apiadmin
Source Server Version : 50536
Source Host           : 192.168.134.128:3306
Source Database       : ppgo_api_admin

Target Server Type    : MYSQL
Target Server Version : 50536
File Encoding         : 65001

Date: 2017-12-21 18:07:31
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `pp_api_detail`
-- ----------------------------
DROP TABLE IF EXISTS `pp_api_detail`;
CREATE TABLE `pp_api_detail` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `source_id` int(11) NOT NULL DEFAULT '0' COMMENT '主表ID',
  `method` tinyint(1) NOT NULL DEFAULT '1' COMMENT '方法名称：1-GET 2-POST 3-PUT 4-PATCH 5-DELETE',
  `api_name` varchar(100) NOT NULL DEFAULT '0' COMMENT '接口名称',
  `api_url` varchar(100) NOT NULL DEFAULT '0' COMMENT '接口地址',
  `protocol_type` varchar(20) NOT NULL DEFAULT '1' COMMENT '协议类型，1-http,2-https',
  `result` text COMMENT '返回结果，正确或错误',
  `example` text COMMENT '接口示例',
  `detail` varchar(1000) NOT NULL DEFAULT '0' COMMENT '注意事项',
  `audit_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '审核时间',
  `audit_id` int(11) NOT NULL DEFAULT '0' COMMENT '审核人',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态：1-正在审核，0-暂停使用，2-审核通过',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`source_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='API附表';

-- ----------------------------
-- Records of pp_api_detail
-- ----------------------------
INSERT INTO `pp_api_detail` VALUES ('1', '1', '1', '获取会员列表5', '/member', '1', '{\nstatus:200,\nmsg:\"成功\",\ndata:{\n}\n}', 'http://example.com/member?key=1001&sign=123', '测试', '0', '1', '3', '1', '1', '1507617867', '1507699008');
INSERT INTO `pp_api_detail` VALUES ('2', '1', '1', '获取单个会员详情', '/member/:id', '1', '{\nstatus:200,\nmsg:\"成功\",\ndata:{\n}\n}', 'http://www.haodaquan.com/member/12', 'ceshi', '0', '1', '1', '1', '1', '1507619939', '1507693579');
INSERT INTO `pp_api_detail` VALUES ('3', '2', '3', '积分增加', '/score', '0', '{\nstatus:200,\nmsg:\"成功\",\ndata{\nmember_id:11\n}\n}', 'http://example.com/score', '这是一个测试的接口', '0', '0', '1', '1', '1', '1507699351', '1507699351');

-- ----------------------------
-- Table structure for `pp_api_param`
-- ----------------------------
DROP TABLE IF EXISTS `pp_api_param`;
CREATE TABLE `pp_api_param` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `detail_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '附表ID',
  `api_key` varchar(100) NOT NULL DEFAULT '0' COMMENT '参数名',
  `api_type` varchar(100) NOT NULL DEFAULT '0' COMMENT '类型',
  `api_value` varchar(500) NOT NULL DEFAULT '0' COMMENT '参数值',
  `api_detail` varchar(500) NOT NULL DEFAULT '0' COMMENT '参数说明',
  `is_null` varchar(10) NOT NULL DEFAULT 'no' COMMENT '是否必填',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态：1-正常，0-删除',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_main_id` (`detail_id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COMMENT='API参数表';

-- ----------------------------
-- Records of pp_api_param
-- ----------------------------
INSERT INTO `pp_api_param` VALUES ('1', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507617867', '1507699008');
INSERT INTO `pp_api_param` VALUES ('2', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507617867', '1507699008');
INSERT INTO `pp_api_param` VALUES ('3', '1', 'page_no', 'int', '1', '非必填，默认', 'no', '0', '1', '1', '1507617867', '1507699008');
INSERT INTO `pp_api_param` VALUES ('4', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507692188', '1507699008');
INSERT INTO `pp_api_param` VALUES ('5', '1', 'page_no', 'int', '1', '非必填，默认', 'no', '0', '1', '1', '1507692188', '1507699008');
INSERT INTO `pp_api_param` VALUES ('6', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507692188', '1507699008');
INSERT INTO `pp_api_param` VALUES ('7', '1', 'page_no', 'int', '1', '非必填，默认', 'no', '0', '1', '1', '1507692427', '1507699008');
INSERT INTO `pp_api_param` VALUES ('8', '1', 'rank_id2', 'int', '1', '等级', 'no', '0', '1', '1', '1507692427', '1507699008');
INSERT INTO `pp_api_param` VALUES ('9', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507692427', '1507699008');
INSERT INTO `pp_api_param` VALUES ('10', '1', 'rank_id2', 'int', '1', '等级', 'no', '0', '1', '1', '1507693216', '1507699008');
INSERT INTO `pp_api_param` VALUES ('11', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507693216', '1507699008');
INSERT INTO `pp_api_param` VALUES ('12', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507693216', '1507699008');
INSERT INTO `pp_api_param` VALUES ('13', '1', 'page_size222', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507693216', '1507699008');
INSERT INTO `pp_api_param` VALUES ('14', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507693216', '1507699008');
INSERT INTO `pp_api_param` VALUES ('15', '1', 'page_no', 'int', '1', '非必填，默认', 'no', '0', '1', '1', '1507693216', '1507699008');
INSERT INTO `pp_api_param` VALUES ('16', '1', 'page_no', 'int', '1', '非必填，默认', 'no', '0', '1', '1', '1507693216', '1507699008');
INSERT INTO `pp_api_param` VALUES ('17', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507693216', '1507699008');
INSERT INTO `pp_api_param` VALUES ('18', '1', 'page_no', 'int', '1', '非必填，默认', 'no', '0', '1', '1', '1507693216', '1507699008');
INSERT INTO `pp_api_param` VALUES ('19', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507693274', '1507699008');
INSERT INTO `pp_api_param` VALUES ('20', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507693274', '1507699008');
INSERT INTO `pp_api_param` VALUES ('21', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507693274', '1507699008');
INSERT INTO `pp_api_param` VALUES ('22', '2', 'D', 'D', 'D', 'D', 'D', '1', '1', '1', '1507693579', '1507693579');
INSERT INTO `pp_api_param` VALUES ('23', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507696615', '1507699008');
INSERT INTO `pp_api_param` VALUES ('24', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507696615', '1507699008');
INSERT INTO `pp_api_param` VALUES ('25', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507696615', '1507699008');
INSERT INTO `pp_api_param` VALUES ('26', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507697203', '1507699008');
INSERT INTO `pp_api_param` VALUES ('27', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507697203', '1507699008');
INSERT INTO `pp_api_param` VALUES ('28', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507697203', '1507699008');
INSERT INTO `pp_api_param` VALUES ('29', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507697237', '1507699008');
INSERT INTO `pp_api_param` VALUES ('30', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507697237', '1507699008');
INSERT INTO `pp_api_param` VALUES ('31', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507697237', '1507699008');
INSERT INTO `pp_api_param` VALUES ('32', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507697271', '1507699008');
INSERT INTO `pp_api_param` VALUES ('33', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507697271', '1507699008');
INSERT INTO `pp_api_param` VALUES ('34', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507697271', '1507699008');
INSERT INTO `pp_api_param` VALUES ('35', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507698950', '1507699008');
INSERT INTO `pp_api_param` VALUES ('36', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '0', '1', '1', '1507698950', '1507699008');
INSERT INTO `pp_api_param` VALUES ('37', '1', 'rank_id', 'int', '1', '等级', 'no', '0', '1', '1', '1507698950', '1507699008');
INSERT INTO `pp_api_param` VALUES ('38', '1', 'rank_id', 'int', '1', '等级', 'no', '1', '1', '1', '1507699008', '1507699008');
INSERT INTO `pp_api_param` VALUES ('39', '1', 'page_size', 'int', '10', '非必填，默认', 'no', '1', '1', '1', '1507699008', '1507699008');
INSERT INTO `pp_api_param` VALUES ('40', '1', 'rank_id', 'int', '1', '等级', 'no', '1', '1', '1', '1507699008', '1507699008');
INSERT INTO `pp_api_param` VALUES ('41', '3', 'score', 'int', '1', '变动积分', 'yes', '1', '1', '1', '1507699351', '1507699351');
INSERT INTO `pp_api_param` VALUES ('42', '3', 'type', 'int', '1', '1-增加，0-减少', 'no', '1', '1', '1', '1507699351', '1507699351');
INSERT INTO `pp_api_param` VALUES ('43', '3', 'member_id', 'int', '0', '会员ID', 'yes', '1', '1', '1', '1507699351', '1507699351');

-- ----------------------------
-- Table structure for `pp_api_source`
-- ----------------------------
DROP TABLE IF EXISTS `pp_api_source`;
CREATE TABLE `pp_api_source` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '分组ID',
  `source_name` varchar(50) NOT NULL DEFAULT '0' COMMENT '接口名称',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态：1-审核通过，0-暂停使用，2-草稿，3-审核中',
  `audit_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '审核人ID',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `audit_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '审核人时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='API主表';

-- ----------------------------
-- Records of pp_api_source
-- ----------------------------
INSERT INTO `pp_api_source` VALUES ('1', '2', '会员基本信息', '2', '0', '1', '1', '1507616276', '1507616276', '0');
INSERT INTO `pp_api_source` VALUES ('2', '2', '会员积分', '2', '0', '1', '1', '1507616329', '1512042651', '0');
INSERT INTO `pp_api_source` VALUES ('3', '2', '会员消费', '2', '0', '1', '1', '1507616394', '1507616394', '0');
INSERT INTO `pp_api_source` VALUES ('4', '1', '商品基本信息', '2', '0', '1', '1', '1507616421', '1507616421', '0');

-- ----------------------------
-- Table structure for `pp_set_code`
-- ----------------------------
DROP TABLE IF EXISTS `pp_set_code`;
CREATE TABLE `pp_set_code` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(50) NOT NULL DEFAULT '0' COMMENT '状态码',
  `desc` varchar(255) NOT NULL DEFAULT '0' COMMENT '描述',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_env_name` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='环境分组表';

-- ----------------------------
-- Records of pp_set_code
-- ----------------------------
INSERT INTO `pp_set_code` VALUES ('1', '200', '返回成功', '请求成功', '1', '0', '0', '1506328003', '1506328037');
INSERT INTO `pp_set_code` VALUES ('2', '300', '返回错误', '请求失败', '1', '0', '0', '1506328023', '1506328023');
INSERT INTO `pp_set_code` VALUES ('3', '302', '请求成功', '错误', '0', '0', '0', '1506328070', '1506334951');

-- ----------------------------
-- Table structure for `pp_set_env`
-- ----------------------------
DROP TABLE IF EXISTS `pp_set_env`;
CREATE TABLE `pp_set_env` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `env_name` varchar(50) NOT NULL DEFAULT '' COMMENT '环境名称',
  `env_host` varchar(255) NOT NULL DEFAULT '0' COMMENT '环境host',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_env_name` (`env_name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='环境分组表';

-- ----------------------------
-- Records of pp_set_env
-- ----------------------------
INSERT INTO `pp_set_env` VALUES ('1', '测试地址', 'http://127.0.0.1:8081', '测试地址', '1', '0', '0', '1506316614', '1506316614');
INSERT INTO `pp_set_env` VALUES ('2', '测试地址3', 'http://127.0.0.1:8083', '测试地址3', '0', '0', '0', '1506316696', '1506316850');

-- ----------------------------
-- Table structure for `pp_set_group`
-- ----------------------------
DROP TABLE IF EXISTS `pp_set_group`;
CREATE TABLE `pp_set_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_name` varchar(50) NOT NULL DEFAULT '' COMMENT '组名',
  `detail` varchar(255) NOT NULL DEFAULT '' COMMENT '说明',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态：1-正常，0-删除',
  `create_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `update_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_create_id` (`create_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of pp_set_group
-- ----------------------------
INSERT INTO `pp_set_group` VALUES ('1', '商品接口', '这是商品的分组', '1', '0', '1', '1506237536', '1507616254');
INSERT INTO `pp_set_group` VALUES ('2', '会员接口', '关于会员的接口', '1', '0', '1', '1506237621', '1507616227');
INSERT INTO `pp_set_group` VALUES ('3', '修改分组', '测试分组333', '0', '0', '0', '1506237655', '1506245311');

-- ----------------------------
-- Table structure for `pp_uc_admin`
-- ----------------------------
DROP TABLE IF EXISTS `pp_uc_admin`;
CREATE TABLE `pp_uc_admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `login_name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `real_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) NOT NULL DEFAULT '0' COMMENT '角色id字符串，如：2,3,4',
  `phone` varchar(20) NOT NULL DEFAULT '0' COMMENT '手机号码',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` char(10) NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_ip` char(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`login_name`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Records of pp_uc_admin
-- ----------------------------
INSERT INTO `pp_uc_admin` VALUES ('1', 'admin', '姚晓丰', '86d9116b545858e9d74e4f261db1db46', '0,2,1', '13661290405', '850630887@qq.com', 'kmcB', '1513784263', '127.0.0.1', '1', '0', '1', '0', '1513784277');
INSERT INTO `pp_uc_admin` VALUES ('2', 'george518', 'georgeHao', 'e5d77ebaffd5e4fe7164b31c6d7f1921', '1,2', '13811558899', '12@11.com', 'ONNy', '1506125048', '127.0.0.1', '1', '0', '0', '0', '1506128425');
INSERT INTO `pp_uc_admin` VALUES ('3', 'haodaquan', '郝大全', 'e9fa9187e7497892c237433aee966cc1', '2,1', '13811559988', 'hao@123.com', '6fWE', '1505960085', '127.0.0.1', '1', '1', '0', '1505919245', '1506128414');
INSERT INTO `pp_uc_admin` VALUES ('4', 'ceshizhanghao', '测试账号', 'aa08f6b4325bcd39538318c21a09cdd4', '1,2', '13988009988', '232@124.com', 'i8Nf', '1513253939', '127.0.0.1', '1', '1', '4', '1506047337', '1513253939');
INSERT INTO `pp_uc_admin` VALUES ('5', 'dangjingjiao', 'dangjingjiao', 'a0177457bea84e8e2b249773e376cbdc', '1,2', '13661260405', '850630887@qq.com', 'fmas', '0', '', '0', '1', '1', '1513780456', '1513782668');

-- ----------------------------
-- Table structure for `pp_uc_auth`
-- ----------------------------
DROP TABLE IF EXISTS `pp_uc_auth`;
CREATE TABLE `pp_uc_auth` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上级ID，0为顶级',
  `auth_name` varchar(64) NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` varchar(255) NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) unsigned NOT NULL DEFAULT '999' COMMENT '排序，越小越前',
  `icon` varchar(255) NOT NULL,
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否显示，0-隐藏，1-显示',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '操作者ID',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态，1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COMMENT='权限因子';

-- ----------------------------
-- Records of pp_uc_auth
-- ----------------------------
INSERT INTO `pp_uc_auth` VALUES ('1', '0', '所有权限', '/', '1', '', '0', '1', '1', '1', '1', '1505620970', '1505620970');
INSERT INTO `pp_uc_auth` VALUES ('2', '1', '权限管理', '/', '999', 'fa-id-card', '1', '1', '0', '1', '1', '0', '1505622360');
INSERT INTO `pp_uc_auth` VALUES ('3', '2', '管理员', '/admin/list', '1', 'fa-user-o', '1', '1', '1', '1', '1', '1505621186', '1505621186');
INSERT INTO `pp_uc_auth` VALUES ('4', '2', '角色管理', '/role/list', '2', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1505621852');
INSERT INTO `pp_uc_auth` VALUES ('5', '3', '新增', '/admin/add', '1', '', '0', '1', '0', '1', '1', '0', '1505621685');
INSERT INTO `pp_uc_auth` VALUES ('6', '3', '修改', '/admin/edit', '2', '', '0', '1', '0', '1', '1', '0', '1505621697');
INSERT INTO `pp_uc_auth` VALUES ('7', '3', '删除', '/admin/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621756', '1505621756');
INSERT INTO `pp_uc_auth` VALUES ('8', '4', '新增', '/role/add', '1', '', '1', '1', '0', '1', '1', '0', '1505698716');
INSERT INTO `pp_uc_auth` VALUES ('9', '4', '修改', '/role/edit', '2', '', '0', '1', '1', '1', '1', '1505621912', '1505621912');
INSERT INTO `pp_uc_auth` VALUES ('10', '4', '删除', '/role/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621951', '1505621951');
INSERT INTO `pp_uc_auth` VALUES ('11', '2', '权限因子', '/auth/list', '3', 'fa-list', '1', '1', '1', '1', '1', '1505621986', '1505621986');
INSERT INTO `pp_uc_auth` VALUES ('12', '11', '新增', '/auth/add', '1', '', '0', '1', '1', '1', '1', '1505622009', '1505622009');
INSERT INTO `pp_uc_auth` VALUES ('13', '11', '修改', '/auth/edit', '2', '', '0', '1', '1', '1', '1', '1505622047', '1505622047');
INSERT INTO `pp_uc_auth` VALUES ('14', '11', '删除', '/auth/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505622111', '1505622111');
INSERT INTO `pp_uc_auth` VALUES ('15', '1', '个人中心', 'profile/edit', '1001', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1506001114');
INSERT INTO `pp_uc_auth` VALUES ('16', '1', 'API管理', '/', '1', 'fa-cubes', '1', '0', '0', '0', '1', '0', '1506125698');
INSERT INTO `pp_uc_auth` VALUES ('17', '16', 'API接口', '/api/list', '1', 'fa-link', '1', '1', '1', '1', '1', '1505622447', '1505622447');
INSERT INTO `pp_uc_auth` VALUES ('19', '16', 'API监控', '/apimonitor/list', '3', 'fa-bar-chart', '1', '1', '0', '1', '1', '0', '1507700851');
INSERT INTO `pp_uc_auth` VALUES ('20', '1', '基础设置', '/', '2', 'fa-cogs', '1', '1', '1', '1', '1', '1505622601', '1505622601');
INSERT INTO `pp_uc_auth` VALUES ('21', '20', '分组设置', '/group/list', '1', 'fa-object-ungroup', '1', '1', '1', '1', '1', '1505622645', '1505622645');
INSERT INTO `pp_uc_auth` VALUES ('22', '20', '环境设置', '/env/list', '2', 'fa-tree', '1', '1', '1', '1', '1', '1505622681', '1505622681');
INSERT INTO `pp_uc_auth` VALUES ('23', '20', '状态码设置', '/code/list', '3', 'fa-code', '1', '1', '1', '1', '1', '1505622728', '1505622728');
INSERT INTO `pp_uc_auth` VALUES ('24', '15', '资料修改', '/user/edit', '1', 'fa-edit', '1', '1', '0', '1', '1', '0', '1506057468');
INSERT INTO `pp_uc_auth` VALUES ('25', '21', '新增', '/group/add', '1', 'n', '1', '0', '0', '0', '1', '1506229739', '1506229739');
INSERT INTO `pp_uc_auth` VALUES ('26', '21', '修改', '/group/edit', '2', 'fa', '0', '0', '0', '0', '1', '1506237920', '1506237920');
INSERT INTO `pp_uc_auth` VALUES ('27', '21', '删除', '/group/ajaxdel', '3', 'fa', '0', '0', '0', '0', '1', '1506237948', '1506237948');
INSERT INTO `pp_uc_auth` VALUES ('28', '22', '新增', '/env/add', '1', 'fa', '0', '0', '0', '0', '1', '1506316506', '1506316506');
INSERT INTO `pp_uc_auth` VALUES ('29', '22', '修改', '/env/edit', '2', 'fa', '0', '0', '0', '0', '1', '1506316532', '1506316532');
INSERT INTO `pp_uc_auth` VALUES ('30', '22', '删除', '/env/ajaxdel', '3', 'fa', '0', '0', '0', '0', '1', '1506316567', '1506316567');
INSERT INTO `pp_uc_auth` VALUES ('31', '23', '新增', '/code/add', '1', 'fa', '0', '0', '0', '0', '1', '1506327812', '1506327812');
INSERT INTO `pp_uc_auth` VALUES ('32', '23', '修改', '/code/edit', '2', 'fa', '0', '0', '0', '0', '1', '1506327831', '1506327831');
INSERT INTO `pp_uc_auth` VALUES ('33', '23', '删除', '/code/ajadel', '3', 'fa', '0', '0', '0', '0', '1', '1506327857', '1506327857');
INSERT INTO `pp_uc_auth` VALUES ('34', '17', '新增资源', '/api/add', '1', 'fa-link', '1', '1', '0', '1', '1', '0', '1507436029');
INSERT INTO `pp_uc_auth` VALUES ('35', '17', '修改资源', '/api/edit', '2', 'fa-link', '1', '1', '0', '1', '1', '0', '1507436042');
INSERT INTO `pp_uc_auth` VALUES ('36', '17', '删除资源', '/api/ajaxdel', '3', 'fa-link', '1', '1', '0', '1', '1', '0', '1507436052');
INSERT INTO `pp_uc_auth` VALUES ('37', '17', '新增接口', '/api/addapi', '4', '', '0', '1', '1', '1', '1', '1507436014', '1507436014');
INSERT INTO `pp_uc_auth` VALUES ('38', '17', '修改接口', '/api/editapi', '5', '', '0', '1', '0', '1', '1', '0', '1507705049');

-- ----------------------------
-- Table structure for `pp_uc_role`
-- ----------------------------
DROP TABLE IF EXISTS `pp_uc_role`;
CREATE TABLE `pp_uc_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改这ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Records of pp_uc_role
-- ----------------------------
INSERT INTO `pp_uc_role` VALUES ('1', 'API管理员', '拥有API所有权限', '0', '2', '1', '1505874156', '1505874156');
INSERT INTO `pp_uc_role` VALUES ('2', '系统管理员', '系统管理员', '0', '0', '1', '1506124114', '1506124114');

-- ----------------------------
-- Table structure for `pp_uc_role_auth`
-- ----------------------------
DROP TABLE IF EXISTS `pp_uc_role_auth`;
CREATE TABLE `pp_uc_role_auth` (
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `auth_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限ID',
  PRIMARY KEY (`role_id`,`auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限和角色关系表';

-- ----------------------------
-- Records of pp_uc_role_auth
-- ----------------------------
INSERT INTO `pp_uc_role_auth` VALUES ('1', '16');
INSERT INTO `pp_uc_role_auth` VALUES ('1', '17');
INSERT INTO `pp_uc_role_auth` VALUES ('1', '18');
INSERT INTO `pp_uc_role_auth` VALUES ('1', '19');
INSERT INTO `pp_uc_role_auth` VALUES ('2', '0');
INSERT INTO `pp_uc_role_auth` VALUES ('2', '1');
INSERT INTO `pp_uc_role_auth` VALUES ('2', '15');
INSERT INTO `pp_uc_role_auth` VALUES ('2', '20');
INSERT INTO `pp_uc_role_auth` VALUES ('2', '21');
INSERT INTO `pp_uc_role_auth` VALUES ('2', '22');
INSERT INTO `pp_uc_role_auth` VALUES ('2', '23');
INSERT INTO `pp_uc_role_auth` VALUES ('2', '24');
