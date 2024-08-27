-- mydatabase.payment definition

CREATE TABLE `payment` (
  `payment_id` varchar(100) NOT NULL,
  PRIMARY KEY (`payment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- mydatabase.billing definition

CREATE TABLE `billing` (
  `billing_id` varchar(100) NOT NULL,
  `fee` int NOT NULL,
  `payment_id` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`billing_id`),
  KEY `billing_payment_FK` (`payment_id`),
  CONSTRAINT `billing_payment_FK` FOREIGN KEY (`payment_id`) REFERENCES `payment` (`payment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;