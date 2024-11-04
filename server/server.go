package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jgbz/product_catalog/service"
)

// Server wraps the service layer and expose the api endpoints.
type Server struct {
	service service.Service
}

// New creates a new server
func New(service service.Service) *Server {
	return &Server{
		service: service,
	}
}

// Listen register the api endpoints and starts listening to requests
func (s *Server) Listen() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		if err := s.service.RepoHealth(); err != nil {
			return c.Status(http.StatusServiceUnavailable).JSON(err)
		}
		return c.SendStatus(http.StatusOK)
	})

	app.Get("/feed", func(c *fiber.Ctx) error {
		var customerID string
		if customerID = c.Query("customer"); customerID == "" {
			return c.SendStatus(http.StatusBadRequest)
		}

		// Get current page
		// If it fails to parse the page default to 1.
		page := 1
		if pageParam := c.Query("page"); pageParam != "" {
			v, err := strconv.Atoi(pageParam)
			if err == nil {
				page = v
			}
		}

		feed, err := s.service.GetFeedByClient(&service.FeedRequest{ClientID: customerID, Page: page})
		if err != nil {
			log.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.Status(http.StatusOK).JSON(feed)
	})

	app.Post("/campaign", func(c *fiber.Ctx) error {
		req := &service.CampaignRequest{}
		if err := c.BodyParser(req); err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		res, err := s.service.CreateCampaign(req)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.Status(http.StatusCreated).JSON(res)
	})

	app.Static("/", "./static")
	app.Listen(":8080")
}
