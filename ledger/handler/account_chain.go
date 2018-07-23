package handler

import (
	"github.com/vitelabs/go-vite/protocols"
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"log"
	"github.com/vitelabs/go-vite/crypto"
	"errors"
	"time"
	"math/big"
)

type AccountChain struct {
	vite Vite
	// Handle block
	acAccess *access.AccountChainAccess
	aAccess *access.AccountAccess
	scAccess *access.SnapshotChainAccess
	uAccess *access.UnconfirmedAccess
	tAccess *access.TokenAccess
}

func NewAccountChain (vite Vite) (*AccountChain) {
	return &AccountChain{
		vite: vite,
		acAccess: access.GetAccountChainAccess(),
		aAccess: access.GetAccountAccess(),
		uAccess: access.GetUnconfirmedAccess(),
	}
}

// HandleBlockHash
func (ac *AccountChain) HandleGetBlocks (msg *protocols.GetAccountBlocksMsg, peer *protocols.Peer) error {
	go func() {
		blocks, err := ac.acAccess.GetBlocksFromOrigin(&msg.Origin, msg.Count, msg.Forward)
		if err != nil {
			log.Println(err)
			return
		}

		// send out
		ac.vite.Pm().SendMsg(peer, &protocols.Msg{
			Code: protocols.AccountBlocksMsgCode,
			Payload: blocks,
		})
	}()
	return nil
}


// HandleBlockHash
func (ac *AccountChain) HandleSendBlocks (msg protocols.AccountBlocksMsg, peer *protocols.Peer) error {
	go func() {
		globalRWMutex.RLock()
		defer globalRWMutex.RUnlock()

		for _, block := range msg {
			// Verify signature
			isVerified, verifyErr := crypto.VerifySig(block.PublicKey, block.Hash.Bytes(), block.Signature)

			if verifyErr != nil {
				log.Println(verifyErr)
				continue
			}

			if !isVerified {
				continue
			}

			// Write block
			writeErr := ac.acAccess.WriteBlock(block, nil)
			if writeErr != nil {
				switch writeErr.(type) {
				case access.AcWriteError:
					err := writeErr.(access.AcWriteError)
					if writeErr.(access.AcWriteError).Code == access.WacPrevHashUncorrectErr {
						errData := err.Data.(ledger.AccountBlock)

						if block.Meta.Height.Cmp(errData.Meta.Height) <= 0 {
							return
						}
						// Download fragment
						count := &big.Int{}
						count.Sub(block.Meta.Height, errData.Meta.Height)
						ac.vite.Pm().SendMsg(peer, &protocols.Msg {
							Code: protocols.GetAccountBlocksMsgCode,
							Payload: &protocols.GetAccountBlocksMsg{
								Origin: *errData.Hash,
								Forward: true,
								Count: count.Uint64(),
							},
						})
						return
					}
				}

				log.Println(writeErr)
				continue
			}
		}
	}()
	return nil
}


// AccAddr = account address
func (ac *AccountChain) GetAccountByAccAddr (addr *types.Address) (*ledger.AccountMeta, error) {
	return ac.aAccess.GetAccountMeta(addr)
}

// AccAddr = account address
func (ac *AccountChain) GetBlocksByAccAddr (addr *types.Address, index, num, count int) (ledger.AccountBlockList, error) {
	return ac.acAccess.GetBlockListByAccountAddress(index, num, count, addr)
}

func (ac *AccountChain) CreateTx (block *ledger.AccountBlock) (error) {
	return ac.CreateTxWithPassphrase(block, "")
}

func (ac *AccountChain) CreateTxWithPassphrase (block *ledger.AccountBlock, passphrase string) error {
	globalRWMutex.RLock()
	defer globalRWMutex.RUnlock()

	accountMeta, err := ac.aAccess.GetAccountMeta(block.AccountAddress)

	if err != nil {
		return err
	}

	if accountMeta == nil {
		return errors.New("CreateTx failed, because account " + block.AccountAddress.String() + " is not existed.")
	}


	// Set addr
	block.AccountAddress = block.AccountAddress

	// Set prevHash
	latestBlock, err := ac.acAccess.GetLatestBlockByAccountAddress(block.AccountAddress)
	if err != nil {
		return err
	}

	if latestBlock != nil {
		block.PrevHash = latestBlock.PrevHash
	}

	// Set Snapshot Timestamp
	currentSnapshotBlock, err := ac.scAccess.GetLatestBlock()
	if err != nil {
		return err
	}

	block.SnapshotTimestamp = currentSnapshotBlock.Hash

	// Set Timestamp
	block.Timestamp = uint64(time.Now().Unix())

	// Set Pow params: Nounce、Difficulty
	block.Nounce = []byte{0, 0, 0, 0, 0}
	block.Difficulty = []byte{0, 0, 0, 0, 0}
	block.FAmount = big.NewInt(0)

	writeErr := ac.acAccess.WriteBlock(block, func(accountBlock *ledger.AccountBlock) (*ledger.AccountBlock, error) {
		var signErr error
		if passphrase == "" {
			accountBlock.Signature, accountBlock.PublicKey, signErr =
				ac.vite.WalletManager().KeystoreManager.SignData(*block.AccountAddress, block.Hash.Bytes())

		} else {

			accountBlock.Signature, accountBlock.PublicKey, signErr =
				ac.vite.WalletManager().KeystoreManager.SignDataWithPassphrase(*block.AccountAddress, passphrase, block.Hash.Bytes())
		}

		return accountBlock, signErr
	})
	if err != nil {
		return writeErr
	}

	// Broadcast
	ac.vite.Pm().SendMsg(nil, &protocols.Msg {
		Code: protocols.AccountBlocksMsgCode,
		Payload: &protocols.AccountBlocksMsg{block},
	})
	return nil
}