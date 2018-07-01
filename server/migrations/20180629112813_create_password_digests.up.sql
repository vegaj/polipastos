CREATE TABLE `password_digests` (
	`id` CHAR(36) NOT NULL,
	`owner_id` CHAR(36) NOT NULL,
	`digest` CHAR(255) NULL DEFAULT NULL,
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	INDEX `disgest_fk_with_user` (`owner_id`),
	CONSTRAINT `disgest_fk_with_user` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;