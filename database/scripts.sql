CREATE IF NOT EXISTS TABLE `orders` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `discount_percentage` decimal(3,2) NOT NULL,
  `total_cost` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `items` (
  `product_id` bigint(20) NOT NULL,
  `order_id` bigint(20) NOT NULL,
  `quantity` bigint(20) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `total_cost` decimal(10,2) NOT NULL,
  KEY `order_id` (`order_id`),
  CONSTRAINT `items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;