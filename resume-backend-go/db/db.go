package db

import (
	"github.com/JamesCCoder/resume_backend_go/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetStudentCollection() *mongo.Collection {
    return config.MongoDBClient.Database("resume-project1").Collection("students")
}

func GetProfessorCollection() *mongo.Collection {
    return config.MongoDBClient.Database("resume-project1").Collection("professors")
}

func GetAdminCollection() *mongo.Collection {
    return config.MongoDBClient.Database("resume-project1").Collection("administrators")
}
