 
-- ----------------------------
-- Table structure for basic_file
-- ----------------------------
DROP TABLE IF EXISTS `basic_file`;
CREATE TABLE `basic_file` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `guid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '唯一id',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '用户id',
  `user_id_sys` bigint(20) DEFAULT '0' COMMENT '系统用户id',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件名',
  `cat_id` bigint(20) DEFAULT '0' COMMENT '栏目id',
  `module` smallint(6) DEFAULT '0' COMMENT '模块id',
  `media_type` smallint(6) DEFAULT '0' COMMENT '媒体类型: 1 图片 2 视频 3音频 4文档 ',
  `driver` smallint(6) DEFAULT '0' COMMENT '存储位置:0本地1阿里云2腾讯云',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件路径',
  `file_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件类型',
  `ext` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '文件类型',
  `size` int(10) unsigned DEFAULT '0' COMMENT '文件大小[k]',
  `md5` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'md5值',
  `sha1` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'sha散列值',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `download` int(10) unsigned DEFAULT '0' COMMENT '下载次数',
  `used_time` int(11) DEFAULT '0' COMMENT '使用次数',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `status` smallint(6) DEFAULT '0' COMMENT '状态: 3=删除 ',
  PRIMARY KEY (`id`),
  KEY `guid` (`guid`),
  KEY `md5` (`md5`),
  KEY `userid` (`user_id`),
  KEY `module` (`module`),
  KEY `sha1` (`sha1`),
  KEY `deleted_at` (`deleted_at`),
  KEY `cat_id` (`cat_id`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1146 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文件附件表';

-- ----------------------------
-- Table structure for mem_logs
-- ----------------------------
DROP TABLE IF EXISTS `mem_logs`;
CREATE TABLE `mem_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) unsigned DEFAULT '0' COMMENT '集趣id',
  `type` smallint(6) DEFAULT '0' COMMENT '类型: 1登录 2退出 3增加用户4 增加任务',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '备注',
  `ip` int(10) unsigned DEFAULT '0' COMMENT 'ip',
  `ip_addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'ip城市',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0=不正常,1=正常',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `type` (`type`),
  KEY `user_id` (`user_id`),
  KEY `deleted_at` (`deleted_at`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户日志';

-- ----------------------------
-- Table structure for mem_user
-- ----------------------------
DROP TABLE IF EXISTS `mem_user`;
CREATE TABLE `mem_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '会员id',
  `user_id` bigint(20) DEFAULT '0' COMMENT '录入人',
  `user_type` smallint(6) DEFAULT '0' COMMENT '类型:1员工 2子账号  3管理员 9 平台号',
  `guid` varchar(32) DEFAULT '' COMMENT '唯一id',
  `username` varchar(40) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) DEFAULT '' COMMENT '密码',
  `password_slat` varchar(32) DEFAULT '' COMMENT '密码盐',
  `nickname` varchar(200) DEFAULT '' COMMENT '昵称',
  `realname` varchar(200) DEFAULT '' COMMENT '真实名',
  `role_list` varchar(255) DEFAULT '' COMMENT '角色list',
  `role` bigint(20) DEFAULT '0' COMMENT '角色',
  `email` varchar(60) DEFAULT '' COMMENT '邮件',
  `mobile` varchar(20) DEFAULT '' COMMENT '手机',
  `card_id` varchar(18) DEFAULT '' COMMENT '身份证',
  `sex` smallint(6) DEFAULT '0' COMMENT '性别: 0保密 1 男 2 女',
  `birthday` datetime DEFAULT NULL COMMENT '生日',
  `avatar` varchar(256) DEFAULT '' COMMENT '头像',
  `job_id` bigint(20) DEFAULT '0' COMMENT '岗位',
  `depart_id` bigint(20) DEFAULT '0' COMMENT '部门',
  `mobile_validated` tinyint(1) DEFAULT '0' COMMENT '验证手机',
  `email_validated` tinyint(1) DEFAULT '0' COMMENT '验证邮箱',
  `cardid_validated` tinyint(1) DEFAULT '0' COMMENT '验证实名',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `status_safe` smallint(6) DEFAULT '0' COMMENT '安全状态:0=正常，1=修改了密码 2=修改了手机号',
  `reg_ip` int(10) unsigned DEFAULT '0' COMMENT '注册ip',
  `login_ip` int(10) unsigned DEFAULT '0' COMMENT '登录ip',
  `login_total` int(11) DEFAULT '0' COMMENT '登录次数',
  `login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `scan` varchar(32) DEFAULT '' COMMENT '扫码',
  `scan_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '扫码',
  `setting` varchar(1000) DEFAULT '' COMMENT '设置值',
  `rtc_model` smallint(6) DEFAULT '0' COMMENT '远程协助模式',
  `status` smallint(6) DEFAULT '0' COMMENT '状态:0=审核中,1=审核通过 2=审核未通过3=禁止登录4=非法信息5:已注销6:非法攻击',
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `guid` (`guid`),
  KEY `mobile` (`mobile`),
  KEY `username` (`username`),
  KEY `email` (`email`),
  KEY `card_id` (`card_id`),
  KEY `realname` (`realname`),
  KEY `sex` (`sex`),
  KEY `job_id` (`job_id`),
  KEY `depart_id` (`depart_id`),
  KEY `scan` (`scan`),
  KEY `user_type` (`user_type`),
  KEY `deleted_at` (`deleted_at`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户';

 