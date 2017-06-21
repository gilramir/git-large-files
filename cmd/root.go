package cmd

import (
	"fmt"
	"github.com/gilramir/argparse"
	"github.com/gilramir/git-large-files/largeFilesLib"
	"os"
	"strconv"
	"strings"
)

var argumentParser = &argparse.ArgumentParser{
	Name:             "git-large-files",
	ShortDescription: "Find large files throught all branches and history",
	Destination:      &Options{},
}

func Execute() {
	err := argumentParser.ParseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type Options struct {
	Size string
	J    int
}

func init() {
	argumentParser.AddArgument(&argparse.Argument{
		Short: "-j",
		Help:  "Number of simultaneous jobs to run while process. Defaults to # of CPUs.",
	})
	argumentParser.AddArgument(&argparse.Argument{
		Name: "size",
		Help: "Minimimum size of files to list, in bytes, or with a suffix " +
			"of KB, MB, or GB",
	})
}

func parseSize(sizeString string) (int, error) {
	if sizeString == "" {
		return 0, argparse.ParseError("The size was not given")
	}

	// Maybe the user gave us just an integer number of bytes
	size, err := strconv.Atoi(sizeString)
	if err == nil {
		return size, nil
	}

	// Surely they didn't give us a floating point number of bytes?
	_, err = strconv.ParseFloat(sizeString, 32)
	if err == nil {
		return 0, argparse.ParseError("The size cannot be a non-integer number of bytes")
	}

	if len(sizeString) < 3 {
		// We need two characters for the units, and at least one for the size.
		return 0, argparse.ParseError("The size should be a number, possibly with a suffix of KB, MB, or GB")
	}

	var multiplier int
	switch strings.ToLower(sizeString[len(sizeString)-2:]) {
	case "kb":
		multiplier = 1024
	case "mb":
		multiplier = 1024 * 1024
	case "gb":
		multiplier = 1024 * 1024 * 1024
	default:
		return 0, argparse.ParseErrorf("Unknown size unit '%s'. Use a suffix of KB, MB, or GB, or no suffix",
			sizeString[len(sizeString)-2:])
	}

	// Try an integer first
	numberString := sizeString[0 : len(sizeString)-2]
	size, err = strconv.Atoi(numberString)
	if err == nil {
		return size * multiplier, nil
	}

	// Try a floating point
	floatSize, err := strconv.ParseFloat(numberString, 32)
	if err != nil {
		return 0, argparse.ParseErrorf("The number before the units (%s) is not parseable: %s",
			sizeString[0:len(sizeString)-2], err.Error())
	}
	size = int(float32(floatSize) * float32(multiplier))
	return size, nil
}

func (self *Options) Run(d []argparse.Destination) error {

	sizeBytes, err := parseSize(self.Size)
	if err != nil {
		return err
	}

	if sizeBytes < 0 {
		return argparse.ParseError("The size cannot be negative.")
	}

	if self.J < 0 {
		return argparse.ParseError("The -j flag cannot be negative.")
	}

	err = largeFilesLib.GetLargeFiles(sizeBytes, options.J)
	if err != nil {
		return err
	}

	return nil
}
