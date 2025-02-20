package polybft

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
)

func TestState_Insert_And_Get_ExitEvents_PerEpoch(t *testing.T) {
	const (
		numOfEpochs         = 11
		numOfBlocksPerEpoch = 10
		numOfEventsPerBlock = 11
	)

	state := newTestState(t)
	insertTestExitEvents(t, state, numOfEpochs, numOfBlocksPerEpoch, numOfEventsPerBlock)

	t.Run("Get events for existing epoch", func(t *testing.T) {
		events, err := state.CheckpointStore.getExitEventsByEpoch(1)

		assert.NoError(t, err)
		assert.Len(t, events, numOfBlocksPerEpoch*numOfEventsPerBlock)
	})

	t.Run("Get events for non-existing epoch", func(t *testing.T) {
		events, err := state.CheckpointStore.getExitEventsByEpoch(12)

		assert.NoError(t, err)
		assert.Len(t, events, 0)
	})
}

func TestState_Insert_And_Get_ExitEvents_ForProof(t *testing.T) {
	const (
		numOfEpochs         = 11
		numOfBlocksPerEpoch = 10
		numOfEventsPerBlock = 10
	)

	state := newTestState(t)
	insertTestExitEvents(t, state, numOfEpochs, numOfBlocksPerEpoch, numOfEventsPerBlock)

	var cases = []struct {
		epoch                  uint64
		checkpointBlockNumber  uint64
		expectedNumberOfEvents int
	}{
		{1, 1, 10},
		{1, 2, 20},
		{1, 8, 80},
		{2, 12, 20},
		{2, 14, 40},
		{3, 26, 60},
		{4, 38, 80},
		{11, 105, 50},
	}

	for _, c := range cases {
		events, err := state.CheckpointStore.getExitEventsForProof(c.epoch, c.checkpointBlockNumber)

		assert.NoError(t, err)
		assert.Len(t, events, c.expectedNumberOfEvents)
	}
}

func TestState_Insert_And_Get_ExitEvents_ForProof_NoEvents(t *testing.T) {
	t.Parallel()

	state := newTestState(t)
	insertTestExitEvents(t, state, 1, 10, 1)

	events, err := state.CheckpointStore.getExitEventsForProof(2, 11)

	assert.NoError(t, err)
	assert.Nil(t, events)
}

func TestState_decodeExitEvent(t *testing.T) {
	t.Parallel()

	const (
		exitID      = 1
		epoch       = 1
		blockNumber = 10
	)

	state := newTestState(t)

	topics := make([]ethgo.Hash, 4)
	topics[0] = ExitEventABI.ID()
	topics[1] = ethgo.BytesToHash([]byte{exitID})
	topics[2] = ethgo.BytesToHash(ethgo.HexToAddress("0x1111").Bytes())
	topics[3] = ethgo.BytesToHash(ethgo.HexToAddress("0x2222").Bytes())
	personType := abi.MustNewType("tuple(string firstName, string lastName)")
	encodedData, err := personType.Encode(map[string]string{"firstName": "John", "lastName": "Doe"})
	require.NoError(t, err)

	log := &ethgo.Log{
		Address: ethgo.ZeroAddress,
		Topics:  topics,
		Data:    encodedData,
	}

	event, err := decodeExitEvent(log, epoch, blockNumber)
	require.NoError(t, err)
	require.Equal(t, uint64(exitID), event.ID)
	require.Equal(t, uint64(epoch), event.EpochNumber)
	require.Equal(t, uint64(blockNumber), event.BlockNumber)

	require.NoError(t, state.CheckpointStore.insertExitEvents([]*ExitEvent{event}))
}

func TestState_decodeExitEvent_NotAnExitEvent(t *testing.T) {
	t.Parallel()

	topics := make([]ethgo.Hash, 4)
	topics[0] = stateTransferEventABI.ID()

	log := &ethgo.Log{
		Address: ethgo.ZeroAddress,
		Topics:  topics,
	}

	event, err := decodeExitEvent(log, 1, 1)
	require.NoError(t, err)
	require.Nil(t, event)
}

func insertTestExitEvents(t *testing.T, state *State,
	numOfEpochs, numOfBlocksPerEpoch, numOfEventsPerBlock int) []*ExitEvent {
	t.Helper()

	var (
		index      = uint64(0)
		block      = uint64(1)
		exitEvents = make([]*ExitEvent, numOfEpochs*numOfBlocksPerEpoch*numOfEventsPerBlock)
	)

	for i := uint64(1); i <= uint64(numOfEpochs); i++ {
		for j := 1; j <= numOfBlocksPerEpoch; j++ {
			for k := 1; k <= numOfEventsPerBlock; k++ {
				exitEvents[index] =
					&ExitEvent{
						ID:          index,
						Sender:      ethgo.ZeroAddress,
						Receiver:    ethgo.ZeroAddress,
						Data:        generateRandomBytes(t),
						EpochNumber: i,
						BlockNumber: block,
					}
				index++
			}
			block++
		}
	}
	require.NoError(t, state.CheckpointStore.insertExitEvents(exitEvents))

	return exitEvents
}
