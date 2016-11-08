// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go build
// ./example3

// Sample program to output the result of the test, based on a critical value.
package main

import (
	"fmt"
	"math"
)

// This example program calculates the chi-squared test statistic relevant to
// the following example problem from here:
//
// http://sphweb.bumc.bu.edu/otlt/MPH-Modules/BS/BS704_HypothesisTesting-ChiSquare/BS704_HypothesisTesting-ChiSquare_print.html
//
// A University conducted a survey of its recent graduates to collect demographic and
// health information for future planning purposes as well as to assess students' satisfaction
// with their undergraduate experiences. The survey revealed that a substantial proportion of
// students were not engaging in regular exercise, many felt their nutrition was poor and a
// substantial number were smoking. In response to a question on regular exercise, 60% of all
// graduates reported getting no regular exercise, 25% reported exercising sporadically and 15%
// reported exercising regularly as undergraduates. The next year the University launched a
// health promotion campaign on campus in an attempt to increase health behaviors among
// undergraduates. The program included modules on exercise, nutrition and smoking cessation.
// To evaluate the impact of the program, the University again surveyed graduates and asked
// the same questions. The survey was completed by 470 graduates and the following data were
// collected on the exercise question:
//
// No Regular Exercise: 255
// Sporadic Exercise: 125
// Regular Exercise: 90
// Total: 470
//
// Based on the data, is there evidence of a shift in the distribution of responses to the
// exercise question following the implementation of the health promotion campaign on campus?
// Run the test at a 5% level of significance.

var (

	// The slice includes the observed frequencies.
	observed = []float64{
		255.0, // This number is the number of observed with no regular exercise.
		125.0, // This number is the number of observed with sporatic exercise.
		90.0,  // This number is the number of observed with regular exercise.
	}

	// This value is the total number of observations.
	totalObserved = 470.0

	// This value is the critical value of chi-squared that we will use to determine the
	// results of our test.   Critical values can be found in a table of probabilities for
	// the chi squared distribution. We have df=k-1=3-1=2 and a 5% level of
	// significance. The appropriate critical value is 5.99, and the decision rule is as
	// follows: Reject H0 if chi squared > 5.99.
	criticalChi = 5.99
)

func main() {

	// Calculate the expected frequencies.
	expected := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	// Calculate the chi-squared test statistic.
	var chiSquared float64
	for idx, val := range observed {
		chiSquared += math.Pow(val-expected[idx], 2.0) / expected[idx]
	}

	// If chi squared is greater than the critical value, then we will reject the
	// null hypothesis.
	var result bool
	if chiSquared > criticalChi {
		result = true
	}

	// Output the test results to standard out.
	fmt.Printf("\nOur hypothesis is: %t\n\n", result)
}
