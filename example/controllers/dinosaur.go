package controllers

import gb "github.com/pav-studio/gorbit"

func ListDinosaurs(c *gb.Ctx) {

	c.OK(gb.JSON{
		"dinosaurs": []string{
			"Tyrannosaurus",
			"Triceratops",
			"Stegosaurus",
		},
	})
}

func GetDinosaur(c *gb.Ctx) {

	c.OK(gb.JSON{
		"id": c.Param("id"),
	})
}

func CreateDinosaur(c *gb.Ctx) {

	c.Created(gb.JSON{
		"message": "Dinosaur created",
	})
}

func UpdateDinosaur(c *gb.Ctx) {

	c.OK(gb.JSON{
		"message": "Dinosaur updated",
	})
}

func DeleteDinosaur(c *gb.Ctx) {

	c.OK(gb.JSON{
		"message": "Dinosaur removed",
	})
}