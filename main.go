package main

// const defaultPort = "3000"

import (
	"gcim/example/config"

	infrastructure "gcim/example/internal/infrastructures"
)

func main() {
	// Load .env file
	config.LoadEnv()

	// e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.GET("/upload-sample/:eventId/:orgCspDocId", uploadexample.UploadExample)
	// e.GET("/get-download-url", getdownloadurlexample.GetDownloadURLExample)
	// e.Static("/static", "static")
	// e.Logger.Fatal(e.Start(":1313"))
	infrastructure.InitRouter()
}

// func main() {
// 	// Load .env file
// 	config.LoadEnv()

// 	e := echo.New()

// 	e.GET("/", createSample)

// 	e.Logger.Fatal(e.Start(":1313"))
// }

// func createSample(ctx echo.Context) error {
// 	projectID := "test-project"

// 	client, err := firestore.NewClient(ctx.Request().Context(), projectID)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}
// 	rep := repository.NewTaskRepository(client)
// 	_, err = rep.Insert(ctx.Request().Context(), &model.Task{
// 		ID: "id223",
// 	})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Printf("sss1")
// 	}
// 	return nil
// }
