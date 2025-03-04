CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the user if it doesn't exist
DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'myuser') THEN
      CREATE USER myuser WITH PASSWORD 'mypassword';
   END IF;
END
$do$;

-- Grant privileges to the user
ALTER USER myuser WITH SUPERUSER;

-- Create the database if it doesn't exist
CREATE DATABASE mydb
    WITH 
    OWNER = myuser
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TEMPLATE template0;

-- Connect to the new database
\c mydb;

-- Create schema and grant privileges
CREATE SCHEMA IF NOT EXISTS public;
GRANT ALL ON SCHEMA public TO myuser;
GRANT ALL ON ALL TABLES IN SCHEMA public TO myuser;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO myuser;

-- Create your tables here
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Add any additional tables or initial data here
