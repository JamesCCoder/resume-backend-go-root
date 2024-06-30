package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/JamesCCoder/resume_backend_go/db"
	"github.com/JamesCCoder/resume_backend_go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AdminLogin(c *gin.Context) {
    adminCollection := db.GetAdminCollection()

    var admin models.Administrator
    var foundAdmin models.Administrator

    if err := c.BindJSON(&admin); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := adminCollection.FindOne(ctx, bson.M{"username": admin.Username, "password": admin.Password}).Decode(&foundAdmin)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
