DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `role`;
DROP TABLE IF EXISTS `user_role`;
DROP TABLE IF EXISTS `permission`;
DROP TABLE IF EXISTS `role_permission`;
DROP TABLE IF EXISTS `competition`;

CREATE TABLE users (
	id int unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(100) DEFAULT NULL,
	email varchar(100) DEFAULT NULL,
	`password` varchar(100) DEFAULT NULL,
	stu_no varchar(100) DEFAULT NULL,
	stu_college varchar(100) DEFAULT NULL,
	stu_grade varchar(100) DEFAULT NULL,
	PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

INSERT INTO users(`id`, `name`, `email`, `password`) VALUES(1, 'zhangsan', 'zhangsan@qq.com', '123456');
INSERT INTO users(`id`, `name`, `email`, `password`) VALUES(2, 'lisi', 'lisi@qq.com', '123456');
INSERT INTO users(`id`, `name`, `email`, `password`) VALUES(3, 'zhaowu', 'zhaowu@qq.com', '123456');


CREATE TABLE role(
	`id` INT PRIMARY KEY auto_increment,
	`role_name` varchar(20)	
);

INSERT INTO role(`id`, `role_name`) VALUES(1, 'admin');
INSERT INTO role(`id`, `role_name`) VALUES(2, 'competition_manager');
INSERT INTO role(`id`, `role_name`) VALUES(3, 'user');


CREATE TABLE permission(
	`id` INT PRIMARY KEY auto_increment,
	`permission_name` varchar(32),
	`url`	varchar(100),
	`method`	varchar(12)
);

INSERT INTO permission(`id`, `permission_name`, `url`, `method`) VALUES(1, '所有权限', '*', '*');	# 超级用户
INSERT INTO permission(`id`, `permission_name`, `url`, `method`) VALUES(2, '比赛管理', '/api/v1/setting/competition*', '*');
INSERT INTO permission(`id`, `permission_name`, `url`, `method`) VALUES(3, '普通查看页面', '/api/v1/index', 'GET');

CREATE TABLE user_role(
	`id` INT PRIMARY KEY auto_increment,
	`user_id` INT,
	`role_id` INT
);

INSERT INTO user_role(`user_id`, `role_id`) VALUES(1, 1);
INSERT INTO user_role(`user_id`, `role_id`) VALUES(2, 2);
INSERT INTO user_role(`user_id`, `role_id`) VALUES(3, 3);

CREATE TABLE role_permission(
	`id` INT PRIMARY KEY auto_increment,
	`role_id` INT,
	`permission_id` INT
);

INSERT INTO role_permission(`role_id`, `permission_id`) VALUES(1, 1);
INSERT INTO role_permission(`role_id`, `permission_id`) VALUES(2, 2);
INSERT INTO role_permission(`role_id`, `permission_id`) VALUES(3, 3);

SET SESSION sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
CREATE TABLE `competition` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(50),
  `description` varchar(255),
  `reward` varchar(255),
  `entry_requirement` varchar(255),
  `work_requirement` varchar(255),
  `signup_deadline` TIMESTAMP NOT NULL,
  `submit_deadline` TIMESTAMP NOT NULL,
  `company_id` int(11)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


