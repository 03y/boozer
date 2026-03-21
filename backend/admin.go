package main

import (
	"context"
	"net/http"
	"strconv"

	"boozer/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (a *App) GetReports(c *gin.Context) {
	rows, err := a.DB.Query(context.Background(), "SELECT report_id, item_id, user_id, reason, created FROM item_reports ORDER BY created DESC")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	reports := make([]models.ItemReport, 0)
	for rows.Next() {
		var report models.ItemReport
		err := rows.Scan(&report.Report_id, &report.Item_id, &report.User_id, &report.Reason, &report.Created)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		reports = append(reports, report)
	}

	c.JSON(http.StatusOK, reports)
}

func (a *App) UpdateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	_, err = a.DB.Exec(context.Background(), "UPDATE items SET name=$1, units=$2 WHERE item_id=$3", item.Name, item.Units, itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (a *App) ClearItemReports(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	_, err = a.DB.Exec(context.Background(), "DELETE FROM item_reports WHERE item_id=$1", itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (a *App) AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := parseJWT(tokenString, a.JWT_KEY)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var isAdmin bool
		err = a.DB.QueryRow(context.Background(), "SELECT admin FROM users WHERE username=$1", claims["username"]).Scan(&isAdmin)
		if err != nil {
			if err == pgx.ErrNoRows {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !isAdmin {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
