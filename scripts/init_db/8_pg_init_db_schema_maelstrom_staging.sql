-- Create the maelstrom_staging database relations.
-- PostgreSQL 16

-- Database roles are created at the database
-- cluster level.
CREATE ROLE maelstrom_staging
WITH
	LOGIN
	NOSUPERUSER
	NOCREATEDB
	NOCREATEROLE
	NOINHERIT
	NOBYPASSRLS
	PASSWORD	NULL;

-- Database schemas are created at the
-- database level.
CREATE SCHEMA
IF NOT EXISTS maelstrom_staging
AUTHORIZATION maelstrom_staging;

-- Grant permissions to role maelstrom_staging.
GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES
IN SCHEMA maelstrom_staging
TO maelstrom_staging;

-- Create the TBD relation in the maelstrom_staging
-- schema in the maelstrom_staging database.
