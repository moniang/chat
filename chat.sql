SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for chat_set
-- ----------------------------
DROP TABLE IF EXISTS `chat_set`;
CREATE TABLE `chat_set` (
  `set_id` int(11) NOT NULL AUTO_INCREMENT,
  `set_user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `set_font_color` char(7) NOT NULL DEFAULT '' COMMENT '字体颜色',
  PRIMARY KEY (`set_id`),
  UNIQUE KEY `ids_user_id` (`set_user_id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for chat_user
-- ----------------------------
DROP TABLE IF EXISTS `chat_user`;
CREATE TABLE `chat_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_nick` varchar(255) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_user` varchar(20) NOT NULL DEFAULT '' COMMENT '用户账号',
  `user_pass` char(32) NOT NULL DEFAULT '' COMMENT '用户密码',
  `user_salt` char(8) NOT NULL DEFAULT '' COMMENT '加密盐值',
  `user_token` char(32) NOT NULL DEFAULT '' COMMENT 'Token凭证',
  `user_vip` int(1) NOT NULL DEFAULT '0' COMMENT '会员等级',
  `user_token_end_time` int(10) DEFAULT NULL COMMENT 'Token过期时间戳',
  `user_add_time` datetime DEFAULT NULL COMMENT '用户注册时间',
  PRIMARY KEY (`user_id`)
) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户表';