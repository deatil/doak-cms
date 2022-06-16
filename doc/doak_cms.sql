DROP TABLE IF EXISTS `cms_art`;
CREATE TABLE `cms_art` (
  `id` char(32) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'id',
  `title` varchar(150) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '标题',
  `keywords` varchar(100) CHARACTER SET utf8mb4 DEFAULT '' COMMENT '关键字',
  `description` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `content` longtext CHARACTER SET utf8mb4 NOT NULL COMMENT '内容',
  `tags` varchar(250) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标签',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '作者',
  `from` varchar(200) CHARACTER SET utf8mb4 DEFAULT '' COMMENT '来源',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章';

DROP TABLE IF EXISTS `cms_cate`;
CREATE TABLE `cms_cate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(10) NOT NULL DEFAULT '0' COMMENT '父级ID',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标志',
  `desc` varchar(200) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '描述',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类';

DROP TABLE IF EXISTS `cms_config`;
CREATE TABLE `cms_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字段',
  `value` text COLLATE utf8mb4_unicode_ci COMMENT '字段值',
  `desc` varchar(200) CHARACTER SET utf8mb4 DEFAULT '' COMMENT '字段说明',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='配置';

DROP TABLE IF EXISTS `cms_tag`;
CREATE TABLE `cms_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `desc` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签';

DROP TABLE IF EXISTS `cms_user`;
CREATE TABLE `cms_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账号，大小写字母数字',
  `password` char(62) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(100) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '昵称',
  `sign` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '签名',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户';

INSERT INTO `cms_user` VALUES (1,'admin','$2a$10$EGQtXJLwsMhY.Cdoqma/IeHruRWtfYYMbqRRd1I0Asgjf6s048PrS','admin','',1653622870,''),(2,'myname2','','','',1653622881,'');
