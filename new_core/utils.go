package new_core

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func IsDirectoryExists(directoryPath string) bool {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDirectory(pathToDir string) {
	fmt.Println(pathToDir)
	// os.Mkdir(fmt.Sprintf("./%s", pathToDir))
}

func getFolderEntries(folderPath string) []os.DirEntry {
	// Creating empty var for return
	folderFilesObjects := make([]os.DirEntry, 0)
	// Reading directory entries
	folderFiles, readingFolderError := os.ReadDir(folderPath)
	if readingFolderError != nil {
		log.Printf("Error reading folder: %e\n", readingFolderError)
		return folderFilesObjects
	}
	// Adding files to folderFilesObjects
	for _, file := range folderFiles {
		folderFilesObjects = append(folderFilesObjects, file)
	}
	return folderFilesObjects
}

func UpdateClusters() {
	databaseEntries := getFolderEntries(DatabaseCore.DatabasePath)
	for _, entry := range databaseEntries {
		if entry.IsDir() {
			// Creating new cluster
			newCluster := Cluster{ClusterName: entry.Name(),
				ClusterPath:  strings.Join([]string{DatabaseCore.DatabasePath, entry.Name()}, ""),
				ClusterFiles: []Document{}}
			// Reading cluster files
			clusterFiles := getFolderEntries(newCluster.ClusterPath)
			for _, file := range clusterFiles {
				if !file.IsDir() {
					// Reading file data
					fileData, documentReadingError := os.ReadFile(strings.Join([]string{newCluster.ClusterPath,
						file.Name()}, ""))
					if documentReadingError != nil {
						log.Printf("Error reading document data: %e\n", documentReadingError)
						continue
					}
					// Creating document object
					newDocument := Document{DocumentName: file.Name(),
						DocumentValue: string(fileData)}
					// Adding document to cluster
					newCluster.ClusterFiles = append(newCluster.ClusterFiles, newDocument)
				}
			}
		}
	}
}
