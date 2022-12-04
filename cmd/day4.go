package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day4"
	"github.com/dbgeek/aoc_2022/internal/day4"
	"github.com/spf13/cobra"
)

var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "day4",
	Long:  `day4`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day4.Day4_1(data.Input)
    part2 := day4.Day4_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
