// paste
package model

import (
	"errors"
	"labix.org/v2/mgo/bson"
	"log"
	"regexp"
	"strings"
	"time"
)

type Paste struct {
	Id         bson.ObjectId `bson:"_id"`
	Title      string
	Content    string
	CreatedOn  time.Time
	LanguageId int64
}

// const DATABASE = "mgopastebin"
const PASTES = "pastes"

func (this Paste) Add() (id bson.ObjectId, err error, createdOn time.Time) {
	session, err1 := GetDB()
	if err1 != nil {
		log.Fatal("Error on database start - Add:", err1)
	}
	collection := session.DB(database).C(PASTES)
	id = bson.NewObjectId()
	createdOn = time.Now()
	this.Id = id
	this.CreatedOn = createdOn
	err = collection.Insert(&this)
	return
}
func ToObjectId(rawId string) bson.ObjectId {
	rawId = strings.TrimLeft(rawId, "ObjectIdHex(")
	rawId = strings.TrimRight(rawId, ")")
	rawId = strings.Trim(rawId, "\"")
	id := bson.ObjectIdHex(rawId)
	return id
}

func GetPaste(id bson.ObjectId) Paste {
	session, err1 := GetDB()
	if err1 != nil {
		log.Fatal("Error on database start - GetPaste():", err1)
	}
	collection := session.DB(database).C(PASTES)
	var paste Paste
	err := collection.FindId(id).One(&paste)
	if err != nil {
		log.Fatal("Error on database get - GetPaste():", err)
	}
	return paste
}

func (this Paste) HighlightKeywords() (result string, err error) {
	var results []string
	var replacer func(string) string
	if this.LanguageId >= int64(len(Languages)) {
		result = ""
		err = errors.New("Invalid language")
		return
	}
	replacer = makeReplacerFunction(Languages[this.LanguageId])

	wordRx := regexp.MustCompile("[A-Za-z]+")
	lines := strings.Split(this.Content, "\n")
	for _, line := range lines {
		newLine := wordRx.ReplaceAllStringFunc(line, replacer)
		results = append(results, newLine)
	}
	result = strings.Join(results, "\n")
	err = nil
	return
}

func makeReplacerFunction(language Language) func(string) string {
	keywords := language.Keywords

	return func(word string) string {
		if keywords == nil {
			return word
		}
		if i := keywords.Search(word); i > 0 {
			return "<b>" + word + "</b>"
		}
		return word
	}
}

/*
func main() {
	fmt.Println("Hello World!")
}
*/
