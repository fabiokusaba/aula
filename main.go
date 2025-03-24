package main

import (
	"github.com/fabiokusaba/aula/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// Estrutura de Student
type Student struct {
	// Ao utilizarmos `json:"id"` estamos dizendo que o campo ID será serializado para JSON com o nome "id", desta forma estamos
	// convertendo os campos para o formato JSON
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

// Slice de Student
var students = []Student{
	{ID: shared.GetUuid(), Name: "João", Age: 20},
	{ID: shared.GetUuid(), Name: "Maria", Age: 22},
	{ID: shared.GetUuid(), Name: "José", Age: 21},
}

// Funções das rotas
func routeHeart(c *gin.Context) {
	// gin.H -> é um map[string]interface{}, ou seja, uma forma genérica de passar dados
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})

	c.Done()
}

func routeGetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func routeGetStudentByID(c *gin.Context) {
	var student Student

	idString := c.Param("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		return
	}

	// Capturando e convertendo o ID da rota para um valor inteiro
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}

	// Percorrendo a lista de students para verificar se o student com o ID informado existe
	// Existindo o student, retornamos o student encontrado
	// Caso não exista, retornamos um erro de "Student not found"
	for _, s := range students {
		if s.ID == id {
			student = s
		}
	}

	if student.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func routePostStudent(c *gin.Context) {
	// Criando uma variável para do tipo Student para receber os dados vindos da requisição
	var student Student

	// Recebendo os dados da requisição e armazenando na variável student (passando o endereço da variável)
	// Fazendo a tratativa caso ocorra algum erro
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		// Utilizamos o return para que a função finalize aqui e não continue a execução
		return
	}

	// Incrementando o ID do novo student
	//student.ID = students[len(students)-1].ID + 1

	// Utilizando UUID
	student.ID = shared.GetUuid()

	// Adicionando o novo student a nossa listagem de students
	students = append(students, student)

	// Retornando uma resposta de sucesso
	c.JSON(http.StatusCreated, student)
}

func routePutStudent(c *gin.Context) {
	var studentReq Student
	var studentFound Student
	var newStudents []Student

	if err := c.ShouldBindJSON(&studentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := c.Param("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Como estamos trabalhando com UUID não precisamos mais fazer a conversão
	// Capturando e convertendo o ID da rota para um valor inteiro
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}

	// Percorrendo a lista de students para verificar se o student com o ID informado existe
	// Existindo o student, atualizamos os dados do mesmo e retornamos o student atualizado
	// Caso não exista, retornamos um erro de "Student not found"
	for _, s := range students {
		if s.ID == id {
			studentFound = s
		}
	}

	if studentFound.ID == shared.GetUuidEmpty() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	studentFound.Name = studentReq.Name
	studentFound.Age = studentReq.Age

	// Atualizando a lista de students com os novos dados do student encontrado
	for _, s := range students {
		if s.ID == id {
			newStudents = append(newStudents, studentFound)
		} else {
			newStudents = append(newStudents, s)
		}
	}

	// Retornando uma resposta de sucesso com os dados atualizados do student
	students = newStudents
	c.JSON(http.StatusOK, studentFound)
}

func routeDeleteStudent(c *gin.Context) {
	var newStudents []Student

	idString := c.Param("id")
	id, err := shared.GetUuidByString(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Capturando e convertendo o ID da rota para um valor inteiro
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}

	// Percorrendo a lista de students para verificar se o student com o ID informado existe
	// Existindo o student, removemos o mesmo da lista e retornamos uma resposta de sucesso
	for _, s := range students {
		if s.ID == id {
			continue
		} else {
			newStudents = append(newStudents, s)
		}
	}

	students = newStudents
	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted",
	})
}

func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run(":8080")
}

// Ponteiros -> trabalhando com o mesmo arquivo em memória através de sua referência, então toda vez que fizer uma modificação recebo essa
// modificação de volta
func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/heart", routeHeart)

	// Criando um agrupamento de rotas
	groupStudents := c.Group("/students")
	groupStudents.GET("/", routeGetStudents)
	groupStudents.GET("/:id", routeGetStudentByID)
	groupStudents.POST("/", routePostStudent)
	groupStudents.PUT("/:id", routePutStudent)
	groupStudents.DELETE("/:id", routeDeleteStudent)

	return c
}
