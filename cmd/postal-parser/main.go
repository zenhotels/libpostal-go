package main

import (
	"log"
	"os"
	"strings"

	"github.com/apcera/termtables"
	"github.com/xlab/closer"
	"github.com/zenhotels/libpostal-go/postal"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	defer closer.Close()

	// check the input
	if len(os.Args) < 2 {
		log.Fatalln("wrong input: expected an address")
	}
	addr := strings.Join(os.Args[1:], " ")

	// init the libpostal
	if !postal.Setup() || !postal.SetupParser() {
		log.Fatalln("error at libpostal init")
	}
	// bind the teardown routine to be called upon exit
	closer.Bind(postal.TeardownParser)

	// call the libpostal
	opt := postal.GetLibpostalAddressParserDefaultOptions()
	resp := postal.ParseAddress([]byte(addr), opt)
	defer postal.AddressParserResponseDestroy(resp)
	resp.Deref()

	// read all the address components from the response
	log.Println("address components:", resp.NumComponents)
	resp.Components = make([][]byte, resp.NumComponents)
	resp.Labels = make([][]byte, resp.NumComponents)
	resp.Deref()

	// print a table
	table := termtables.CreateTable()
	table.UTF8Box()
	table.AddTitle("libpostal parse")
	for i, label := range resp.Labels {
		table.AddRow(toString(label, 255), toString(resp.Components[i], 1024))
	}
	log.Println("\n" + table.Render())
}
