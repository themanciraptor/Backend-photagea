-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: photagea
-- ------------------------------------------------------
-- Server version	8.0.19

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account` (
  `AccountID` int NOT NULL AUTO_INCREMENT,
  `Email` varchar(255) NOT NULL,
  `Password` varchar(64) NOT NULL,
  `Created` datetime DEFAULT CURRENT_TIMESTAMP,
  `Updated` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Deleted` datetime DEFAULT NULL,
  PRIMARY KEY (`AccountID`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account`
--

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;
INSERT INTO `account` VALUES (10,'maxpunk@hotmail.com','','2020-03-11 16:10:04',NULL,NULL),(12,'Duudgotgroove@snacksnakc.com','','2020-03-11 16:10:04',NULL,NULL),(14,'maxk@hotmail.com','','2020-03-20 09:41:13',NULL,NULL),(15,'smeagol@gollum.net','','2020-03-24 13:39:18',NULL,NULL),(16,'dood','$2a$12$uelnn2gP1HC7NPLrr1zhX.E4PYTqir3uwUzVDg6N.OcKHVUeljgGi','2020-03-24 16:05:23',NULL,NULL);
/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `image`
--

DROP TABLE IF EXISTS `image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `image` (
  `ImageID` int NOT NULL AUTO_INCREMENT,
  `AccountID` int DEFAULT NULL,
  `URL` varchar(255) DEFAULT NULL,
  `Created` datetime DEFAULT CURRENT_TIMESTAMP,
  `Updated` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Deleted` datetime DEFAULT NULL,
  PRIMARY KEY (`ImageID`),
  KEY `AccountID` (`AccountID`),
  CONSTRAINT `image_ibfk_1` FOREIGN KEY (`AccountID`) REFERENCES `account` (`AccountID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `image`
--

LOCK TABLES `image` WRITE;
/*!40000 ALTER TABLE `image` DISABLE KEYS */;
INSERT INTO `image` VALUES (1,16,'https://upload.wikimedia.org/wikipedia/commons/thumb/d/d9/Via_appia.jpg/1200px-Via_appia.jpg','2020-03-25 15:34:24',NULL,NULL),(2,16,'https://www.telegraph.co.uk/content/dam/Travel/Destinations/Europe/France/Nice/nice-promenade-attractions-xlarge.jpg','2020-03-26 16:14:59',NULL,NULL);
/*!40000 ALTER TABLE `image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `imagedata`
--

DROP TABLE IF EXISTS `imagedata`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `imagedata` (
  `ImageDataID` int NOT NULL AUTO_INCREMENT,
  `AccountID` int DEFAULT NULL,
  `mimetype` varchar(64) NOT NULL,
  `Created` datetime DEFAULT CURRENT_TIMESTAMP,
  `Updated` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Deleted` datetime DEFAULT NULL,
  PRIMARY KEY (`ImageDataID`),
  KEY `AccountID` (`AccountID`),
  CONSTRAINT `imagedata_ibfk_1` FOREIGN KEY (`AccountID`) REFERENCES `account` (`AccountID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `imagedata`
--

LOCK TABLES `imagedata` WRITE;
/*!40000 ALTER TABLE `imagedata` DISABLE KEYS */;
INSERT INTO `imagedata` VALUES (1,16,'image/jpeg','2020-03-25 17:39:36',NULL,NULL);
/*!40000 ALTER TABLE `imagedata` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `UserID` int NOT NULL AUTO_INCREMENT,
  `Alias` varchar(64) DEFAULT NULL,
  `FirstName` varchar(64) DEFAULT NULL,
  `LastName` varchar(64) DEFAULT NULL,
  `AccountID` int DEFAULT NULL,
  `Created` datetime DEFAULT CURRENT_TIMESTAMP,
  `Updated` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Deleted` datetime DEFAULT NULL,
  PRIMARY KEY (`UserID`),
  KEY `AccountID` (`AccountID`),
  CONSTRAINT `user_ibfk_1` FOREIGN KEY (`AccountID`) REFERENCES `account` (`AccountID`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (10,'mammasboy','ana','matopea',10,'2020-03-11 16:10:04',NULL,NULL),(17,'With','love,','johnny',14,'2020-03-20 09:41:29','2020-03-24 13:43:31',NULL),(26,'Lovely','blended','Corn',15,'2020-03-24 13:40:17',NULL,NULL),(32,'Barkabark','Dude','Smile',16,'2020-03-24 16:31:51',NULL,NULL);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-26 17:27:40
