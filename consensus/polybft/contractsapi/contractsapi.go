// Code generated by scapi/gen. DO NOT EDIT.
package contractsapi

import (
	"math/big"

	"github.com/0xPolygon/polygon-edge/consensus/polybft/contractsapi/artifact"
	"github.com/0xPolygon/polygon-edge/types"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
)

var (
	StateReceiverContract     = &StateReceiverContractImpl{Artifact: StateReceiver}
	ChildValidatorSetContract = &ChildValidatorSetContractImpl{Artifact: ChildValidatorSet}
	StateSenderContract       = &StateSenderContractImpl{Artifact: StateSender}
	CheckpointManagerContract = &CheckpointManagerContractImpl{Artifact: CheckpointManager}
)

type StateReceiverContractImpl struct {
	Artifact *artifact.Artifact

	Commit  Commit
	Execute Execute
}

var CommitmentABIType = abi.MustNewType("tuple(uint256 startId,uint256 endId,bytes32 root)")

type Commitment struct {
	StartID *big.Int   `abi:"startId"`
	EndID   *big.Int   `abi:"endId"`
	Root    types.Hash `abi:"root"`
}

func (c *Commitment) EncodeAbi() ([]byte, error) {
	return CommitmentABIType.Encode(c)
}

func (c *Commitment) DecodeAbi(buf []byte) error {
	return decodeStruct(CommitmentABIType, buf, &c)
}

type Commit struct {
	Commitment *Commitment `abi:"commitment"`
	Signature  []byte      `abi:"signature"`
	Bitmap     []byte      `abi:"bitmap"`
}

func (c *Commit) EncodeAbi() ([]byte, error) {
	return StateReceiver.Abi.Methods["commit"].Encode(c)
}

func (c *Commit) DecodeAbi(buf []byte) error {
	return decodeMethod(StateReceiver.Abi.Methods["commit"], buf, c)
}

var ObjABIType = abi.MustNewType("tuple(uint256 id,address sender,address receiver,bytes data)")

type Obj struct {
	ID       *big.Int      `abi:"id"`
	Sender   types.Address `abi:"sender"`
	Receiver types.Address `abi:"receiver"`
	Data     []byte        `abi:"data"`
}

func (o *Obj) EncodeAbi() ([]byte, error) {
	return ObjABIType.Encode(o)
}

func (o *Obj) DecodeAbi(buf []byte) error {
	return decodeStruct(ObjABIType, buf, &o)
}

type Execute struct {
	Proof []types.Hash `abi:"proof"`
	Obj   *Obj         `abi:"obj"`
}

func (e *Execute) EncodeAbi() ([]byte, error) {
	return StateReceiver.Abi.Methods["execute"].Encode(e)
}

func (e *Execute) DecodeAbi(buf []byte) error {
	return decodeMethod(StateReceiver.Abi.Methods["execute"], buf, e)
}

var (
	StateSyncResultEventType = abi.MustNewEvent("event StateSyncResult(uint256 indexed counter,bool indexed status,bytes message)") //nolint:all
)

type StateSyncResultEvent struct {
	Counter *big.Int `abi:"counter"`
	Status  bool     `abi:"status"`
	Message []byte   `abi:"message"`
}

func (S *StateSyncResultEvent) ParseLog(log *ethgo.Log) error {
	return decodeEvent(StateSyncResultEventType, log, S)
}

var (
	NewCommitmentEventType = abi.MustNewEvent("event NewCommitment(uint256 indexed startId,uint256 indexed endId,bytes32 root)") //nolint:all
)

type NewCommitmentEvent struct {
	StartID *big.Int   `abi:"startId"`
	EndID   *big.Int   `abi:"endId"`
	Root    types.Hash `abi:"root"`
}

func (N *NewCommitmentEvent) ParseLog(log *ethgo.Log) error {
	return decodeEvent(NewCommitmentEventType, log, N)
}

type ChildValidatorSetContractImpl struct {
	Artifact *artifact.Artifact

	CommitEpoch CommitEpoch
}

var EpochABIType = abi.MustNewType("tuple(uint256 startBlock,uint256 endBlock,bytes32 epochRoot)")

type Epoch struct {
	StartBlock *big.Int   `abi:"startBlock"`
	EndBlock   *big.Int   `abi:"endBlock"`
	EpochRoot  types.Hash `abi:"epochRoot"`
}

func (e *Epoch) EncodeAbi() ([]byte, error) {
	return EpochABIType.Encode(e)
}

func (e *Epoch) DecodeAbi(buf []byte) error {
	return decodeStruct(EpochABIType, buf, &e)
}

var UptimeABIType = abi.MustNewType("tuple(uint256 epochId,tuple(address validator,uint256 signedBlocks)[] uptimeData,uint256 totalBlocks)")

var UptimeDataABIType = abi.MustNewType("tuple(address validator,uint256 signedBlocks)")

type UptimeData struct {
	Validator    types.Address `abi:"validator"`
	SignedBlocks *big.Int      `abi:"signedBlocks"`
}

func (u *UptimeData) EncodeAbi() ([]byte, error) {
	return UptimeDataABIType.Encode(u)
}

func (u *UptimeData) DecodeAbi(buf []byte) error {
	return decodeStruct(UptimeDataABIType, buf, &u)
}

type Uptime struct {
	EpochID     *big.Int      `abi:"epochId"`
	UptimeData  []*UptimeData `abi:"uptimeData"`
	TotalBlocks *big.Int      `abi:"totalBlocks"`
}

func (u *Uptime) EncodeAbi() ([]byte, error) {
	return UptimeABIType.Encode(u)
}

func (u *Uptime) DecodeAbi(buf []byte) error {
	return decodeStruct(UptimeABIType, buf, &u)
}

type CommitEpoch struct {
	ID     *big.Int `abi:"id"`
	Epoch  *Epoch   `abi:"epoch"`
	Uptime *Uptime  `abi:"uptime"`
}

func (c *CommitEpoch) EncodeAbi() ([]byte, error) {
	return ChildValidatorSet.Abi.Methods["commitEpoch"].Encode(c)
}

func (c *CommitEpoch) DecodeAbi(buf []byte) error {
	return decodeMethod(ChildValidatorSet.Abi.Methods["commitEpoch"], buf, c)
}

type StateSenderContractImpl struct {
	Artifact *artifact.Artifact

	SyncState SyncState
}

type SyncState struct {
	Receiver types.Address `abi:"receiver"`
	Data     []byte        `abi:"data"`
}

func (s *SyncState) EncodeAbi() ([]byte, error) {
	return StateSender.Abi.Methods["syncState"].Encode(s)
}

func (s *SyncState) DecodeAbi(buf []byte) error {
	return decodeMethod(StateSender.Abi.Methods["syncState"], buf, s)
}

var (
	StateSyncedEventType = abi.MustNewEvent("event StateSynced(uint256 indexed id,address indexed sender,address indexed receiver,bytes data)") //nolint:all
)

type StateSyncedEvent struct {
	ID       *big.Int      `abi:"id"`
	Sender   types.Address `abi:"sender"`
	Receiver types.Address `abi:"receiver"`
	Data     []byte        `abi:"data"`
}

func (S *StateSyncedEvent) ParseLog(log *ethgo.Log) error {
	return decodeEvent(StateSyncedEventType, log, S)
}

type CheckpointManagerContractImpl struct {
	Artifact *artifact.Artifact

	Submit Submit
}

var CheckpointMetadataABIType = abi.MustNewType("tuple(bytes32 blockHash,uint256 blockRound,bytes32 currentValidatorSetHash)")

type CheckpointMetadata struct {
	BlockHash               types.Hash `abi:"blockHash"`
	BlockRound              *big.Int   `abi:"blockRound"`
	CurrentValidatorSetHash types.Hash `abi:"currentValidatorSetHash"`
}

func (c *CheckpointMetadata) EncodeAbi() ([]byte, error) {
	return CheckpointMetadataABIType.Encode(c)
}

func (c *CheckpointMetadata) DecodeAbi(buf []byte) error {
	return decodeStruct(CheckpointMetadataABIType, buf, &c)
}

var CheckpointABIType = abi.MustNewType("tuple(uint256 epoch,uint256 blockNumber,bytes32 eventRoot)")

type Checkpoint struct {
	Epoch       *big.Int   `abi:"epoch"`
	BlockNumber *big.Int   `abi:"blockNumber"`
	EventRoot   types.Hash `abi:"eventRoot"`
}

func (c *Checkpoint) EncodeAbi() ([]byte, error) {
	return CheckpointABIType.Encode(c)
}

func (c *Checkpoint) DecodeAbi(buf []byte) error {
	return decodeStruct(CheckpointABIType, buf, &c)
}

var NewValidatorSetABIType = abi.MustNewType("tuple(address _address,uint256[4] blsKey,uint256 votingPower)")

type NewValidatorSet struct {
	_address    types.Address `abi:"_address"`
	BlsKey      [4]*big.Int   `abi:"blsKey"`
	VotingPower *big.Int      `abi:"votingPower"`
}

func (n *NewValidatorSet) EncodeAbi() ([]byte, error) {
	return NewValidatorSetABIType.Encode(n)
}

func (n *NewValidatorSet) DecodeAbi(buf []byte) error {
	return decodeStruct(NewValidatorSetABIType, buf, &n)
}

type Submit struct {
	ChainID            *big.Int            `abi:"chainId"`
	CheckpointMetadata *CheckpointMetadata `abi:"checkpointMetadata"`
	Checkpoint         *Checkpoint         `abi:"checkpoint"`
	Signature          [2]*big.Int         `abi:"signature"`
	NewValidatorSet    []*NewValidatorSet  `abi:"newValidatorSet"`
	Bitmap             []byte              `abi:"bitmap"`
}

func (s *Submit) EncodeAbi() ([]byte, error) {
	return CheckpointManager.Abi.Methods["submit"].Encode(s)
}

func (s *Submit) DecodeAbi(buf []byte) error {
	return decodeMethod(CheckpointManager.Abi.Methods["submit"], buf, s)
}
