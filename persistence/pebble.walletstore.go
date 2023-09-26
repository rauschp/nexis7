package persistence

import (
	"encoding/json"
	"errors"
	"github.com/cockroachdb/pebble"
	"github.com/rs/zerolog/log"
	"nexis7/types"
	"nexis7/util"
)

type PersistentWalletStore struct {
	DB *pebble.DB
}

type WalletRecord struct {
	Balance   float32
	PublicKey *util.PublicKey
	Nonce     int64
}

func CreatePersistentWalletStore() *PersistentWalletStore {
	db, err := pebble.Open("data/wallet-store", &pebble.Options{})
	if err != nil {
		log.Error().Err(err).Msg("Error opening store")
	}

	return &PersistentWalletStore{
		DB: db,
	}
}

func (w *PersistentWalletStore) GetByPublicKey(pk *util.PublicKey) (*types.Wallet, error) {
	addr := pk.GetAddress()

	return w.GetByAddress(addr)
}

func (w *PersistentWalletStore) GetByAddress(address util.Address) (*types.Wallet, error) {
	record, err := w.getWalletRecord(address)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching wallet")
		return nil, err
	}

	return &types.Wallet{
		Address:   address,
		PublicKey: record.PublicKey,
		Balance:   record.Balance,
		Nonce:     record.Nonce,
	}, nil
}

func (w *PersistentWalletStore) DepositCurrency(address util.Address, amount float32) error {
	record, err := w.getWalletRecord(address)
	if err != nil {
		log.Error().Err(err).Msg("Error getting wallet")
		return err
	}

	newRecord := WalletRecord{
		Balance:   record.Balance + amount,
		PublicKey: record.PublicKey,
		Nonce:     record.Nonce + 1,
	}

	if err := w.saveWalletRecord(address, newRecord); err != nil {
		return err
	}

	return nil
}

func (w *PersistentWalletStore) WithdrawCurrency(address util.Address, amount float32) error {
	record, err := w.getWalletRecord(address)
	if err != nil {
		log.Error().Err(err).Msg("Error getting wallet")
		return err
	}

	if record.Balance < amount {
		err = errors.New("insufficient funds")
		log.Error().Err(err)
		return err
	}

	newRecord := WalletRecord{
		Balance:   record.Balance - amount,
		PublicKey: record.PublicKey,
		Nonce:     record.Nonce + 1,
	}

	if err := w.saveWalletRecord(address, newRecord); err != nil {
		log.Error().Err(err).Msg("error saving wallet")
		return err
	}

	return nil
}

func (w *PersistentWalletStore) getWalletRecord(address util.Address) (*WalletRecord, error) {
	value, closer, err := w.DB.Get(address.ToBytes())
	if err != nil {
		log.Error().Err(err).Msg("Error getting height value")
		return nil, err
	}

	if err := closer.Close(); err != nil {
		log.Error().Err(err).Msg("Error closing reader")
		return nil, err
	}

	var record = &WalletRecord{}
	if err = json.Unmarshal(value, record); err != nil {
		log.Error().Err(err).Msg("Error parsing kvs result")
		return nil, err
	}

	return record, nil
}

func (w *PersistentWalletStore) saveWalletRecord(address util.Address, record WalletRecord) error {
	key := address.ToBytes()
	value, err := json.Marshal(record)
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling wallet record to save")
		return err
	}

	if err := w.DB.Set(key, value, pebble.Sync); err != nil {
		log.Error().Err(err).Msg("Error saving record to kvs")
		return err
	}

	return nil
}
