/*
 Navicat Premium Data Transfer

 Source Server         : MySQL-Docker
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3310
 Source Schema         : db_demo

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 15/07/2023 20:36:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sort` bigint NULL DEFAULT 0,
  `status` smallint NULL DEFAULT 1,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `id_card_number` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `full_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `legal_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `place_of_birth` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `date_of_birth` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `salary` double NULL DEFAULT NULL,
  `id_card_photo` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `selfie_photo` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_customer_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_customer_id_card_number`(`id_card_number` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customer
-- ----------------------------
INSERT INTO `customer` VALUES ('71e11445-8f94-493d-bd19-d7e43a1e576c', 1, 1, '2023-07-14 14:46:02.066', '2023-07-14 14:46:02.067', NULL, '3211111111111111', 'Budi', 'Budi', 'Palembang', '1987-11-22', 3000000, 'budi-ktp-11223344.jpg', 'budi-selfie-11223355.jpg');
INSERT INTO `customer` VALUES ('7a49c137-84d4-433f-810c-3188ddef783f', 2, 1, '2023-07-14 14:46:02.066', '2023-07-14 14:46:02.067', NULL, '3222222222222222', 'Annisa', 'Annisa', 'Bandung', '1992-01-27', 12000000, 'annisa-ktp-11223366.jpg', 'budi-selfie-11223377.jpg');

-- ----------------------------
-- Table structure for installment
-- ----------------------------
DROP TABLE IF EXISTS `installment`;
CREATE TABLE `installment`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sort` bigint NULL DEFAULT 0,
  `status` smallint NULL DEFAULT 1,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `contract_number` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `otr_amt` double NULL DEFAULT NULL,
  `admin_fee` double NULL DEFAULT NULL,
  `installment_amt` double NOT NULL,
  `interest_amt` double NULL DEFAULT NULL,
  `total_amt` double NULL DEFAULT NULL,
  `asset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `payment_status` smallint NULL DEFAULT 0,
  `customer_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_installment_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_installment_contract_number`(`contract_number` ASC) USING BTREE,
  INDEX `fk_installment_customer`(`customer_id` ASC) USING BTREE,
  CONSTRAINT `fk_installment_customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of installment
-- ----------------------------
INSERT INTO `installment` VALUES ('1ed6b5b8-6bae-45b2-b702-ef9d42da4090', 2, 1, '2023-07-14 20:22:25.485', '2023-07-14 20:22:25.485', NULL, '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 0, '7a49c137-84d4-433f-810c-3188ddef783f');
INSERT INTO `installment` VALUES ('2a276436-8820-4a8b-b67f-8002b4c23e60', 2, 1, '2023-07-14 20:22:25.483', '2023-07-14 20:22:25.485', NULL, '2307-09377109', 500000, 2500, 166667, 1389, 170556, 'LED TV', 0, '71e11445-8f94-493d-bd19-d7e43a1e576c');
INSERT INTO `installment` VALUES ('4c876e95-0101-44a9-b245-e1bc7774674c', 3, 1, '2023-07-14 20:22:25.483', '2023-07-14 20:22:25.485', NULL, '2307-09377109', 500000, 2500, 166667, 1389, 170556, 'LED TV', 0, '71e11445-8f94-493d-bd19-d7e43a1e576c');
INSERT INTO `installment` VALUES ('4f8ff573-b574-4a45-ad22-cfd2e4a53d6d', 3, 1, '2023-07-14 20:22:25.485', '2023-07-14 20:22:25.485', NULL, '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 0, '7a49c137-84d4-433f-810c-3188ddef783f');
INSERT INTO `installment` VALUES ('96f3ebf6-e5e2-4639-85d8-a7461daec608', 1, 1, '2023-07-14 20:22:25.483', '2023-07-14 20:22:25.485', NULL, '2307-09377109', 500000, 2500, 166667, 1389, 170556, 'LED TV', 0, '71e11445-8f94-493d-bd19-d7e43a1e576c');
INSERT INTO `installment` VALUES ('f62bb008-4628-4b7d-8793-e6ce30c73a21', 4, 1, '2023-07-14 20:22:25.485', '2023-07-14 20:22:25.485', NULL, '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 0, '7a49c137-84d4-433f-810c-3188ddef783f');
INSERT INTO `installment` VALUES ('f78079a4-de98-4be9-b424-55416fe5fd6d', 1, 1, '2023-07-14 20:22:25.485', '2023-07-15 15:05:40.725', NULL, '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 1, '7a49c137-84d4-433f-810c-3188ddef783f');

-- ----------------------------
-- Table structure for loan_application
-- ----------------------------
DROP TABLE IF EXISTS `loan_application`;
CREATE TABLE `loan_application`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sort` bigint NULL DEFAULT 0,
  `status` smallint NULL DEFAULT 1,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `limit` double NULL DEFAULT NULL,
  `asset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `approved_date` datetime NULL DEFAULT NULL,
  `approved_by` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `customer_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_loan_application_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_loan_application_customer`(`customer_id` ASC) USING BTREE,
  CONSTRAINT `fk_loan_application_customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of loan_application
-- ----------------------------
INSERT INTO `loan_application` VALUES ('3644a03c-b69c-4854-828e-4178fdac987d', 1, 4, '2023-07-14 18:47:35.354', '2023-07-14 20:22:25.488', NULL, 500000, 'LED TV', '2023-07-14 18:47:35', '966dca3a-a2c4-42c9-b4d9-905abf5df40d', '71e11445-8f94-493d-bd19-d7e43a1e576c');
INSERT INTO `loan_application` VALUES ('91907819-cf4a-4a16-9e2d-b11a100ac1ea', 1, 1, '2023-07-14 18:47:35.354', '2023-07-14 18:47:35.360', NULL, 10000000, 'Mobil Daihatsu Xenia', NULL, NULL, '71e11445-8f94-493d-bd19-d7e43a1e576c');
INSERT INTO `loan_application` VALUES ('a33f2062-a637-454d-bbcc-2e2f3f2c4c8d', 2, 4, '2023-07-14 18:47:35.354', '2023-07-14 20:22:25.493', NULL, 2000000, 'Laptop', '2023-07-14 18:47:35', '966dca3a-a2c4-42c9-b4d9-905abf5df40d', '7a49c137-84d4-433f-810c-3188ddef783f');

-- ----------------------------
-- Table structure for transaction
-- ----------------------------
DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `sort` bigint NULL DEFAULT 0,
  `status` smallint NULL DEFAULT 1,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `installment_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `contract_number` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `otr_amt` double NULL DEFAULT NULL,
  `admin_fee` double NULL DEFAULT NULL,
  `installment_amt` double NOT NULL,
  `interest_amt` double NULL DEFAULT NULL,
  `total_amt` double NULL DEFAULT NULL,
  `asset_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `payment_status` smallint NULL DEFAULT 0,
  `paymet_expired_at` datetime NULL DEFAULT NULL,
  `refnum` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `customer_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `full_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `legal_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `id_card_number` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `short_description` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `long_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_transaction_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_transaction_contract_number`(`contract_number` ASC) USING BTREE,
  INDEX `fk_installment_transaction`(`installment_id` ASC) USING BTREE,
  CONSTRAINT `fk_installment_transaction` FOREIGN KEY (`installment_id`) REFERENCES `installment` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of transaction
-- ----------------------------
INSERT INTO `transaction` VALUES ('01df979b-7a47-4983-82ce-e78644e2f73e', 0, 1, '2023-07-15 14:56:14.430', '2023-07-15 14:56:26.928', '2023-07-15 14:56:26.932', 'f78079a4-de98-4be9-b424-55416fe5fd6d', '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 1, '2023-07-16 02:56:14', NULL, '7a49c137-84d4-433f-810c-3188ddef783f', 'Annisa', 'Annisa', '3222222222222222', '$2a$04$MiZwgElZeRLSk83SMi1kQeGK385Au6RJzs5BTBRQG.TIKVoe.d5x2', 'Bayar Via Dealer XYZ', 'Mismatch Token');
INSERT INTO `transaction` VALUES ('048113e8-e8b5-43ef-9554-29ea8ef84e41', 0, 1, '2023-07-15 15:05:40.720', '2023-07-15 15:08:00.898', NULL, 'f78079a4-de98-4be9-b424-55416fe5fd6d', '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 2, '2023-07-16 03:05:41', NULL, '7a49c137-84d4-433f-810c-3188ddef783f', 'Annisa', 'Annisa', '3222222222222222', '$2a$04$ea0bPX1FG3Z1YlRR/jFK6eNBTp8fOsoLpyuSg0Aq75mzYuIJ67woa', 'Bayar Via Dealer XYZ', NULL);
INSERT INTO `transaction` VALUES ('679b36e7-0c59-4e97-8538-c22d576be52a', 0, 1, '2023-07-15 14:58:07.466', '2023-07-15 14:58:15.488', '2023-07-15 14:58:15.490', 'f78079a4-de98-4be9-b424-55416fe5fd6d', '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 1, '2023-07-16 02:58:07', NULL, '7a49c137-84d4-433f-810c-3188ddef783f', 'Annisa', 'Annisa', '3222222222222222', '$2a$04$JhR63maSyHoeVxRVPSmmR.hBB25h/dqq7q0zJ6jqwjgTbuEYL07I2', 'Bayar Via Dealer XYZ', 'Mismatch Token');
INSERT INTO `transaction` VALUES ('afd54f97-6a90-4247-bae3-afbf2e2ec91c', 0, 1, '2023-07-15 15:00:03.834', '2023-07-15 15:00:15.117', '2023-07-15 15:00:15.122', 'f78079a4-de98-4be9-b424-55416fe5fd6d', '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 1, '2023-07-16 03:00:04', NULL, '7a49c137-84d4-433f-810c-3188ddef783f', 'Annisa', 'Annisa', '3222222222222222', '$2a$04$2jXHovGohKTY5z5uRtDL8.GgPeo9CVLhw9U/XL4IRVX0iCuwL99Qm', 'Bayar Via Dealer XYZ', 'Mismatch Token');
INSERT INTO `transaction` VALUES ('d1a9a6ec-4827-4b5b-9f3e-333ee111b173', 0, 1, '2023-07-15 14:59:41.409', '2023-07-15 14:59:47.980', '2023-07-15 14:59:47.987', 'f78079a4-de98-4be9-b424-55416fe5fd6d', '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 1, '2023-07-16 02:59:41', NULL, '7a49c137-84d4-433f-810c-3188ddef783f', 'Annisa', 'Annisa', '3222222222222222', '$2a$04$chi8ZuecVEsElyuIgBh96uBfEMd2.STJf0vnuOAmWE2O40e1FZMZW', 'Bayar Via Dealer XYZ', 'Mismatch Token');
INSERT INTO `transaction` VALUES ('e2484f73-ff8e-4fd1-bef5-88e5bca59c4d', 0, 1, '2023-07-15 14:48:26.090', '2023-07-15 14:55:26.052', '2023-07-15 14:55:26.056', 'f78079a4-de98-4be9-b424-55416fe5fd6d', '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 1, '2023-07-16 02:48:26', NULL, '7a49c137-84d4-433f-810c-3188ddef783f', 'Annisa', 'Annisa', '3222222222222222', '$2a$04$jsYvaK/2tZju8mXLKsdEp.73Z7mIoFm6u1R3wcU7mQPVF1dYJYPxK', 'Bayar Via Dealer XYZ', 'Mismatch Token');
INSERT INTO `transaction` VALUES ('f044df84-9898-4384-8db0-fb0214c93519', 0, 1, '2023-07-15 14:55:50.266', '2023-07-15 14:55:59.126', '2023-07-15 14:55:59.128', 'f78079a4-de98-4be9-b424-55416fe5fd6d', '2307-48234933', 2000000, 2500, 500000, 4167, 506667, 'Laptop', 1, '2023-07-16 02:55:50', NULL, '7a49c137-84d4-433f-810c-3188ddef783f', 'Annisa', 'Annisa', '3222222222222222', '$2a$04$ebMoNrMCwEuoir9JycGoXuxG72e6XrQOBZbXW.3ZFx0ZZM4jJh0Xa', 'Bayar Via Dealer XYZ', 'Mismatch Token');

SET FOREIGN_KEY_CHECKS = 1;
