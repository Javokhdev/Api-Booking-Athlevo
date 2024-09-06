package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config struct holds the configuration settings.
type Config struct {
	HTTPPort       string
	BookingSvcAddr string
	// Kafka Configuration
	KafkaBrokers                   []string
	KafkaBrokersTest               []string
	KafkaMedicalRecordTopic        string
	KafkaGeneticDataTopic          string
	KafkaLifestyleDataTopic        string
	KafkaWearableDataTopic         string
	KafkaHealthRecommendationTopic string

	// JWT
	JWTSecretKey string
	JWTExpiry    int

	LOG_PATH        string
	TimelineSvcAddr string
	MemorySvcAddr   string
}

// Load loads the configuration from environment variables.
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTPPort = cast.ToString(coalesce("HTTP_PORT", ":8081"))
	config.BookingSvcAddr = cast.ToString(coalesce("BOOKING_PORT", "booking:8082"))

	config.KafkaBrokers = cast.ToStringSlice(coalesce("KAFKA_BROKERS", []string{"localhost:9092"}))
	config.KafkaBrokersTest = cast.ToStringSlice(coalesce("KAFKA_BROKERS_Test", []string{"localhost:9092"}))

	// Kafka Topics
	config.KafkaMedicalRecordTopic = cast.ToString(coalesce("KAFKA_MEDICAL_RECORD_TOPIC", "medical_record_topic"))
	config.KafkaGeneticDataTopic = cast.ToString(coalesce("KAFKA_GENETIC_DATA_TOPIC", "genetic_data_topic"))
	config.KafkaLifestyleDataTopic = cast.ToString(coalesce("KAFKA_LIFESTYLE_DATA_TOPIC", "lifestyle_data_topic"))
	config.KafkaWearableDataTopic = cast.ToString(coalesce("KAFKA_WEARABLE_DATA_TOPIC", "wearable_data_topic"))
	config.KafkaHealthRecommendationTopic = cast.ToString(coalesce("KAFKA_HEALTH_RECOMMENDATION_TOPIC", "health_recommendation_topic"))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	// JWT Configuration
	config.JWTSecretKey = cast.ToString(coalesce("JWT_SECRET_KEY", "your_secret_key"))
	config.JWTExpiry = cast.ToInt(coalesce("JWT_EXPIRY", 60))

	config.TimelineSvcAddr = cast.ToString(coalesce("TIME_LINE_SERVICE_port", "timeline:9091"))
	config.MemorySvcAddr = cast.ToString(coalesce("MEMORY_SERVICE_port", "memory:9090"))
	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
