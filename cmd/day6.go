package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day6"
	"github.com/dbgeek/aoc_2022/internal/day6"
	"github.com/spf13/cobra"
)

var day6Cmd = &cobra.Command{
	Use:   "day6",
	Short: "day6",
	Long:  `day6`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day6.Day6_1(data.Input)
    part2 := day6.Day6_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day6Cmd)
}
