package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/JamesCCoder/resume_backend_go/db"
	"github.com/JamesCCoder/resume_backend_go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfessors(c *gin.Context) {
    professorCollection := db.GetProfessorCollection()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var professors []models.Professor
    cursor, err := professorCollection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var professor models.Professor
        cursor.Decode(&professor)
        professors = append(professors, professor)
    }

    c.JSON(http.StatusOK, professors)
}

func CreateProfessor(c *gin.Context) {
    professorCollection := db.GetProfessorCollection()

    var professor models.Professor
    if err := c.BindJSON(&professor); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    professor.ID = primitive.NewObjectID()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := professorCollection.InsertOne(ctx, professor)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, professor)
}

func GetProfessor(c *gin.Context) {
    professorCollection := db.GetProfessorCollection()
    studentCollection := db.GetStudentCollection()

    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var professor models.Professor
    err := professorCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&professor)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Professor not found"})
        return
    }

    var students []models.Student
    for _, studentID := range professor.Students {
        var student models.Student
        err := studentCollection.FindOne(ctx, bson.M{"_id": studentID}).Decode(&student)
        if err == nil {
            students = append(students, student)
        }
    }

    c.JSON(http.StatusOK, gin.H{"professor": professor, "students": students})
}

func UpdateProfessor(c *gin.Context) {
    professorCollection := db.GetProfessorCollection()

    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)

    var professor models.Professor
    if err := c.BindJSON(&professor); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    update := bson.M{"name": professor.Name, "sex": professor.Sex, "email": professor.Email, "students": professor.Students}
    _, err := professorCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, professor)
}

func DeleteProfessor(c *gin.Context) {
    professorCollection := db.GetProfessorCollection()

    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := professorCollection.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Professor deleted"})
}
