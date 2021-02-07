package main

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) {
	entryCount++
	j.entries = append(j.entries, text)
}

func (j *Journal) RemoveEntry(index string) {
	// ...
}

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile()  {

}

// Separation of Concerns
// God Object - NOPE!

func main()  {
	j := Journal{}
	j.AddEntry("Player - 1")
	j.AddEntry("Player - 2")

	p := Persistence{
		",",
	}
	p.SaveToFile()
}
