-- -------------------------------------------------------------
-- TablePlus 5.3.8(500)
--
-- https://tableplus.com/
--
-- Database: food_delivery_data
-- Generation Time: 2023-07-27 11:03:53.6920
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `restaurants_200lab` (
  `id` int NOT NULL AUTO_INCREMENT,
  `owner_id` int DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `addr` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `type` enum('normal','vip','premium') NOT NULL DEFAULT 'normal',
  `city_id` int DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `cover` json DEFAULT NULL,
  `logo` json DEFAULT NULL,
  `shipping_fee_per_km` double DEFAULT '0',
  `status` int NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `owner_id` (`owner_id`) USING BTREE,
  KEY `city_id` (`city_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb3;

INSERT INTO `restaurants_200lab` (`id`, `owner_id`, `name`, `addr`, `type`, `city_id`, `lat`, `lng`, `cover`, `logo`, `shipping_fee_per_km`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 'Nha hang Hai Au', '796/17 Le Duc Tho', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-07-14 09:59:24', '2023-07-19 15:06:34'),
(2, 2, 'Place Holder', '796/17 Le Dai Hanh', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-14 13:57:10', '2023-07-24 10:10:26'),
(4, 1, 'Nha hang Hai Au', '1122 Quang Trung', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-07-14 14:06:45', '2023-07-19 15:06:34'),
(5, 3, 'Place Holder', '', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-14 14:07:31', '2023-07-19 15:06:34'),
(6, 2, 'Dong Que Place', '11 Le Hoang Phai, P15, Q. Go Vap', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-07-14 14:31:07', '2023-07-19 15:06:34'),
(7, 2, 'Place Holder', '566 Nguyen Thai Son', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-14 16:32:32', '2023-07-19 15:06:34'),
(8, NULL, 'Place Holder', '566 Nguyen Thai Son', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-07-14 16:42:15', '2023-07-14 16:42:15'),
(9, 1, 'Diamond Plaza 2', '', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-19 07:28:50', '2023-07-19 15:06:34'),
(10, NULL, 'Diamond Plaza 2', '123 Hoang Van Thu, Phuong 13, Tan Binh', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-19 07:29:49', '2023-07-19 07:29:49'),
(11, 3, 'Kingdom Diamond', '', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-19 07:31:50', '2023-07-21 04:26:13'),
(12, NULL, 'Diamond Plaza 2', '123 Hoang Van Thu, Phuong 13, Tan Binh', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-19 07:37:38', '2023-07-19 07:37:38'),
(13, NULL, 'Dong Que Place', '', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-19 07:39:29', '2023-07-21 04:27:01'),
(14, 2, 'Diamond Plaza 2', '123 Hoang Van Thu, Phuong 13, Tan Binh', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-07-19 07:40:58', '2023-07-19 16:50:07'),
(15, NULL, 'Ghanh Hang Rong', '147, Au Co, P2, Q.Tan Phu', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-19 11:44:54', '2023-07-19 11:44:54'),
(16, NULL, 'Ghanh Hang Rong 2', '147, Au Co, P2, Q.Tan Phu', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-07-25 02:37:32', '2023-07-25 02:37:32'),
(17, NULL, 'Ghanh Hang Rong CN 3', '147, Au Co, P2, Q.Tan Phu', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(18, NULL, 'Ghanh Hang Rong CN 4', '147, Au Co, P2, Q.Tan Phu', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(19, NULL, 'Thuan Kieu Plaza', '288 Hung vUong Quan 5', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(20, NULL, 'Thuan Kieu Plaza', '288 Hung vUong Quan 5', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(21, NULL, 'Thuan Kieu Plaza 1', '288 Hung vUong Quan 1', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, NULL, '2023-07-25 03:02:22'),
(22, NULL, 'Thuan Kieu Plaza 2', '288 Hung vUong Quan 3', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(23, NULL, 'Thuan Kieu Plaza 1', '288 Hung vUong Quan 1', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(24, NULL, 'Thuan Kieu Plaza 1', '288 Hung vUong Quan 1', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(25, NULL, 'Thuan Kieu Plaza 1', '288 Hung vUong Quan 1', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(26, NULL, 'Thuan Kieu Plaza 1', '288 Hung vUong Quan 1', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(27, NULL, 'Thuan Kieu Plaza 1', '288 Hung vUong Quan 1', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(28, NULL, 'Thuan Kieu Plaza 1', '288 Hung vUong Quan 1', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-25 02:39:52', '2023-07-25 02:39:52'),
(29, NULL, 'Thuan Kieu Plaza 1', '288 Hung vUong Quan 1', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, NULL, NULL),
(30, NULL, 'Chung Cu Khong Co Ma', 'Tan Hung Thuan Quan 7', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 0, '2023-07-25 02:49:18', '2023-07-25 02:58:19'),
(31, NULL, 'Thuan Kieu Plaza Co Ma', '288 Hung vUong Quan 5', 'normal', NULL, NULL, NULL, NULL, NULL, 0, 1, '2023-07-25 02:53:28', '2023-07-25 02:53:28'),
(32, NULL, 'Phoenix Restaurant and Coffee', 'RM76+GX8, Phường 2, Tân Bình, Thành phố Hồ Chí Minh, Vietnam', 'normal', NULL, NULL, NULL, NULL, '{\"id\": 0, \"url\": \"https://images.g2crowd.com/uploads/attachment/file/197354/Retrospective.png\", \"width\": 100, \"height\": 100}', 0, 1, '2023-07-25 09:44:23', '2023-07-25 09:44:23'),
(33, NULL, 'HOME Saigon - HOME Vietnamese Restaurant', '216/4 Đ. Điện Biên Phủ, Phường 4, Quận 3, Thành phố Hồ Chí Minh 700000, Vietnam', 'normal', NULL, NULL, NULL, '[{\"id\": 0, \"url\": \"https://images.g2crowd.com/uploads/attachment/file/197354/Retrospective.png\", \"width\": 100, \"height\": 100}, {\"id\": 0, \"url\": \"https://images.g2crowd.com/uploads/attachment/file/197354/Retrospective.png\", \"width\": 100, \"height\": 100}]', '{\"id\": 0, \"url\": \"https://images.g2crowd.com/uploads/attachment/file/197354/Retrospective.png\", \"width\": 100, \"height\": 100}', 0, 1, '2023-07-25 10:01:57', '2023-07-25 10:01:57'),
(34, NULL, 'HOME Saigon - HOME Vietnamese Restaurant', '216/4 Đ. Điện Biên Phủ, Phường 4, Quận 3, Thành phố Hồ Chí Minh 700000, Vietnam', 'normal', NULL, NULL, NULL, '[{\"id\": 0, \"url\": \"https://images.g2crowd.com/uploads/attachment/file/197354/Retrospective.png\", \"width\": 100, \"height\": 100}, {\"id\": 0, \"url\": \"https://images.g2crowd.com/uploads/attachment/file/197354/Retrospective.png\", \"width\": 100, \"height\": 100}]', '{\"id\": 0, \"url\": \"https://images.g2crowd.com/uploads/attachment/file/197354/Retrospective.png\", \"width\": 100, \"height\": 100}', 0, 1, '2023-07-25 10:02:57', '2023-07-25 10:02:57');


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;