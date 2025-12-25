// Package cmd contains the CLI commands for the application.
package cmd

import (
	"bufio"
	"log"
	"os"

	sort "github.com/lastlife77/sort/sortutil"
	"github.com/spf13/cobra"
)

var sortCmd = &cobra.Command{
	Use:   "Usage: sort [-nru] [-k START] [FILE]...",
	Short: "Usage: sort [-nru] [-k START] [FILE]...",
	Long: `Usage: sort [-nru] [-k START] [FILE]...

Sort lines of text

        -n      Sort numbers
		-M      Sort month
		-H      Sort human readable numbers (2K 1G)
        -k N[,M] Sort by Nth field
        -r      Reverse sort order
        -u      Suppress duplicate lines
		-b      Ignore leading blanks
		-c      Check whether input is sorted

		`,
	Run: func(cmd *cobra.Command, args []string) {
		var file *os.File

		switch len(args) {
		case 0:
			file = os.Stdin
		case 1:
			var err error
			file, err = os.Open(args[0])
			if err != nil {
				log.Fatal("The system cannot find the file specified.")
			}
			defer file.Close()
		default:
			log.Fatal("Input file specified two times.")
		}

		scanner := bufio.NewScanner(file)

		s := sort.New()

		for scanner.Scan() {
			line := scanner.Text()
			s.Append(line)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		// args
		isSelectSortType := false
		isSelectSortType = selectSortType("n", isSelectSortType, cmd, s.AsNums)
		isSelectSortType = selectSortType("M", isSelectSortType, cmd, s.AsMonths)
		isSelectSortType = selectSortType("H", isSelectSortType, cmd, s.AsHumanNums)

		addSortOption("r", cmd, s.Reverse)
		addSortOption("u", cmd, s.Unique)
		addSortOption("b", cmd, s.IgnoreLeadingBlanks)
		addSortOption("c", cmd, s.IsSorted)

		k, err := cmd.Flags().GetInt("k")
		if err != nil {
			log.Fatal(err)
		}

		s.Sort(k)
		s.Print()
	},
}

func selectSortType(flag string, isSelect bool, cmd *cobra.Command, sortType func()) bool {
	res, err := cmd.Flags().GetBool(flag)
	if err != nil {
		log.Fatal(err)
	}
	if res {
		if isSelect {
			log.Fatal("sort: unknown sort type")
		}
		sortType()

		return true
	}

	return isSelect
}

func addSortOption(flag string, cmd *cobra.Command, option func()) {
	res, err := cmd.Flags().GetBool(flag)
	if err != nil {
		log.Fatal(err)
	} else if res {
		option()
	}
}

// Execute runs the sort CLI command.
func Execute() {
	err := sortCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// type of sorts
	sortCmd.Flags().BoolP("n", "n", false, "Sort numbers")
	sortCmd.Flags().BoolP("M", "M", false, "Sort month")
	sortCmd.Flags().BoolP("H", "H", false, "Sort human readable numbers (2K 1G)")
	// sorts options
	sortCmd.Flags().IntP("k", "k", 1, "Sort by Nth field")
	sortCmd.Flags().BoolP("r", "r", false, "Reverse sort order")
	sortCmd.Flags().BoolP("u", "u", false, "Suppress duplicate lines")
	sortCmd.Flags().BoolP("b", "b", false, "Ignore leading blanks")
	sortCmd.Flags().BoolP("c", "c", false, "Check whether input is sorted")
}
