package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var MaelstromInstallDb *sql.DB
var MaelstromDb *sql.DB
var MaelstromArchiveDb *sql.DB
var MaelstromAuditDb *sql.DB
var MaelstromDwDb *sql.DB
var MaelstromLogDb *sql.DB
var MaelstromStagingDb *sql.DB

func ConnectDatabases() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error has occurred on .env file.")
	}

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	mi_db_name := os.Getenv("MAELSTROM_INSTALL_DB_NAME")
	mi_db_role := os.Getenv("MAELSTROM_INSTALL_DB_ROLE")
	mi_db_pass := os.Getenv("MAELSTROM_INSTALL_DB_PASSWORD")

	m_db_name := os.Getenv("MAELSTROM_DB_NAME")
	m_db_role := os.Getenv("MAELSTROM_DB_ROLE")
	m_db_pass := os.Getenv("MAELSTROM_DB_PASSWORD")

	mar_db_name := os.Getenv("MAELSTROM_ARCHIVE_DB_NAME")
	mar_db_role := os.Getenv("MAELSTROM_ARCHIVE_DB_ROLE")
	mar_db_pass := os.Getenv("MAELSTROM_ARCHIVE_DB_PASSWORD")

	mau_db_name := os.Getenv("MAELSTROM_AUDIT_DB_NAME")
	mau_db_role := os.Getenv("MAELSTROM_AUDIT_DB_ROLE")
	mau_db_pass := os.Getenv("MAELSTROM_AUDIT_DB_PASSWORD")

	mdw_db_name := os.Getenv("MAELSTROM_DW_DB_NAME")
	mdw_db_role := os.Getenv("MAELSTROM_DW_DB_ROLE")
	mdw_db_pass := os.Getenv("MAELSTROM_DW_DB_PASSWORD")

	mlg_db_name := os.Getenv("MAELSTROM_LOG_DB_NAME")
	mlg_db_role := os.Getenv("MAELSTROM_LOG_DB_ROLE")
	mlg_db_pass := os.Getenv("MAELSTROM_LOG_DB_PASSWORD")

	mst_db_name := os.Getenv("MAELSTROM_STAGING_DB_NAME")
	mst_db_role := os.Getenv("MAELSTROM_STAGING_DB_ROLE")
	mst_db_pass := os.Getenv("MAELSTROM_STAGING_DB_PASSWORD")

	mi_cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		mi_db_role, mi_db_pass, host, port, mi_db_name)

	m_cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		m_db_role, m_db_pass, host, port, m_db_name)

	mar_cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		mar_db_role, mar_db_pass, host, port, mar_db_name)

	mau_cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		mau_db_role, mau_db_pass, host, port, mau_db_name)

	mdw_cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		mdw_db_role, mdw_db_pass, host, port, mdw_db_name)

	mlg_cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		mlg_db_role, mlg_db_pass, host, port, mlg_db_name)

	mst_cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		mst_db_role, mst_db_pass, host, port, mst_db_name)

	mi_db, errSql := sql.Open("postgres", mi_cs)
	if errSql != nil {
		slog.Error("Cannot connect to the maelstrom_install database: ", err)
	} else {
		MaelstromInstallDb = mi_db
		slog.Info("Successfully connected to the maelstrom_install database.")
	}

	m_db, errSql := sql.Open("postgres", m_cs)
	if errSql != nil {
		fmt.Println("Cannot connect to the maelstrom database: ", err)
		panic(err)
	} else {
		MaelstromDb = m_db
		fmt.Println("Successfully connected to the maelstrom database.")
	}

	mar_db, errSql := sql.Open("postgres", mar_cs)
	if errSql != nil {
		fmt.Println("Cannot connect to the maelstrom_archive database: ", err)
		panic(err)
	} else {
		MaelstromArchiveDb = mar_db
		fmt.Println("Successfully connected to the maelstrom_archive database.")
	}

	mau_db, errSql := sql.Open("postgres", mau_cs)
	if errSql != nil {
		fmt.Println("Cannot connect to the maelstrom_audit database: ", err)
		panic(err)
	} else {
		MaelstromAuditDb = mau_db
		fmt.Println("Successfully connected to the maelstrom_audit database.")
	}

	mdw_db, errSql := sql.Open("postgres", mdw_cs)
	if errSql != nil {
		fmt.Println("Cannot connect to the maelstrom_dw database: ", err)
		panic(err)
	} else {
		MaelstromDwDb = mdw_db
		fmt.Println("Successfully connected to the maelstrom_dw database.")
	}

	mlg_db, errSql := sql.Open("postgres", mlg_cs)
	if errSql != nil {
		fmt.Println("Cannot connect to the maelstrom_log database: ", err)
		panic(err)
	} else {
		MaelstromLogDb = mlg_db
		fmt.Println("Successfully connected to the maelstrom_log database.")
	}

	mst_db, errSql := sql.Open("postgres", mst_cs)
	if errSql != nil {
		fmt.Println("Cannot connect to the maelstrom_staging database: ", err)
		panic(err)
	} else {
		MaelstromStagingDb = mst_db
		fmt.Println("Successfully connected to the maelstrom_staging database.")
	}

	sqla := fmt.Sprintf("ALTER ROLE maelstrom WITH PASSWORD '%s';", m_db_pass)
	_, err = MaelstromInstallDb.Exec(sqla)
	if err != nil {
		slog.Error("Error updating password for role maelstrom." + err.Error())
	} else {
		slog.Info("Successfully updated password for role maelstrom.")
	}

	sqla = fmt.Sprintf("ALTER ROLE maelstrom_archive WITH PASSWORD '%s';", mar_db_pass)
	_, err = MaelstromInstallDb.Exec(sqla)
	if err != nil {
		slog.Error("Error updating password for role maelstrom_archive.")
	} else {
		slog.Info("Successfully updated password for role maelstrom_archive.")
	}

	sqla = fmt.Sprintf("ALTER ROLE maelstrom_audit WITH PASSWORD '%s';", mau_db_pass)
	_, err = MaelstromInstallDb.Exec(sqla)
	if err != nil {
		slog.Error("Error updating password for role maelstrom_audit.")
	} else {
		slog.Info("Successfully updated password for role maelstrom_audit.")
	}

	sqla = fmt.Sprintf("ALTER ROLE maelstrom_dw WITH PASSWORD '%s';", mdw_db_pass)
	_, err = MaelstromInstallDb.Exec(sqla)
	if err != nil {
		slog.Error("Error updating password for role maelstrom_dw.")
	} else {
		slog.Info("Successfully updated password for role maelstrom_dw.")
	}

	sqla = fmt.Sprintf("ALTER ROLE maelstrom_log WITH PASSWORD '%s';", mlg_db_pass)
	_, err = MaelstromInstallDb.Exec(sqla)
	if err != nil {
		slog.Error("Error updating password for role maelstrom_log.")
	} else {
		slog.Info("Successfully updated password for role maelstrom_log.")
	}

	sqla = fmt.Sprintf("ALTER ROLE maelstrom_staging WITH PASSWORD '%s';", mst_db_pass)
	_, err = MaelstromInstallDb.Exec(sqla)
	if err != nil {
		slog.Error("Error updating password for role maelstrom_staging.")
	} else {
		slog.Info("Successfully updated password for role maelstrom_staging.")
	}

}
