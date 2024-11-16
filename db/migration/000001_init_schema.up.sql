-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Nov 16, 2024 at 07:28 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.0.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_crud`
--

-- --------------------------------------------------------

--
-- Table structure for table `tickets`
--

CREATE TABLE `tickets` (
  `id` varchar(36) NOT NULL,
  `ticket_title` varchar(255) DEFAULT NULL,
  `ticket_message` varchar(255) DEFAULT NULL,
  `user_id` varchar(50) DEFAULT NULL,
  `status_code` varchar(3) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `tickets`
--

INSERT INTO `tickets` (`id`, `ticket_title`, `ticket_message`, `user_id`, `status_code`, `created_at`, `updated_at`) VALUES
('382ae672-621e-40e5-ae28-ac2bf2d21341', 'ticket number three', ' Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum', '1', 'opn', '2024-11-15 13:11:58', '2024-11-15 13:11:58'),
('5975d42a-b74b-4f75-a43b-37e9d74ba224', 'ticket number three', ' Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum', '4', 'opn', '2024-11-15 13:12:02', '2024-11-15 13:12:02'),
('a54bef09-9cf3-4c4f-95dc-1613effa88c1', 'ticket number three', ' Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum', '5', 'opn', '2024-11-15 13:12:05', '2024-11-15 13:12:05'),
('ad4277c5-8752-4a4c-951e-30cd4011491e', 'ticket number three', ' Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum', '1', 'cld', '2024-11-15 13:11:49', '2024-11-15 13:24:29'),
('bd2d5682-250a-40e7-85a2-66c49d634388', 'ticket number three', ' Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum Lorem Ipsum', '2', 'opn', '2024-11-15 13:11:54', '2024-11-15 13:11:54');

-- --------------------------------------------------------

--
-- Table structure for table `ticket_statuses`
--

CREATE TABLE `ticket_statuses` (
  `id` varchar(36) NOT NULL,
  `code` varchar(3) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ticket_statuses`
--

INSERT INTO `ticket_statuses` (`id`, `code`, `status`, `created_at`, `updated_at`) VALUES
('a0d901e0-a319-11ef-a474-ba3b71d7a374', 'opn', 'Open', '2024-11-15 13:19:53', '2024-11-15 13:19:53'),
('a0d90c26-a319-11ef-a474-ba3b71d7a374', 'cld', 'Closed', '2024-11-15 13:19:53', '2024-11-15 13:19:53'),
('a90ddc46-a319-11ef-a474-ba3b71d7a374', 'asn', 'Assigned', '2024-11-15 13:20:07', '2024-11-15 13:20:07');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `username` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `status` varchar(50) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `firstname`, `lastname`, `phone`, `avatar`, `email`, `username`, `password`, `status`, `created_at`, `updated_at`) VALUES
('079af6eb-43be-4cce-b9bd-53a395021188', 'banam', 'anam', '', '', 'anam7', 'anam7', '$2a$14$gzkW8bZD4WifZsCCfEcI3OnJt/hV0zh5aDlD.k4UAkSvQMUDW3BX2', '', '2024-11-13 10:22:39', '2024-11-13 10:22:39');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tickets`
--
ALTER TABLE `tickets`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ticket_statuses`
--
ALTER TABLE `ticket_statuses`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `username` (`username`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;