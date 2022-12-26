package cmd

import (
	"fmt"

	data "github.com/dbgeek/aoc_2022/internal/data/day24"
	"github.com/dbgeek/aoc_2022/internal/day24"
	"github.com/spf13/cobra"
)

var day24Cmd = &cobra.Command{
	Use:   "day24",
	Short: "day24",
	Long:  `day24`,
	Run: func(cmd *cobra.Command, args []string) {
    part1 := day24.Day24_1(data.Input)
    part2 := day24.Day24_2(data.Input)

    fmt.Println(part1, part2)
	},
}

func init() {
	rootCmd.AddCommand(day24Cmd)
}
