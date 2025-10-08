package middleware

import (
	"backendgo/utils"
	"strings"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Middleware untuk verifikasi JWT
func AuthRequired() fiber.Handler {
    return func(c *fiber.Ctx) error {
		log.Println("AuthRequired dijalankan")
        tokenString := c.Get("Authorization")
        if tokenString == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
        }

        // Hapus prefix Bearer jika ada
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")


        claims, err := utils.ValidateToken(tokenString)
        if err != nil {
            log.Println("Token invalid:", err)
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
        }

        log.Printf("Claims: %+v\n", claims)


        // Ambil nilai dari MapClaims
        userID := int(claims["user_id"].(float64))
        username := claims["username"].(string)
        role := claims["role"].(string)

        // Simpan ke context
        c.Locals("user_id", userID)
        c.Locals("username", username)
        c.Locals("role", role)

        return c.Next()
    }
}


// Middleware tambahan untuk cek role admin
func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role")
		log.Println("AdminOnly dijalankan, role:", role) 
		if role == nil || role.(string) != "admin" {
			return c.Status(403).JSON(fiber.Map{"error": "Akses ditolak, hanya admin yang diizinkan"})
		}
		return c.Next()
	}
}
