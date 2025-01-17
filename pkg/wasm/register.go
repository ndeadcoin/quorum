//go:build js && wasm
// +build js,wasm

package wasm

import (
	"encoding/json"
	"strings"
	"syscall/js"

	"github.com/rumsystem/quorum/pkg/wasm/api"
	quorumAPI "github.com/rumsystem/quorum/pkg/wasm/api"
	"github.com/rumsystem/quorum/pkg/wasm/logger"
)

// quit channel
var qChan chan struct{} = nil

func RegisterJSFunctions() {
	js.Global().Set("SetDebug", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		enableDebug := args[0].Bool()
		logger.SetDebug(enableDebug)
		return true
	}))

	js.Global().Set("StartQuorum", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if qChan == nil {
			qChan = make(chan struct{}, 0)
		}
		if len(args) < 2 {
			return nil
		}
		password := args[0].String()
		bootAddrsStr := args[1].String()
		bootAddrs := strings.Split(bootAddrsStr, ",")

		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			ok, err := StartQuorum(qChan, password, bootAddrs)
			ret["ok"] = ok
			if err != nil {
				return ret, err
			}
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("IsQuorumRunning", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ret := qChan != nil
		return js.ValueOf(ret).Bool()
	}))

	js.Global().Set("StopQuorum", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if qChan != nil {
			close(qChan)
			qChan = nil
		}
		return js.ValueOf(true).Bool()
	}))

	js.Global().Set("StartSync", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		groupId := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.StartSync(groupId)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("Announce", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsonStr := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.Announce([]byte(jsonStr))
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetGroupProducers", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		groupId := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetGroupProducers(groupId)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetAnnouncedGroupProducers", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		groupId := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetAnnouncedGroupProducers(groupId)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetAnnouncedGroupUsers", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		groupId := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetAnnouncedGroupUsers(groupId)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GroupProducer", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsonStr := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GroupProducer([]byte(jsonStr))
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("AddPeers", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return nil
		}
		peersStr := args[0].String()
		peers := strings.Split(peersStr, ",")

		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := api.AddPeers(peers)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("CreateGroup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsonStr := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.CreateGroup([]byte(jsonStr))
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("MgrGrpBlkList", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsonStr := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.MgrGrpBlkList([]byte(jsonStr))
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetDeniedUserList", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		groupId := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetDeniedUserList(groupId)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("Ping", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		peer := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.Ping(peer)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("UpdateProfile", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsonStr := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.UpdateProfile([]byte(jsonStr))
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetTrx", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			// TODO: return a Promise.reject
			return nil
		}
		groupId := args[0].String()
		trxId := args[1].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetTrx(groupId, trxId)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("PostToGroup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsonStr := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.PostToGroup([]byte(jsonStr))
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetNodeInfo", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetNodeInfo()
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetNetwork", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetNetwork()
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetContent", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 4 {
			return nil
		}
		groupId := args[0].String()
		num := args[1].Int()
		startTrx := args[2].String()
		reverse := args[3].Bool()
		senders := []string{}
		for i := 4; i < len(args); i += 1 {
			sender := args[i].String()
			senders = append(senders, sender)
		}

		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetContent(groupId, num, startTrx, reverse, senders)
			if err != nil {
				return ret, err
			}
			retBytes, _ := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("JoinGroup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		seed := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.JoinGroup([]byte(seed))
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("LeaveGroup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		groupId := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.LeaveGroup(groupId)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("ClearGroupData", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		groupId := args[0].String()
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.ClearGroupData(groupId)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetGroups", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetGroups()
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetBlockById", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		groupId := args[0].String()
		blockId := args[1].String()

		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetBlockById(groupId, blockId)
			if err != nil {
				return ret, err
			}
			retBytes, err := json.Marshal(res)
			json.Unmarshal(retBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("GetDecodedBlockById", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		groupId := args[0].String()
		blockId := args[1].String()

		handler := func() (map[string]interface{}, error) {
			ret := make(map[string]interface{})
			res, err := quorumAPI.GetDecodedBlockById(groupId, blockId)
			if err != nil {
				return ret, err
			}
			resBytes, err := json.Marshal(res)
			json.Unmarshal(resBytes, &ret)
			return ret, nil
		}
		return Promisefy(handler)
	}))

	js.Global().Set("IndexDBTest", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go IndexDBTest()
		return js.ValueOf(true).Bool()
	}))
}
