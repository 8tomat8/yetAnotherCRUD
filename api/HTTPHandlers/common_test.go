package HTTPHandlers

import (
	"io/ioutil"

	"time"

	"github.com/8tomat8/yetAnotherCRUD/entity"
	"github.com/sirupsen/logrus"
)

var logger = &logrus.Logger{
	Out:       ioutil.Discard,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}

var storageUsers = []entity.User{
	{
		UserID:    431,
		Birthdate: time.Date(1990, time.December, 15, 0, 0, 0, 0, time.UTC),
		Firstname: "Firstname431",
		Lastname:  "Lastname431",
		Username:  "Username431",
		Password:  "SuperMegaPassword431",
		Sex:       "male",
	},
	{
		UserID:    432,
		Birthdate: time.Date(1991, time.December, 15, 0, 0, 0, 0, time.UTC),
		Firstname: "Firstname432",
		Lastname:  "Lastname432",
		Username:  "Username432",
		Password:  "SuperMegaPassword432",
		Sex:       "female",
	},
	{
		UserID:    433,
		Birthdate: time.Date(1992, time.December, 15, 0, 0, 0, 0, time.UTC),
		Firstname: "Firstname433",
		Lastname:  "Lastname433",
		Username:  "Username433",
		Password:  "SuperMegaPassword433",
		Sex:       "female",
	},
}
