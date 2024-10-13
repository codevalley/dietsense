package handlers

import (
	"dietsense/internal/repositories"
	"dietsense/internal/services"
	"dietsense/pkg/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnalyzeFood(factory *services.ServiceFactory, db repositories.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Stage 1: Determine input type
		fileHeader, _ := c.FormFile("image")
		inputText := c.PostForm("context")

		var inputType services.InputType
		var err error

		if fileHeader == nil {
			inputType = services.InputTypeText
		} else {
			// Use the image classifier service to determine input type
			classifierService, err := factory.GetImageClassifierService()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get image classifier service", "details": err.Error()})
				return
			}

			file, err := fileHeader.Open()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot open uploaded file"})
				return
			}
			defer file.Close()

			inputType, err = classifierService.ClassifyImage(file)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to classify image", "details": err.Error()})
				return
			}
		}

		// Stage 2: Analyze input based on type
		analyzerService, err := factory.GetAnalyzerService(inputType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get analyzer service", "details": err.Error()})
			return
		}

		var result *services.AnalysisResult
		context := fmt.Sprintf("%s\n%s", inputText, config.Config.ContextString)

		if inputType == services.InputTypeText {
			result, err = analyzerService.AnalyzeFoodText(context)
		} else {
			file, _ := fileHeader.Open()
			defer file.Close()
			result, err = analyzerService.AnalyzeFood(file, context, inputType)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to analyze", "details": err.Error()})
			return
		}

		// Stage 3: Compile and send response
		response := map[string]interface{}{
			"nutrition_info": result.NutritionInfo,
			"summary":        result.Summary,
			"confidence":     result.Confidence,
			"input_type":     result.InputType,
			"service":        result.Service,
		}

		c.JSON(http.StatusOK, response)
	}
}
