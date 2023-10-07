# Deploy the Maelstrom Platform API databases.

# This script will deploy and configure the
# databases for a new install of the Maelstrom
# Platform API.  Upgrades of an existing
# installation are handled with another script.

# This script assumes a PostgreSQL server is installed
# and running, and that a role maelstrom_install and a
# database maelstrom_install have been created manually.
# Please see the document INSTALL located in the root
# of this repository.

echo Reset the database...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom_install -f 0_pg_init_db_reset.sql

echo Create the tablespaces...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom_install -f 1_pg_init_db_tablespaces.sql

echo Create the databases...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom_install -f 2_pg_init_db.sql

echo Create the database schemas...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom -f 3_pg_init_db_schema_maelstrom.sql

echo Create the database schemas...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom_archive -f 4_pg_init_db_schema_maelstrom_archive.sql

echo Create the database schemas...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom_audit -f 5_pg_init_db_schema_maelstrom_audit.sql

echo Create the database schemas...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom_dw -f 6_pg_init_db_schema_maelstrom_dw.sql

echo Create the database schemas...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom_log -f 7_pg_init_db_schema_maelstrom_log.sql

echo Create the database schemas...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom_staging -f 8_pg_init_db_schema_maelstrom_staging.sql

echo Configure all databases...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom_install -f 8a_pg_init_db_configure_all_databases.sql

echo Create the functions and procedures...
PGPASSWORD=mi psql -U maelstrom_install -d maelstrom -f 9_pg_init_db_maelstrom_procedures.sql
