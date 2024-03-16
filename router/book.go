 package router

 import (
  "github.com/bmdavis419/fiber-mongo-example/common"
	"github.com/bmdavis419/fiber-mongo-example/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
 )

 func AddBookGroup(app *fiber.App) {
   bookGroup := app.Group("/books")

   bookGroup.Get("/", getBooks)
   bookGroup.Get("/:id", getBook)
   bookGroup.Post("/", createBook)
   bookGroup.Put("/:id", updateBook)
   bookGroup.Delete("/:id", deleteBook)
 }

 func getBooks(c *fiber.Ctx) error {
   coll := common.GetDBCollection("books")

   // find all books
   books := make([]models.Book, 0)
   cursor, err := coll.Find(c.Context(), bson.M{})

   if err != nil {
     return c.Status(500).JSON(fiber.Map{
       "error": err.Error(),
     })
   }
   

   //Iterate over the cursor
   for cursor.Next(c.Context()) {
     book := models.Book{}
     err := cursor.Decode(&book)
     if err != nil {
       return c.Status(500).JSON(fiber.Map{
         "error": err.Error(),
       })
     }
     books = append(books, book)
   }

   return c.Status(200).JSON(fiber.Map{"data": books})
 }
