package greak

import (
	"bytes"
	"fmt"
	"regexp"
	"runtime"
	"sort"
	"strconv"
)

var idPattern = regexp.MustCompile(`^goroutine (\d+)`)

// Size of buffer to use to dump stack, very large apps may need to increase this
var MaxStackTextSize = 500 * 1024

type Entry struct {
	Id          int
	Description []byte
}

func (e Entry) String() string { return string(e.Description) }

// Creates a new snapshot that represents the current set of goroutines
func New() Entries {
	// Check against empty entries results in current goroutine list
	return Entries{}.Check()
}

// A set of goroutines.
type Entries []Entry

func (es Entries) contains(id int) bool {
	for _, e := range es {
		if e.Id == id {
			return true
		}
	}
	return false
}

// Takes a new snapshot of goroutines and returns a new Entries that contains any routines only in the new snapshot.
func (es Entries) Check() Entries {
	var ret Entries

	buf := make([]byte, MaxStackTextSize)
	i := runtime.Stack(buf, true)
	for _, s := range bytes.Split(buf[:i], []byte("\n\n")) {
		x := idPattern.FindSubmatch(s)
		if id, err := strconv.Atoi(string(x[1])); err != nil {
			panic(err.Error())
		} else if !es.contains(id) {
			ret = append(ret, Entry{Id: id, Description: s})
		}
	}

	sort.Sort(esort(ret))
	return ret
}

func (es Entries) String() string {
	var b = bytes.NewBuffer(nil)
	if len(es) == 1 {
		fmt.Fprint(b, "1 Entry")
	} else {
		fmt.Fprintf(b, "%d Entries", len(es))

	}
	for _, e := range es {
		fmt.Fprintf(b, "\n%s", e)
	}
	return b.String()
}

type esort Entries

func (es esort) Len() int           { return len(es) }
func (es esort) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }
func (es esort) Less(i, j int) bool { return es[i].Id < es[j].Id }
