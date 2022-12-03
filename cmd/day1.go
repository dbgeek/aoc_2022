package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day1"
	"github.com/dbgeek/aoc_2022/internal/day1"
	"github.com/spf13/cobra"
)

var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1",
	Long:  `day1`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day1.Day1_1(data.Part)
    part2 := day1.Day1_2(data.Part)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}
