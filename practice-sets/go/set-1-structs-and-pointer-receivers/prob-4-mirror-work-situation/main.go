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
/*
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
*/
// End Approach 1.

// Approach 2: create ReportGenerator locally in the method below

type Service struct {
	name      string
}

func NewService(name string) *Service {
	return &Service{name: name}
}

func (s *Service) ProcessData() error {
	// Try using s.reportGen:
	// s.reportGen.Generate("data-analysis")

	// Try creating a local ReportGenerator
	r := NewReportGenerator("over_there")
	r.Generate("data-analysis")
	return nil
}

func main() {
	// Approach 1:
	/*
		r := NewReportGenerator("right_here")
		s := NewService("spencer-service", r)
		s.ProcessData()
	*/

	// Approach 2:
	// /*
	s := NewService("spencer-service")
	s.ProcessData()
	// */
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
//
//               Update: 2025-07-06
//               I'm now thinking Approach 1 was meant to be the version where I added
//               ReportGenerator to teh Service struct and NewService constructor and
//               then used s.reportGen within ProcessData()
//
//               And that Approach 2 is where Service struct only has the name field,
//               and ProcessData creates a local ReportGenerator.
//               When I do things this way, less work is required up front to create
//               the ReportGenerator the Service uses, but the disadvantage is that the
//               Service is then less self contained.

// Answering the questions:
// 1. Which approach feels more like your existing codebase patterns?
//   - Definitely Approach 1. We follow dependency injection to ensure things are
//     easy to test and less tightly coupled.
//
// 2. What are the pros/cons of each approach?
//   - Approach 1 (dependency injection):
//     - Pros:
//       - Things remain easier to test and less tightly coupled.
//       - If we wanted to, we could change ReportGenerator's instantiation logic
//         without affecting any of the Service methods which use it.
//       - Each method of Service will know about ReportGenerator without having
//         to create their own.
//     - Cons:
//       - Have to instantiate a ReportGenerator to use the Service. Service doesn't
//         create a ReportGenerator on its own.
//
//   - Approach 2:
//     - Pros:
//       - Don't have to instantiate a ReportGenerator to use the Service.
//       - If Service only has 1 method that needs to use ReportGenerator and
//         we don't need to do anything with it other than generate the report
//         one time, then Approach 2 would be better.
//     - Cons:
//       - Would need to instantiate ReportGenerator multiple times if used in
//         multiple Service methods--or pass it around between methods.
//       - If we were passing ReportGenerator around, then we'd need to be
//         careful to ensure that if we had concurrency going on, or multiple
//         different functions calling ReportGenerator with different params,
//         it's possible we'd get race conditions or otherwise need to be careful
//         to ensure the wrong instance isn't used in an operation. The dependency
//         injection pattern would ensure that each application instance only creates
//         one ReportGenerator, which provides greater safety.
//
// 3. When might you choose one over the other?
//   - Choose Approach 1 whenever you need to unit test components, or when you'll be
//     using ReportGenerator in a variety of Service methods.
//   - Choose Approach 2 if only 1 of Service struct's methods need a temporary instance
//     of ReportGenerator.
