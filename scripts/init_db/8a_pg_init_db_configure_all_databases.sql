-- Configure the Maelstrom Platform API databases.

-- Configure the maelstrom database.
-- Configure the database owner.
ALTER DATABASE maelstrom
OWNER TO maelstrom;

-- Configure the database default tablespace.
ALTER DATABASE maelstrom
SET TABLESPACE maelstrom;


-- Configure the maelstrom_archive database.
-- Configure the database owner.
ALTER DATABASE maelstrom_archive
OWNER TO maelstrom_archive;

-- Configure the database default tablespace.
ALTER DATABASE maelstrom_archive
SET TABLESPACE maelstrom_archive;


-- Configure the maelstrom_audit database.
-- Configure the database owner.
ALTER DATABASE maelstrom_audit
OWNER TO maelstrom_audit;

-- Configure the database default tablespace.
ALTER DATABASE maelstrom_audit
SET TABLESPACE maelstrom_audit;


-- Configure the maelstrom_dw database.
-- Configure the database owner.
ALTER DATABASE maelstrom_dw
OWNER TO maelstrom_dw;

-- Configure the database default tablespace.
ALTER DATABASE maelstrom_dw
SET TABLESPACE maelstrom_dw;


-- Configure the maelstrom_log database.
-- Configure the database owner.
ALTER DATABASE maelstrom_log
OWNER TO maelstrom_log;

-- Configure the database default tablespace.
ALTER DATABASE maelstrom_log
SET TABLESPACE maelstrom_log;


-- Configure the maelstrom_staging database.
-- Configure the database owner.
ALTER DATABASE maelstrom_staging
OWNER TO maelstrom_staging;

-- Configure the database default tablespace.
ALTER DATABASE maelstrom_staging
SET TABLESPACE maelstrom_staging;
