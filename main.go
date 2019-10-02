package main

import (
	"archive/zip"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"strings"
	"time"
	"path/filepath"
)

func main() {

	// Get the directory of the currently running file
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	// Set path to .env file and get values
	err := godotenv.Load(dir+"/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get .env variables
	logDirectory := os.Getenv("ROOT_COMPRESSED_LOG_DIRECTORY")
	strLogFiles := os.Getenv("LOG_FILES")

	// Get prefix file name with date for zip file
	t := time.Now()
	dateFilePrefix := fmt.Sprintf("%d_%02d_%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	if _, err := os.Stat(logDirectory); os.IsNotExist(err) {
		os.Mkdir(logDirectory, 0666)
	}

	items := strings.Split(strLogFiles, "; ")
	for _, e := range items {
		parts := strings.Split(e, " = ")
		logFilePath := parts[0]
		logFileName := parts[1]
		newFileName := dateFilePrefix+"_"+logFileName

		// Rename original file for compressing
		if err := os.Rename(logFilePath+"/"+logFileName, logDirectory+"/"+newFileName+".log"); err != nil {
			log.Print(err.Error())
		}

		// Create new file
		f, e := os.Create(logFilePath+"/"+logFileName)
		if e != nil {
			log.Print(err.Error())
		}
		defer f.Close()

		if err := os.Chmod(logFilePath+"/"+logFileName, 0666); err != nil {
			log.Println(err)
		}

		// Compress file to zip
		newZipFile, err := os.Create(logDirectory+"/"+newFileName+".zip")
		if err != nil {
			log.Print(err.Error())
		}
		defer newZipFile.Close()

		zipWriter := zip.NewWriter(newZipFile)
		defer zipWriter.Close()

		if err = AddFileToZip(zipWriter, logDirectory, newFileName); err != nil {
			log.Print(err.Error())
		}

		// Remove log file after compressing
		if err := os.Remove(logDirectory+"/"+newFileName+".log"); err != nil {
			log.Print(err.Error())
		}
	}

}

func AddFileToZip(zipWriter *zip.Writer, directory string, filename string) error {

	fileToZip, err := os.Open(directory+"/"+filename+".log")
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
