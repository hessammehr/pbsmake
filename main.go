package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Job struct {
	Name         string
	FileName     string
	Nodes        int
	Processors   int
	Hours        int
	TemplateFile string
}

var cwd, _ = filepath.Abs(".")

func findJobs(dir string, ext string) []string {
	results := []string{}
	entries, err := ioutil.ReadDir(".")
	if err != nil {
		log.Panic("Cannot read contents of current directory!")
	}
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ext {
			results = append(results, filepath.Base(entry.Name()))
		}
	}
	return results
}

func process(job Job) {
	fmt.Printf("Processing input file %s...\n", job.FileName)
	pbsTemplate, err := ioutil.ReadFile(job.TemplateFile)
	if err != nil {
		log.Panicf("Cannot open template file %s", job.TemplateFile)
	}

	templ, err := template.New("PBS").Parse(string(pbsTemplate))
	if err != nil {
		log.Panicf("Cannot not parse template file %s", job.TemplateFile)
	}
	if dirErr := os.Mkdir(job.Name, os.ModeDir|0755); dirErr != nil {
		log.Panicf("Cannot create output directory %s, does it already exist?", job.Name)
	}
	pbsFileName := filepath.Join("./"+job.Name, filepath.Base(cwd)+"-"+job.Name+".pbs")
	pbsFile, err := os.Create(pbsFileName)
	defer pbsFile.Close()
	if err != nil {
		log.Panicf("Cannot create PBS file %s", pbsFileName)
	}
	if templ.Execute(pbsFile, job) != nil {
		log.Panicf("Cannot write to submission file %s", pbsFileName)
	}
	if os.Rename(job.FileName, filepath.Join(job.Name, job.FileName)) != nil {
		log.Panicf("Failed to move job file %s to folder %s", job.Name, job.FileName)
	}
}

func main() {
	templateFile := flag.String("template", "template.pbs", "Template file, default: template.pbs")
	hours := flag.Int("h", 5, "Requested walltime in hours")
	nodes := flag.Int("n", 1, "Number of nodes to require")
	processors := flag.Int("p", 4, "Number of processors to require")
	extension := flag.String("extension", "com", "File extension for input files")
	flag.Parse()

	jobFiles := flag.Args()
	if flag.NArg() == 0 {
		jobFiles = findJobs(".", "."+*extension)
	}

	for _, jobFile := range jobFiles {
		jobName := strings.Split(jobFile, ".")[0]
		job := Job{jobName, jobFile, *nodes, *processors, *hours, *templateFile}
		process(job)
	}

}
