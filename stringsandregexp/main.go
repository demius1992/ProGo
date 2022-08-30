package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	Mapping()
	//r := regexp.MustCompile("p([a-z]*)")
	//fmt.Println(r.MatchString("peach"))
}

func Mapping() {
	text := "It was a boat. A small boat."
	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		} else if r == '.' {
			return ':'
		}
		return r
	}
	mapped := strings.Map(mapper, text)
	fmt.Println("Mapped:", mapped)
}

func replaceStringFunc() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	replaced := pattern.ReplaceAllStringFunc(description, func(s string) string { return "This is the replacement content" })
	fmt.Println(replaced)
}

func replaceAllStrings() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	template := "(type: ${type}, capacity: ${capacity})"
	replaced := pattern.ReplaceAllString(description, template)
	fmt.Println(replaced)
}

func subExpNames() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	subs := pattern.FindStringSubmatch(description)
	for _, name := range []string{"type", "capacity"} {
		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
	}
}

func REgexp2() {
	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	firstIndex := pattern.FindString(description)
	allIndices := pattern.FindAllString(description, -1)
	fmt.Println("First str", firstIndex)
	for i, idx := range allIndices {
		fmt.Println("str", i, "=", idx)
	}
	//First index 0 - 5 = Kayak
	//Index 0 = 0 - 5 = Kayak
	//Index 1 = 9 - 13 = boat
}

func regexpCompile() {
	pattern, compileErr := regexp.Compile("[A-z]oat")
	description := "A boat for one person"
	question := "Is that a goat?"
	preference := "I like oats"
	if compileErr == nil {
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}
	//Description: false
	//Question: false
	//Preference: false

}

func builder() {
	text := "It was a boat. A small boat."

	var builder strings.Builder

	for _, sub := range strings.Fields(text) {
		if sub == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteRune(' ')
	}

	fmt.Println("String:", builder.String())
	//String: It was a boat. A very small boat.
}

func joings() {
	text := "It was a boat. A small boat."
	elements := strings.Fields(text)
	joined := strings.Join(elements, " ")
	fmt.Println("Joined:", joined)

	//Joined: It was a boat. A small boat.
}

func newReplacer() {
	text := "It was a boat. A small boat."
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	replaced := replacer.Replace(text)

	fmt.Println("Replaced:", replaced)
}

func replaceString() {
	text := "It was a boat. A small boat."
	replace := strings.Replace(text, "boat", "canoe", 1)
	replaceAll := strings.ReplaceAll(text, "boat", "truck")
	fmt.Println("Replace:", replace)
	fmt.Println("Replace All:", replaceAll)
}

func trimFunc() {
	description := "A boat for one person"

	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}

	trimmed := strings.TrimFunc(description, trimmer)
	fmt.Println("Trimmed:", trimmed)
}

func filding() {
	description := "This  is  double  spaced"

	splitter := func(r rune) bool {
		return r == ' '
	}

	splits := strings.FieldsFunc(description, splitter)
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
}

func Splits() {
	description := "A boat for one person"
	splits := strings.Split(description, " ")
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}
}

func indexFunc() {
	description := "A boat for one person"
	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'e'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))
}

func funcName2() {
	description := "A boat for one person"

	fmt.Println("Count:", strings.Count(description, "o"))
	fmt.Println("Index:", strings.Index(description, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("IndexAny:", strings.IndexAny(description, "cde"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcdn"))
}

func funcName() {
	description := "A boat for sailing"

	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.ToTitle(description))
	product := "Kayak"

	for _, char := range product {
		fmt.Println(string(char), "is Upper:", unicode.IsUpper(char))
	}
	fmt.Println("Product:", product)
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAyAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
}
