CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,              -- Primary key with auto-increment
    email VARCHAR(255) NOT NULL UNIQUE,                         -- Email field, unique index
    password VARCHAR(255) NOT NULL,                             -- Password field
    remember_token VARCHAR(100) DEFAULT NULL,                   -- Nullable remember token
    status SMALLINT NOT NULL DEFAULT 1,                         -- Status with default value of 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,             -- Timestamp for creation
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp for last update
    deleted_at TIMESTAMP NULL DEFAULT NULL,                     -- Soft delete field
    INDEX idx_deleted_at (deleted_at)                           -- Index for deleted_at (to support soft deletes)
);  