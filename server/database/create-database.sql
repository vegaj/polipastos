-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Versión del servidor:         10.3.7-MariaDB - mariadb.org binary distribution
-- SO del servidor:              Win64
-- HeidiSQL Versión:             9.4.0.5125
-- --------------------------------------------------------


-- Volcando estructura de base de datos para polipastos
DROP DATABASE IF EXISTS `polipastos`;
CREATE DATABASE IF NOT EXISTS `polipastos` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `polipastos`;

-- Volcando estructura para tabla polipastos.hashes
DROP TABLE IF EXISTS `hashes`;
CREATE TABLE IF NOT EXISTS `hashes` (
  `owner_id` int(10) NOT NULL,
  `digest` char(255) DEFAULT NULL,
  PRIMARY KEY (`owner_id`),
  CONSTRAINT `fk_digest_owner_id` FOREIGN KEY (`owner_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Here we store the digestion of the user password';

-- La exportación de datos fue deseleccionada.
-- Volcando estructura para tabla polipastos.resources
DROP TABLE IF EXISTS `resources`;
CREATE TABLE IF NOT EXISTS `resources` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_id` int(11) NOT NULL DEFAULT 0,
  `resource_type` enum('uns','txt','doc','img','vid') NOT NULL DEFAULT 'uns',
  `title` varchar(50) NOT NULL,
  `url` text DEFAULT NULL,
  `editedAt` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `fk_resource_owner_id` (`owner_id`),
  CONSTRAINT `fk_resource_owner_id` FOREIGN KEY (`owner_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='A register of a certain resource that belongs to a user and a url that says where it is stored';

-- La exportación de datos fue deseleccionada.
-- Volcando estructura para tabla polipastos.user
DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UniqueUsername` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='Polipastos User';
