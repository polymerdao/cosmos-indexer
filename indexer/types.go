package indexer

import (
	"github.com/DefiantLabs/cosmos-indexer/config"
	"github.com/DefiantLabs/cosmos-indexer/core"
	dbTypes "github.com/DefiantLabs/cosmos-indexer/db"
	"github.com/DefiantLabs/cosmos-indexer/db/models"
	"github.com/DefiantLabs/cosmos-indexer/filter"
	"github.com/DefiantLabs/cosmos-indexer/parsers"
	"github.com/DefiantLabs/probe/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	"gorm.io/gorm"
)

type Indexer struct {
	Config                              *config.IndexConfig
	DryRun                              bool
	DB                                  *gorm.DB
	ChainClient                         *client.ChainClient
	BlockEnqueueFunction                func(chan *core.EnqueueData) error
	CustomModuleBasics                  []module.AppModuleBasic // Used for extending the AppModuleBasics registered in the probe ChainClientient
	BlockEventFilterRegistries          BlockEventFilterRegistries
	MessageTypeFilters                  []filter.MessageTypeFilter
	CustomBeginBlockEventParserRegistry map[string][]parsers.BlockEventParser // Used for associating parsers to block event types in BeginBlock events
	CustomEndBlockEventParserRegistry   map[string][]parsers.BlockEventParser // Used for associating parsers to block event types in EndBlock events
	CustomBeginBlockParserTrackers      map[string]models.BlockEventParser    // Used for tracking block event parsers in the database
	CustomEndBlockParserTrackers        map[string]models.BlockEventParser    // Used for tracking block event parsers in the database
	CustomMessageParserRegistry         map[string][]parsers.MessageParser    // Used for associating parsers to message types
	CustomMessageParserTrackers         map[string]models.MessageParser       // Used for tracking message parsers in the database
	CustomModels                        []any
}

type BlockEventFilterRegistries struct {
	BeginBlockEventFilterRegistry *filter.StaticBlockEventFilterRegistry
	EndBlockEventFilterRegistry   *filter.StaticBlockEventFilterRegistry
}

type DBData struct {
	txDBWrappers []dbTypes.TxDBWrapper
	block        models.Block
}

type BlockEventsDBData struct {
	blockDBWrapper *dbTypes.BlockDBWrapper
}
