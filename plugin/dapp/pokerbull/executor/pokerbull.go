package executor

import (
	"fmt"

	log "github.com/inconshreveable/log15"
	pkt "gitlab.33.cn/chain33/chain33/plugin/dapp/pokerbull/types"
	drivers "gitlab.33.cn/chain33/chain33/system/dapp"
	"gitlab.33.cn/chain33/chain33/types"
)

var logger = log.New("module", "execs.pokerbull")

func Init(name string, sub []byte) {
	drivers.Register(newPBGame().GetName(), newPBGame, 0)
}

var driverName = pkt.PokerBullX

func init() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&PokerBull{}))
}

type PokerBull struct {
	drivers.DriverBase
}

func newPBGame() drivers.Driver {
	t := &PokerBull{}
	t.SetChild(t)
	t.SetExecutorType(types.LoadExecutorType(driverName))
	return t
}

func GetName() string {
	return newPBGame().GetName()
}

func (g *PokerBull) GetDriverName() string {
	return pkt.PokerBullX
}

func calcPBGameStatusKey(status int32, player int32, index int64) []byte {
	key := fmt.Sprintf("PBgame-status:%d:%d:%d", status, player, index)
	return []byte(key)
}

func calcPBGameStatusAndPlayerPrefix(status int32, player int32) []byte {
	key := fmt.Sprintf("PBgame-status:%d:%d:", status, player)
	return []byte(key)
}

func addPBGameStatus(status int32, player int32, index int64, gameId string) *types.KeyValue {
	kv := &types.KeyValue{}
	kv.Key = calcPBGameStatusKey(status, player, index)
	record := &pkt.PBGameRecord{
		GameId: gameId,
	}
	kv.Value = types.Encode(record)
	return kv
}

func delPBGameStatus(status int32, player int32, index int64) *types.KeyValue {
	kv := &types.KeyValue{}
	kv.Key = calcPBGameStatusKey(status, player, index)
	kv.Value = nil
	return kv
}