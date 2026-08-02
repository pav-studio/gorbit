package controllers

import gb "github.com/pav-studio/gorbit"

func ListTrains(c *gb.Ctx) {

	c.OK(gb.JSON{
		"trains": []string{
			"Comet Express",
			"Jurassic Flyer",
		},
	})
}

func BoardTrain(c *gb.Ctx) {

	c.OK(gb.JSON{
		"message": "You boarded the train!",
	})
}