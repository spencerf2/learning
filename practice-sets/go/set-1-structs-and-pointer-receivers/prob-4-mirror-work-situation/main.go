package main

import "fmt"

type ReportGenerator struct {
	outputPath string
}

func NewReportGenerator(outputPath string) *ReportGenerator {
	return &ReportGenerator{
		outputPath: outputPath,
	}
}

func (r *ReportGenerator) Generate(reportName string) error {
	fmt.Printf("Generating %s report to %s\n", reportName, r.outputPath)
	return nil
}

// Try both approaches.

// Approach 1: add reportGen *ReportGenerator to Service struct and NewService constructor

type Service struct {
	name      string
	reportGen *ReportGenerator
}

func NewService(name string, reportGen *ReportGenerator) *Service {
	return &Service{
		name: name,
		reportGen: reportGen,
	}
}

// End Approach 1.

// Approach 2: create ReportGenerator locally in the method below

func (s *Service) ProcessData() error {
	// Try using s.reportGen:
	s.reportGen.Generate("data-analysis")

	// Try creating a local ReportGenerator
	// r := NewReportGenerator("over_there")
	// r.Generate("data-analysis")
	return nil
}

func main() {
	// Approach 1:
	// /*
		r := NewReportGenerator("right_here")
		s := NewService("spencer-service", r)
		s.ProcessData()
	// */

	// Approach 2:
	/*
	r := NewReportGenerator("right_here")
	s := NewService("spencer-service", r)
	s.ProcessData()
	*/
}

// Note to self: I'm finding that when I add reportGen *ReportGenerator to the
//               Service struct and NewService constructor, then I have to
//               instantiate a ReportGenerator to have a Service to work with.
//               That means that I end up repeating the work that could take
//               place locally within ProcessData (if creating a local
//               ReportGenerator there). Also, the local ReportGenerator is
//               an entirely different instance from the one used to create
//               the service, so that's some wasted / non-sensical code.
//
//               If instead, I use s.reportGen within ProcessData(), then I still
//               have to instantiate a ReportGenerator to pass into the Service
//               in main(), but then we actually use it to generate the report,
//               rather than creating an additional reporter in ProcessData()
//               and leaving the original unused.
//
//               I'm thinking my next move is to see what happens when I omit
//               ReportGenerator from the Service struct and only create it
//               locally in methods.
//               I'm also thinking that this practice problem wasn't fully polished
//               when I accepted it from Claude. That's an opportunity for
//               me to explore more and see where I lead me!

// Answering the questions:
