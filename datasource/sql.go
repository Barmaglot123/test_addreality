package datasource

import (
    log "github.com/Sirupsen/logrus"
    "github.com/jinzhu/gorm"
    "untitled3/config"
    . "untitled3/utils"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "strconv"
)

var Sql *gorm.DB

func SetupSql() {
    connect()
    migrate()
}

func SetupTestSql() error {
     return connectToTest()
}

const (
    NO_VERSION 	= 0
)

func connect() {
    var err error
    Sql, err = gorm.Open("postgres",
        " host=" + config.Sql.Host() +
            " dbname="   + config.Sql.Name() +
            " user="     + config.Sql.User() +
            " password=" + config.Sql.Password() +
            " sslmode="  + config.Sql.SSL())

    if err != nil {
        log.Fatal(err)
    }
}

func connectToTest() error {
    var err error
    Sql, err = gorm.Open("postgres",
        " host=" + config.Sql.Host() +
            " dbname="   + config.Sql.Name() +
            " user="     + config.Sql.User() +
            " port=9668" +
            " password=" + config.Sql.Password() +
            " sslmode="  + config.Sql.SSL())

    if err != nil {
        return err
    }
    return nil

}

func migrate() {
    currentVersion := currentSchemaVersion()
    maxVersionFromScripts := SqlScriptMaxVersion()

    if maxVersionFromScripts > currentVersion {
        for i := currentVersion + 1; i <= maxVersionFromScripts; i++ {
            versionString := strconv.Itoa(i)
            log.Info("Migrating to version " + versionString)
            runMigrationScriptToVersion(versionString)
        }
    }
}

func currentSchemaVersion() int {
    var version string = "0"
    query := "select exists (select table_name from information_schema.tables where table_schema = 'public' and table_name = 'system')"
    e := false
    Sql.Raw(query).Row().Scan(&e)
    if !e { return NO_VERSION }
    Sql.Raw("select value from system where name = ?", "version").Row().Scan(&version)
    if v, err := strconv.Atoi(version); err != nil {
        return NO_VERSION
    } else {
        return v
    }
}

func runMigrationScriptToVersion(version string) {
    if err := Sql.Exec(SqlScript(version)).Error; err != nil {
        log.Fatal(err)
    }

    if err := Sql.
    Exec(`
			insert into system (name, value) values ('version', ?)
			on conflict (name) do update
			set value = excluded.value;`, version).
        Error; err != nil {
        log.Fatal(err)
    }
}