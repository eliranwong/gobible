package bible

import (
	"fmt"
	"strings"

	"github.com/eliranwong/gobible/internal/check"
	"github.com/eliranwong/gobible/internal/parser"
)

// read multiple verses
func GetVerseBlock(module string, bcv []int) (string, string, string) {
	reference := parser.BcvToVerseReference(bcv)
	var block strings.Builder

	db := getDb(module)
	defer db.Close()

	var b, c, cs, vs, ce, ve int
	b, cs, vs = bcv[0], bcv[1], bcv[2]

	addContent := func(b, c, vs, ve int, f string) {
		var q, ref, content string
		if vs == ve {
			q = fmt.Sprintf("SELECT DISTINCT Chapter, Verse, Scripture FROM Verses WHERE Book=%v AND Chapter=%v AND Verse=%v", b, c, vs)
		} else if vs <= 0 && ve <= 0 {
			q = fmt.Sprintf("SELECT DISTINCT Chapter, Verse, Scripture FROM Verses WHERE Book=%v AND Chapter=%v ORDER BY Verse", b, c)
		} else if vs > 0 && ve > 0 {
			q = fmt.Sprintf("SELECT DISTINCT Chapter, Verse, Scripture FROM Verses WHERE Book=%v AND Chapter=%v AND Verse BETWEEN %v AND %v ORDER BY Verse", b, c, vs, ve)
		} else if vs > 0 {
			q = fmt.Sprintf("SELECT DISTINCT Chapter, Verse, Scripture FROM Verses WHERE Book=%v AND Chapter=%v AND Verse>=%v ORDER BY Verse", b, c, vs)
		}
		results, err := db.Query(q)
		check.DbErr(err)
		defer results.Close()

		for results.Next() {
			var chapter, verse int
			var scripture string
			err = results.Scan(&chapter, &verse, &scripture)
			check.DbErr(err)

			if f == "cv" {
				ref = fmt.Sprintf("%v:%v ", chapter, verse)
			} else if f == "v" {
				ref = fmt.Sprintf("%v ", verse)
			} else {
				ref = ""
			}
			content = fmt.Sprintf("%v%v ", ref, formatVerseText(scripture))
			block.WriteString(content)
		}
		err = results.Err()
		check.DbErr(err)
	}

	if len(bcv) == 3 {
		addContent(b, cs, vs, vs, "")
	} else if len(bcv) == 4 {
		ve = bcv[3]
		addContent(b, cs, vs, ve, "v")
	} else if len(bcv) == 5 {
		ce, ve = bcv[3], bcv[4]

		if cs > ce {
			// skip invalid range
		} else if cs == ce {
			addContent(b, cs, vs, ve, "v")
		} else {
			c = cs
			addContent(b, c, vs, 0, "cv")
			c += 1
			for c < ce {
				addContent(b, c, 0, 0, "cv")
				c += 1
			}
			if c == ce {
				addContent(b, c, 1, ve, "cv")
			}
		}
	}
	return module, reference, block.String()
}
