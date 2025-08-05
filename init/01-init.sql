-- Database initialization script
-- This script runs automatically when the PostgreSQL container starts for the first time

-- Ensure the database is properly set up
CREATE DATABASE IF NOT EXISTS tododb;

-- Grant all privileges to todo_admin user
GRANT ALL PRIVILEGES ON SCHEMA public TO todo_admin;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO todo_admin;

GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO todo_admin;

-- Set default privileges for future objects
ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT ALL ON TABLES TO todo_admin;

ALTER DEFAULT PRIVILEGES IN SCHEMA public
GRANT ALL ON SEQUENCES TO todo_admin;

-- Log completion
SELECT 'Database initialization completed successfully' AS status;