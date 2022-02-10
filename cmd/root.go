/*
Package cmd

Copyright © 2022 Tim Hinz <timhinz16@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"
	"go-photo-exporter/internal/db"
	"go-photo-exporter/internal/fs"
	"os"
	"path/filepath"
)

var pathFlagString string
var outputFlagString string
var nameFlagString string

var rootCmd = &cobra.Command{
	Use:   "ireaver",
	Short: "Export iCloud photos to a specified folder on your computer",
	Long:  `Export iCloud photos to a specified folder on your computer`,
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&pathFlagString, "path", "p", "$HOME/Pictures/Photos Library.photoslibrary/", "Path to your local iCloud photos library")
	rootCmd.Flags().StringVarP(&outputFlagString, "output", "o", "", "Directory path to the folder you want to output images to")
	rootCmd.Flags().StringVarP(&nameFlagString, "name", "n", "original", "Renames output images (options: \"sequential\", \"original\", \"guid\")")
	rootCmd.MarkFlagRequired("output")
}

func processFiles(assets []db.Asset) {
	bar := pb.StartNew(len(assets))

	fmt.Println("Copying files...")
	for _, s := range assets {
		var name string
		if nameFlagString == "sequential" {
			name = fmt.Sprintf("%d%s", s.Id, filepath.Ext(s.Filename))
		} else if nameFlagString == "original" {
			name = s.OriginalFileName
		} else if nameFlagString == "guid" {
			name = s.Filename
		} else {
			fmt.Println("Invalid name option \"" + nameFlagString + "\", valid options: \"sequential\", \"original\", \"guid\"")
			os.Exit(1)
		}

		fmt.Println(name)
		var src = fmt.Sprintf("%soriginals/%s/%s", pathFlagString, s.Directory, s.Filename)
		var dest = fmt.Sprintf("%s%s", outputFlagString, name)
		fs.Copy(src, dest)

		bar.Increment()
	}

	bar.Finish()
}

func main() {
	pathFlagString = fs.CleanPath(pathFlagString)
	outputFlagString = fs.CleanPath(outputFlagString)

	connection := db.Connect(pathFlagString)
	assets := db.FetchImageRecords(connection)

	processFiles(assets)
}
