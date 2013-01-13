package messages

import (
	"encoding/hex"
)

var (
	KnownMessageNames = map[byte]string{
		ASSIGN_CHANNEL_MSG_ID:              ASSIGN_CHANNEL_MSG_NAME,
		CAPABILITIES_MSG_ID:                CapabilitiesCommandName,
		OPEN_CHANNEL_MSG_ID:                OpenChannelCommandName,
		RECV_BROADCAST_DATA_MSG_ID:         RecvBroadcastDataCommandName,
		REQUEST_MESSAGE_MSG_ID:             RequestMessageCommandName,
		RESET_SYSTEM_MSG_ID:                ResetSystemCommandName,
		SEND_BROADCAST_DATA_MSG_ID:         SendBroadcastDataCommandName,
		SET_CHANNEL_ID_MSG_ID:              SetChannelIdCommandName,
		SET_CHANNEL_PERIOD_MSG_ID:          SetChannelPeriodCommandName,
		SET_CHANNEL_RF_FREQ_MSG_ID:         SetChannelRfFrequentyCommandName,
		SET_CHANNEL_SEARCH_TIMEOUT_MSG_ID:  SetChannelSearchTimeoutCommandName,
		SET_CHANNEL_SEARCH_WAVEFORM_MSG_ID: SetChannelSearchWaveformCommandName,
		SET_NETWORK_KEY_MSG_ID:             SetNetworkKeyCommandName,
	}
)

func CommandIdToName(id byte) string {
	name, ok := KnownMessageNames[id]

	if ok {
		return name
	}

	input := []byte{id}
	return "UNKNOWN_" + hex.EncodeToString(input)
}
