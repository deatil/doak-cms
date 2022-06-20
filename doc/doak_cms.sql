DROP TABLE IF EXISTS `cms_art`;
CREATE TABLE `cms_art` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT 'id',
  `cate_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '分类ID',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '作者',
  `cover` char(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '封面',
  `title` varchar(150) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '标题',
  `keywords` varchar(100) CHARACTER SET utf8mb4 DEFAULT '' COMMENT '关键字',
  `description` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `content` longtext CHARACTER SET utf8mb4 NOT NULL COMMENT '内容',
  `tags` varchar(250) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标签',
  `from` varchar(200) CHARACTER SET utf8mb4 DEFAULT '' COMMENT '来源',
  `is_top` tinyint(1) unsigned DEFAULT '0' COMMENT '1-置顶',
  `status` tinyint(1) DEFAULT '1' COMMENT '1-启用，0-禁用',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章';

DROP TABLE IF EXISTS `cms_attach`;
CREATE TABLE `cms_attach` (
  `id` char(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文件名',
  `path` varchar(250) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '文件路径',
  `ext` varchar(10) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '文件类型',
  `size` int(11) NOT NULL DEFAULT '0' COMMENT '文件大小',
  `md5` char(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件md5',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='附件表';

DROP TABLE IF EXISTS `cms_cate`;
CREATE TABLE `cms_cate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(10) NOT NULL DEFAULT '0' COMMENT '父级ID',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标志',
  `desc` varchar(200) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '描述',
  `sort` int(5) DEFAULT '100' COMMENT '排序',
  `status` tinyint(1) DEFAULT '1' COMMENT '1-启用，0-禁用',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类';

DROP TABLE IF EXISTS `cms_config`;
CREATE TABLE `cms_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '字段',
  `value` text COLLATE utf8mb4_unicode_ci COMMENT '字段值',
  `desc` varchar(200) CHARACTER SET utf8mb4 DEFAULT '' COMMENT '字段说明',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='配置';

DROP TABLE IF EXISTS `cms_page`;
CREATE TABLE `cms_page` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '作者',
  `slug` varchar(150) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '标志',
  `title` varchar(200) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '标题',
  `keywords` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '关键字',
  `description` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `content` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
  `status` tinyint(1) DEFAULT '1' COMMENT '1-启用，0-禁用',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='单页';

DROP TABLE IF EXISTS `cms_tag`;
CREATE TABLE `cms_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `desc` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '描述',
  `sort` int(5) DEFAULT '100' COMMENT '排序',
  `status` tinyint(1) DEFAULT '1' COMMENT '1-启用，0-禁用',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签';

DROP TABLE IF EXISTS `cms_user`;
CREATE TABLE `cms_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账号，大小写字母数字',
  `password` char(62) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(100) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` char(32) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '头像',
  `sign` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '签名',
  `status` tinyint(1) DEFAULT '1' COMMENT '1-启用，0-禁用',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `add_ip` varchar(50) CHARACTER SET utf8mb4 NOT NULL DEFAULT '' COMMENT '添加IP',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户';

INSERT INTO `cms_art` VALUES (2,'38592cb2f1ff93980b3470201c9cff8d',1,1,'859de1506917e0d3f93a07c7c725e858','文章测试123','','','<p>文章测试123123</p><p><img src=\"/upload/images/78a1b859e4ce392cd5e8fa430382f90a.jpg\" alt=\"/upload/images/78a1b859e4ce392cd5e8fa430382f90a.jpg\" data-href=\"/upload/images/78a1b859e4ce392cd5e8fa430382f90a.jpg\" style=\"\"/></p>','测试,这是对的,你好啊','',0,1,1655610440,'127.0.0.1');
INSERT INTO `cms_attach` VALUES ('6c4c5b96aab258792559794be4ae6aa8','Koala.jpg','images/8898f771f216847361029fbae115e04d.jpg','jpg',780831,'2b04df3ecc1d94afddff082d139c6f15',1,1655714827,'127.0.0.1'),('859de1506917e0d3f93a07c7c725e858','Tulips.jpg','images/853b804e53393a08fad5c0b8d2283b8f.jpg','jpg',620888,'fafa5efeaf3cbe3b23b2748d13e629a1',1,1655715756,'127.0.0.1');
INSERT INTO `cms_cate` VALUES (1,0,'热门八卦','rmbg','热门八卦',100,1,1653622870,''),(6,0,'影视综合','video','影视综合',100,1,1655522965,'127.0.0.1');
INSERT INTO `cms_config` VALUES (1,'website_name','网站名称','网站名称'),(2,'website_keywords','网站关键字2','网站关键字'),(3,'website_description','网站描述1','网站描述'),(4,'website_copyright','版权1','版权'),(5,'website_status','1','网站关闭状态'),(6,'website_beian','网站备案2','网站备案');
INSERT INTO `cms_page` VALUES (1,1,'aboutme','关于我们','','','<p>关于我们</p>',1,1655711798,'127.0.0.1');
INSERT INTO `cms_tag` VALUES (1,'标签1','标签1',101,1,1655520561,'127.0.0.1');
INSERT INTO `cms_user` VALUES (1,'admin','$2a$10$EGQtXJLwsMhY.Cdoqma/IeHruRWtfYYMbqRRd1I0Asgjf6s048PrS','管理员','6c4c5b96aab258792559794be4ae6aa8','签名数据123',1,1653622870,''),(2,'myname2','$2a$10$6qUnH6qnYFdeSLlxKHcCseki3xgt6qbBzjMABZWkm0FIOUE5W5Axm','myname',NULL,'myname',1,1653622881,'');
