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

package knowledge

import (
	"log"

	"github.com/dgraph-io/dgraph/client"
	"github.com/dgraph-io/dgraph/protos/graphp"
	"google.golang.org/grpc"
)

var (
	dgraph = "127.0.0.1:8080" // Address to communicate with dgraph server.
)

// Sync pushes data to graph store by transforming
// data into desirable format of graph store.
func Sync() {
	log.Println("This will sync data to graphdb!")

	// Just starting a basic syncer for songs
	// Later I will make it generic.

	// Creating connecction with dgraph.
	conn, err := grpc.Dial(dgraph, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	// Creating a new client.
	graphp.NewDgraphClient(conn)
	// Starting a new request.
	req := client.Req{}
	// _:person1 tells Dgraph to assign a new Uid and is the preferred way of creating new nodes.
	// See https://wiki.dgraph.io/index.php?title=Query_Language&redirect=no#Assigning_UID for more details.
	nq := graphp.NQuad{
		Subject:   "_:person1",
		Predicate: "name",
	}
	// Str is a helper function to add a string value.
	client.Str("Steven Spielberg", &nq)
	// Adding a new mutation.
	req.AddMutation(nq, client.SET)
}

// InitSchema initializes schema based on meta onto graphdb.
func InitSchema() {
	log.Println("Initializing Schema on graphdb!")

	/* TODO:
	1. Add code for Initializing schema on graphdb based on metadata.
	2. Add code for making a meta graph on graphdb.
	3. Make a generelized code for step 1 which will be applicable for all
		metas.
	*/

}
