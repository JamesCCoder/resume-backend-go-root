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

func GetStudents(c *gin.Context) {
    studentCollection := db.GetStudentCollection()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var students []models.Student
    cursor, err := studentCollection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var student models.Student
        cursor.Decode(&student)
        students = append(students, student)
    }

    c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
    studentCollection := db.GetStudentCollection()

    var student models.Student
    if err := c.BindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    student.ID = primitive.NewObjectID()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := studentCollection.InsertOne(ctx, student)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, student)
}

func GetStudent(c *gin.Context) {
    studentCollection := db.GetStudentCollection()
    professorCollection := db.GetProfessorCollection()

    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var student models.Student
    err := studentCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&student)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
        return
    }

    var professors []models.Professor
    for _, professorID := range student.Professors {
        var professor models.Professor
        err := professorCollection.FindOne(ctx, bson.M{"_id": professorID}).Decode(&professor)
        if err == nil {
            professors = append(professors, professor)
        }
    }

    c.JSON(http.StatusOK, gin.H{"student": student, "professors": professors})
}

func UpdateStudent(c *gin.Context) {
    studentCollection := db.GetStudentCollection()

    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)

    var student models.Student
    if err := c.BindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    update := bson.M{"name": student.Name, "sex": student.Sex, "email": student.Email, "professors": student.Professors}
    _, err := studentCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
    studentCollection := db.GetStudentCollection()

    id := c.Param("id")
    objID, _ := primitive.ObjectIDFromHex(id)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := studentCollection.DeleteOne(ctx, bson.M{"_id": objID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
