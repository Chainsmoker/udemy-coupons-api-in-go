package main

import (
	"strings"
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
)

func main(){
	var contador = 0
	courses := []string{}
	courses_udemy := []string{}
	app := fiber.New()
	domains := []string{"www.cursosdev.com/","www.cursosdev.com","cursosdev.com","cursosdev.com/"}
	
	ctx := colly.NewCollector(
		colly.AllowedDomains(
			domains...,
		),
	)
	ctx1 := colly.NewCollector(
		colly.AllowedDomains(
			domains...,
		),
	)
	ctx.OnHTML("a[href]", func(e *colly.HTMLElement){
		if contador == 16 {
			ctx1.OnHTML("a[href]", func(e *colly.HTMLElement){
				if strings.Contains(e.Attr("href"), "/ad.admitad.com/g/05dgete24s372c5c98e4b3e3b7aadc/") {
					courses_udemy = append(courses_udemy, e.Attr("href"))
				}
			})
			for _, course := range courses {
				ctx1.Visit(course)
			}
		}
		if contador < 17 {
			if strings.Contains(e.Attr("href"), "/coupons-udemy/"){
				courses = append(courses, e.Attr("href"))
				contador++
			}
		}
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		ctx.Visit("https://www.cursosdev.com/coupons")
		response , _ := json.Marshal(courses_udemy)
		return c.SendString(string(response))
    })

    app.Listen(":3000")
}