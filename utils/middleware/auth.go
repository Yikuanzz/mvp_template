package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"mvp/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware 认证中间件
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "请求头缺少 Authorization")
			c.Abort()
			return
		}

		// Authorization 格式必须是 Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "Authorization 格式错误，应为 'Bearer <token>'")
			c.Abort()
			return
		}

		// 解析并验证 token
		tokenStr := parts[1]
		claims, err := ParseToken(tokenStr, secretKey)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Token 无效或已过期")
			c.Abort()
			return
		}

		// 将用户ID写入上下文
		userID := claims.UserID
		c.Set("userID", userID)
		// 将角色信息写入上下文

		roleIDsStr := claims.RoleIDs
		roleIDs := strings.Split(roleIDsStr, ",")
		var roleIDsUint []uint
		for _, roleIDStr := range roleIDs {
			roleID, err := strconv.ParseUint(roleIDStr, 10, 32)
			if err != nil {
				response.Error(c, http.StatusInternalServerError, "角色ID转换失败")
				c.Abort()
				return
			}
			roleIDsUint = append(roleIDsUint, uint(roleID))
		}
		c.Set("roleIDs", roleIDsUint)

		// 继续处理后续请求
		c.Next()
	}
}

// Claims 认证信息
type Claims struct {
	UserID  uint   `json:"user_id"`  // 用户ID
	RoleIDs string `json:"role_ids"` // 角色ID列表，逗号分隔
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token，带上角色信息
func GenerateToken(userID uint, roleIDs string, secretKey string, expires int) (string, error) {
	claims := Claims{
		UserID:  userID,
		RoleIDs: roleIDs,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expires) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(secretKey)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		// fmt.Println("Error signing token:", err)
		return "", err
	}
	// fmt.Println("Generated Token:", signedToken) // 打印生成的 token
	return signedToken, nil
}

// ParseToken 解析JWT token，强制验证算法一致并检查是否有效
func ParseToken(tokenString string, secretKey string) (*Claims, error) {
	// fmt.Println("Token to parse:", tokenString)                // 打印传入的 token
	// fmt.Println("Secret Key (Parse):", cfg.AuthInfo.SecretKey) // 打印密钥
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 强制验证签名算法一致
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// 返回秘钥
		return []byte(secretKey), nil
	})
	// 如果解析失败，返回错误
	if err != nil {
		fmt.Println("Error parsing token:", err) // 打印解析错误
		return nil, err
	}

	// 如果 Token 是有效的，进一步检查 Token 是否有效（例如是否过期或未生效）
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// 检查过期时间
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, errors.New("token expired")
		}

		// 检查生效时间
		if claims.NotBefore != nil && claims.NotBefore.Time.After(time.Now()) {
			return nil, errors.New("token not active yet")
		}

		return claims, nil
	}

	// 如果 Token 无效，返回错误
	return nil, errors.New("invalid token")
}
