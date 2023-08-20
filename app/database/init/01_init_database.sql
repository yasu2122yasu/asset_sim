CREATE TABLE sample_table (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO sample_table (name, created_at, updated_at) VALUES
('Sample Name 1', NOW(), NOW()),
('Sample Name 2', NOW(), NOW()),
('Sample Name 3', NOW(), NOW()),
('Sample Name 4', NOW(), NOW()),
('Sample Name 5', NOW(), NOW());