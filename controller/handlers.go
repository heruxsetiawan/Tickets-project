package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"tickets-project/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = time.Now()
	_, err := models.DB.Exec("INSERT INTO Users (Username, Email, PasswordHash, CreatedAt) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.PasswordHash, user.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser models.User
	row := models.DB.QueryRow("SELECT UserID, Username, Email, PasswordHash FROM Users WHERE PasswordHash = $1 AND Username=$2 ", user.PasswordHash, user.Username)
	if err := row.Scan(&storedUser.UserID, &storedUser.Username, &storedUser.Email, &storedUser.PasswordHash); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username or password"})
		return
	}
	fmt.Println(user.PasswordHash + " ==" + storedUser.PasswordHash)
	if user.PasswordHash != storedUser.PasswordHash {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username or password"})
		return
	}

	token, err := CreateToken(storedUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetTickets(c *gin.Context) {
	rows, err := models.DB.Query("SELECT TicketID, CreatorID, Title, Description, Status, StartTime, EndTime, CreatedAt, UpdatedAt FROM Tickets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var ticket models.Ticket
		if err := rows.Scan(&ticket.TicketID, &ticket.CreatorID, &ticket.Title, &ticket.Description, &ticket.Status, &ticket.StartTime, &ticket.EndTime, &ticket.CreatedAt, &ticket.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tickets = append(tickets, ticket)
	}

	c.JSON(http.StatusOK, tickets)
}

func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()
	_, err := models.DB.Exec("INSERT INTO Tickets (CreatorID, Title, Description, Status, StartTime, EndTime, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", ticket.CreatorID, ticket.Title, ticket.Description, ticket.Status, ticket.StartTime, ticket.EndTime, ticket.CreatedAt, ticket.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ticket)
}

func UpdateTicket(c *gin.Context) {
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticket.UpdatedAt = time.Now()
	_, err = models.DB.Exec("UPDATE Tickets SET Title = $1, Description = $2, Status = $3, StartTime = $4, EndTime = $5, UpdatedAt = $6 WHERE TicketID = $7", ticket.Title, ticket.Description, ticket.Status, ticket.StartTime, ticket.EndTime, ticket.UpdatedAt, ticketID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func DeleteTicket(c *gin.Context) {
	ticketID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	_, err = models.DB.Exec("DELETE FROM Tickets WHERE TicketID = $1", ticketID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted"})
}

func GetTasks(c *gin.Context) {
	ticketID, err := strconv.Atoi(c.Param("ticketID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	rows, err := models.DB.Query("SELECT TaskID, TicketID, Title, Description, Status, StartTime, EndTime, CreatedAt, UpdatedAt FROM Tasks WHERE TicketID = $1", ticketID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.TaskID, &task.TicketID, &task.Title, &task.Description, &task.Status, &task.StartTime, &task.EndTime, &task.CreatedAt, &task.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	_, err := models.DB.Exec("INSERT INTO Tasks (TicketID, Title, Description, Status, StartTime, EndTime, CreatedAt, UpdatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", task.TicketID, task.Title, task.Description, task.Status, task.StartTime, task.EndTime, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.UpdatedAt = time.Now()
	_, err = models.DB.Exec("UPDATE Tasks SET Title = $1, Description = $2, Status = $3, StartTime = $4, EndTime = $5, UpdatedAt = $6 WHERE TaskID = $7", task.Title, task.Description, task.Status, task.StartTime, task.EndTime, task.UpdatedAt, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	_, err = models.DB.Exec("DELETE FROM Tasks WHERE TaskID = $1", taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func CreateTicketAssignment(c *gin.Context) {
	var ticketAssignment models.TicketAssignment
	if err := c.ShouldBindJSON(&ticketAssignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticketAssignment.AssignedAt = time.Now()

	query := `INSERT INTO TicketAssignments (TicketID, AssigneeID, AssignedAt) VALUES ($1, $2, $3) RETURNING AssignmentID`
	err := models.DB.QueryRow(query, ticketAssignment.TicketID, ticketAssignment.AssigneeID, ticketAssignment.AssignedAt).Scan(&ticketAssignment.AssignmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticketAssignment)
}

func GetTicketAssignments(c *gin.Context) {
	rows, err := models.DB.Query("SELECT AssignmentID, TicketID, AssigneeID, AssignedAt FROM TicketAssignments")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var ticketAssignments []models.TicketAssignment
	for rows.Next() {
		var ticketAssignment models.TicketAssignment
		if err := rows.Scan(&ticketAssignment.AssignmentID, &ticketAssignment.TicketID, &ticketAssignment.AssigneeID, &ticketAssignment.AssignedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ticketAssignments = append(ticketAssignments, ticketAssignment)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticketAssignments)
}

func GetTicketAssignmentsByAssigneeID(c *gin.Context) {
	assigneeID, err1 := strconv.Atoi(c.Param("assigneeID"))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid assignee ID"})
		return
	}

	var ticketAssignments []models.TicketAssignment
	query := "SELECT AssignmentID, TicketID, AssigneeID, AssignedAt FROM TicketAssignments WHERE AssigneeID = $1"
	rows, err := models.DB.Query(query, assigneeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var ticketAssignment models.TicketAssignment
		if err := rows.Scan(&ticketAssignment.AssignmentID, &ticketAssignment.TicketID, &ticketAssignment.AssigneeID, &ticketAssignment.AssignedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ticketAssignments = append(ticketAssignments, ticketAssignment)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ticketAssignments)
}
func UpdateTicketAssignment(c *gin.Context) {
	id := c.Param("id")
	var ticketAssignment models.TicketAssignment
	if err := c.ShouldBindJSON(&ticketAssignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE TicketAssignments SET TicketID = $1, AssigneeID = $2, AssignedAt = $3 WHERE AssignmentID = $4"
	_, err := models.DB.Exec(query, ticketAssignment.TicketID, ticketAssignment.AssigneeID, ticketAssignment.AssignedAt, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "TicketAssignment updated"})
}

func DeleteTicketAssignment(c *gin.Context) {
	id := c.Param("id")
	query := "DELETE FROM TicketAssignments WHERE AssignmentID = $1"
	_, err := models.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "TicketAssignment deleted"})
}
