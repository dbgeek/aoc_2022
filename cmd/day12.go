package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day12"
	"github.com/dbgeek/aoc_2022/internal/day12"
	"github.com/spf13/cobra"
)

var day12Cmd = &cobra.Command{
	Use:   "day12",
	Short: "day12",
	Long:  `day12`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day12.Day12_1(data.Input)
    part2 := day12.Day12_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day12Cmd)
}
