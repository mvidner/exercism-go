package tree

import (
	"fmt"
	"sort"
)

// input to Build
type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	nrecords := len(records)
	if nrecords == 0 {
		return nil, nil
	}

	// Hold the nodes in a slice,
	nodes := make([]Node, nrecords)
	// Helper to detect duplicate Records
	seen := make([]bool, nrecords)

	// Sort the input by ID so that the resulting Children slices are
	// also sorted
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	// Now give them IDs and child references
	for _, r := range records {
		if err := CheckParent(&r); err != nil {
			return nil, err
		}
		id := r.ID
		if id >= nrecords {
			return nil, fmt.Errorf("id %d is too high", id)
		}
		if seen[id] {
			return nil, fmt.Errorf("duplicate record with id %d", id)
		}
		seen[id] = true
		nodes[id].ID = id

		if id != 0 {
			parent := r.Parent
			nodes[parent].Children = append(nodes[parent].Children, &nodes[id])
		}
	}

	return &nodes[0], nil
}

func CheckParent(r *Record) error {
	if r.ID == 0 {
		if r.Parent == 0 {
			return nil
		}
		return fmt.Errorf("root node has a parent: %d", r.Parent)
	}
	if r.Parent < r.ID {
		return nil
	}
	return fmt.Errorf("node %d has a parent with too high ID: %d", r.ID, r.Parent)
}
