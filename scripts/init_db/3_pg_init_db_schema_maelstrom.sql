-- Create the maelstrom database schema.
-- PostgreSQL 16

-- Database roles are created at the database
-- cluster level.
CREATE ROLE maelstrom
WITH
	LOGIN
	SUPERUSER
	NOCREATEDB
	NOCREATEROLE
	NOINHERIT
	NOBYPASSRLS
	PASSWORD	NULL;

-- Database schemas are created at the
-- database level.
CREATE SCHEMA
IF NOT EXISTS maelstrom
AUTHORIZATION maelstrom;

-- Grant permissions to role maelstrom.
GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES
IN SCHEMA maelstrom
TO maelstrom;

-- Extensions are created at the
-- database level.  This extension is
-- needed for uuid column support.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the Issue relation in the maelstrom
-- schema in the maelstrom database.
CREATE TABLE IF NOT EXISTS maelstrom.issue (
              sysid                       UUID                        DEFAULT UUID_generate_v4(),        -- System identifier
              identifier                  VARCHAR(16)                 NOT NULL,                          --
              summary_brief               VARCHAR(50)                 NOT NULL,                          --
              summary_long                VARCHAR(72)                 NOT NULL,                          --
              problem_description         VARCHAR(2048)               NOT NULL,                          --
              work_around                 VARCHAR(4096)               NOT NULL,                          --
              steps_to_reproduce          VARCHAR(512)                NOT NULL,                          --
              kind                        UUID                        NOT NULL,                          --
              date_found                  TIMESTAMP WITH TIME ZONE    NOT NULL,                          --
              date_reported               TIMESTAMP WITH TIME ZONE    NOT NULL,                          --
              date_input                  TIMESTAMP WITH TIME ZONE    NOT NULL,                          --
              found_by_primary            UUID                        NOT NULL,                          --
              found_by_team_primary       UUID                        NOT NULL,                          --
              reported_by_primary         UUID                        NOT NULL,                          --
              reported_by_team_primary    UUID                        NOT NULL,                          --
              input_by_primary            UUID                        NOT NULL,                          --
              input_by_team_primary       UUID                        NOT NULL,                          --
              severity                    UUID                        NOT NULL,                          --
              priority                    UUID                        NOT NULL,                          --
              organization_value          UUID                        NOT NULL,                          --
              current_status              UUID                        NOT NULL,                          --
              current_state               UUID                        NOT NULL,                          --
              is_resolved                 BOOLEAN                     NOT NULL,                          --
              date_resolved               TIMESTAMP WITH TIME ZONE    NOT NULL,                          --
              resolved_by_primary         UUID                        NOT NULL,                          --
              resolved_by_team_primary    UUID                        NOT NULL,                          --
              resolution_due_date         TIMESTAMP WITH TIME ZONE    NOT NULL,                          --
              resolution_effort_unit      UUID                        NOT NULL,                          --
              resolution_effort           VARCHAR(10)                 NOT NULL,                          --
              estimated_resolution_date   TIMESTAMP WITH TIME ZONE    NOT NULL,                          --
              target_resolution_date      TIMESTAMP WITH TIME ZONE    NOT NULL,                          --
              root_cause_analysis         VARCHAR(2048)               NOT NULL,                          --
              fix_description             VARCHAR(2048)               NOT NULL,                          --
              assigned_to_primary         UUID                        NOT NULL,                          --
              assigned_to_team_primary    UUID                        NOT NULL,                          --
              target_original_build       UUID                        NOT NULL,                          --
              estimated_original_build    UUID                        NOT NULL,                          --
              actual_original_build       UUID                        NOT NULL,                          --
              target_original_release     UUID                        NOT NULL,                          --
              estimated_original_release  UUID                        NOT NULL,                          --
              actual_original_release     UUID                        NOT NULL,                          --
)
