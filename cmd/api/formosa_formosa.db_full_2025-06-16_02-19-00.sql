-- MySQL dump 10.13  Distrib 8.0.42, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: formosa_formosa.db
-- ------------------------------------------------------
-- Server version	8.0.42-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `formosa_formosa.db`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `formosa_formosa.db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `formosa_formosa.db`;

--
-- Table structure for table `attendances`
--

DROP TABLE IF EXISTS `attendances`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `attendances` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `employee_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `attendance` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `hours` bigint NOT NULL,
  `date` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `amount` float NOT NULL,
  `is_holiday` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_employees_attendances` (`employee_id`),
  CONSTRAINT `fk_employees_attendances` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attendances`
--

LOCK TABLES `attendances` WRITE;
/*!40000 ALTER TABLE `attendances` DISABLE KEYS */;
/*!40000 ALTER TABLE `attendances` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `clients`
--

DROP TABLE IF EXISTS `clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `clients` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cuil` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `dni` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_clients_cuil` (`cuil`),
  UNIQUE KEY `uni_clients_dni` (`dni`),
  UNIQUE KEY `uni_clients_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `clients`
--

LOCK TABLES `clients` WRITE;
/*!40000 ALTER TABLE `clients` DISABLE KEYS */;
/*!40000 ALTER TABLE `clients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `employees`
--

DROP TABLE IF EXISTS `employees`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `employees` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_employees_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employees`
--

LOCK TABLES `employees` WRITE;
/*!40000 ALTER TABLE `employees` DISABLE KEYS */;
/*!40000 ALTER TABLE `employees` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `expenses`
--

DROP TABLE IF EXISTS `expenses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `expenses` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `details` longtext COLLATE utf8mb4_unicode_ci,
  `supplier_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `movement_type_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `amount` float NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_expenses_supplier` (`supplier_id`),
  KEY `fk_expenses_movement_type` (`movement_type_id`),
  CONSTRAINT `fk_expenses_movement_type` FOREIGN KEY (`movement_type_id`) REFERENCES `movement_types` (`id`),
  CONSTRAINT `fk_expenses_supplier` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `expenses`
--

LOCK TABLES `expenses` WRITE;
/*!40000 ALTER TABLE `expenses` DISABLE KEYS */;
/*!40000 ALTER TABLE `expenses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `income_services`
--

DROP TABLE IF EXISTS `income_services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `income_services` (
  `service_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `income_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`service_id`,`income_id`),
  KEY `fk_income_services_income` (`income_id`),
  CONSTRAINT `fk_income_services_income` FOREIGN KEY (`income_id`) REFERENCES `incomes` (`id`),
  CONSTRAINT `fk_income_services_service` FOREIGN KEY (`service_id`) REFERENCES `services` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `income_services`
--

LOCK TABLES `income_services` WRITE;
/*!40000 ALTER TABLE `income_services` DISABLE KEYS */;
/*!40000 ALTER TABLE `income_services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `incomes`
--

DROP TABLE IF EXISTS `incomes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `incomes` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ticket` longtext COLLATE utf8mb4_unicode_ci,
  `details` longtext COLLATE utf8mb4_unicode_ci,
  `client_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `vehicle_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `employee_id` varchar(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `amount` float NOT NULL,
  `movement_type_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_incomes_vehicle` (`vehicle_id`),
  KEY `fk_incomes_employee` (`employee_id`),
  KEY `fk_incomes_movement_type` (`movement_type_id`),
  KEY `fk_incomes_client` (`client_id`),
  CONSTRAINT `fk_incomes_client` FOREIGN KEY (`client_id`) REFERENCES `clients` (`id`),
  CONSTRAINT `fk_incomes_employee` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `fk_incomes_movement_type` FOREIGN KEY (`movement_type_id`) REFERENCES `movement_types` (`id`),
  CONSTRAINT `fk_incomes_vehicle` FOREIGN KEY (`vehicle_id`) REFERENCES `vehicles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `incomes`
--

LOCK TABLES `incomes` WRITE;
/*!40000 ALTER TABLE `incomes` DISABLE KEYS */;
/*!40000 ALTER TABLE `incomes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `members`
--

DROP TABLE IF EXISTS `members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `members` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `role_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_members_username` (`username`),
  UNIQUE KEY `uni_members_email` (`email`),
  KEY `idx_members_deleted_at` (`deleted_at`),
  KEY `fk_members_role` (`role_id`),
  CONSTRAINT `fk_members_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `members`
--

LOCK TABLES `members` WRITE;
/*!40000 ALTER TABLE `members` DISABLE KEYS */;
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movement_types`
--

DROP TABLE IF EXISTS `movement_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movement_types` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_income` tinyint(1) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movement_types`
--

LOCK TABLES `movement_types` WRITE;
/*!40000 ALTER TABLE `movement_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `movement_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permissions`
--

DROP TABLE IF EXISTS `permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `permissions` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` longtext COLLATE utf8mb4_unicode_ci,
  `details` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `group` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permissions`
--

LOCK TABLES `permissions` WRITE;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
INSERT INTO `permissions` VALUES ('030b2566-4e09-47dc-9244-f40bfbc34565','AOV','Ver ajuste stock','Compras'),('0314b48b-6758-4baf-83f4-7396c5336526','PAR','Ver resumen ejecutivo','Panel'),('07567b70-c0d2-4356-a7fc-d81d5c9b19a8','EGA','Agregar egreso','Egresos/Ingresos'),('08811052-fcf0-4f9e-9653-6a63b6a55b1d','ING','Generales (ventas, compras, cajas)','Informes'),('0a4ca795-2ab8-4129-aec7-7f19cd0a5327','CXX','Arqueo de caja','Caja Rápida'),('0af24027-98df-4f86-8159-cb2e35815f85','PRE','Editar proveedor','Proveedores'),('0e235cd0-94f6-4006-8f41-7517da306bfc','CSP','Enviar email proveedor','Compras'),('1bb0be28-f880-4c6f-93ee-9a85d6a5dc7c','PDD','Eliminar producto','Productos'),('1c24a664-7b0e-4ed0-997f-463a6b4be613','INI','Ingresos','Informes'),('1d1835fa-cf03-45de-9d75-18390c918f23','VED2','Editar descuento','Ventas'),('2317c91b-449e-4c01-829e-ec07622ccf11','IGE','Editar ingreso','Egresos/Ingresos'),('258f288c-73d6-4d8c-baec-0d38b88b35e0','VEV','Ver ventas','Ventas'),('2920eaeb-68d8-4f8a-9ad8-6d1673ba1f73','ARC','Crear factura ARCA','ARCA'),('2ae92b12-f820-4dc2-9e0f-4f5b80976309','IGD','Eliminar Ingreso','Egresos/Ingresos'),('31e3b6e0-6162-493d-9200-9ed0e0de3f55','USE','Editar usuarios','Usuarios'),('3be3a3dc-dbe0-426a-b1e9-70b7a4c580de','CEP','Editar pago proveedor','Compras'),('3d346478-67ec-43ab-9d56-f96d0a903f82','INS','Stock','Informes'),('415d38af-b2ec-4855-a2b5-ac96d3c11dc8','PDP','Ver precio compra','Productos'),('44b7e96c-4ef2-479e-8d97-db21d89d94f7','IGA','Agregar ingreso','Egresos/Ingresos'),('4536a01f-cfd0-45b6-b1ae-13e0ebcc4b94','RLE','Editar roles','Roles'),('476ca74d-d7e7-4fc0-a5aa-5733ad4da294','VEE','Editar venta','Ventas'),('4a112cd8-187f-451e-9979-6719afe0c5da','VMP','Modificar precios','Ventas'),('4cb5a2b7-25f0-4673-af3c-3377301b9793','AOA','Agregar ajuste','Compras'),('570cff36-b350-4e2f-a192-a5c1c8a4156d','VSC','Enviar email a cliente','Ventas'),('5abb0f80-3390-4792-bbec-0b63ab609d64','PDS','Añadir stock inicial','Productos'),('620542f7-5ea7-421e-89ee-d1dd70521938','EGD','Eliminar egreso','Egresos/Ingresos'),('671f0129-6409-46e4-b2ba-02ffc9afe87d','AOE','Editar ajuste','Compras'),('6ad7ca97-577b-48c0-87c7-692301f17aed','RLV','Ver roles','Roles'),('6ae577a0-ca42-46af-bc83-8051ef93b2d9','ARV','Ver factura ARCA','ARCA'),('6c819bd7-d7d5-4f4f-8152-f1c2c2792209','RLD','Eliminar roles','Roles'),('6d51e7c6-9091-4532-9476-ab49a961378c','MCA','Agregar marcas/categorías','Marcas/Categorías'),('71db38bf-c38b-4f48-bae0-514ac89fb3e6','VMD','Modificar descuentos','Ventas'),('7512be16-3dc3-4150-8b65-55143ae431bf','CTX','Acceso módulo contable','Cuentas'),('75ddfc75-87e8-47f2-9219-6add452eb616','MCD','Eliminar marcas/categorías','Marcas/Categorías'),('79269555-5e73-4863-9af4-ffe24b2d0c38','EGE','Editar egreso','Egresos/Ingresos'),('7963c013-e5c3-416b-bf3a-be2f0e3f8b99','PAV','Ver ventas en dashboard','Panel'),('7a0b5c30-46d3-46b0-87cb-c02dacfb5ba4','MCV','Ver marcas/categorías','Marcas/Categorías'),('7cc6b647-9a2f-4098-9602-8f29aec744b5','INR','Ranking','Informes'),('8302d859-412d-4459-868b-03f2d2c11d37','RLA','Agregar roles','Roles'),('89a68f0f-b2a9-40d2-8328-23afbd5f62f2','COA','Agregar compra','Compras'),('8d64af70-0169-4a55-852a-e9f3f93436a5','VED','Eliminar venta','Ventas'),('97007492-2f33-4f82-a71e-edbca73be2cc','CLE','Editar cliente','Clientes'),('9afd88c3-98da-421a-a672-8023adc04b84','PDV','Ver producto','Productos'),('9eb1e876-714a-47e3-be98-03a3dab76597','CLD','Eliminar cliente','Clientes'),('a3755985-340c-4367-874d-2be97416fc4e','PDE','Editar producto','Productos'),('a83034af-62ec-4cf6-8798-12e20abb3ca8','COV','Ver compras','Compras'),('a8c3f000-5a77-4ffc-b7f5-62a670168daa','CLV','Ver cliente','Clientes'),('aa5e039e-b276-4ab7-b196-85ce55fbe24a','PDA','Agregar producto','Productos'),('baaf9de3-50f0-445f-a1bb-8295ebeef277','VDP','Eliminar pago','Ventas'),('c217403c-66af-48fb-b878-08cad5a3b7c6','USD','Eliminar usuarios','Usuarios'),('c2351daa-a1ed-4993-824c-53b339d906ca','CDP','Eliminar pago proveedor','Compras'),('c4bac39c-368d-4a55-9fa0-8ab7594d6dd2','PRV','Ver proveedor','Proveedores'),('c5658da5-0f74-4ce0-a582-ddfddcdf1c9d','PDI','Editar IVA','Productos'),('c7a9d74e-24e9-4bbb-a703-1fbb423190a2','PRD','Eliminar proveedor','Proveedores'),('ca80e20f-04f6-48fa-a840-592405276f03','USV','Ver usuarios','Usuarios'),('cb7e8c5b-79c5-40dc-8b61-30c77c9a3131','INP','Reporte pagos/cobros','Informes'),('cd5838c3-d255-4628-8f63-463ec216afe7','VAP','Agregar pago','Ventas'),('cecf99b0-8d9b-4f64-8cec-42f137de05c9','USA','Agregar usuarios','Usuarios'),('cedbd9be-9f0b-4bc6-bad8-88b2b1da29ed','INE','Egresos','Informes'),('d069b6f1-9bd5-40d9-b128-6487975010b9','VAD','Agregar descuento','Ventas'),('d79d7e12-a44a-48c8-9723-be8d9a2144f7','CAP','Agregar pago proveedor','Compras'),('d982e46e-cb2c-444e-b7e0-1c211bc4a218','PRA','Agregar proveedor','Proveedores'),('e0f7c621-dd44-4ca0-81e9-d540664dd49b','CLA','Agregar cliente','Clientes'),('e571a9e8-fab8-4fd1-8f96-302c95421aff','AOD','Eliminar ajuste','Compras'),('ed9f2d2a-7737-4cf0-a5e6-8b331fb2ebae','COD','Eliminar compra','Compras'),('edcfc32d-949a-4723-a097-69d05b784f74','VDD','Eliminar descuento','Ventas'),('ee873afe-f896-4bec-bddb-9f7f46d87045','VEP','Editar pago','Ventas'),('f04873be-1961-4fc9-a26a-78dc65f34722','COE','Editar compra','Compras'),('f3b260b3-2b3c-4d4f-8fc7-93f1ee03c15c','VEA','Agregar venta','Ventas'),('f7ba72bd-99a1-410d-8a6d-2fc19a864ee2','INC','Cuenta corriente','Informes'),('fdff04fa-0c54-4fe0-9bb7-452b27aa98a7','VEX','Acceso a Caja Rápida','Ventas'),('fef2488a-88ee-4b5c-8489-cd2e484d9577','MCE','Editar marcas/categorías','Marcas/Categorías');
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `identifier` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `stock` float NOT NULL DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_products_identifier` (`identifier`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `purchase_orders`
--

DROP TABLE IF EXISTS `purchase_orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `purchase_orders` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `order_number` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `order_date` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `amount` float NOT NULL,
  `supplier_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_purchase_orders_supplier` (`supplier_id`),
  CONSTRAINT `fk_purchase_orders_supplier` FOREIGN KEY (`supplier_id`) REFERENCES `suppliers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `purchase_orders`
--

LOCK TABLES `purchase_orders` WRITE;
/*!40000 ALTER TABLE `purchase_orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `purchase_orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `purchase_products`
--

DROP TABLE IF EXISTS `purchase_products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `purchase_products` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `product_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `purchase_order_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `expired_at` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `unit_price` float NOT NULL,
  `quantity` bigint NOT NULL,
  `total_price` float NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_purchase_products_product` (`product_id`),
  KEY `fk_purchase_orders_purchase_products` (`purchase_order_id`),
  CONSTRAINT `fk_purchase_orders_purchase_products` FOREIGN KEY (`purchase_order_id`) REFERENCES `purchase_orders` (`id`),
  CONSTRAINT `fk_purchase_products_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `purchase_products`
--

LOCK TABLES `purchase_products` WRITE;
/*!40000 ALTER TABLE `purchase_products` DISABLE KEYS */;
/*!40000 ALTER TABLE `purchase_products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `resume_expenses`
--

DROP TABLE IF EXISTS `resume_expenses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `resume_expenses` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date` datetime(3) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `resume_expenses`
--

LOCK TABLES `resume_expenses` WRITE;
/*!40000 ALTER TABLE `resume_expenses` DISABLE KEYS */;
/*!40000 ALTER TABLE `resume_expenses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `resume_incomes`
--

DROP TABLE IF EXISTS `resume_incomes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `resume_incomes` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date` datetime(3) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `resume_incomes`
--

LOCK TABLES `resume_incomes` WRITE;
/*!40000 ALTER TABLE `resume_incomes` DISABLE KEYS */;
/*!40000 ALTER TABLE `resume_incomes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_permissions`
--

DROP TABLE IF EXISTS `role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_permissions` (
  `role_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `permission_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`role_id`,`permission_id`),
  KEY `fk_role_permissions_permission` (`permission_id`),
  CONSTRAINT `fk_role_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_permissions`
--

LOCK TABLES `role_permissions` WRITE;
/*!40000 ALTER TABLE `role_permissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `role_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_roles_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `services`
--

DROP TABLE IF EXISTS `services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `services` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_services_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `services`
--

LOCK TABLES `services` WRITE;
/*!40000 ALTER TABLE `services` DISABLE KEYS */;
/*!40000 ALTER TABLE `services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `suppliers`
--

DROP TABLE IF EXISTS `suppliers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `suppliers` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `suppliers`
--

LOCK TABLES `suppliers` WRITE;
/*!40000 ALTER TABLE `suppliers` DISABLE KEYS */;
/*!40000 ALTER TABLE `suppliers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `vehicles`
--

DROP TABLE IF EXISTS `vehicles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `vehicles` (
  `id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `brand` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `model` longtext COLLATE utf8mb4_unicode_ci,
  `color` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `year` longtext COLLATE utf8mb4_unicode_ci,
  `domain` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `client_id` varchar(36) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_vehicles_domain` (`domain`),
  KEY `fk_clients_vehicles` (`client_id`),
  CONSTRAINT `fk_clients_vehicles` FOREIGN KEY (`client_id`) REFERENCES `clients` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `vehicles`
--

LOCK TABLES `vehicles` WRITE;
/*!40000 ALTER TABLE `vehicles` DISABLE KEYS */;
/*!40000 ALTER TABLE `vehicles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'formosa_formosa.db'
--

--
-- Dumping routines for database 'formosa_formosa.db'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-06-16  2:19:00
