package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day9"
	"github.com/dbgeek/aoc_2022/internal/day9"
	"github.com/spf13/cobra"
)

var day9Cmd = &cobra.Command{
	Use:   "day9",
	Short: "day9",
	Long:  `day9`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day9.Day9_1(data.Input)
    part2 := day9.Day9_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)
}
