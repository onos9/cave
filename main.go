package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/cave/cmd/api"
	cfg "github.com/cave/config"
)

func main() {
	service := flag.String("s", "monolith", "service api, chat or stream")
	flag.Parse()

	cfg.LoadConfig()
	db := cfg.NewDBConnection()
	apiServer := api.Init(db)

	switch *service {
	case "monolith":
		apiServer.Run()
	case "api":
		apiServer.Run()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	apiServer.Shutdown()
}

// function exportTableToCSV(filename) {
//     var csv = [];
//     var tables = document.querySelectorAll("table");

// 	for (var i = 0; i < tables.length; i++) {
// 		var table = [], rows = tables[i].querySelectorAll("tr");

// 		for (var i = 0; i < rows.length; i++) {
// 			var row = [], cols = rows[i].querySelectorAll("td, th");

// 			for (var j = 0; j < cols.length; j++)
// 				row.push(cols[j].innerText);

// 			csv.push(row.join(","));
// 		}
// 		csv.push(table.join(","));
//     }

//     csv.join("\n")

// 	var csvFile;
//     var downloadLink;
//     csvFile = new Blob([csv], {type: "text/csv"});

//     downloadLink = document.createElement("a");
//     downloadLink.download = "filename.csv";
//     downloadLink.href = window.URL.createObjectURL(csvFile);
//     downloadLink.style.display = "none";

//     document.body.appendChild(downloadLink);
//     downloadLink.click();

// }
