package media

type Socialmedia interface {
	Feed() []string
	Fame() int
}

type Facebook struct {
	URL     string
	Name    string
	Friends int
}

type Twitter struct {
	URL       string
	Username  string
	Followers int
}

type Linkedin struct {
	URL         string
	Name        string
	Connections int
}

//Feed returns the laest facebook posts
func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here's my cool new selfie",
		"what is code?",
	}
}

//Fame tell how faous a useris. The higher
// he number of friends the more famous
func (f *Facebook) Fame() int {
	return f.Friends
}

//Fees returns the latest Linkedin posts
func (l *Linkedin) Feed() []string {
	return []string{
		"LinkedIn feeds",
		"Hey, I just started a new positin at Hotels.ng",
	}
}

//Fame tell how famous a useris. The higher
// the number of friends the more famous
func (l *Linkedin) Fame() int {
	return l.Connections
}

//Fees returns the latest Linkedin posts
func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Coding is basically copying other people's work",
	}
}

//Fame tell how famous a useris. The higher
// the number of friends the more famous
func (t *Twitter) Fame() int {
	return t.Followers
}
