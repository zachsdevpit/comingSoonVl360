package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"
)

func main() {
	app, err := initApp()
	if err != nil {
		log.Fatal(err)
	}
	app.Static("/static", "./static")
	app.Get("/", handleGetHome)

	log.Fatal(app.Listen(os.Getenv("HTTP_LISTEN_ADDR")))
}

func initApp() (*fiber.App, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	engine := django.New("./views", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
		PassLocalsToViews:     true,
		Views:                 engine,
	})
	return app, nil
}

type demoCard struct {
	ID             int
	Package        string
	Classification string
	Description    string
	RedirectLink   string
}

var demoCards = []demoCard{
	{
		ID:             1,
		Package:        "Inspection Plus",
		Classification: "Standards of Practice",
		Description: `Reinforce your understanding and application of the Texas Administrative Code (TAC) 
    Title 22 Part 23 Chapter 535 Subchapter R Standards of Practice 535.227 - 535.233 with the comprehensive Standards of Practice module. 
    With over 350 items, learn, inspect, and quiz yourself. 
    Elevate your inspection knowledge today!`,
		RedirectLink: "http://Inspection-Sop-Demo.vl360-interactive.com",
	},
	{
		ID:             2,
		Package:        "Inspection Plus",
		Classification: "Condominium",
		Description: `Introducing the Condominium Interactive Module, the ultimate tool for mastering residential property inspection.
    With photos, videos, and graphics, evaluate systems and components. 
    Add commentary based on observations to enhance inspection skills. Over 90 items to inspect. Elevate your expertise today!`,
		RedirectLink: "http://Inspection-Condo-Demo.vl360-interactive.com",
	},
	{
		ID:             3,
		Package:        "Framing Plus",
		Classification: "Lesson",
		Description: `Discover the groundbreaking Framing Interactive Plus Module—an immersive learning experience like no other.
    Explore a two-story residential construction site at your own pace with 42 scenes. 
    Covering over 128 learning points from the International Residential Code Chapter 3, this module seamlessly merges theory with practicality through interactive graphics. 
    Elevate your framing skills today!`,
		RedirectLink: "http://Framing-Lesson-Demo.vl360-interactive.com",
	},
	{
		ID:             4,
		Package:        "Framing Plus",
		Classification: "Quiz",
		Description: `Discover the Framing Interactive Quiz Module—a revolutionary tool to gauge your framing expertise.
    Put your knowledge from the Framing Interactive Plus Module to the test with 163 carefully crafted questions.
    Showcase your mastery, elevate your skills, and become a framing authority.
    Embrace the quiz today!`,
		RedirectLink: "http://Framing-Quiz-Demo.vl360-interactive.com",
	},
	{
		ID:             5,
		Package:        "Plumbing Plus",
		Classification: "Lesson",
		Description: `Experience the Plumbing Interactive Module—an immersive learning tool for mastering plumbing.
    Discover over 65 learning points from the International Residential Code Part 4.
    Explore a single-family residence under construction, revealing what's typically concealed.
    With 15 scenes and interactive graphic overlays, gain practical insights at your own pace.
    Elevate your plumbing skills today!`,
		RedirectLink: "http://Plumbing-Lesson-Demo.vl360-interactive.com",
	},
	{
		ID:             6,
		Package:        "Plumbing Plus",
		Classification: "Quiz",
		Description: `Put your plumbing knowledge to the test with the Plumbing Interactive Quiz Module!
    With over 130 questions from the International Residential Code Part 4's 65 learning points, this quiz module reinforces your expertise.
    Challenge yourself and elevate your plumbing inspection skills today!`,
		RedirectLink: "http://Plumbing-Quiz-Demo.vl360-interactive.com",
	},
	{
		ID:             7,
		Package:        "Electrical Plus",
		Classification: "Lesson",
		Description: `Master the carefully selected portions of Part 8 Electrical International Residential codes with the Electrical Interactive Module!
    Explore 42 scenes and 90 learning points at your own pace.
    Experience an immersive learning environment that eliminates confusion and brings clarity to complex codes.
    Elevate your electrical inspection knowledge today!`,
		RedirectLink: "http://Electrical-Lesson-Demo.vl360-interactive.com",
	},
	{
		ID:             8,
		Package:        "Electrical Plus",
		Classification: "Quiz",
		Description: `Test your electrical inspection knowledge with the Electrical Interactive Quiz Module!
    Experience 131 questions derived from the 90 learning points in the Electrical Interactive Module.
    Prove your expertise with questions from select parts of Part 8 Electrical from the International Residential Code.
    Elevate your skills today!`,
		RedirectLink: "http://Electrical-Quiz-Demo.vl360-interactive.com",
	},
	{
		ID:             9,
		Package:        "Mechanical Plus",
		Classification: "Lesson",
		Description: `Dive into the future of International Residential Codes Part V Mechanical.
    Embark on an immersive journey into residential mechanical systems.
    Discover intricacies of cooling, heating, and exhaust.
    Move beyond traditional learning, and focus on vital code systems.
    Transform your understanding of codes!`,
		RedirectLink: "http://Mechanical-Lesson-Demo.vl360-interactive.com",
	},
	{
		ID:             10,
		Package:        "Mechanical Plus",
		Classification: "Quiz",
		Description: `Test your knowledge on cooling, heating, and exhaust systems.
    Dive deep into vital code aspects, challenge your understanding.
    Elevate your expertise with just a few clicks!`,
		RedirectLink: "http://Mechanical-Quiz-Demo.vl360-interactive.com",
	},
}

func handleGetHome(c *fiber.Ctx) error {
	data := fiber.Map{
		"DemoCards": demoCards,
	}
	return c.Render("home/index", data)
}
