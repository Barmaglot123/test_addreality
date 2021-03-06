package utils

import (
    "path/filepath"
    "io/ioutil"
    "strconv"
    "regexp"
    "github.com/Sirupsen/logrus"
    "os"
)

var sqlScriptRegexp = regexp.MustCompile(`(\d+?)_.+`)

func SqlScript(version string) string {
    path, _ := filepath.Abs(findMigrationFileNameByVersion(version))
    logrus.Info("Using migration script ", path)
    dat, err := ioutil.ReadFile(path)
    if err != nil {
        logrus.Fatal(err)
    }
    return string(dat)
}

func SqlScriptMaxVersion() int {
    var dir []os.FileInfo
    var e bool = true
    var err error
    dir, err = ioutil.ReadDir("./resources/sql/")

    if err != nil {
        e = false
        err = nil
    }

    if !e {
        dir, err = ioutil.ReadDir("../resources/sql/")
    }

    if err != nil {
        logrus.Fatal(err)
    }

    max := 0

    for _, file := range dir {
        name := file.Name()
        sVersion := sqlScriptRegexp.FindStringSubmatch(name)[1]
        if version, err := strconv.Atoi(sVersion); err == nil && version > max {
            max = version
        }
    }

    return max
}

func findMigrationFileNameByVersion(version string) string {
    var dir []os.FileInfo
    var e bool = true
    var err error
    var parentFolder string

    var salScriptFindRegexp = regexp.MustCompile(`0*?` + version + `_.+?\.sql`)
    parentFolder = "./resources/sql/"

    dir, err = ioutil.ReadDir("./resources/sql/")

    if err != nil {
        e = false
        err = nil
        parentFolder = "../resources/sql/"
    }

    if !e {
        dir, err = ioutil.ReadDir(parentFolder)
    }

    if err != nil {
        logrus.Fatal(err)
    }

    for _, file := range dir {
        name := file.Name()
        if salScriptFindRegexp.MatchString(name) {
            return parentFolder + name
        }
    }

    return ""
}