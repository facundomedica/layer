package flags

import (
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// List of CLI flags for Server and Client.
const (
	// Flag names
	FlagUnixSocketAddress           = "unix-socket-address"
	FlagPanicOnDaemonFailureEnabled = "panic-on-daemon-failure-enabled"
	FlagMaxDaemonUnhealthySeconds   = "max-daemon-unhealthy-seconds"

	FlagPriceDaemonEnabled     = "price-daemon-enabled"
	FlagPriceDaemonLoopDelayMs = "price-daemon-loop-delay-ms"
)

// Shared flags contains configuration flags shared by all daemons.
type SharedFlags struct {
	// SocketAddress is the location of the unix socket to communicate with the daemon gRPC service.
	SocketAddress string
	// PanicOnDaemonFailureEnabled toggles whether the daemon should panic on failure.
	PanicOnDaemonFailureEnabled bool
	// MaxDaemonUnhealthySeconds is the maximum allowable duration for which a daemon can be unhealthy.
	MaxDaemonUnhealthySeconds uint32
}

// PriceFlags contains configuration flags for the Price Daemon.
type PriceFlags struct {
	// Enabled toggles the price daemon on or off.
	Enabled bool
	// LoopDelayMs configures the update frequency of the price daemon.
	LoopDelayMs uint32
}

// DaemonFlags contains the collected configuration flags for all daemons.
type DaemonFlags struct {
	Shared SharedFlags
	Price  PriceFlags
}

var defaultDaemonFlags *DaemonFlags

// GetDefaultDaemonFlags returns the default values for the Daemon Flags using a singleton pattern.
func GetDefaultDaemonFlags() DaemonFlags {
	if defaultDaemonFlags == nil {
		defaultDaemonFlags = &DaemonFlags{
			Shared: SharedFlags{
				SocketAddress:               "/tmp/daemons.sock",
				PanicOnDaemonFailureEnabled: true,
				MaxDaemonUnhealthySeconds:   5 * 60, // 5 minutes.
			},
			Price: PriceFlags{
				Enabled:     true,
				LoopDelayMs: 3_000,
			},
		}
	}
	return *defaultDaemonFlags
}

// AddDaemonFlagsToCmd adds the required flags to instantiate a server and client for
// price updates. These flags should be applied to the `start` command LAYER Cosmos application.
// E.g. `layerd start --price-daemon-enabled=true --unix-socket-address $(unix_socket_address)`
func AddDaemonFlagsToCmd(
	cmd *cobra.Command,
) {
	//
	df := GetDefaultDaemonFlags()

	// Shared Flags.
	cmd.Flags().String(
		FlagUnixSocketAddress,
		df.Shared.SocketAddress,
		"Socket address for the daemons to send updates to, if not set "+
			"will establish default location to ingest daemon updates from",
	)
	cmd.Flags().Bool(
		FlagPanicOnDaemonFailureEnabled,
		df.Shared.PanicOnDaemonFailureEnabled,
		"Enables panicking when a daemon fails.",
	)
	cmd.Flags().Uint32(
		FlagMaxDaemonUnhealthySeconds,
		df.Shared.MaxDaemonUnhealthySeconds,
		"Maximum allowable duration for which a daemon can be unhealthy.",
	)

	// Price Daemon.
	cmd.Flags().Bool(
		FlagPriceDaemonEnabled,
		df.Price.Enabled,
		"Enable Price Daemon. Set to false for non-validator nodes.",
	)
	cmd.Flags().Uint32(
		FlagPriceDaemonLoopDelayMs,
		df.Price.LoopDelayMs,
		"Delay in milliseconds between sending price updates to the application.",
	)
}

// GetDaemonFlagValuesFromOptions gets all daemon flag values from the `AppOptions` struct.
func GetDaemonFlagValuesFromOptions(
	appOpts servertypes.AppOptions,
) DaemonFlags {
	// Default value
	result := GetDefaultDaemonFlags()

	// Shared Flags
	if option := appOpts.Get(FlagUnixSocketAddress); option != nil {
		if v, err := cast.ToStringE(option); err == nil {
			result.Shared.SocketAddress = v
		}
	}
	if option := appOpts.Get(FlagPanicOnDaemonFailureEnabled); option != nil {
		if v, err := cast.ToBoolE(option); err == nil {
			result.Shared.PanicOnDaemonFailureEnabled = v
		}
	}
	if option := appOpts.Get(FlagMaxDaemonUnhealthySeconds); option != nil {
		if v, err := cast.ToUint32E(option); err == nil {
			result.Shared.MaxDaemonUnhealthySeconds = v
		}
	}

	// Price Daemon.
	if option := appOpts.Get(FlagPriceDaemonEnabled); option != nil {
		if v, err := cast.ToBoolE(option); err == nil {
			result.Price.Enabled = v
		}
	}
	if option := appOpts.Get(FlagPriceDaemonLoopDelayMs); option != nil {
		if v, err := cast.ToUint32E(option); err == nil {
			result.Price.LoopDelayMs = v
		}
	}

	return result
}
