-- Create the Maelstrom Platform API tablespaces.
-- PostgreSQL 16

-- Tablespaces are created at the database
-- cluster level.

CREATE TABLESPACE maelstrom
       LOCATION '/var/pg/data1';
       
CREATE TABLESPACE maelstrom_archive
       LOCATION '/var/pg/data2';
       
CREATE TABLESPACE maelstrom_audit
       LOCATION '/var/pg/data3';
       
CREATE TABLESPACE maelstrom_dw
       LOCATION '/var/pg/data4';
       
CREATE TABLESPACE maelstrom_log
       LOCATION '/var/pg/data5';

CREATE TABLESPACE maelstrom_staging
       LOCATION '/var/pg/data6';
