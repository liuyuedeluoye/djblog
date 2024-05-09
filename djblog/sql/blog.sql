DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `user_id` bigint(20) NOT NULL,
                        `role` bigint(20) not null ,
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `create_time` timestamp not null default current_timestamp,
                        `update_time` timestamp not null default current_timestamp on update current_timestamp,
                        PRIMARY KEY (`id`),

                        UNIQUE KEY `idx_username` (`username`) USING BTREE,
                        UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET= utf8mb4 COLLATE=utf8mb4_general_ci;
INSERT INTO `user` (`user_id`, `role`, `username`, `password`) VALUES (1, 1, 'dsj', '123456');


##post帖子
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `article_id` bigint(28) not null comment '帖子id',
                        `title` varchar(128) COLLATE utf8mb4_general_ci not null comment '标题',
                        `content` varchar(8192) collate utf8mb4_general_ci not null comment '内容',
#                         `category_id` bigint(20) not null comment '所属分类',
                        `create_time` timestamp not null default current_timestamp comment '创建时间',
                        `update_time` timestamp not null default current_timestamp comment '更新时间',
                        primary key (`id`),
                        unique key `idx_article_id` (`article_id`)
#                         key `idx_category_id` (`category_id`)
)engine = InnoDB default charset = utf8mb4 collate = utf8mb4_general_ci;

DROP TABLE IF EXISTS `category`;



