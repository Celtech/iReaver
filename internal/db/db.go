package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"log"
)

type Asset struct {
	Id               int
	Directory        string
	Filename         string
	OriginalFileName string
}

func Connect(path string) (db *sql.DB) {
	db, err := sql.Open("sqlite3", path+"/database/Photos.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func FetchImageRecords(db *sql.DB) []Asset {
	row, err := db.Query("SELECT ZASSET.Z_PK, ZDIRECTORY, ZFILENAME, ZORIGINALFILENAME FROM ZASSET LEFT JOIN ZCLOUDMASTER ON ZASSET.ZMASTER = ZCLOUDMASTER.Z_PK")
	var assets []Asset

	if err != nil {
		log.Fatal(err)
	}

	for row.Next() { // Iterate and fetch the records from result cursor
		var Z_PK int
		var ZDIRECTORY string
		var ZFILENAME string
		var ZORIGINALFILENAME string
		row.Scan(&Z_PK, &ZDIRECTORY, &ZFILENAME, &ZORIGINALFILENAME)

		assets = append(assets, []Asset{
			{
				Id:               Z_PK,
				Directory:        ZDIRECTORY,
				Filename:         ZFILENAME,
				OriginalFileName: ZORIGINALFILENAME,
			},
		}...)
	}

	defer row.Close()

	return assets
}
