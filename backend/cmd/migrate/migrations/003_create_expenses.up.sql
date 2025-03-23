
DROP TABLE IF EXISTS expenses;
DROP TABLE IF EXISTS expense_splits;

-- Create expenses table
CREATE TABLE IF NOT EXISTS `expenses` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `description` VARCHAR(255) NOT NULL,
    `amount` DECIMAL(10,2) NOT NULL,
    `paid_by` BIGINT UNSIGNED NOT NULL,
    `group_id` BIGINT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`paid_by`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`group_id`) REFERENCES `groups`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create expense_splits table
CREATE TABLE IF NOT EXISTS `expense_splits` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `expense_id` BIGINT NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,    
    `amount` DECIMAL(10,2) NOT NULL,
    `status` ENUM('pending', 'paid') NOT NULL DEFAULT 'pending',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`expense_id`) REFERENCES `expenses`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
