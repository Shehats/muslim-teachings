package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTeachings gets all teachings
func GetTeachings(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

// AddTeachings adds teaching
func AddTeachings(context *gin.Context) {

}

// UpdateTeachings updates teachings
func UpdateTeachings(context *gin.Context) {

}

// DeleteTeachings deletes teachings
func DeleteTeachings(context *gin.Context) {

}

// GetTeaching gets teaching
func GetTeaching(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}

// AddTeaching adds teaching
func AddTeaching(context *gin.Context) {

}

// UpdateTeaching updates teaching
func UpdateTeaching(context *gin.Context) {

}

// DeleteTeaching deletes teaching
func DeleteTeaching(context *gin.Context) {

}

// SearchTeaching searches context via params
func SearchTeaching(context *gin.Context) {

}
