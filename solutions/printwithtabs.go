package solutions

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func printWithTabs(slice [][]int) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	for _, l := range slice {
		for _, r := range l {
			fmt.Fprintf(w, "\t%d", r)
		}
		fmt.Fprint(w, "\n")
	}
	w.Flush()
}
