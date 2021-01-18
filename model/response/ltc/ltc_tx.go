package ltc

/**
 *
 * @export
 * @interface LtcTx
 */
type Tx struct {

	/**
	 * Transaction hash.
	 * @type {string}
	 * @memberof LtcTx
	 */
	Hash string
	/**
	 * Witness hash in case of a SegWit transaction.
	 * @type {string}
	 * @memberof LtcTx
	 */
	WitnessHash string
	/**
	 * Fee paid for this transaction, in LTC.
	 * @type {string}
	 * @memberof LtcTx
	 */
	Fee string
	/**
	 *
	 * @type {string}
	 * @memberof LtcTx
	 */
	Rate string
	/**
	 *
	 * @type {number}
	 * @memberof LtcTx
	 */
	Ps uint32
	/**
	 * Height of the block this transaction belongs to.
	 * @type {number}
	 * @memberof LtcTx
	 */
	Height uint32
	/**
	 * Hash of the block this transaction belongs to.
	 * @type {string}
	 * @memberof LtcTx
	 */
	Block string
	/**
	 * Time of the transaction.
	 * @type {number}
	 * @memberof LtcTx
	 */
	Ts uint64
	/**
	 * Index of the transaction in the block.
	 * @type {number}
	 * @memberof LtcTx
	 */
	Index uint32
	/**
	 * Index of the transaction.
	 * @type {number}
	 * @memberof LtcTx
	 */
	Version uint32
	/**
	 *
	 * @type {number}
	 * @memberof LtcTx
	 */
	Flag uint32
	/**
	 *
	 * @type {Array<LtcTxInputs>}
	 * @memberof LtcTx
	 */
	Inputs []TxInputs
	/**
	 *
	 * @type {Array<LtcTxOutputs>}
	 * @memberof LtcTx
	 */
	Outputs TxOutputs
	/**
	 * Block this transaction was included in.
	 * @type {number}
	 * @memberof LtcTx
	 */
	Locktime uint64
}

/**
 *
 * @export
 * @interface LtcTxCoin
 */
type TxCoin struct {

	/**
	 *
	 * @type {number}
	 * @memberof LtcTxCoin
	 */
	Version uint32
	/**
	 *
	 * @type {number}
	 * @memberof LtcTxCoin
	 */
	Height uint32
	/**
	 *
	 * @type {string}
	 * @memberof LtcTxCoin
	 */
	Value string
	/**
	 *
	 * @type {string}
	 * @memberof LtcTxCoin
	 */
	Script string
	/**
	 * Sender address.
	 * @type {string}
	 * @memberof LtcTxCoin
	 */
	Address string
	/**
	 * Coinbase transaction - miner fee.
	 * @type {boolean}
	 * @memberof LtcTxCoin
	 */
	Coinbase bool
}

/**
 *
 * @export
 * @interface LtcTxInputs
 */
type TxInputs struct {

	/**
	 *
	 * @type {LtcTxPrevout}
	 * @memberof LtcTxInputs
	 */
	Prevout TxPrevout
	/**
	 * Data generated by a spender which is almost always used as variables to satisfy a pubkey script.
	 * @type {string}
	 * @memberof LtcTxInputs
	 */
	Script string
	/**
	 * Transaction witness.
	 * @type {string}
	 * @memberof LtcTxInputs
	 */
	Witness string
	/**
	 *
	 * @type {number}
	 * @memberof LtcTxInputs
	 */
	Sequence uint32
	/**
	 *
	 * @type {LtcTxCoin}
	 * @memberof LtcTxInputs
	 */
	Coin TxCoin
}

/**
 *
 * @export
 * @interface LtcTxOutputs
 */
type TxOutputs struct {

	/**
	 * Sent amount in LTC.
	 * @type {string}
	 * @memberof LtcTxOutputs
	 */
	Value string
	/**
	 * Transaction script.
	 * @type {string}
	 * @memberof LtcTxOutputs
	 */
	Script string
	/**
	 * Recipient address.
	 * @type {string}
	 * @memberof LtcTxOutputs
	 */
	Address string
}

/**
 *
 * @export
 * @interface LtcTxPrevout
 */
type TxPrevout struct {

	/**
	 * Transaction hash of the input.
	 * @type {string}
	 * @memberof LtcTxPrevout
	 */
	Hash string
	/**
	 * Transaction index of the input.
	 * @type {number}
	 * @memberof LtcTxPrevout
	 */
	Index uint32
}