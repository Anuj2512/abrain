/*
* Copyright (C) 2017 Abrain, Ankur Yadav <ankurayadav@gmail.com>
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU Affero General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU Affero General Public License for more details.
*
* You should have received a copy of the GNU Affero General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
)

// BrainMeta should be a generic structure for ArBrain.
// Every resource should be usable and accesible via brainQL only.
// should be in such a way that its generic enough with graph schema
// architecture.
// It should also capture out concepts action associated with the data
type BrainMeta struct {
	nodes []NodeMeta
}

// NodeMeta structure for generic node type.
type NodeMeta struct {
	Node       string `json:"node"`
	Properties []PropsMeta
	EdgesOut   []EdgesMeta
}

// PropsMeta structure for generic properties type for each node.
type PropsMeta struct {
	Property   string `json:"property"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Constraint string `json:"constraint"`
}

// EdgesMeta structure for generic edgesOut type for each node.
type EdgesMeta struct {
	Edge       string `json:"edge"`
	To         string `json:"to"`
	Properties []PropsMeta
}

func check() {
	fmt.Printf("This is and internal function")
}
