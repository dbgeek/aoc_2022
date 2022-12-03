package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day3"
	"github.com/dbgeek/aoc_2022/internal/day3"
	"github.com/spf13/cobra"
)

var day3Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day3",
	Long:  `day3`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day3.Day3_1(data.Input)
    part2 := day3.Day3_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)
}
