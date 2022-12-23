package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day10"
	"github.com/dbgeek/aoc_2022/internal/day10"
	"github.com/spf13/cobra"
)

var day10Cmd = &cobra.Command{
	Use:   "day10",
	Short: "day10",
	Long:  `day10`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day10.Day10_part1(data.Input)
    part2 := day10.Day10_part2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}
