package controllor

import "github.com/gin-gonic/gin"

func AdminInfo(c *gin.Context) (interface{}, error) {
	return map[string]interface{}{
		"name":         "admin",
		"avatar":       "",
		"roles":        "",
		"introduction": "",
	}, nil
}
