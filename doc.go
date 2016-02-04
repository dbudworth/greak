// Goroutine Leak Detector.
//
// A simple library that allows you to compare the current list of goroutines against a previously recorded set.
//
// Example usecase:
//
//  - Setup main application loop, create a snapshot with New()
//  - Spawn worker routines
//  - Wait for work to complete
//  - Call Check on previously recorded snapshot to see if there are any lingering new goroutines.
package greak
