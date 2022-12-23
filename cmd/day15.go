package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day15"
	"github.com/dbgeek/aoc_2022/internal/day15"
	"github.com/spf13/cobra"
)

var day15Cmd = &cobra.Command{
	Use:   "day15",
	Short: "day15",
	Long:  `day15`,
	Run: func(cmd *cobra.Command, args []string) {
		part1 := day15.Day15_1(data.Input, 2000000)
		part2 := day15.Day15_2(data.Input)

		fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day15Cmd)
}
