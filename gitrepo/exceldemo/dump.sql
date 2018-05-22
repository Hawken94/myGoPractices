-- MySQL dump 10.13  Distrib 5.7.20, for macos10.12 (x86_64)
--
-- Host: 127.0.0.1    Database: ncu
-- ------------------------------------------------------
-- Server version	5.7.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `student_info`
--

DROP TABLE IF EXISTS `student_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `student_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `examinee_no` varchar(55) DEFAULT '' COMMENT '考生号',
  `unknown_code` int(11) DEFAULT '0' COMMENT '不知道的代号',
  `student_no` varchar(55) DEFAULT '' COMMENT '学号',
  `stu_name` varchar(55) NOT NULL DEFAULT '' COMMENT '姓名',
  `gender` int(11) DEFAULT '0' COMMENT '性别 0:男 1:女',
  `birthday` datetime DEFAULT NULL COMMENT '出生年月日',
  `indentity` varchar(55) NOT NULL COMMENT '身份证',
  `politic_status` int(11) DEFAULT '0' COMMENT '政治面貌 0:共青团员 1:群众 2:中共预备党员 3:中共党员',
  `nation` varchar(55) DEFAULT '' COMMENT '民族',
  `degree_no` int(11) DEFAULT '0' COMMENT '学位证编号',
  `school` varchar(55) DEFAULT '' COMMENT '学校名称',
  `major_code` int(11) DEFAULT '0' COMMENT '专业代码',
  `major_name` varchar(55) DEFAULT '' COMMENT '专业名称',
  `college` varchar(55) DEFAULT '' COMMENT '学院名称',
  `major_direction` varchar(55) DEFAULT '' COMMENT '专业方向',
  `class_name` varchar(55) DEFAULT '' COMMENT '班级名称',
  `length_of_schooling` int(11) DEFAULT '0' COMMENT '学年制',
  `university_mode` int(11) DEFAULT '0' COMMENT '培养形式 0:普通全日制 1:其他',
  `school_begin_date` datetime DEFAULT NULL COMMENT '入学时间',
  `school_year` int(11) DEFAULT '0' COMMENT '入学年份',
  `register_school_roll` varchar(55) DEFAULT '' COMMENT '注册学籍',
  `school_leave_date` datetime DEFAULT NULL COMMENT '离校时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学信网学生信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student_info`
--

LOCK TABLES `student_info` WRITE;
/*!40000 ALTER TABLE `student_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `student_info` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-05-22 23:05:02
