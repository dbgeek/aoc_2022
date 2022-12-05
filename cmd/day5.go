package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day5"
	"github.com/dbgeek/aoc_2022/internal/day5"
	"github.com/spf13/cobra"
)

var day5Cmd = &cobra.Command{
	Use:   "day5",
	Short: "day5",
	Long:  `day5`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day5.Day5_1(data.Input)
    part2 := day5.Day5_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}
