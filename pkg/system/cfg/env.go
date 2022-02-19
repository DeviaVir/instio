package cfg

import (
	"log"
	"os"
	"path/filepath"
)

func getWd() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Couldn't find working directory", err)
	}
	return wd
}

func DataDir() string {
	dataDir := os.Getenv("INSTIO_DATA_DIR")
	if dataDir == "" {
		return getWd()
	}
	return dataDir
}

func TlsDir() string {
	tlsDir := os.Getenv("INSTIO_TLS_DIR")
	if tlsDir == "" {
		tlsDir = filepath.Join(getWd(), "system", "tls")
	}
	return tlsDir
}

func AdminStaticDir() string {
	staticDir := os.Getenv("INSTIO_ADMINSTATIC_DIR")
	if staticDir == "" {
		staticDir = filepath.Join(getWd(), "system", "admin", "static")
	}
	return staticDir
}

func UploadDir() string {
	uploadDir := os.Getenv("INSTIO_UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = filepath.Join(DataDir(),"uploads")
	}
	return uploadDir
}

func SearchDir() string {
	searchDir := os.Getenv("INSTIO_SEARCH_DIR")
	if searchDir == "" {
		searchDir = filepath.Join(DataDir(),"search")
	}
	return searchDir
}