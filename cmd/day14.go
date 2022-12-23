package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day14"
	"github.com/dbgeek/aoc_2022/internal/day14"
	"github.com/spf13/cobra"
)

var day14Cmd = &cobra.Command{
	Use:   "day14",
	Short: "day14",
	Long:  `day14`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day14.Day14_1(data.Input)
    part2 := day14.Day14_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day14Cmd)
}
