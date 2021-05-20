package controllor

import "github.com/gin-gonic/gin"

func AdminInfo(c *gin.Context) (interface{}, error) {
	return map[string]interface{}{
		"name":         "admin",
		"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"roles":        "",
		"introduction": "",
	}, nil
}
