package api

import (
	infra_controller "github.com/fabiokusaba/aula/api/controller/infra"
	student_controller "github.com/fabiokusaba/aula/api/controller/students"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", infra_controller.Heart)

	groupStudents := s.Engine.Group("/students")
	groupStudents.GET("/", student_controller.List)
	groupStudents.GET("/:id", student_controller.Details)
	groupStudents.POST("/", student_controller.Create)
	groupStudents.PUT("/:id", student_controller.Update)
	groupStudents.DELETE("/:id", student_controller.Delete)
}
