package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day2"
	"github.com/dbgeek/aoc_2022/internal/day2"
	"github.com/spf13/cobra"
)

var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day2",
	Long:  `day2`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day2.Day2_1(data.Input)
    part2 := day2.Day2_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day2Cmd)
}
