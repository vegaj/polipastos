-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Versión del servidor:         10.3.7-MariaDB - mariadb.org binary distribution
-- SO del servidor:              Win64
-- HeidiSQL Versión:             9.4.0.5125
-- --------------------------------------------------------


-- Volcando estructura de base de datos para polipastos-conf
CREATE DATABASE IF NOT EXISTS `polipastos-conf` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `polipastos-conf`;

-- Volcando estructura para tabla polipastos-conf.map
CREATE TABLE IF NOT EXISTS `map` (
  `name` varchar(64) NOT NULL,
  `value` varchar(256) NOT NULL,
  PRIMARY KEY (`name`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Un Diccionario clave valor de tipo texto 64 -> 256 Bytes';
