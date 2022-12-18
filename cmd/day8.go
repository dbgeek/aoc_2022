package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day8"
	"github.com/dbgeek/aoc_2022/internal/day8"
	"github.com/spf13/cobra"
)

var day8Cmd = &cobra.Command{
	Use:   "day8",
	Short: "day8",
	Long:  `day8`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day8.Day8_1(data.Input)
    part2 := day8.Day8_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day8Cmd)
}
