-- Create the Maelstrom Platform API databases.
-- PostgreSQL 16

-- To avoid a chicken-and-the-egg problem,
-- we create the databases now, and subsequent
-- scripts will handle the final configuration.

-- Databases are created at the database
-- cluster level.

-- Create the primary database.
CREATE DATABASE maelstrom;

-- Create the archive database which will store
-- old data from the primary database.
CREATE DATABASE maelstrom_archive;

-- Create the audit database which will store
-- who changed what and when in the primary
-- database.
CREATE DATABASE maelstrom_audit;

-- Create the data warehouse database.
CREATE DATABASE maelstrom_dw;

-- Create the log database which will store
-- activity from the primary database.
CREATE DATABASE maelstrom_log;

-- Create the staging database which will
-- store temporary data used for a variety
-- of purposes.
CREATE DATABASE maelstrom_staging;
