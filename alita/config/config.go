package config

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	AllowedUpdates     []string
	ValidLangCodes     []string
	BotToken           string
	DatabaseURI        string
	MainDbName         string
	BotVersion         string = "2.1.3"
	ApiServer          string
	WorkingMode        = "worker"
	Debug              = false
	DropPendingUpdates = true
	OwnerId            int64
	MessageDump        int64
	RedisAddress       string
	RedisPassword      string
	RedisDB            int
)

// init initializes the config variables.
func init() {
	// set logger config
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(
		&log.JSONFormatter{
			DisableHTMLEscape: true,
			PrettyPrint:       true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				return f.Function, fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
			},
		},
	)

	// load goenv config
	godotenv.Load()

	// set necessary variables
	Debug = typeConvertor{str: os.Getenv("DEBUG")}.Bool()
	DropPendingUpdates = typeConvertor{str: os.Getenv("DROP_PENDING_UPDATES")}.Bool()
	DatabaseURI = os.Getenv("DB_URI")
	MainDbName = os.Getenv("DB_NAME")
	OwnerId = typeConvertor{str: os.Getenv("OWNER_ID")}.Int64()
	MessageDump = typeConvertor{str: os.Getenv("MESSAGE_DUMP")}.Int64()
	BotToken = os.Getenv("BOT_TOKEN")

	AllowedUpdates = typeConvertor{str: os.Getenv("ALLOWED_UPDATES")}.StringArray()
	// if allowed updates is not set, set it to receive all updates
	if (len(AllowedUpdates) == 1 && AllowedUpdates[0] == "") || (len(AllowedUpdates) == 0) {
		AllowedUpdates = []string{
			"message",
			"edited_message",
			"channel_post",
			"edited_channel_post",
			"inline_query",
			"chosen_inline_result",
			"callback_query",
			"shipping_query",
			"pre_checkout_query",
			"poll",
			"poll_answer",
			"my_chat_member",
			"chat_member",
			"chat_join_request",
		}
	}

	ValidLangCodes = typeConvertor{str: os.Getenv("ENABLED_LOCALES")}.StringArray()
	// if valid lang codes is not set, set it to 'en' only
	if (len(ValidLangCodes) == 1 && ValidLangCodes[0] == "") || (len(ValidLangCodes) == 0) {
		ValidLangCodes = []string{"en"}
	}

	ApiServer = os.Getenv("API_SERVER")
	// set as default api server if not set
	if ApiServer == "" {
		ApiServer = "https://api.telegram.org"
	}
	// set default db_name
	if MainDbName == "" {
		MainDbName = "Alita_Robot"
	}

	// redis config
	RedisAddress = os.Getenv("REDIS_ADDRESS")
	if os.Getenv("REDIS_ADDRESS") == "" {
		RedisAddress = "localhost:6379"
	}
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	if os.Getenv("REDIS_PASSWORD") == "" {
		RedisPassword = ""
	}
	RedisDB = typeConvertor{str: os.Getenv("REDIS_DB")}.Int()
	if os.Getenv("REDIS_DB") == "" {
		RedisDB = 0
	}
}
