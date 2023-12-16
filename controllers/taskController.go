package controllers

import (
	"Go-ToDo/models"
	"Go-ToDo/postgres"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	postgres.DB.Find(&tasks)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"result":  tasks,
	})
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("taskId")
	var task models.Task
	result := postgres.DB.Where("id=?", id).First(&task)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Task not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"result":  task,
	})
}

func CreateTask(c *fiber.Ctx) error {

	var req models.TaskRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	newTask := models.Task{
		TaskItem:  req.Item,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	if err := postgres.DB.Create(&newTask).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create task",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "New task successfully created",
		"result":  newTask,
	})
}

func CompleteTask(c *fiber.Ctx) error {
	id := c.Params("taskId")
	result := postgres.DB.Model(&models.Task{}).Where("id=?", id).Update("Completed", true).Update("UpdatedAt", time.Now())
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Task not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task Updated Successfully",
	})
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("taskId")
	var task models.Task
	result := postgres.DB.Model(&models.Task{}).Where("id=?", id).Delete(&task)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Task not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task Deleted Successfully",
	})
}
