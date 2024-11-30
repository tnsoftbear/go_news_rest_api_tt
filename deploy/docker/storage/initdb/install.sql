CREATE TABLE `News` (
  `Id` bigint NOT NULL,
  `Title` tinytext NOT NULL,
  `Content` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `NewsCategories` (
  `NewsId` bigint NOT NULL,
  `CategoryId` bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE `News`
  ADD PRIMARY KEY (`Id`);

ALTER TABLE `NewsCategories`
  ADD PRIMARY KEY (`NewsId`,`CategoryId`);

ALTER TABLE `News`
  MODIFY `Id` bigint NOT NULL AUTO_INCREMENT;