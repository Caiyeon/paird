package store

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Initializing db store file", t, func() {
		err := Initialize("test.db")
		So(err, ShouldBeNil)

		Convey("Inserting a new team", func() {
			err := AddTeamIfNotExists("sampleteam")
			So(err, ShouldBeNil)

			Convey("The new team should be exist", func() {
				exists, err := DoesTeamExist("sampleteam")
				So(err, ShouldBeNil)
				So(exists, ShouldBeTrue)

			})

			Convey("But nonexistent teams should return false", func() {
				exists, err := DoesTeamExist("notanexistingteam")
				So(err, ShouldBeNil)
				So(exists, ShouldBeFalse)
			})
		})

		Reset(func() {
			CloseDB()
			_, err = os.Stat("test.db")
			So(err, ShouldBeNil)
			err = os.Remove("test.db")
			So(err, ShouldBeNil)
		})
	})
}
