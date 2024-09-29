package utils

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func testRoute1(c *gin.Context) {
	testInfo := map[string]interface{}{
		"name": "Kyle",
		"id":   1,
	}

	newToken, err := CreateToken(testInfo)

	if err != nil {
		c.IndentedJSON(401, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"token": newToken})
}

func testRoute2(c *gin.Context) {
	type content struct {
		Email string
		Name  string
	}

	// Check correct type of receiving data
	if c.ContentType() != "application/json" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "content type should by 'application/json'"})
		return
	}

	var body content

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid body of request" + err.Error()})
		return
	}

	// Get data from token
	temp, err := c.Get("claims")

	if !err {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "your token does not contain information needed"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"1": temp, "2": body})
}

func getRecipeInformation(c *gin.Context) {

	recipe_id, err := strconv.Atoi(strings.Trim(c.Param("recipe_id"), "/"))
	if err != nil {
		c.IndentedJSON(404, gin.H{"err": err.Error()})
		return
	}

	returnData := map[string]interface{}{}

	recipe, err := Dbx.Query(context.Background(), "select * from recipes where recipe_id = $1", recipe_id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"err": err.Error()})
		return
	}

	isMore := recipe.Next()

	recipeFields := recipe.FieldDescriptions()
	recipeValues, err := recipe.Values()
	if err != nil {
		c.IndentedJSON(500, gin.H{"err": err.Error()})
		return
	}

	for i, recipeField := range recipeFields {
		returnData[recipeField.Name] = recipeValues[i]
	}

	if isMore {
		fmt.Println("there is more than one items with this id: ", recipe_id)
	}

	recipe.Close()

	//
	// Adding ingredients
	//

	ingredients, err := Dbx.Query(context.Background(),
		"select a.ingredient_id, a.quantity, a.unit, b.name from recipe_ingredients a join "+
			"ingredients b on a.recipe_id = $1 and "+
			"a.ingredient_id = b.ingredient_id", recipe_id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"err": err.Error()})
		return
	}

	ingredientsSlice := []map[string]interface{}{}

	for ingredients.Next() {

		ingredientsFields := ingredients.FieldDescriptions()
		ingredientsValues, err := ingredients.Values()
		if err != nil {
			c.IndentedJSON(500, gin.H{"err": err.Error()})
			return
		}

		tempMap := map[string]interface{}{}

		for i, ingredient := range ingredientsFields {
			tempMap[ingredient.Name] = ingredientsValues[i]
		}

		ingredientsSlice = append(ingredientsSlice, tempMap)

	}

	returnData["proportions"] = ingredientsSlice

	ingredients.Close()
	//
	// End of adding ingredients
	//

	//
	// Start adding steps
	//

	stepSlice := []map[string]interface{}{}

	steps, err := Dbx.Query(context.Background(), "select * from instructions where recipe_id = $1", recipe_id)
	if err != nil {
		c.IndentedJSON(500, gin.H{"err": err.Error()})
		return
	}

	for steps.Next() {
		tempStep := map[string]interface{}{}

		stepFields := steps.FieldDescriptions()
		stepValues, err := steps.Values()
		if err != nil {
			c.IndentedJSON(500, gin.H{"err": err.Error()})
			return
		}

		for i, step := range stepFields {
			tempStep[step.Name] = stepValues[i]
		}

		stepSlice = append(stepSlice, tempStep)
	}

	returnData["steps"] = stepSlice

	steps.Close()

	//
	// End adding steps
	//

	c.IndentedJSON(200, returnData)
}

// Register new user in the system and return JWT-token if
// all went well
func registerUser(c *gin.Context) {
	type content struct {
		Password string
		Name     string
		Username string
	}

	if c.ContentType() != "application/json" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "content type should by 'application/json'"})
		return
	}

	var body = content{}

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUuid := uuid.New().ID()

	hashedPassword, err := Argon2Hasher.HashGen([]byte(body.Password), []byte{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsToClose, errr := Dbx.Query(context.Background(), "INSERT INTO users "+
		"(user_id, username, user_name, user_password) "+
		"VALUES($1, $2, $3, $4);", newUuid, body.Username, body.Name, hashedPassword.stringToStore)

	if errr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errr.Error()})
		return
	}

	rowsToClose.Close()

	newToken, err := CreateToken(map[string]interface{}{"id": newUuid, "name": body.Name})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", newToken)
	c.IndentedJSON(200, gin.H{"newToken": newToken})
}

// Returns all available diets and ids of them
func getAllDiets(c *gin.Context) {
	diets, err := Dbx.Query(context.Background(), "select * from diets")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sliceOfDiets := []map[string]interface{}{}

	for diets.Next() {

		temp := map[string]interface{}{}

		dietsRows := diets.FieldDescriptions()
		dietsValues, err := diets.Values()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for i, value := range dietsRows {
			temp[value.Name] = dietsValues[i]
		}

		sliceOfDiets = append(sliceOfDiets, temp)

	}

	c.IndentedJSON(http.StatusOK, sliceOfDiets)
}

// Returns slice of cuisines and they ids
func getAllCuisines(c *gin.Context) {
	cuisines, err := Dbx.Query(context.Background(), "select * from cuisines")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cuisinesSlice := []map[string]interface{}{}

	for cuisines.Next() {

		cuisinesFields := cuisines.FieldDescriptions()
		cuisinesValues, err := cuisines.Values()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		temp := map[string]interface{}{}

		for i, cuisine := range cuisinesFields {
			temp[cuisine.Name] = cuisinesValues[i]
		}

		cuisinesSlice = append(cuisinesSlice, temp)

	}

	c.IndentedJSON(http.StatusOK, cuisinesSlice)

}

// Returns all ingredients with they ids
func getAllIngredients(c *gin.Context) {
	cuisines, err := Dbx.Query(context.Background(), "select * from ingredients")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ingredientsSlice := []map[string]interface{}{}

	for cuisines.Next() {

		ingredientsFields := cuisines.FieldDescriptions()
		ingredientsValues, err := cuisines.Values()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		temp := map[string]interface{}{}

		for i, cuisine := range ingredientsFields {
			temp[cuisine.Name] = ingredientsValues[i]
		}

		ingredientsSlice = append(ingredientsSlice, temp)

	}

	c.IndentedJSON(http.StatusOK, ingredientsSlice)

}

func getRecipesByIngredient(c *gin.Context) {
	ingredientId, err := strconv.Atoi(strings.Trim(c.Param("ingredientId"), "/"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	recipes, err := Dbx.Query(context.Background(),
		"select a.recipe_id, a.name, a.description, a.servings, a.prep_time, a.cook_time, a.image_url, a.source_url, a.author from recipes a join (select * from ingredients "+
			" b join recipe_ingredients c on b.ingredient_id = $1 where b.ingredient_id = c.ingredient_id) d "+
			" on a.recipe_id = d.recipe_id",
		ingredientId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	recipesSlice := []map[string]interface{}{}

	for recipes.Next() {
		temp := map[string]interface{}{}

		recipeFields := recipes.FieldDescriptions()
		recipeValues, err := recipes.Values()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for i, field := range recipeFields {
			temp[field.Name] = recipeValues[i]
		}

		recipesSlice = append(recipesSlice, temp)

	}

	c.IndentedJSON(http.StatusOK, recipesSlice)
}

func getRecipesByCuisine(c *gin.Context) {
	cuisineId, err := strconv.Atoi(strings.Trim(c.Param("cuisineId"), "/"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	recipes, err := Dbx.Query(context.Background(),
		"select a.recipe_id, a.name, a.description, a.servings, a.prep_time, a.cook_time, a.image_url, a.source_url, a.author from recipes a join (select * from cuisines "+
			" b join recipe_cuisines c on b.cuisine_id = $1 where b.cuisine_id = c.cuisine_id) d "+
			" on a.recipe_id = d.recipe_id",
		cuisineId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cuisinesSlice := []map[string]interface{}{}

	for recipes.Next() {
		temp := map[string]interface{}{}

		recipeFields := recipes.FieldDescriptions()
		recipeValues, err := recipes.Values()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for i, field := range recipeFields {
			temp[field.Name] = recipeValues[i]
		}

		cuisinesSlice = append(cuisinesSlice, temp)

	}

	c.IndentedJSON(http.StatusOK, cuisinesSlice)
}

func conditionalRollback(b *bool) {
	if *b {
		_, err := Dbx.Exec(context.Background(),
			"ROLLBACK;")
		if err != nil {
			fmt.Println("Failed to rollback. Data in database could be invalid.")
			return
		}
		fmt.Println(RedColor + strings.ToUpper("Some things went wrong. Rollback.") + ResetColor)
	}
}

func addRecipe(c *gin.Context) {
	type ingredient struct {
		IngredientId uint32
		Name         string
		Quantity     float64
		Unit         string
	}

	type instruction struct {
		InstructionId   uint32
		Step_number     uint16
		InstructionText string
		InstructionTime string
	}

	type cuisine struct {
		CuisineId uint32
		Name      string
	}

	type diet struct {
		DietId uint32
		Name   string
	}

	type content struct {

		// Recipe information we receive from client
		Name        string
		Description string
		Servings    uint8
		Prep_time   uint16
		Cook_time   uint16
		Image_url   string
		Source_url  string

		// Ingredient list
		IngredientList []ingredient // Done

		// Instruction list
		InstructionList []instruction // Done

		// Other information
		Cuisines []cuisine // Client side
		Diets    []diet    // Client side
	}

	usrID, idExists := c.Get("id")
	if !idExists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing required data in token"})
		return
	}

	// Id of user that created recipe
	userId, err := strconv.Atoi(usrID.(string))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing required data in body of the request"})
		return
	}

	var body content

	err = c.BindJSON(&body)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing required data in body of the request " + err.Error()})
		return
	}

	recipeUuid := uuid.New()

	// New recipe uuid
	var recipeUuidInt = recipeUuid.ID()

	// Set author ID as user ID request came from

	// Set new uuid for recipe

	// Ingredient to be added to the db
	newIngredients := []ingredient{}

	needToRollback := false
	_, err = Dbx.Exec(context.Background(),
		"BEGIN;")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to begin transaction"})
		return
	}
	defer conditionalRollback(&needToRollback)

	// Check if ingredient already in database
	// if not we give it a uuid and add to slice to add it to db later
	// Also we change name of ingredient to lower case
	for i, ingredient := range body.IngredientList {
		temp, err := Dbx.Query(context.Background(),
			"select find_id_in_the_table_by_name_ingredient($1)",
			strings.ToLower(ingredient.Name))

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			needToRollback = true
			return
		}

		temp.Next()
		if t, err := temp.Values(); err == nil && t[0].(int64) < 0 {

			uid := uuid.New()
			body.IngredientList[i].IngredientId = uid.ID()
			body.IngredientList[i].Name = strings.ToLower(body.IngredientList[i].Name)
			newIngredients = append(newIngredients, body.IngredientList[i])

		} else if err == nil && t[0].(int64) > 0 {

			ttemp := t[0].(int64)
			body.IngredientList[i].IngredientId = uint32(ttemp)
			body.IngredientList[i].Name = strings.ToLower(body.IngredientList[i].Name)

		} else if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error in loop": err.Error()})
			needToRollback = true
			return
		}
		temp.Close()
	}

	// Give every instruction uuid
	for i := range body.InstructionList {
		body.InstructionList[i].InstructionId = uuid.New().ID()
	}

	// Adding new ingredients to db
	for _, newIngredient := range newIngredients {
		_, err := Dbx.Exec(context.Background(),
			"INSERT INTO ingredients "+
				`(ingredient_id, "name") `+
				"VALUES($1, $2);",
			newIngredient.IngredientId, newIngredient.Name)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			needToRollback = true
			return
		}

	}

	// Adding recipe with it information
	_, err = Dbx.Exec(context.Background(),
		"INSERT INTO recipes "+
			`(recipe_id, "name", description, servings, prep_time, cook_time, image_url, source_url, author) `+
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);",
		recipeUuidInt, body.Name, body.Description, body.Servings, body.Prep_time, body.Cook_time, body.Image_url, body.Source_url, userId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error adding recipe": err.Error() + fmt.Sprint(userId)})
		needToRollback = true
		return
	}

	var counter = 0
	// Adding recipe steps to db
	for _, instruction := range body.InstructionList {
		counter++
		_, err := Dbx.Exec(context.Background(),
			"INSERT INTO instructions "+
				"(instruction_id, recipe_id, step_number, instruction_text, instruction_time) "+
				"VALUES($1, $2, $3, $4, $5);",
			instruction.InstructionId, recipeUuidInt, instruction.Step_number, instruction.InstructionText, instruction.InstructionTime)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error adding recipe step": err.Error() + fmt.Sprint(counter)})
			needToRollback = true
			return
		}
	}

	// Link this recipe with stuff

	// Link recipe with ingredients
	for _, ingredient := range body.IngredientList {
		_, err := Dbx.Exec(context.Background(),
			"INSERT INTO recipe_ingredients "+
				"(recipe_id, ingredient_id, quantity, unit) "+
				"VALUES($1, $2, $3, $4);",
			recipeUuidInt, ingredient.IngredientId, ingredient.Quantity, ingredient.Unit)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error() + fmt.Sprint(recipeUuidInt, ingredient.IngredientId)})
			needToRollback = true
			return
		}
	}

	// Link with cuisines
	for _, cuisine := range body.Cuisines {
		_, err := Dbx.Exec(context.Background(),
			"INSERT INTO recipe_cuisines "+
				"(recipe_id, cuisine_id) "+
				"VALUES($1, $2);",
			recipeUuidInt, cuisine.CuisineId)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			needToRollback = true
			return
		}
	}

	// Link with diets
	for _, diet := range body.Diets {
		_, err := Dbx.Exec(context.Background(),
			"INSERT INTO recipe_diets "+
				"(recipe_id, diet_id) "+
				"VALUES($1, $2);",
			recipeUuidInt, diet.DietId)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			needToRollback = true
			return
		}
	}
	//
	//Todo: Add transactions, so bad could not pass to db

	_, err = Dbx.Exec(context.Background(),
		"COMMIT;")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to COMMIT " + err.Error()})
		needToRollback = true
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "all good"})
}

func findRecipesByIngredients(c *gin.Context) {
	type content struct {
		Ingredients []string
	}

	if c.ContentType() != "application/json" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "content_type is not 'application/json'"})
		return
	}

	var body content

	if err := c.BindJSON(&body); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error 1": `Invalid body of the request`})
		return
	}

	recipe_ids, err := Dbx.Query(context.Background(),
		`select recipe_id from 
			recipe_ingredients ri 
			join ingredients i on i.ingredient_id = ri.ingredient_id
			where i.name in ('`+strings.Join(body.Ingredients[:], "','")+`')
			group by recipe_id
			having count(distinct name) = `+strconv.Itoa(len(body.Ingredients)))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}

	ids := []uint32{}

	for recipe_ids.Next() {
		values, err := recipe_ids.Values()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error 3": err.Error()})
			return
		}

		ids = append(ids, uint32(values[0].(int64)))
	}

	c.IndentedJSON(http.StatusOK, gin.H{"ids": ids})
}

func getRandomRecipes(c *gin.Context) {
	recipes, err := Dbx.Query(context.Background(),
		`select * from recipes`)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	recipeNumber := 15
	recipesToReturn := make([]map[string]interface{}, recipeNumber)

	recipeCounter := 0
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for recipes.Next() {
		b := r.Intn(10)

		if b >= 9 {
			temp := map[string]interface{}{}

			fields := recipes.FieldDescriptions()
			val, err := recipes.Values()

			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			for i, field := range fields {
				temp[field.Name] = val[i]
			}
			recipesToReturn[recipeCounter] = temp

			recipeCounter++

			if recipeCounter >= recipeNumber {
				break
			}
		}

	}

	recipes.Close()

	c.IndentedJSON(http.StatusOK, recipesToReturn)
}

func isAuth(c *gin.Context) {
	_, err := ValidateToken(c.GetHeader("Authorization"))

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "valid token"})
}

func addFavoriteRecipe(c *gin.Context) {

	usrID, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing required data in token"})
		return
	}

	recipeId, err := strconv.Atoi(strings.Trim(c.Param("recipeId"), "/"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing recipe id"})
		return
	}

	_, err = Dbx.Exec(context.Background(), `INSERT INTO user_favorite_recipe
	(user_id, recipe_id)
	VALUES($1, $2);`, usrID, recipeId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error/inserting": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "all good"})

}

func removeFavoriteRecipe(c *gin.Context) {

	usrID, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing required data in token"})
		return
	}

	recipeId, err := strconv.Atoi(strings.Trim(c.Param("recipeId"), "/"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing recipe id"})
		return
	}

	_, err = Dbx.Exec(context.Background(), `DELETE from user_favorite_recipe where 
	user_id = $1 and recipe_id = $2`, usrID, recipeId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error/inserting": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "all good"})

}

func getAllFavoriteRecipes(c *gin.Context) {
	usrID, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing required data in token"})
		return
	}

	recipeIdsRows, err := Dbx.Query(context.Background(), `select * from recipes r 
	join user_favorite_recipe ufr on r.recipe_id = ufr.recipe_id 
	where user_id = $1`, usrID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error/selecting": err.Error()})
		return
	}

	recipeIds := make([]any, 0)
	for recipeIdsRows.Next() {
		fieldNames := recipeIdsRows.FieldDescriptions()
		temp, err := recipeIdsRows.Values()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error/selecting": err.Error()})
			return
		}
		tempMap := map[string]any{}
		for i, fieldName := range fieldNames {
			tempMap[fieldName.Name] = temp[i]
		}
		recipeIds = append(recipeIds, tempMap)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"values": recipeIds})
}

func isFavorite(c *gin.Context) {
	usrID, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "missing required data in token"})
		return
	}

	recipeId, err := strconv.Atoi(strings.Trim(c.Param("recipeId"), "/"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, err := Dbx.Query(context.Background(), `SELECT user_id, recipe_id
	FROM user_favorite_recipe where user_id = $1 and recipe_id = $2`, usrID, recipeId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error/selecting": err.Error()})
		return
	}

	flag := false
	for rows.Next() {
		flag = true
	}

	c.IndentedJSON(http.StatusOK, flag)
}

func getRecipesByName(c *gin.Context) {
	recipeName := strings.Trim(c.Param("recipeName"), "/")

	rows, err := Dbx.Query(context.Background(), `SELECT *
	FROM recipes where "name" LIKE $1;`, fmt.Sprint(`%`, recipeName, `%`))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error in select (getRecipesByName)": err.Error()})
		return
	}

	recipesToReturn := []map[string]any{}

	for rows.Next() {
		fields := rows.FieldDescriptions()
		values, err := rows.Values()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		recipeInfo := map[string]any{}
		for i, field := range fields {
			recipeInfo[field.Name] = values[i]
		}

		recipesToReturn = append(recipesToReturn, recipeInfo)
	}

	c.IndentedJSON(http.StatusOK, recipesToReturn)
}

func deleteRecipe(c *gin.Context) {
	fmt.Println(c.Param("recipe_id"))
	recipe_id, err := strconv.Atoi(strings.Trim(c.Param("recipe_id"), "/"))
	if err != nil {
		c.IndentedJSON(404, gin.H{"err": err.Error()})
		return
	}

	userID, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(404, gin.H{"err": "no valid token provided"})
		return
	}

	var count int

	Dbx.QueryRow(context.Background(), `select count(*) from recipes where author = $1 and recipe_id = $2`, userID, recipe_id).Scan(&count)

	if count > 0 {
		_, err := Dbx.Exec(context.Background(), `DELETE FROM recipes
		WHERE recipe_id = $1`, recipe_id)

		if err != nil {
			c.IndentedJSON(404, gin.H{"err": err.Error()})
			return
		}
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Could not delete recipe. ID: " + strconv.Itoa(recipe_id)})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully deleted recipe. ID: " + strconv.Itoa(recipe_id)})
}

func getAllUserRecipes(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		c.IndentedJSON(404, gin.H{"err": "no valid token provided"})
		return
	}

	rows, err := Dbx.Query(context.Background(), `select * from recipes where author = $1`, userID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	recipesToReturn := []map[string]any{}

	for rows.Next() {
		fields := rows.FieldDescriptions()
		values, err := rows.Values()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		recipeInfo := map[string]any{}
		for i, field := range fields {
			recipeInfo[field.Name] = values[i]
		}

		recipesToReturn = append(recipesToReturn, recipeInfo)
	}

	c.IndentedJSON(http.StatusOK, recipesToReturn)

}

func SetupRotes(g *gin.Engine) error {
	g.Use(CORSMiddleware())

	freeGroup := g.Group("/")

	// Free group routes
	freeGroup.GET("recipe_info/:recipe_id", getRecipeInformation)
	freeGroup.GET("get_all_diets", getAllDiets)
	freeGroup.GET("get_all_cuisines", getAllCuisines)
	freeGroup.GET("get_all_ingredients", getAllIngredients)
	freeGroup.GET("get_recipes_by_ingredient/:ingredientId", getRecipesByIngredient)
	freeGroup.GET("get_recipes_by_cuisine/:cuisineId", getRecipesByCuisine)
	freeGroup.GET("get_recipes_by_name/:recipeName", getRecipesByName)
	freeGroup.GET("getRandomRecipes", getRandomRecipes)
	freeGroup.GET("isAuth", isAuth)
	freeGroup.POST("register_user", registerUser)
	freeGroup.POST("login", login)
	freeGroup.POST("get_recipes_by_ingredients", findRecipesByIngredients)

	authGroup := g.Group("/")
	authGroup.Use(AuthMiddleware())

	// Auth group routes
	authGroup.GET("test1", testRoute1)
	authGroup.GET("addToFavorite/:recipeId", addFavoriteRecipe)
	authGroup.GET("removeFromFavorite/:recipeId", removeFavoriteRecipe)
	authGroup.GET("getAllFavorite", getAllFavoriteRecipes)
	authGroup.GET("isFavorite/:recipeId", isFavorite)
	authGroup.GET("getAllUserRecipes", getAllUserRecipes)
	authGroup.POST("test2", testRoute2)
	authGroup.POST("add_recipe", addRecipe)
	authGroup.DELETE("deleteRecipe/:recipe_id", deleteRecipe)

	return nil
}

func login(c *gin.Context) {
	type content struct {
		Nickname string
		Password string
	}

	var body content

	err := c.BindJSON(&body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error 1st": err.Error()})
		return
	}

	hash, err := Dbx.Query(context.Background(),
		"select * from users where username = $1;",
		body.Nickname)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error 2nd": err.Error()})
		return
	}

	hash.Next()
	values, err := hash.Values()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user_id := strconv.FormatInt(values[0].(int64), 10)
	user_name := values[2].(string)
	user_hash := values[3].(string)
	hash.Close()

	user_hash_, salt, err := DecodeArgon2String(user_hash)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = Argon2Hasher.Compare(user_hash_, salt, []byte(body.Password))
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user_info := map[string]interface{}{
		"id":   user_id,
		"name": user_name,
	}

	token, err := CreateToken(user_info)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", token)
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Token, Uid, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		claims, err := ValidateToken(token)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error auth": err.Error()})
			c.Abort()
			return
		}

		// Claims -> data stored in token
		c.Set("id", claims["id"])
		c.Set("claims", claims)
		c.Next()

	}
}
