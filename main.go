package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	ps "github.com/mitchellh/go-ps"
	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
)

func work() []map[string]string {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed, are you using windows?")
	}
	var out []map[string]string
	// temp := make(map[string]string)
	// var xd []int

	// map ages
	// var iter int
	for x := range processList {
		var process ps.Process
		temp := make(map[string]string)
		process = processList[x]
		fmt.Printf("%d\t%s\n", process.Pid(), fmt.Sprintf("%v", process.PPid()))
		// do os.* stuff on the pid

		// temp["pid"] = fmt.Sprintf("%v", process.Pid())
		temp["pid"] = strconv.Itoa(process.Pid())
		temp["ppid"] = strconv.Itoa(process.PPid())
		// fmt.Println(temp)
		// iter += 1
		out = append(out, temp)
		// xd = append(xd, process.Pid())
	}

	fmt.Println(out)
	// out = append(out, temp)
	// fmt.Println(xd)
	return out
}

func main() {
	socket := flag.String("socket", "", "Path to osquery socket file")
	flag.Parse()
	if *socket == "" {
		log.Fatalf(`Usage: %s --socket SOCKET_PATH`, os.Args[0])
	}

	server, err := osquery.NewExtensionManagerServer("process_parent", *socket)
	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	// Create and register a new table plugin with the server.
	// table.NewPlugin requires the table plugin name,
	// a slice of Columns and a Generate function.
	server.RegisterPlugin(table.NewPlugin("process_parent", FoobarColumns(), FoobarGenerate))
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}

// FoobarColumns returns the columns that our table will return.
func FoobarColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("pid"),
		table.TextColumn("ppid"),
	}
}

// FoobarGenerate will be called whenever the table is queried. It should return
// a full table scan.
func FoobarGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	return work(), nil
}
