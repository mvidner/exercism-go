package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name   string
	played uint
	won    uint
	drawn  uint
	lost   uint
	points uint
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams := map[string]Team{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		fields := strings.Split(line, ";")
		if len(fields) != 3 {
			return fmt.Errorf("Expected 3 semicolon-separated fields in: %q", line)
		}

		team1 := fields[0]
		team2 := fields[1]
		result_s := fields[2]
		var result int
		switch result_s {
		case "win":
			result = 1
		case "loss":
			result = -1
		case "draw":
			result = 0
		default:
			return fmt.Errorf("Unexpected match result: %q", result)
		}
		record_result(teams, team1, result)
		record_result(teams, team2, -result)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	sorted_keys := make([]string, 0, len(teams))
	for k := range teams {
		sorted_keys = append(sorted_keys, k)
	}
	cmp := func(i, j int) bool {
		ki, kj := sorted_keys[i], sorted_keys[j]
		pi, pj := teams[ki].points, teams[kj].points
		if pi != pj {
			return pi > pj
		}
		return ki < kj
	}
	sort.Slice(sorted_keys, cmp)

	fmt.Fprintf(writer, "%-30v | MP |  W |  D |  L |  P\n", "Team")
	for _, k := range sorted_keys {
		t := teams[k]
		fmt.Fprintf(writer,
			"%-30v | %2v | %2v | %2v | %2v | %2v\n",
			t.name, t.played, t.won, t.drawn, t.lost, t.points)
	}

	return nil
}

// update(or create) teams[team_name] according to the match result (-1/0/+1)
func record_result(teams map[string]Team, team_name string, result int) {
	t := teams[team_name]
	t.name = team_name
	t.played += 1
	switch result {
	case -1:
		t.lost += 1
	case 0:
		t.drawn += 1
		t.points += 1
	case 1:
		t.won += 1
		t.points += 3
	default:
		panic("want -1, 0, or +1")
	}
	teams[team_name] = t
}
