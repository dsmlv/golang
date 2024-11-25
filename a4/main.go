package main

// Dosmailova Dinara

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"expvar"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Dosmailova Dinara
// Define a simple user struct
type User struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var logger = logrus.New()
var requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_request_duration_seconds",
	Help:    "Duration of HTTP requests in seconds.",
	Buckets: prometheus.DefBuckets,
}, []string{"path", "method"})
var validate = validator.New()
var requestCount = expvar.NewInt("request_count")

// Dosmailova Dinara
func init() {
	// Initialize Prometheus metrics
	prometheus.MustRegister(requestDuration)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "./log/app.log",
		MaxSize:    10,   // Max megabytes before log is rotated
		MaxBackups: 3,    // Max number of old log files to keep
		MaxAge:     28,   // Max number of days to retain old log files
		Compress:   true, // Compress the old log files
	})
}

// Dosmailova Dinara
func main() {
	// Set up router
	r := mux.NewRouter()

	// Set up CSRF protection
	csrfProtection := csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(true))

	// Public routes
	r.HandleFunc("/signup", SignupHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(JWTMiddleware)
	protected.HandleFunc("/profile", UserProfileHandler).Methods("GET")
	protected.HandleFunc("/admin", AdminHandler).Methods("GET")
	protected.Use(RoleMiddleware("admin"))

	// Dosmailova Dinara
	// Middleware for security headers
	r.Use(SecurityHeadersMiddleware)
	// Middleware for request logging
	r.Use(RequestLoggingMiddleware)
	// Middleware for monitoring
	r.Use(MetricsMiddleware)

	// Metrics endpoint
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")

	// Dosmailova Dinara
	// Start server with TLS
	httpsServer := &http.Server{
		Addr:      ":8443",
		Handler:   csrfProtection(r),
		TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
	}
	logger.Info("Starting server on :8443")
	httpsServer.ListenAndServeTLS("server.crt", "server.key")
}

// Dosmailova Dinara
// SignupHandler handles user registration
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate user input
	if err := validate.Struct(user); err != nil {
		handleError(w, http.StatusBadRequest, "Validation failed: "+err.Error())
		return
	}

	// Dosmailova Dinara
	// Hash the user's password
	hash := sha256.New()
	hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))

	logger.Info("User registered: ", user.Username)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Dosmailova Dinara
// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handleError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate user input
	if err := validate.Struct(user); err != nil {
		handleError(w, http.StatusBadRequest, "Validation failed: "+err.Error())
		return
	}

	// Dosmailova Dinara
	// Users credentials
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// UserProfileHandler handles user profile requests
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for user profile
}

// Dosmailova Dinara
// AdminHandler handles admin requests
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for admin actions
}

// Dosmailova Dinara
// JWTMiddleware checks for a valid JWT token
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			handleError(w, http.StatusUnauthorized, "Missing token")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			handleError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Dosmailova Dinara
// RoleMiddleware restricts access based on user roles
func RoleMiddleware(role string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Header.Get("UserRole")

			if userRole != role {
				handleError(w, http.StatusForbidden, "Access denied for user role")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Dosmailova Dinara
// SecurityHeadersMiddleware adds security headers to each response
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")

		next.ServeHTTP(w, r)
	})
}

// Dosmailova Dinara
// RequestLoggingMiddleware logs each incoming HTTP request
func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		logger.WithFields(logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"duration": duration,
			"status":   w.Header().Get("status"),
		}).Info("Handled request")
	})
}

// Dosmailova Dinara
// MetricsMiddleware tracks metrics for each HTTP request
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(r.URL.Path, r.Method).Observe(duration)
		requestCount.Add(1)
	})
}

// handleError handles errors consistently and logs them
func handleError(w http.ResponseWriter, status int, message string) {
	logger.Error(message)
	http.Error(w, http.StatusText(status), status)
}

// Dosmailova Dinara
// Prometheus alerting rules
// Add the following alerting rule to your Prometheus server configuration:
//
// groups:
// - name: alert.rules
//   rules:
//   - alert: HighErrorRate
//     expr: rate(http_request_duration_seconds_sum[1m]) > 0.05
//     for: 2m
//     labels:
//       severity: warning
//     annotations:
//       summary: "High error rate detected"
//       description: "High error rate for {{ $labels.path }} on {{ $labels.instance }}"
