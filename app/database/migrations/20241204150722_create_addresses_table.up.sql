CREATE TABLE addresses (
    id INT AUTO_INCREMENT PRIMARY KEY,              
    user_id INT NOT NULL,                           
    street_address VARCHAR(255) NOT NULL,         
    city VARCHAR(100) NOT NULL,                      
    state VARCHAR(100),                             
    postal_code VARCHAR(20) NOT NULL,                
    country VARCHAR(100) NOT NULL DEFAULT 'INA',   
    is_primary BOOLEAN DEFAULT FALSE,               
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    deleted_at TIMESTAMP NULL DEFAULT NULL
);
